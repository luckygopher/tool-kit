package config

import (
	"github.com/jinzhu/configor"
	"github.com/qingyunjun/tool-kit/pkg/db"
	"github.com/qingyunjun/tool-kit/pkg/httpclient"
	"github.com/qingyunjun/tool-kit/pkg/vaccine"
	"go.uber.org/zap"
)

var C = Config{}

type Config struct {
	ENV        string            `toml:"ENV" yaml:"env"`
	Debug      bool              `toml:"Debug" yaml:"debug"`
	LogLevel   string            `toml:"LogLevel" yaml:"log_level"`
	HTTPClient httpclient.Config `toml:"HTTPClient" yaml:"http_client"`
	Database   db.Config         `toml:"Database" yaml:"database"`
	Vaccine    vaccine.Config    `toml:"Vaccine" yaml:"vaccine"`
}

func ParseConfig(filePath string) {
	if filePath != "" {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(&C, filePath); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	} else {
		if err := configor.New(&configor.Config{AutoReload: false}).Load(&C); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	}
}
