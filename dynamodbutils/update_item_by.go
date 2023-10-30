package dynamodbutils

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mansio-gmbh/goapiutils/awsif"
	"github.com/mansio-gmbh/goapiutils/must"
)

type ItemKey struct {
	PK string `dynamodbav:"pk"`
	SK string `dynamodbav:"sk"`
}

func UpdateItemBy(ctx context.Context, tableName string, client awsif.DynamoDBClient, key ItemKey, updated any) error {
	item, err := attributevalue.MarshalMap(updated)
	if err != nil {
		return err
	}

	updateExpression := make([]string, 0)
	expressionAttributeNames := make(map[string]string)
	expressionAttributeValues := make(map[string]types.AttributeValue)
	for fieldName, fieldValue := range item {
		updateExpression = append(updateExpression, fmt.Sprintf("#%s = :%s", fieldName, fieldName))
		expressionAttributeValues[":"+fieldName] = fieldValue
		expressionAttributeNames["#"+fieldName] = fieldName
	}

	if len(updateExpression) == 0 {
		return nil
	}
	updateExpression = append(updateExpression, "#updatedAt = :updatedAt")
	expressionAttributeNames["#updatedAt"] = "updatedAt"
	expressionAttributeValues[":updatedAt"] = must.Must(attributevalue.Marshal(time.Now()))

	if _, err :=
		client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
			TableName:                 &tableName,
			Key:                       must.Must(attributevalue.MarshalMap(key)),
			ConditionExpression:       aws.String("attribute_exists(sk)"),
			ExpressionAttributeValues: expressionAttributeValues,
			ExpressionAttributeNames:  expressionAttributeNames,
			UpdateExpression:          aws.String("SET " + strings.Join(updateExpression, ", ")),
		}); err != nil && strings.Contains(err.Error(), "ConditionalCheckFailed") {
		return errors.New("not found")
	}
	return err
}
