package in_memory

import (
	"context"
	"fmt"

	tasktracker "task_tracker/pkg/task_tracker"
)

func (db *InMemoryDB) SelectTaskByID(_ context.Context, id string) (*tasktracker.Task, error) {
	task, ok := db.tasksByID[id]
	if !ok {
		return nil, fmt.Errorf("selecting task by ID: task %q does not exist", id)
	}
	return task, nil
}
