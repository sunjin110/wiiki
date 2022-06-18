package logger

import (
	"wiiki_server/common/config"
	"wiiki_server/common/logger"
	"wiiki_server/common/wiikierr"

	"go.uber.org/zap"
)

func New(conf *config.WiikiConfig) (logger.WiikiLogger, error) {
	var zapLogger *zap.Logger
	switch conf.Env {
	case config.EnvProduction:
		zapLoggerPrd, err := zap.NewProduction()
		if err != nil {
			return nil, wiikierr.Bind(err, wiikierr.FailedNewLogger, "")
		}
		zapLogger = zapLoggerPrd
	default:
		zapLoggerDevelop, err := zap.NewDevelopment()
		if err != nil {
			return nil, wiikierr.Bind(err, wiikierr.FailedNewLogger, "")
		}
		zapLogger = zapLoggerDevelop
	}

	return zapLogger.Sugar(), nil
}
