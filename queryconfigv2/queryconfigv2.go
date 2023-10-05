package queryconfigv2

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

type QueryConfig struct {
	TableName                 *string
	IndexName                 *string
	Limit                     *int32
	KeyConditionExpression    *string
	FilterExpression          *string
	ExpressionAttributeValues map[string]types.AttributeValue
	ExclusiveStartKey         map[string]types.AttributeValue
	ReverseOrder              bool
}
