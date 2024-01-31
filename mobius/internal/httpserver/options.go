package httpserver

import (
	"fmt"
	"time"
)

type Option func(*Server)

func WithAddr(port int) Option {
	return func(s *Server) {
		s.server.Addr = fmt.Sprintf(":%d", port)
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
