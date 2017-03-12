package main

import (
	"fmt"
	"os"
	"time"

	"github.com/qianlnk/config"
)

type conf struct {
	Input  ioConfig
	Output ioConfig
}

func main() {
	var conf conf
	checkError(config.Parse(&conf))
	forwarder, err := newForwarder(&conf)
	checkError(err)
	checkError(forwarder.Forward())
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		time.Sleep(time.Second)
		os.Exit(-1)
	}
}
