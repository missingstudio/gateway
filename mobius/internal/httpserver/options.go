package httpserver

import (
	"fmt"
	"time"
)

type Option func(*Server)

func WithAddr(host string, port int) Option {
	return func(s *Server) {
		s.server.Addr = fmt.Sprintf("%s:%d", host, port)
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

func WithShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
