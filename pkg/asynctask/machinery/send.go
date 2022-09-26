//go:generate mockgen --source send.go --destination mock/send.mock.go
package machinery

import (
	"context"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/backends/result"
	"github.com/RichardKnop/machinery/v1/tasks"
)

type Send interface {
	SendTask(ctx context.Context, signature *tasks.Signature) (*result.AsyncResult, error)
}

type SendImpl struct {
	server *machinery.Server
}

func (m SendImpl) SendTask(ctx context.Context, signature *tasks.Signature) (*result.AsyncResult, error) {
	return m.server.SendTaskWithContext(ctx, signature)
}

func NewSendServer(server *machinery.Server) Send {
	return &SendImpl{
		server: server,
	}
}
