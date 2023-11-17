package awsif

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/location"
)

type LocationClient interface {
	CreateTracker(ctx context.Context, params *location.CreateTrackerInput, optFns ...func(*location.Options)) (*location.CreateTrackerOutput, error)
}
