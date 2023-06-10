package boostrap

import (
	"github.com/sonuos1501/go-app/worker/config"
	"github.com/sonuos1501/go-app/worker/handlers"
	"github.com/sonuos1501/go-app/worker/queues"
	"go.uber.org/fx"
)

func All() []fx.Option {
	return []fx.Option{
		fx.Provide(config.LoadConfig),

		fx.Provide(queues.NewQueueServer),

		fx.Provide(handlers.NewQueueHandler),

		fx.Invoke(queues.RegisterHandlers),
	}
}
