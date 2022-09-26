package machinery

import "time"

type MachineConfig struct {
	Broker          string      `toml:"Broker" yaml:"broker"`                                // 消息代理 redis://[password@]host[port][/db_num]
	DefaultQueue    string      `toml:"DefaultQueue" yaml:"default_queue"`                   // 默认队列的名字
	ResultBackend   string      `toml:"ResultBackend" yaml:"result_backend"`                 // 保存任务的状态和结果集 redis://[password@]host[port][/db_num]
	ResultsExpireIn int         `toml:"ResultsExpireIn" yaml:"results_expire_in"`            // 任务结果的存储时间，以秒为单位。默认值为3600(1小时)
	TaskRetryTime   string      `toml:"TaskRetryTime" yaml:"task_retry_time" default:"180s"` // 任务重试时间
	DelayTaskTime   string      `toml:"DelayTaskTime" yaml:"delay_task_time" default:"300s"` // 延时任务多久之后执行
	Redis           RedisConfig `toml:"Redis" yaml:"redis"`
}

type RedisConfig struct {
	MaxIdle                int  `toml:"MaxIdle" yaml:"max_idle"`
	MaxActive              int  `toml:"MaxActive" yaml:"max_active"`
	IdleTimeout            int  `toml:"IdleTimeout" yaml:"idle_timeout"`
	Wait                   bool `toml:"Wait" yaml:"wait"`
	ReadTimeout            int  `toml:"ReadTimeout" yaml:"read_timeout"`
	WriteTimeout           int  `toml:"WriteTimeout" yaml:"write_timeout"`
	ConnectTimeout         int  `toml:"ConnectTimeout" yaml:"connect_timeout"`
	NormalTasksPollPeriod  int  `toml:"NormalTasksPollPeriod" yaml:"normal_tasks_poll_period"`
	DelayedTasksPollPeriod int  `toml:"DelayedTasksPollPeriod" yaml:"delayed_tasks_poll_period"`
}

// GetTaskRetryTime 任务重试时间
func (cfg MachineConfig) GetTaskRetryTime() time.Duration {
	t, err := time.ParseDuration(cfg.TaskRetryTime)
	if err != nil || t.Seconds() == 0 {
		t = time.Second * 180
	}

	return t
}

// GetDelayTaskTime 延时任务多久之后执行
func (cfg MachineConfig) GetDelayTaskTime() time.Duration {
	t, err := time.ParseDuration(cfg.DelayTaskTime)
	if err != nil || t.Seconds() == 0 {
		t = time.Second * 1
	}

	return t
}
