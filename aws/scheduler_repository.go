package aws

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

func createSchedule(client *scheduler.Client) {

	// clientToken := "token"
	// arn := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
	// roleArn := "arn:aws:iam::123456789012:role/MyRole"
	// timezone := "JST"

	// // ref. https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/scheduler#CreateScheduleInput
	// input := &scheduler.CreateScheduleInput{
	// 	ClientToken: &clientToken,

	// 	Name:        aws.String("MySchedule"),
	// 	GroupName:   aws.String("MyGroup"),
	// 	Description: aws.String("MyDescription"),

	// 	ScheduleExpression:         aws.String("at(2021-12-31T23:59:59)"),
	// 	ScheduleExpressionTimezone: &timezone,
	// 	FlexibleTimeWindow: &types.FlexibleTimeWindow{
	// 		Mode: types.FlexibleTimeWindowModeOff,
	// 	},

	// 	Target: &types.Target{
	// 		Arn:     &arn,
	// 		RoleArn: &roleArn,
	// 	},

	// 	ActionAfterCompletion: types.ActionAfterCompletionDelete,
	// }

	// result, err := client.CreateSchedule(context.TODO(), input)
	// if err != nil {
	// 	log.Fatalf("failed to create schedule, %v", err)
	// }

	// fmt.Printf("Created schedule: %s\n", *result.ScheduleArn)
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
