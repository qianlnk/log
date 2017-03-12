package main

import (
	"agithub.com/qianlnk/log/log-forwarder/redis"
	"github.com/qianlnk/log/log-forwarder/std"
)

func newInput(config *ioConfig) (input, error) {
	switch config.Type {
	case "redis":
		return redis.NewInput(&config.Redis)
	}
	return std.NewInput(), nil
}
