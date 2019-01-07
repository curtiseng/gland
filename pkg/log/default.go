package log

import (
	"fmt"

	"go.uber.org/zap"
)

func Error(msg string, args ...interface{}) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	logger.Error(msg)
}