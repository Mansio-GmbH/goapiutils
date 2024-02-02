package queryconfigv2

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type QueryConfig struct {
	TableName                 *string
	IndexName                 *string
	Limit                     *int32
	KeyConditionExpression    *string
	FilterExpression          *string
	ExpressionAttributeValues map[string]types.AttributeValue
	ExpressionAttributeNames  map[string]string
	ExclusiveStartKey         map[string]types.AttributeValue
	ReverseOrder              bool
	ConsistentRead            bool
	ProjectionExpression      *string
}

func (qc QueryConfig) ToQueryInput() *dynamodb.QueryInput {
	return &dynamodb.QueryInput{
		TableName:                 qc.TableName,
		IndexName:                 qc.IndexName,
		KeyConditionExpression:    qc.KeyConditionExpression,
		ExpressionAttributeValues: qc.ExpressionAttributeValues,
		ExpressionAttributeNames:  qc.ExpressionAttributeNames,
		ExclusiveStartKey:         qc.ExclusiveStartKey,
		FilterExpression:          qc.FilterExpression,
		Limit:                     qc.Limit,
		ScanIndexForward:          aws.Bool(!qc.ReverseOrder),
		ConsistentRead:            aws.Bool(qc.ConsistentRead),
		ProjectionExpression:      qc.ProjectionExpression,
	}
}
