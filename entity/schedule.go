package entity

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
)

type Schedule struct {
	Name        string
	ScheduledAt time.Time
	LambdaInput string
}

func (s *Schedule) ToCreateScheduleInput() *scheduler.CreateScheduleInput {
	// 本当はもっと適切な場所で設定する
	lambdaArn := os.Getenv("LAMBDA_ARN")
	scheduleGroupName := os.Getenv("SCHEDULE_GROUP_NAME")
	schedulerRoleArn := os.Getenv("SCHEDULER_ROLE_ARN")
	timezone := "Asia/Tokyo"

	// at(yyyy-mm-ddThh:mm:ss)
	expression := fmt.Sprintf("at(%s)", s.ScheduledAt.Format("2006-01-02T15:04:05"))

	// ref. https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/scheduler#CreateScheduleInput
	return &scheduler.CreateScheduleInput{
		// 必須パラメータ
		Name: aws.String(s.Name),
		Target: &types.Target{
			Arn:     &lambdaArn,
			RoleArn: &schedulerRoleArn,
			Input:   aws.String(s.LambdaInput), // 本当は lambda を呼び出す際のパラメータ
			RetryPolicy: &types.RetryPolicy{
				MaximumRetryAttempts: aws.Int32(0), // リトライしない
			},
		},
		ScheduleExpression: aws.String(expression),

		// 必須ではないが設定するパラメータ
		ActionAfterCompletion: types.ActionAfterCompletionDelete, // 実行後に削除
		GroupName:             aws.String(scheduleGroupName),     // 設定しない場合は default になる
		FlexibleTimeWindow: &types.FlexibleTimeWindow{
			Mode: types.FlexibleTimeWindowModeOff, // 無効化
		},
		ScheduleExpressionTimezone: &timezone, // トラブったら困るので明示しておく
	}
}

func NewScheduleFromGetScheduleOutput(output *scheduler.GetScheduleOutput) (*Schedule, error) {
	expression := *output.ScheduleExpression

	// expression must be at expression - at(yyyy-mm-ddThh:mm:ss)
	if expression[:2] != "at" {
		return nil, fmt.Errorf("invalid expression: %s", expression)
	}

	scheduledAt, err := timeFromExpression(expression)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expression: %w", err)
	}

	return &Schedule{
		Name:        *output.Name,
		ScheduledAt: scheduledAt,
		LambdaInput: *output.Target.Input,
	}, nil
}

func timeFromExpression(expression string) (time.Time, error) {
	// at(yyyy-mm-ddThh:mm:ss)
	layout := "2006-01-02T15:04:05"
	return time.Parse(layout, expression[3:22])
}
