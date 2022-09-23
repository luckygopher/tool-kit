package httpclient

import (
	"time"
)

type Config struct {
	Debug               bool   `toml:"Debug" yaml:"debug"`
	Verbose             bool   `toml:"Verbose" yaml:"verbose" default:"false"`                           // 请求输出详细的调试信息
	Timeout             string `toml:"Timeout" yaml:"timeout" default:"5s"`                              // 请求超时时间
	DialTimeout         string `toml:"DialTimeout" yaml:"dial_timeout" default:"30s"`                    // 连接池DialTimeout
	IdleConnTimeout     string `toml:"IdleConnTimeout" yaml:"idle_conn_timeout" default:"90s"`           // 连接空闲最长时间
	MaxIdleConns        int    `toml:"MaxIdleConns" yaml:"max_idle_conns" default:"0"`                   // 最大空闲
	MaxIdleConnsPerHost int    `toml:"MaxIdleConnsPerHost" yaml:"max_idle_conns_per_host" default:"500"` // 单个host最大空闲
	MaxConnsPerHost     int    `toml:"MaxConnsPerHost" yaml:"max_conns_per_host" default:"2000"`         // 单个host允许最大连接
}

func (cfg Config) GetHTTPTimeout() time.Duration {
	t, err := time.ParseDuration(cfg.Timeout)
	if err != nil || t.Seconds() == 0 {
		t = DefaultTimeout
	}

	return t
}

func (cfg Config) GetHTTPDialTimeout() time.Duration {
	t, err := time.ParseDuration(cfg.DialTimeout)
	if err != nil || t.Seconds() == 0 {
		t = DefaultDialTimeout
	}

	return t
}

func (cfg Config) GetHTTPIdleConnTimeout() time.Duration {
	t, err := time.ParseDuration(cfg.IdleConnTimeout)
	if err != nil || t.Seconds() == 0 {
		t = DefaultIDleConnTimeout
	}

	return t
}
