package redis

import (
	"bufio"
	"io"
	"os"

	"appcoachs.net/x/log/log-forwarder/redis"
)

type Config struct {
	Addr     string
	Password string
	ListKey  string
}

func New(config *Config) (io.Writer, error) {
	output, err := redis.NewOutput(&redis.Config{
		Addr:     config.Addr,
		Password: config.Password,
		ListKey:  config.ListKey,
	})
	if err != nil {
		return nil, err
	}
	r, w := io.Pipe()
	writer := &redisWriter{
		reader: r,
		writer: w,
		output: output,
	}
	go writer.Forward()
	return writer, nil
}

type redisWriter struct {
	reader *io.PipeReader
	writer *io.PipeWriter
	output *redis.Output
}

func (r *redisWriter) Close() error {
	if r.reader != nil {
		r.reader.Close()
		r.reader = nil
	}
	if r.writer != nil {
		r.writer.Close()
		r.writer = nil
	}
	if r.output != nil {
		r.output.Close()
		r.output = nil
	}
	return nil
}

func (r *redisWriter) Write(b []byte) (int, error) {
	return r.writer.Write(b)
}

func (w *redisWriter) Forward() {
	s := bufio.NewScanner(w.reader)
	for s.Scan() {
		os.Stderr.Write(s.Bytes())
		os.Stderr.Write([]byte{'\n'})
		if err := w.output.Send(s.Bytes()); err != nil {
			os.Stderr.Write([]byte(err.Error()))
		}
	}
}
