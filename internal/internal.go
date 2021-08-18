package internal

import "go.uber.org/zap"

var sugar *zap.SugaredLogger

func init() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any

	sugar = logger.Sugar()
}
