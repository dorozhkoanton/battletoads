// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dorozhkoanton/battletoads/config"
	amqprpc "github.com/dorozhkoanton/battletoads/internal/controller/amqp_rpc"
	v1 "github.com/dorozhkoanton/battletoads/internal/controller/http"
	"github.com/dorozhkoanton/battletoads/internal/repo/persistent"
	"github.com/dorozhkoanton/battletoads/internal/repo/webapi"
	"github.com/dorozhkoanton/battletoads/internal/usecase/interval_command"
	"github.com/dorozhkoanton/battletoads/internal/usecase/scheduled_command"
	"github.com/dorozhkoanton/battletoads/internal/usecase/translation"
	"github.com/dorozhkoanton/battletoads/pkg/httpserver"
	"github.com/dorozhkoanton/battletoads/pkg/logger"
	"github.com/dorozhkoanton/battletoads/pkg/postgres"
	"github.com/dorozhkoanton/battletoads/pkg/rabbitmq/rmq_rpc/server"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	translationUseCase := translation.New(
		persistent.NewTranslationRepo(pg),
		webapi.New(),
	)
	intervalCommandUseCase := interval_command.New(
		persistent.NewIntervalCommandRepo(pg),
	)
	scheduledCommandUseCase := scheduled_command.New(
		persistent.NewScheduledCommandRepo(pg),
	)

	// RabbitMQ RPC Server
	rmqRouter := amqprpc.NewRouter(translationUseCase)

	rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	v1.NewRouter(httpServer.App, cfg, l, translationUseCase, intervalCommandUseCase, scheduledCommandUseCase)

	// Start servers
	rmqServer.Start()
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-rmqServer.Notify():
		l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = rmqServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}
}
