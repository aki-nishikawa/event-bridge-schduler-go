package main

import (
	"log"
	"time"

	"github.com/aki-nishikawa/event-bridge-scheduler-go/entity"
	"github.com/aki-nishikawa/event-bridge-scheduler-go/repository"
	"github.com/aki-nishikawa/event-bridge-scheduler-go/repository/driver"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load .env, %v", err)
	}

	schedulerClient := driver.NewScheduler()
	schedulerRepository := repository.NewSchedulerRepository(schedulerClient)

	// Create a schedule
	createSchedule(schedulerRepository)

	// List all schedules
	// listAllSchedule(schedulerRepository)

	// Get a schedule
	// getSchedule(schedulerRepository)

	// TODO: Update a schedule

	// TODO: Delete a schedule
}

func createSchedule(r *repository.SchedulerRepository) {
	// Create a schedule
	schedule := &entity.Schedule{
		Name:        "nishikawa-test-from-golang",
		ScheduledAt: time.Date(2025, 7, 15, 23, 59, 59, 0, time.UTC),
	}

	scheduleArn, err := r.Create(schedule)
	if err != nil {
		log.Fatalf("failed to create schedule, %v", err)
	}

	log.Printf("Created ScheduleArn: %s\n", scheduleArn)
}

func listAllSchedule(r *repository.SchedulerRepository) {
	schedules, err := r.ListAll()
	if err != nil {
		log.Fatalf("failed to list schedules, %v", err)
	}

	log.Println("Schedules:")
	for _, schedule := range schedules {
		log.Printf(" Name: %s, Arn: %s\n", *schedule.Name, *schedule.Arn)
	}
}

func getSchedule(r *repository.SchedulerRepository) {
	scheduleName := "nishikawa-test-from-golang"
	scheduleGroupName := "nishikawa-test-schedules"

	schedule, err := r.Get(scheduleName, scheduleGroupName)
	if err != nil {
		log.Fatalf("failed to get schedule, %v", err)
	}

	log.Printf("Schedule: Name: %s, ScheduledAt: %s\n", schedule.Name, schedule.ScheduledAt)
}
