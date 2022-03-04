package internalhttp

import (
	"context"
	"net"
	"net/http"
)

type Server struct {
	host   string
	port   string
	logger Logger
	server *http.Server
}

type Logger interface {
	Info(message string, params ...interface{})
	Error(message string, params ...interface{})
	LogHTTPRequest(r *http.Request, code, length int)
}

type Application interface { // TODO
}

func NewServer(logger Logger, app Application, host, port string) *Server {
	server := &Server{
		host:   host,
		port:   port,
		logger: logger,
		server: nil,
	}

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: loggingMiddleware(http.HandlerFunc(server.handleHTTP), logger),
	}

	server.server = httpServer

	return server
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("HTTP server listen and serve %s:%s", s.host, s.port)
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Server) handleHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
