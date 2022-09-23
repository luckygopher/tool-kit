package cmd

import (
	"github.com/qingyunjun/tool-kit/config"
	"github.com/qingyunjun/tool-kit/pkg/httpclient"
	"go.uber.org/zap"
)

type SetupFunc func() error
type Setup []SetupFunc

func (s Setup) apply() {
	for _, fn := range s {
		if err := fn(); err != nil {
			zap.L().Fatal("apply setup failed", zap.Error(err))
		}
	}
}

// SetLogger 设置日志
func SetLogger() error {
	var (
		logger *zap.Logger
		conf   zap.Config
		level  = zap.NewAtomicLevel()
		err    error
	)
	if config.C.ENV == ENV_PROD {
		conf = zap.NewProductionConfig()
	} else {
		conf = zap.NewDevelopmentConfig()
	}
	if err = level.UnmarshalText([]byte(config.C.LogLevel)); err != nil {
		return err
	}
	conf.Level = level

	if logger, err = conf.Build(); err != nil {
		return err
	}

	zap.RedirectStdLog(logger)
	zap.ReplaceGlobals(logger)
	return nil
}

// SetupHTTP 设置 httpclient
func SetupHTTP() error {
	httpclient.InitHTTPClient(config.C.HTTPClient, zap.L())
	return nil
}
