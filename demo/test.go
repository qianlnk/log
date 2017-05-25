package main

import (
	"github.com/qianlnk/log"
)

type Test struct {
	ID       string
	MchID    string
	OrderID  string
	OrderFee int64
}

func main() {
	t := Test{
		ID:       "1001",
		MchID:    "100010",
		OrderID:  "10000000001",
		OrderFee: 100,
	}

	log.Struct(t).Info("test")
}
