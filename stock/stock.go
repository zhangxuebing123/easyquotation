package stock

const (
	TICK_DAY_NUM = 4*60 + 2 // 分时数据量
)

type Base struct {
	Symbol string // 代码
	Name   string // 全称
	Short  string // 简称
}

type Entry struct {
	Qty   float64 `json:"qty"`   // 数量
	Price float64 `json:"price"` // 价格
}

type Tick struct {
	Time      string  // 时间
	LastPrice float64 // 最新价
	AvgPrice  float64 // 均价
	Volumn    float64 // 成交量
	Value     float64 // 成交额
}

type Market struct {
	Name          string  `json:"name"`            // 名称
	Open          float64 `json:"open"`            // 开盘
	PreClose      float64 `json:"pre_close"`       // 昨收
	LastPrice     float64 `json:"last_price"`      // 最新价
	High          float64 `json:"high"`            // 最高价
	Low           float64 `json:"low"`             // 最低价
	BidPice       float64 `json:"bid_pice"`        // 竞买价
	OfferPice     float64 `json:"offer_pice"`      // 竞卖价
	Volumn        float64 `json:"volumn"`          // 成交量(股)
	Value         float64 `json:"value"`           // 成交额(元)
	BuyEntryList  []Entry `json:"buy_entry_list"`  // 买5档
	SellEntryList []Entry `json:"sell_entry_list"` // 卖5档
	Date          string  `json:"date"`            // 日期
	Time          string  `json:"time"`            // 时间
	Flag          string  `json:"flag"`            // 状态
}

type Stock struct {
	Base
	Market
}
