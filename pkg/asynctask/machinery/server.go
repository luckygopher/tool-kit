package machinery

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

var MachineServer *machinery.Server

// StartServer 启动machinery服务
func StartServer(conf *MachineConfig) error {
	cnf := &config.Config{
		Broker:          conf.Broker,
		DefaultQueue:    conf.DefaultQueue,
		ResultBackend:   conf.ResultBackend,
		ResultsExpireIn: conf.ResultsExpireIn,
		Redis: &config.RedisConfig{
			MaxIdle:                conf.Redis.MaxIdle,
			MaxActive:              conf.Redis.MaxActive,
			IdleTimeout:            conf.Redis.IdleTimeout,
			Wait:                   conf.Redis.Wait,
			ReadTimeout:            conf.Redis.ReadTimeout,
			WriteTimeout:           conf.Redis.WriteTimeout,
			ConnectTimeout:         conf.Redis.ConnectTimeout,
			NormalTasksPollPeriod:  conf.Redis.NormalTasksPollPeriod,
			DelayedTasksPollPeriod: conf.Redis.DelayedTasksPollPeriod,
		},
	}
	server, err := machinery.NewServer(cnf)
	if err != nil {
		return err
	}
	MachineServer = server
	return nil
}
