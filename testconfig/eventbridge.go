package testconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

type MockedEventBridge struct {
	PutEventsFunc func(ctx context.Context, params *eventbridge.PutEventsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error)
}

func (m MockedEventBridge) PutEvents(ctx context.Context, params *eventbridge.PutEventsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error) {
	return nil, nil
}
