package logging

import (
	"sync"

	"github.com/avayayu/micro/logging"
	"github.com/avayayu/quant_data/internal/configs"
	"go.uber.org/zap"
)

var logger *zap.Logger
var once sync.Once

func newLogger() *zap.Logger {

	configs := configs.GetConfigs()
	logPath := configs.Get("logger.logPath")
	logLevel := configs.Get("logger.logLevel")

	logger := logging.NewProdLoggger(logPath, logLevel)

	return logger
}

func GetLogger() *zap.Logger {
	once.Do(func() {
		if logger == nil {
			logger = newLogger()
		}
	})
	return logger
}
