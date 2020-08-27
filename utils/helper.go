package utils

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
	"errors"
)

func UpdateStockCodes() {
	path, err := os.Getwd()
	if err != nil {
		path = "./"
	}
	req, err := http.NewRequest("GET", "http://www.shdjt.com/js/lib/astock.js", nil)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	list := strings.Split(string(respBody), "=")
	list = strings.Split(list[1][1:len(list[1])-1], "~")
	if fileObj, err := os.OpenFile(path+"/stock_codes", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); err == nil {
		defer fileObj.Close()

		realList := list[1:]
		realLen := len(realList)
		for index0, item := range realList {
			l := strings.Split(item, "`")
			for _, s := range l {
				fileObj.WriteString(s + " ")
			}
			if index0 != realLen-1 {
				fileObj.WriteString("\n")
			}

		}
	}
}

func StartSwitch(s string, prefixs []string) bool {
	for _, prefix := range prefixs {
		if strings.HasPrefix(s, prefix) == true {
			return true
		}
	}
	return false
}

func EndSwitch(s string, suffixs []string) bool {
	for _, suffix := range suffixs {
		if strings.HasSuffix(s, suffix) == true {
			return true
		}
	}
	return false
}

// 生成随机串
// ref:https://cloud.tencent.com/developer/article/1580647
var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func Random(n int) (s string) {

	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func Min(a ...int) int {
	m := int(^uint(0) >> 1)
	for _, i := range a {
		if i < m {
			m = i
		}
	}
	return m
}

func SplitString(s string, f func(rune) bool) []string {
	a := strings.FieldsFunc(s, f)
	return a
}

func DecodeStock(list []string, data reflect.Value) {
	length := Min(len(list), data.NumField())
	for i := 0; i < length; i++ {
		switch data.Field(i).Kind() {
		case reflect.String:
			data.Field(i).SetString(list[i])
		case reflect.Int:
			i64, _ := strconv.ParseInt(list[i], 10,64)
			data.Field(i).SetInt(i64)
		case reflect.Float64:
			f, _ := strconv.ParseFloat(list[i], 64)
			data.Field(i).SetFloat(f)
		default:
			panic(errors.New("Decode error"))
		}
	}
}