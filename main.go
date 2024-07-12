package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := scheduler.NewFromConfig(cfg)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	jobName := "MyJob"
	at := "at(2021-12-31T23:59:59)"

	createSchedule(client, jobName, at)

	// listSchedules(client)

	// deleteSchedule(client)

}
func createSchedule(client *scheduler.Client, jobName string, at string) {

	clientToken := "token"
	arn := "arn:aws:lambda:us-west-2:123456789012:function:my-function"
	roleArn := "arn:aws:iam::123456789012:role/MyRole"
	timezone := "JST"
	groupName := jobName

	// ref. https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/scheduler#CreateScheduleInput
	input := &scheduler.CreateScheduleInput{
		ClientToken: &clientToken,

		Name:        &jobName,
		GroupName:   &groupName,
		Description: &jobName,

		ScheduleExpression:         &at,
		ScheduleExpressionTimezone: &timezone,
		FlexibleTimeWindow: &types.FlexibleTimeWindow{
			Mode: types.FlexibleTimeWindowModeOff,
		},

		Target: &types.Target{
			Arn:     &arn,
			RoleArn: &roleArn,
		},

		ActionAfterCompletion: types.ActionAfterCompletionDelete,
	}

	result, err := client.CreateSchedule(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to create schedule, %v", err)
	}

	fmt.Printf("Created schedule: %s\n", *result.ScheduleArn)
}

func listSchedules(client *scheduler.Client) {
	panic("unimplemented")
}

func deleteSchedule(client *scheduler.Client) {
	panic("unimplemented")
}
