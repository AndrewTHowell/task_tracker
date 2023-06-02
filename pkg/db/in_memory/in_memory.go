// Package in_memory facilitates storing tasks in memory.
package in_memory

import tasktracker "task_tracker/pkg/task_tracker"

type InMemoryDB struct {
	tasksByID map[string]*tasktracker.Task
}

func New() *InMemoryDB {
	return &InMemoryDB{
		tasksByID: make(map[string]*tasktracker.Task),
	}
}
