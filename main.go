package main

import (
	"log"

	"github.com/aki-nishikawa/event-bridge-scheduler-go/aws"
	"github.com/aki-nishikawa/event-bridge-scheduler-go/aws/driver"
)

func main() {
	schedulerClient := driver.NewScheduler()
	schedulerRepository := aws.NewSchedulerRepository(schedulerClient)

	// TODO: Create a schedule

	// List all schedules
	schedules, err := schedulerRepository.ListAll()
	if err != nil {
		log.Fatalf("failed to list schedules, %v", err)
	}

	log.Println("Schedules:")
	for _, schedule := range schedules {
		log.Printf(" Name: %s, Arn: %s\n", *schedule.Name, *schedule.Arn)
	}

	// TODO: Get a schedule by ARN

	// TODO: Update a schedule

	// TODO: Delete a schedule
}
