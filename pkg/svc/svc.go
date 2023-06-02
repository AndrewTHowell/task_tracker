// Package svc facilitates serving tasks.
package svc

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
}

type Svc struct {
	clock Clock
	db    DB
}
