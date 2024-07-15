package entity

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/scheduler"
)

type Schedule struct {
	Name        *string
	ScheduledAt time.Time
	// 本当は lambda を呼び出す際のパラメータ
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
		Name:        output.Name,
		ScheduledAt: scheduledAt,
	}, nil
}

func timeFromExpression(expression string) (time.Time, error) {
	// at(yyyy-mm-ddThh:mm:ss)
	layout := "2006-01-02T15:04:05"
	return time.Parse(layout, expression[3:22])
}
