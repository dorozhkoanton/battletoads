package scheduled_command

import (
	"context"
	"fmt"

	"github.com/dorozhkoanton/battletoads/internal/entity"
	"github.com/dorozhkoanton/battletoads/internal/repo"
)

// UseCase =.
type UseCase struct {
	repo repo.ScheduledCommandRepo
}

// New -.
func New(r repo.ScheduledCommandRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

// Create -.
func (uc *UseCase) Create(ctx context.Context, sc entity.ScheduledCommand) error {
	err := uc.repo.Create(ctx, sc)
	if err != nil {
		return fmt.Errorf("ScheduledCommandUseCase - Create - uc.repo.Store: %w", err)
	}

	return nil
}
