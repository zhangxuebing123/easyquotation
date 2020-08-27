package utils

import (
	"fmt"
	"testing"
)

func TestUpdateStockCodes(t *testing.T) {
	UpdateStockCodes()
}

func TestStockMarket(t *testing.T) {

}

func TestRandom(t *testing.T) {
	for i:=0; i< 100; i++  {
		fmt.Println(Random(5))
	}
}