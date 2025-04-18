package v1

import (
	"net/http"
	"time"

	"github.com/dorozhkoanton/battletoads/internal/entity"
	"github.com/dorozhkoanton/battletoads/internal/usecase"
	"github.com/dorozhkoanton/battletoads/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type intervalCommandRoutes struct {
	ic usecase.IntervalCommand
	l  logger.Interface
	v  *validator.Validate
}

func NewIntervalCommandRoutes(apiV1Group fiber.Router, ic usecase.IntervalCommand, l logger.Interface) {
	r := &intervalCommandRoutes{ic, l, validator.New(validator.WithRequiredStructEnabled())}

	intervalCommandGroup := apiV1Group.Group("interval_command")
	{
		intervalCommandGroup.Post("/", r.create)
	}
}

type createIntervalCommandRequest struct {
	Name     string `json:"name" validate:"required" example:"Покормить жабу"`
	Interval string `json:"interval" validate:"required" example:"6h30m50s"`
}

// @Summary     Create
// @Description Create Interval Command
// @ID          create_interval_command
// @Tags  	    Interval Command
// @Accept      json
// @Produce     json
// @Param       request body createIntervalCommandRequest true "Set up interval command"
// @Success     200 {object} string
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /interval_command/ [post]
func (r *intervalCommandRoutes) create(ctx *fiber.Ctx) error {
	var request createIntervalCommandRequest

	if err := ctx.BodyParser(&request); err != nil {
		r.l.Error(err, "http - v1 - interval_command - create")

		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	if err := r.v.Struct(request); err != nil {
		r.l.Error(err, "http - v1 - interval_command - create")

		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	duration, err := time.ParseDuration(request.Interval)
	if err != nil {
		r.l.Error(err, "http - v1 - interval_command - create")

		return errorResponse(ctx, http.StatusBadRequest, "invalid interval value")
	}

	err = r.ic.Create(
		ctx.UserContext(),
		entity.IntervalCommand{
			Command:  entity.Command{Name: request.Name},
			Interval: duration,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - interval_command - create")

		return errorResponse(ctx, http.StatusInternalServerError, "interval_command service problems")
	}

	return ctx.Status(http.StatusOK).JSON("OK")
}
