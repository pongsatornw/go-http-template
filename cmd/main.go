package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"pongsatorn/go-http-template/di"
	"pongsatorn/go-http-template/pkg/logger"
	"syscall"
)

func main() {
	err := logger.InitLogger()
	if err != nil {
		log.Fatalf("fail to init logger: %v", err)
	}

	defer func() { _ = logger.Logger.Sync() }()

	container, cleanUpFn, err := di.InitApplication()
	defer cleanUpFn()

	if err != nil {
		logger.Logger.Error(fmt.Errorf("failed to start container, %v", err))
		return
	}

	logger.Logger.Info("start container successfully...")
	stopSignal := make(chan os.Signal, 1)
	httpServerErrCh := make(chan error, 1)
	signal.Notify(stopSignal, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		container.HttpServer.Start(httpServerErrCh)
	}()

	select {
	case <-stopSignal:
		logger.Logger.Info("received interupt/terminate signal. prepare for gracefully shutdown...")
	case err = <-httpServerErrCh:
		logger.Logger.Error(fmt.Errorf("failed to start server: %v", err))
	}
}
