package svc

import (
	"context"
	"fmt"

	tasktracker "task_tracker/pkg/task_tracker"
)

func (s *Svc) GetTask(ctx context.Context, id string) (*tasktracker.Task, error) {
	if len(id) == 0 {
		return nil, nil
	}

	task, err := s.db.SelectTaskByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting task: %w", err)
	}
	return task, nil
}

func (s *Svc) TaskExists(ctx context.Context, id string) (bool, error) {
	_, err := s.GetTask(ctx, id)
	if err != nil {
		return false, fmt.Errorf("checking task exists: %w", err)
	}
	return true, nil
}
