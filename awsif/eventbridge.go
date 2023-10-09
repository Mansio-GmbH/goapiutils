package awsif

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

type EventBridgeClient interface {
	PutEvents(ctx context.Context, params *eventbridge.PutEventsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error)
}
