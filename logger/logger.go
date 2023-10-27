package logger

import (
	"go.uber.org/zap"
	"log"
)

// Log logger
var Log Logger

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	Log.SugaredLogger = logger.Sugar()
}

type Logger struct {
	*zap.SugaredLogger
}
