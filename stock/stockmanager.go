package stock

import (
	"bufio"
	"easyquotation/utils"
	"fmt"
	"os"
	"reflect"
	"strings"
)

var G_STOCK_MANAGER *StockManager

type StockManager struct {
	StockList   map[string]*Stock //个股
	IndexList   map[string]*Stock //指数
	BKkList     map[string]*Stock //板块
	ZZIndexList map[string]*Stock //中证指数
}

func (s *StockManager) SetBase(list []string) *Stock {
	b := Base{}
	eleVals := reflect.ValueOf(&b).Elem()
	utils.DecodeStock(list, eleVals)
	return &Stock{Base: b}
}

// 加载股票信息
func (s *StockManager) Load() {
	path, err := os.Getwd()
	if err != nil {
		path = "./"
	}
	if file, err := os.Open(path + "/stock_codes"); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			eleList := strings.Split(line, " ")
			if len(eleList) < 3 {
				continue
			} else {
				if strings.HasPrefix(eleList[0], "zz") {
					s.ZZIndexList[eleList[0]] = s.SetBase(eleList)
				} else {
					if utils.EndSwitch(eleList[1], []string{".沪指数", ".深指数"}) {
						if !strings.HasPrefix(eleList[0], "sh") {
							eleList[0] = fmt.Sprintf("sz%s", eleList[0])
						}
						s.IndexList[eleList[0]] = s.SetBase(eleList)
					} else if strings.HasSuffix(eleList[1], ".板块") {
						s.BKkList[eleList[0]] = s.SetBase(eleList)
					} else {
						if utils.StartSwitch(eleList[0], []string{"5", "6", "9"}) {
							eleList[0] = fmt.Sprintf("sh%s", eleList[0])
						} else {
							eleList[0] = fmt.Sprintf("sz%s", eleList[0])
						}
						s.StockList[eleList[0]] = s.SetBase(eleList)
					}
				}
			}
		}
	}
}
