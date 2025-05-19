package logger

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func InitLogger() error {
	var logger *zap.Logger
	var err error

	logger, err = zap.NewProduction()
	if err != nil {
		return err
	}

	Logger = logger.Sugar()

	return nil
}
