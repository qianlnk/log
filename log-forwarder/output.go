package main

import (
	"errors"

	"appcoachs.net/x/log/log-forwarder/redis"
)

func newOutput(config *ioConfig) (output, error) {
	switch config.Type {
	case "redis":
		return redis.NewOutput(&config.Redis)
	}
	return nil, errors.New("unsupported output type: " + config.Type)
}