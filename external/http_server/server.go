package http_server

import (
	"net/http"
	"time"
)

type Option func(server *Server)

type Server struct {
	httpServer http.Server
}

func New(handler http.Handler, opt ...Option) *Server {
	const (
		defaultReadWriteTimeout = 10 * time.Second
		defaultAddr             = "8080"
		defaultIdleTimeout      = 30 * time.Second
	)

	server := Server{httpServer: http.Server{
		Addr:         defaultAddr,
		ReadTimeout:  defaultReadWriteTimeout,
		WriteTimeout: defaultReadWriteTimeout,
		IdleTimeout:  defaultIdleTimeout,
		Handler:      handler,
	}}

	for _, option := range opt {
		option(&server)
	}

	return &server
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func WithIdleTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.httpServer.IdleTimeout = timeout
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.httpServer.WriteTimeout = timeout
		s.httpServer.ReadTimeout = timeout
	}
}
