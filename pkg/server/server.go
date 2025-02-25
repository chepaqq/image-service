package server

import (
	"context"
	"net/http"
	"time"

	"github.com/chepaqq/image-service/pkg/logger"
)

const (
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultShutdownTimeout = 15 * time.Second
	defaultIdleTimeout     = 1 * time.Minute
)

// Server represents http server.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New creates instance of new http server.
func New(handler http.Handler, port string) *Server {
	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		IdleTimeout:  defaultIdleTimeout,
	}

	srv := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: defaultShutdownTimeout,
	}
	srv.start()
	return srv
}

// Start bootstraps http server.
func (s *Server) start() {
	logger.Infof("Starting HTTP server on port %s", s.server.Addr)
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify returns error notification channel.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown shuts down http server gracefully.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
