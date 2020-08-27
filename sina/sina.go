package sina

import (
	"github.com/gocolly/colly"
	"time"
)

type ISpiderIntface interface {
	Url() string
	Collector(c *colly.Collector)
	OnRequest(r *colly.Request)
	OnResponse(res *colly.Response)
	Start()
}

type Spider struct {
	C     *colly.Collector
	Timer *time.Ticker // 个股
}
