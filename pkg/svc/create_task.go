package svc

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"

	tasktracker "task_tracker/pkg/task_tracker"
)

func (s *Svc) CreateTask(ctx context.Context, task *tasktracker.Task) (*tasktracker.Task, error) {
	if task == nil {
		return nil, nil
	}

	if task.ID == uuid.Nil {
		task.ID = uuid.NewV4()
	}
	if err := task.Validate(ctx, s); err != nil {
		return nil, fmt.Errorf("creating task: %w", err)
	}

	now, err := s.clock.Now(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating task: populating fields with current time: %w", err)
	}
	task.CreateTime, task.UpdateTime = now, now

	if err := s.db.InsertTask(ctx, task); err != nil {
		return nil, fmt.Errorf("creating task: %w", err)
	}
	return task, nil
}
