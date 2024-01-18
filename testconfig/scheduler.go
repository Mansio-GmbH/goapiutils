package testconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/scheduler"
)

type SchedulerMock struct {
	CreateScheduleFunc func(ctx context.Context, params *scheduler.CreateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error)
}

func (m SchedulerMock) CreateSchedule(ctx context.Context, params *scheduler.CreateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error) {
	if m.CreateScheduleFunc != nil {
		return m.CreateScheduleFunc(ctx, params, optFns...)
	}
	return &scheduler.CreateScheduleOutput{}, nil
}
