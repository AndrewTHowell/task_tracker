package task_tracker

import (
	"context"
	"fmt"
	"time"
)

// DateTime represents a Date. A Time component on the given Date is optional.
type DateTime struct {
	dateTime     time.Time
	includesTime bool
}

func (d DateTime) Validate(_ context.Context) error {
	if d.includesTime {
		return nil
	}
	if d.dateTime.Hour() != 0 ||
		d.dateTime.Minute() != 0 ||
		d.dateTime.Second() != 0 ||
		d.dateTime.Nanosecond() != 0 {
		return fmt.Errorf("validating date time: %q should not have time component set", d.dateTime.String())
	}
	return nil
}
