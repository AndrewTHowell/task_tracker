package in_memory

import (
	"context"

	tasktracker "task_tracker/pkg/task_tracker"
)

func (db *InMemoryDB) InsertTask(_ context.Context, task *tasktracker.Task) error {
	db.tasksByID[task.ID.String()] = task
	return nil
}
