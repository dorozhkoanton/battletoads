package interval_command

import (
	"context"
	"fmt"

	"github.com/dorozhkoanton/battletoads/internal/entity"
	"github.com/dorozhkoanton/battletoads/internal/repo"
)

// UseCase - .
type UseCase struct {
	repo repo.IntervalCommandRepo
}

// New -.
func New(r repo.IntervalCommandRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

// Create -.
func (uc *UseCase) Create(ctx context.Context, ic entity.IntervalCommand) error {
	err := uc.repo.Create(ctx, ic)
	if err != nil {
		return fmt.Errorf("IntervalCommandUseCase - Create - uc.repo.Store: %w", err)
	}

	return nil
}
