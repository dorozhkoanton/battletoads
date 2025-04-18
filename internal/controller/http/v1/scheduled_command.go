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

type scheduledCommandRoutes struct {
	sc usecase.ScheduledCommand
	l  logger.Interface
	v  *validator.Validate
}

func NewScheduledCommandsRoutes(apiV1Group fiber.Router, sc usecase.ScheduledCommand, l logger.Interface) {
	r := &scheduledCommandRoutes{sc, l, validator.New(validator.WithRequiredStructEnabled())}

	scheduledCommandGroup := apiV1Group.Group("/scheduled_command")
	{
		scheduledCommandGroup.Post("/", r.create)
	}
}

type createScheduledCommandRequest struct {
	Name string    `json:"name" validate:"required" example:"Покормить жабу"`
	Time time.Time `json:"time" validate:"required" example:"2023-10-01T12:00:00Z"`
}

// @Summary     Create
// @Description Create Scheduled Command
// @ID          create_scheduled_command
// @Tags  	    Scheduled Command
// @Accept      json
// @Produce     json
// @Param       request body createScheduledCommandRequest true "Set up scheduled command"
// @Success     200 {object} string
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /scheduled_command/ [post]
func (r *scheduledCommandRoutes) create(ctx *fiber.Ctx) error {
	var request createScheduledCommandRequest

	if err := ctx.BodyParser(&request); err != nil {
		r.l.Error(err, "http - v1 - schedule_command - create")

		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	if err := r.v.Struct(request); err != nil {
		r.l.Error(err, "http - v1 - schedule_command - create")

		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	err := r.sc.Create(
		ctx.UserContext(),
		entity.ScheduledCommand{
			Command: entity.Command{Name: request.Name},
			Time:    request.Time,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - schedule_command - create")

		return errorResponse(ctx, http.StatusInternalServerError, "scheduled_command service problems")
	}

	return ctx.Status(http.StatusOK).JSON("OK")
}
