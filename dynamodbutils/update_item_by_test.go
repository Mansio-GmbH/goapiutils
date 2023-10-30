package dynamodbutils_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mansio-gmbh/goapiutils/dynamodbutils"
	"github.com/mansio-gmbh/goapiutils/testconfig"
	"github.com/stretchr/testify/require"
)

type TestUpdate struct {
	ValuePresent    *string `dynamodbav:"valuePresent,omitempty"`
	ValueNotPresent *string `dynamodbav:"valueNotPresent,omitempty"`
}

func TestUpdateItemBy(t *testing.T) {
	updateCalled := false
	dbclient := testconfig.DynamodbMock{
		UpdateItemFunc: func(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
			updateCalled = true
			require.Equal(t, "testpk", params.Key["pk"].(*types.AttributeValueMemberS).Value)
			require.Equal(t, "testsk", params.Key["sk"].(*types.AttributeValueMemberS).Value)
			require.Contains(t, *params.UpdateExpression, "valuePresent")
			require.NotContains(t, *params.UpdateExpression, "valueNoPresent")
			return &dynamodb.UpdateItemOutput{}, nil
		},
	}

	err := dynamodbutils.UpdateItemBy(context.Background(), "test-table", dbclient, dynamodbutils.ItemKey{
		PK: "testpk",
		SK: "testsk",
	}, &TestUpdate{
		ValuePresent: aws.String("present"),
	})
	require.NoError(t, err)
	require.True(t, updateCalled)
}
