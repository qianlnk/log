package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type Config struct {
	Addr     string
	Password string
	ListKey  string
}

type Output struct {
	pool    *redis.Pool
	listKey string
}

func NewOutput(config *Config) (*Output, error) {
	return &Output{
		pool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", config.Addr)
				if err != nil {
					return nil, err
				}
				if _, err := c.Do("AUTH", config.Password); err != nil {
					c.Close()
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
		listKey: config.ListKey,
	}, nil
}
func (config *Config) Connect() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", config.Addr)
	if err != nil {
		return nil, err
	}
	if config.Password != "" {
		if _, err := conn.Do("AUTH", config.Password); err != nil {
			conn.Close()
			return nil, err
		}
	}
	return conn, nil
}

func (r *Output) Send(item []byte) error {
	c := r.pool.Get()
	defer c.Close()
	err := c.Send("RPUSH", r.listKey, item)
	if err != nil {
		return err
	}
	return nil
}

func (r *Output) Close() error {
	return r.pool.Close()
}
