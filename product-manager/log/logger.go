package log

import "go.uber.org/zap"

var Logger *zap.Logger

func CreateLogger(cfg Config) {

	switch cfg.Environment {
	case prod:
		Logger, _ = zap.NewProduction()
	case dev:
		Logger, _ = zap.NewDevelopment()
	default:
		Logger, _ = zap.NewProduction()
	}
	defer Logger.Sync()
}
