package main

import (
	"appcoachs.net/x/log/log-forwarder/redis"
	"appcoachs.net/x/log/log-forwarder/std"
)

func newInput(config *ioConfig) (input, error) {
	switch config.Type {
	case "redis":
		return redis.NewInput(&config.Redis)
	}
	return std.NewInput(), nil
}
