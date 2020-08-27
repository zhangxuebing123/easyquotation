package easyquotation

import (
	"easyquotation/sina"
	"easyquotation/stock"
	"fmt"
	"github.com/gocolly/colly"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	Init()

	c := colly.NewCollector()
	SinaStock_spider := sina.NewSinaStock(c)
	SinaStock_spider.Start()

	tt := time.NewTicker(time.Second * 3)
	for {
		select {
		case <- tt.C:
			fmt.Println(stock.G_STOCK_MANAGER.StockList["sh600000"])
		}
	}
}
