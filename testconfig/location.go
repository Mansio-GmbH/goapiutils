package testconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/location"
)

type LocationClientMock struct {
	CreateTrackerFunc func(ctx context.Context, params *location.CreateTrackerInput, optFns ...func(*location.Options)) (*location.CreateTrackerOutput, error)
}

func (m LocationClientMock) CreateTracker(ctx context.Context, params *location.CreateTrackerInput, optFns ...func(*location.Options)) (*location.CreateTrackerOutput, error) {
	if m.CreateTrackerFunc != nil {
		return m.CreateTrackerFunc(ctx, params, optFns...)
	}
	return &location.CreateTrackerOutput{}, nil
}
