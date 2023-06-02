// Package service facilitates serving tasks.
package service

import (
	"context"
	"time"

	tasktracker "task_tracker/pkg/task_tracker"
)

type Clock interface {
	Now(context.Context) (time.Time, error)
}

type DB interface {
	InsertTask(ctx context.Context, task *tasktracker.Task) error
	SelectTaskByID(ctx context.Context, id string) (*tasktracker.Task, error)
}

type Service struct {
	clock Clock
	db    DB
}

func New(clock Clock, db DB) *Service {
	return &Service{
		clock: clock,
		db:    db,
	}
}
