package easyquotation

import (
	. "easyquotation/stock"
	"easyquotation/utils"
	"log"
	"os"
)

func Init() {
	path, err := os.Getwd()
	file := path + "/message" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[quotation]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	// 获取新浪股票代码信息
	utils.UpdateStockCodes()
	//
	G_STOCK_MANAGER = &StockManager{
		make(map[string]*Stock),
		make(map[string]*Stock),
		make(map[string]*Stock),
		make(map[string]*Stock),
	}
	G_STOCK_MANAGER.Load()
}
