package sina

import (
	"easyquotation/stock"
	"easyquotation/utils"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"time"
)

type SinaIndex struct {
	Spider
}

func (s *SinaIndex) Url() string {
	return fmt.Sprintf("http://hq.sinajs.cn/rn=%s&list=", utils.Random(5))
}

func NewSinaIndex(c *colly.Collector) *SinaIndex {
	s := &SinaIndex{Spider{Timer:time.NewTicker(time.Second * 10)}}
	s.Collector(c)
	s.C.OnRequest(s.OnRequest)
	s.C.OnResponse(s.OnResponse)
	return s
}
func (s *SinaIndex) Collector(c *colly.Collector) {
	s.Spider.C = c.Clone()
}

func (s *SinaIndex) OnRequest(r *colly.Request) {
}

func (s *SinaIndex) OnResponse(res *colly.Response) {
	log.Println(string(res.Body))
}

func (s *SinaIndex) Start() {
	go func() {
		for {
			select {
			case <- s.Spider.Timer.C:
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