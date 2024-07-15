package repository

import (
	"context"
	"fmt"

	"github.com/aki-nishikawa/event-bridge-scheduler-go/entity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
)

type SchedulerRepository struct {
	client *scheduler.Client
}

func NewSchedulerRepository(client *scheduler.Client) *SchedulerRepository {
	return &SchedulerRepository{client: client}
}

func (r *SchedulerRepository) Create(schedule *entity.Schedule) (string, error) {
	input := schedule.ToCreateScheduleInput()

	output, err := r.client.CreateSchedule(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to create schedule, %w", err)
	}

	return *output.ScheduleArn, nil
}

func (r *SchedulerRepository) ListAll() ([]*types.ScheduleSummary, error) {
	input := &scheduler.ListSchedulesInput{
		MaxResults: aws.Int32(1),
	}

	schedules := make([]*types.ScheduleSummary, 0)
	for {
		output, err := r.client.ListSchedules(context.TODO(), input)
		if err != nil {
			return nil, fmt.Errorf("failed to list schedules, %w", err)
		}

		for _, schedule := range output.Schedules {
			schedules = append(schedules, &schedule)
		}

		if output.NextToken == nil {
			break
		}
	}

	return schedules, nil
}

func (r *SchedulerRepository) Get(name, groupName string) (*entity.Schedule, error) {
	input := &scheduler.GetScheduleInput{
		Name:      aws.String(name),
		GroupName: aws.String(groupName),
	}

	output, err := r.client.GetSchedule(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to get schedule, %w", err)
	}

	schedule, err := entity.NewScheduleFromGetScheduleOutput(output)
	if err != nil {
		return nil, fmt.Errorf("failed to create schedule entity, %w", err)
	}

	return schedule, nil
}

// TODO: Update

// TODO: Delete
