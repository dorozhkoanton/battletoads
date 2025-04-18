// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import (
	"time"
)

// Command -.
type Command struct {
	Name string `json:"name" example:"Покормить жабу"`
}

// IntervalCommand -.
type IntervalCommand struct {
	Command
	Interval time.Duration `json:"interval" example:"9223372036854775807"`
}

// ScheduledCommand -.
type ScheduledCommand struct {
	Command
	Time time.Time `json:"time" example:"2023-10-01T12:00:00Z"`
}
