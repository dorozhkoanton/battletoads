package persistent

import (
	"context"
	"fmt"

	"github.com/dorozhkoanton/battletoads/internal/entity"
	"github.com/dorozhkoanton/battletoads/pkg/postgres"
)

// const _defaultEntityCap = 64

// ScheduledCommandRepo -.
type ScheduledCommandRepo struct {
	*postgres.Postgres
}

// NewScheduledCommandRepo -.
func NewScheduledCommandRepo(pg *postgres.Postgres) *ScheduledCommandRepo {
	return &ScheduledCommandRepo{pg}
}

// Create -.
func (r *ScheduledCommandRepo) Create(ctx context.Context, sc entity.ScheduledCommand) error {
	sql, args, err := r.Builder.
		Insert("scheduled_command").
		Columns("name", "time").
		Values(sc.Name, sc.Time).
		ToSql()

	if err != nil {
		return fmt.Errorf("ScheduledCommandRepo - Create - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ScheduledCommandRepo - Create - r.Pool.Exec: %w", err)
	}

	return nil
}
