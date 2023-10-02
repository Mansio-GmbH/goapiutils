package testconfig

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type (
	TestServices struct {
		DynamodbClient *dynamodb.Client
		S3Client       *s3.Client
		SQSClient      *sqs.Client
	}
)

var configResolver = aws.EndpointResolverWithOptionsFunc(
	func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		switch service {
		case dynamodb.ServiceID:
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           "http://localhost:8008",
				SigningRegion: region,
			}, nil
		case s3.ServiceID:
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           "http://localhost:9000",
				SigningRegion: region,
			}, nil
		case sqs.ServiceID:
			return aws.Endpoint{
				PartitionID: "aws",
				URL:         "http://localhost:9324",
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	},
)

func AllServices() TestServices {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("localhost"),
		config.WithEndpointResolverWithOptions(configResolver),
	)
	if err != nil {
		log.Fatal(err)
	}

	return TestServices{
		DynamodbClient: dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
			o.Credentials = credentials.NewStaticCredentialsProvider("b59xng", "b2sc60", "")
		}),
		S3Client: s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.Credentials = credentials.NewStaticCredentialsProvider("admin", "mansio_admin", "")
		}),
		SQSClient: sqs.NewFromConfig(cfg),
	}
}

func Configure() (*dynamodb.Client, *s3.Client, error) {
	ts := AllServices()

	return ts.DynamodbClient, ts.S3Client, nil
}
