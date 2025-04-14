package logger

import (
	"go.uber.org/zap"
)

var Logger = zap.NewNop()

func init() {
	err := Init(Config{})
	if err != nil {
		panic(err)
	}
}

func Init(cfg Config) error {
	var config zap.Config
	config = zap.NewDevelopmentConfig()

	var err error
	Logger, err = config.Build(zap.AddCallerSkip(1))
	return err
}
