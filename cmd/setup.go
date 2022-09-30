package cmd

import (
	"github.com/luckygopher/tool-kit/config"
	"github.com/luckygopher/tool-kit/pkg/httpclient"
	"go.uber.org/zap"
)

// 初始化设置函数
type setupFunc func() error

// Setup 初始化设置函数切片
type Setup []setupFunc

// 执行初始化设置
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
