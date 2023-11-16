package testconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

type MockedEventBridge struct {
	PutEventsFunc func(ctx context.Context, params *eventbridge.PutEventsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error)
}

func (m MockedEventBridge) PutEvents(ctx context.Context, params *eventbridge.PutEventsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error) {
	if m.PutEventsFunc != nil {
		return m.PutEventsFunc(ctx, params, optFns...)
	}
	return &eventbridge.PutEventsOutput{}, nil
}
