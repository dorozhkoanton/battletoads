package amqprpc

import (
	"github.com/dorozhkoanton/battletoads/internal/usecase"
	"github.com/dorozhkoanton/battletoads/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
