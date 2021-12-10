package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func init() {
	var logger *zap.Logger
	logger, _ = zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	Log = logger.Sugar()
}
