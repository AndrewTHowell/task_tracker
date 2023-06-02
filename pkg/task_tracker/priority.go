package task_tracker

import (
	"context"
	"fmt"
)

// Priority represents how important a task is.
type Priority int

const (
	NoPriority Priority = iota
	LowPriority
	MediumPriority
	HighPriority
)

func (p Priority) Validate(_ context.Context) error {
	if p < NoPriority || p > HighPriority {
		return fmt.Errorf("validating priority: invalid priority %q", p)
	}
	return nil
}
