package http

import (
	"context"
	"errors"
	"net/http"
	"pongsatorn/go-http-template/internal/adapter/config"
	"pongsatorn/go-http-template/pkg/logger"
	"time"
)

const shutDownTimeout = 10 * time.Second

type Server struct {
	srv *http.Server
}

func NewServer(cfg config.HttpServerConfig, echoWrapper *EchoWrapper) (*Server, func()) {
	server := &Server{}

	srvCfg := config.ToServerConfig(cfg)
	server.srv = &http.Server{
		Addr:         srvCfg.Port,
		Handler:      echoWrapper,
		ReadTimeout:  srvCfg.ReadTimeout,
		WriteTimeout: srvCfg.WriteTimeout,
		IdleTimeout:  srvCfg.IdleTimeout,
	}

	return server, func() {
		server.Shutdown()
	}
}

func (s *Server) Start(serverErrCh chan<- error) {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Logger.Error(err)
			}

			serverErrCh <- err
		}
	}()
}

func (s *Server) Shutdown() {
	logger.Logger.Info("prepare to shutdown http server")
	ctx, cancelFn := context.WithTimeout(context.Background(), shutDownTimeout)
	defer cancelFn()

	err := s.srv.Shutdown(ctx)
	if err != nil {
		logger.Logger.Error(err)
	}
}
