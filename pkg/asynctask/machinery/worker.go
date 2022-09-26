package machinery

import (
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
)

func Worker() {
	defer func() {
		if r := recover(); r != nil {
			log.ERROR.Println("recover outer panic:", r)
		}
	}()
	consumerTag := "integration_grant_coupon_worker"

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	worker := MachineServer.NewWorker(consumerTag, 0)

	// Here we inject some custom code for error handling,
	// start and end of task hooks, useful for metrics for example.
	errorhandler := func(err error) {
		log.ERROR.Println("I am an error handler:", err)
	}

	pretaskhandler := func(signature *tasks.Signature) {
		log.INFO.Println("I am a start of task handler for:", signature.Name)
	}

	posttaskhandler := func(signature *tasks.Signature) {
		log.INFO.Println("I am an end of task handler for:", signature.Name)
	}

	worker.SetPostTaskHandler(posttaskhandler)
	worker.SetErrorHandler(errorhandler)
	worker.SetPreTaskHandler(pretaskhandler)

	errorsChan := make(chan error)
	worker.LaunchAsync(errorsChan)

	if <-errorsChan != nil {
		log.INFO.Println("启动worker失败:", <-errorsChan)
		panic("启动异步任务队列的worker失败")
	}
}
