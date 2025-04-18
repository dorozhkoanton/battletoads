// Package v1 implements routing paths. Each services in own file.
package http

import (
	"net/http"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/dorozhkoanton/battletoads/config"
	_ "github.com/dorozhkoanton/battletoads/docs" // Swagger docs.
	"github.com/dorozhkoanton/battletoads/internal/controller/http/middleware"
	v1 "github.com/dorozhkoanton/battletoads/internal/controller/http/v1"
	"github.com/dorozhkoanton/battletoads/internal/usecase"
	"github.com/dorozhkoanton/battletoads/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// NewRouter -.
// Swagger spec:
// @title       Battletoads API
// @description Automated bot for the game [ToadBot](https://toadbot.info/)
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(app *fiber.App, cfg *config.Config, l logger.Interface, t usecase.Translation, ic usecase.IntervalCommand, sc usecase.ScheduledCommand) {
	// Options
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recovery(l))

	// Prometheus metrics
	if cfg.Metrics.Enabled {
		prometheus := fiberprometheus.New("battletoads")
		prometheus.RegisterAt(app, "/metrics")
		app.Use(prometheus.Middleware)
	}

	// Swagger
	if cfg.Swagger.Enabled {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	// K8s probe
	app.Get("/healthz", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })

	// Routers
	apiV1Group := app.Group("/v1")
	{
		v1.NewTranslationRoutes(apiV1Group, t, l)
		v1.NewIntervalCommandRoutes(apiV1Group, ic, l)
		v1.NewScheduledCommandsRoutes(apiV1Group, sc, l)
	}
}
