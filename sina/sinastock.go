package sina

import (
	"easyquotation/stock"
	"easyquotation/utils"
	"fmt"
	"github.com/gocolly/colly"
	"reflect"
	"strings"
	"time"
)

type SinaStock struct {
	Spider
}

func NewSinaStock(c *colly.Collector) *SinaStock {
	s := &SinaStock{Spider{Timer: time.NewTicker(time.Second * 3)}}
	s.Collector(c)
	s.C.OnRequest(s.OnRequest)
	s.C.OnResponse(s.OnResponse)
	return s
}
func (s *SinaStock) Collector(c *colly.Collector) {
	s.Spider.C = c.Clone()
}
func (s *SinaStock) Url() string {
	return fmt.Sprintf("http://hq.sinajs.cn/rn=%d&list=", time.Now().Unix())
}

func (s *SinaStock) OnRequest(r *colly.Request) {
	//	fmt.Println(r.URL)
}

func (s *SinaStock) Split(r rune) bool {
	return r == ' ' || r == '_' || r == '='
}
func (s *SinaStock) DecodeMarket(str string) {
	l := utils.SplitString(str, s.Split)
	if (len(l) < 5){
		return
	}
	symbol := l[3]					 // 代码
	market := l[4][1 : len(l[4])-1]	 // 去除左右" "
	marketList := strings.Split(market, ",")
	if (len(marketList) < 33) {
		return
	}
	a := stock.Market{}
	eleVals := reflect.ValueOf(&a).Elem()
	utils.DecodeStock(marketList[:10], eleVals)
	// 买五档
	for i := 0; i < 10; i += 2 {
		bEntry := stock.Entry{}
		vals := reflect.ValueOf(&bEntry).Elem()
		utils.DecodeStock(marketList[10+i:10+i+2], vals)
		a.BuyEntryList = append(a.BuyEntryList, bEntry)
	}
	// 卖五档
	for i := 0; i < 10; i += 2 {
		bEntry := stock.Entry{}
		vals := reflect.ValueOf(&bEntry).Elem()
		utils.DecodeStock(marketList[20+i:20+i+2], vals)
		a.SellEntryList = append(a.SellEntryList, bEntry)
	}
	a.Date = marketList[30]
	a.Time = marketList[31]
	a.Flag = marketList[32]
	stock.G_STOCK_MANAGER.StockList[symbol].Market = a
}

func (s *SinaStock) OnResponse(res *colly.Response) {
	l := strings.Split(string(res.Body), ";")
	for _, str := range l {
		s.DecodeMarket(str)
	}
}

func (s *SinaStock) Start() {
	go func() {
		for {
			select {
			case <-s.Spider.Timer.C:
				ipos := 1
				params := make([]string, 0)
				param := ""
				for k, _ := range stock.G_STOCK_MANAGER.StockList {
					if ipos%800 == 0 {
						params = append(params, param)
						param = ""
					}
					param += k
					param += ","
					ipos++
				}
				params = append(params, param)
				for _, str := range params {
					url := s.Url() + str
					go s.C.Visit(url)
				}
			}
		}
	}()
}
