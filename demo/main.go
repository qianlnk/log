package main

import (
	"time"

	"github.com/qianlnk/log"
)

var order map[string]string

func init() {
	order["mch_no"] = ""
	order["trade_no"] = ""
	order["order_fee"] = ""
}

//run it as follow: nohub ./demo > a &
func main() {
	for {
		log.Info("test")
		log.Fields{
			"name": "qianlnk",
			"age":  "27",
		}.Info("test fields")
		log.Error("test error")
		time.Sleep(time.Second * 2)
		test := make(map[string]interface{})
		test["lala"] = "haha"
		test["number"] = 123
		log.Fields{
			"aaa": "aaa",
			"bbb": "bbb",
		}.Add(log.Fields{
			"ddd": "ddd",
			"eee": "eee",
		}).Info("ccc")

		log.Fields{}.Add(log.Fields(test)).Del("number", "lala").Info("bababaa")
		log.Fields(test).Error("testErr")
	}
}
