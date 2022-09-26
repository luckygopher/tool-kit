package task

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/qingyunjun/tool-kit/pkg/asynctask/machinery"
)

const (
	TaskDemo string = "task_demo" // 发放优惠券任务
)

// RegisterTask 注册异步任务
func RegisterTask() error {
	tasks := map[string]interface{}{
		TaskDemo: func() {},
	}
	return machinery.MachineServer.RegisterTasks(tasks)
}

// RegisterPeriodicTask 注册定时任务
func RegisterPeriodicTask(spec, name string, signature *tasks.Signature) error {
	return machinery.MachineServer.RegisterPeriodicTask(spec, name, signature)
}
