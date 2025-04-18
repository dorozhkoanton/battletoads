// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/dorozhkoanton/battletoads/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test

type (
	// Translation -.
	Translation interface {
		Translate(context.Context, entity.Translation) (entity.Translation, error)
		History(context.Context) ([]entity.Translation, error)
	}
)

type (
	// IntervalCommand -.
	IntervalCommand interface {
		Create(context.Context, entity.IntervalCommand) error
	}

	// ScheduledCommand -.
	ScheduledCommand interface {
		Create(context.Context, entity.ScheduledCommand) error
	}
)
