package main

import (
	"context"
	"fmt"
	"time"

	inmemorydb "task_tracker/pkg/db/in_memory"
	"task_tracker/pkg/service"
	tasktracker "task_tracker/pkg/task_tracker"
)

type liveClock struct{}

func (l liveClock) Now(context.Context) (time.Time, error) { return time.Now().UTC(), nil }

func main() {
	fmt.Println("Setting up Task Tracker CLI")
	db := inmemorydb.New()
	svc := service.New(liveClock{}, db)
	fmt.Println("Task Tracker CLI ready")

	today := time.Now().UTC().Truncate(24 * time.Hour)
	tomorrow := today.AddDate(0, 0, 1)

	shoppingTripTask := &tasktracker.Task{
		Title:    "Go to the shops",
		Note:     "Buy some essentials",
		Priority: tasktracker.MediumPriority,
		ScheduleDateTime: tasktracker.DateTime{
			Date: today,
		},
		DueDateTime: tasktracker.DateTime{
			Date: tomorrow,
		},
	}
	shoppingTripTask, err := svc.CreateTask(context.Background(), shoppingTripTask)
	if err != nil {
		panic(any(err))
	}
	fmt.Println(fmt.Printf("shoppingTripTask: %+v", shoppingTripTask))

	buyMilkTask := &tasktracker.Task{
		Title:    "Buy milk",
		Note:     "Semi-skimmed",
		Priority: tasktracker.HighPriority,
		ParentID: shoppingTripTask.ID.String(),
		ScheduleDateTime: tasktracker.DateTime{
			Date: today,
		},
	}
	buyMilkTask, err = svc.CreateTask(context.Background(), buyMilkTask)
	if err != nil {
		panic(any(err))
	}
	fmt.Println(fmt.Printf("buyMilkTask: %+v", buyMilkTask))
	buyEggsTask := &tasktracker.Task{
		Title:    "Buy eggs",
		Note:     "Half a dozen",
		Priority: tasktracker.MediumPriority,
		ParentID: shoppingTripTask.ID.String(),
		ScheduleDateTime: tasktracker.DateTime{
			Date: today,
		},
	}
	buyEggsTask, err = svc.CreateTask(context.Background(), buyEggsTask)
	if err != nil {
		panic(any(err))
	}
	fmt.Println(fmt.Printf("buyEggsTask: %+v", buyEggsTask))
	buyDogTask := &tasktracker.Task{
		Title:    "Buy dog",
		Note:     "Dachshund",
		Priority: tasktracker.LowPriority,
		ParentID: shoppingTripTask.ID.String(),
		ScheduleDateTime: tasktracker.DateTime{
			Date: today,
		},
	}
	buyDogTask, err = svc.CreateTask(context.Background(), buyDogTask)
	if err != nil {
		panic(any(err))
	}
	fmt.Println(fmt.Printf("buyDogTask: %+v", buyDogTask))
}
