package main

import (
	"errors"

	"github.com/qianlnk/log/forwarder/redis"
	"github.com/qianlnk/log/forwarder/std"
)

func newOutput(config *ioConfig) (output, error) {
	switch config.Type {
	case "redis":
		return redis.NewOutput(&config.Redis)
	case "std":
		return std.NewOutput(), nil
	}

	return nil, errors.New("unsupported output type: " + config.Type)
}
