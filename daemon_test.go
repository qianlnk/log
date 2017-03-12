package log_test

import (
	"testing"
	"time"

	"github.com/qianlnk/log"
)

func TestServe(t *testing.T) {
	if testing.Short() {
		t.Skip("skip deamon test")
	}
	for i := 0; i < 10; i++ {
		log.Error("")
		log.Warn("")
		log.Info("")
		log.Debug("")
		time.Sleep(time.Second * 2)
	}
}
