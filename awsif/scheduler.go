package awsif

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/scheduler"
)

type SchedulerClient interface {
	CreateSchedule(ctx context.Context, params *scheduler.CreateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error)
	DeleteSchedule(ctx context.Context, params *scheduler.DeleteScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.DeleteScheduleOutput, error)
}
