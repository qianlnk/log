package main

import (
	"github.com/qianlnk/log/forwarder/redis"
	"github.com/qianlnk/log/forwarder/std"
)

func newInput(config *ioConfig) (input, error) {
	switch config.Type {
	case "redis":
		return redis.NewInput(&config.Redis)
	}
	return std.NewInput(), nil
}
