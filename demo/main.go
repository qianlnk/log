package main

import (
	"time"

	"github.com/qianlnk/log"
)

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
	}
}
