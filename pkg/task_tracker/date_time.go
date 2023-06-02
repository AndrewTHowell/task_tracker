package task_tracker

import (
	"context"
	"fmt"
	"time"
)

// DateTime represents a Date. A Time component on the given Date is optional.
type DateTime struct {
	Date         time.Time
	IncludesTime bool
}

func (d DateTime) Validate(_ context.Context) error {
	if d.IncludesTime {
		return nil
	}
	if d.Date.Hour() != 0 ||
		d.Date.Minute() != 0 ||
		d.Date.Second() != 0 ||
		d.Date.Nanosecond() != 0 {
		return fmt.Errorf("validating date time: %q should not have time component set", d.Date.String())
	}
	return nil
}
