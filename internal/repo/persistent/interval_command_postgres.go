package persistent

import (
	"context"
	"fmt"

	"github.com/dorozhkoanton/battletoads/internal/entity"
	"github.com/dorozhkoanton/battletoads/pkg/postgres"
)

// const _defaultEntityCap = 64

// IntervalCommand -.
type IntervalCommandRepo struct {
	*postgres.Postgres
}

// NewIntervalCommandRepo -.
func NewIntervalCommandRepo(pg *postgres.Postgres) *IntervalCommandRepo {
	return &IntervalCommandRepo{pg}
}

// Create -.
func (r *IntervalCommandRepo) Create(ctx context.Context, ic entity.IntervalCommand) error {
	sql, args, err := r.Builder.
		Insert("interval_command").
		Columns("name, interval").
		Values(ic.Name, ic.Interval).
		ToSql()

	if err != nil {
		return fmt.Errorf("IntervalCommandRepo - Create - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("IntervalCommandRepo - Create - r.Pool.Exec: %w", err)
	}

	return nil
}
