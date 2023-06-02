package task_tracker

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/satori/go.uuid"
)

// Task is something that needs to be tracked.
type Task struct {
	ID                            uuid.UUID
	Title, Note, ParentID         string
	Priority                      Priority
	ScheduleDateTime, DueDateTime DateTime
	Done                          bool
	CreateTime, UpdateTime        time.Time
}

type TaskExistenceChecker interface {
	TaskExists(ctx context.Context, id string) (bool, error)
}

func (t *Task) Validate(ctx context.Context, existenceChecker TaskExistenceChecker) error {
	var validationErr error

	// Static validation.
	if t.ID == uuid.Nil {
		validationErr = errors.Join(validationErr, fmt.Errorf("validating task: ID required"))
	}
	if len(t.Title) == 0 {
		validationErr = errors.Join(validationErr, fmt.Errorf("validating task: Title required"))
	}
	for _, validatable := range []interface{ Validate(context.Context) error }{
		t.Priority, t.ScheduleDateTime, t.DueDateTime,
	} {
		if err := validatable.Validate(ctx); err != nil {
			validationErr = errors.Join(validationErr, fmt.Errorf("validating task: %w", err))
		}
	}

	// Stateful validation.
	if len(t.ParentID) == 0 {
		parentExists, err := existenceChecker.TaskExists(ctx, t.ParentID)
		if err != nil {
			return errors.Join(validationErr, fmt.Errorf("validating task: validating parent: %w", err))
		}
		if !parentExists {
			validationErr = errors.Join(
				validationErr,
				fmt.Errorf("validating task: validating parent: %q does not exist", t.ParentID),
			)
		}
	}
	return validationErr
}
