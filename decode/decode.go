package decode

import (
	"easyquotation/stock"
	"easyquotation/utils"
	"fmt"
	"reflect"
	"strings"
)

var str string = "var hq_str_sz161505=\"银河通利,0.000,1.258,0.000,0.000,0.000,1.239,1.280,0,0.000,3000,1.239,8100,1.234,1000,1.233,1000,1.198,900,1.188,5548,1.280,900,1.300,1000,1.314,1000,1.326,1600,1.330,2020-07-24,10:58:12,00\""

func Split(r rune) bool {
	return r == ' ' || r == '_' || r == '='
}
func DecodeMarket() {
	l := utils.SplitString(str, Split)
	symbol := l[3]

	market := l[4][1 : len(l[4])-1]
	marketList := strings.Split(market, ",")

	a := stock.Market{}
	eleVals := reflect.ValueOf(&a).Elem()
	utils.DecodeStock(marketList[:10], eleVals)
	for i := 0; i < 10; i += 2 {
		bEntry := stock.Entry{}
		vals := reflect.ValueOf(&bEntry).Elem()
		utils.DecodeStock(marketList[10+i:10+i+2], vals)
		a.BuyEntryList = append(a.BuyEntryList, bEntry)
	}
	for i := 0; i < 10; i += 2 {
		bEntry := stock.Entry{}
		vals := reflect.ValueOf(&bEntry).Elem()
		utils.DecodeStock(marketList[20+i:20+i+2], vals)
		a.SellEntryList = append(a.SellEntryList, bEntry)
	}
	a.Date = marketList[30]
	a.Time = marketList[31]
	a.Flag = marketList[32]

	fmt.Println(a)
}
