package dynamobatchwrite

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mansio-gmbh/goapiutils/awsif"
)

func Write(ctx context.Context, client awsif.DynamoDBClient, requests map[string][]types.WriteRequest) error {
	writeBatchesSplitted := make([]map[string][]types.WriteRequest, 0)

	const chunkSize = 25
	for table, items := range requests {
		for currentItemCount := 0; currentItemCount < len(items); currentItemCount += chunkSize {
			end := currentItemCount + chunkSize
			if end > len(items) {
				end = len(items)
			}
			writeBatchesSplitted = append(writeBatchesSplitted, map[string][]types.WriteRequest{
				table: items[currentItemCount:end],
			})
		}
	}

	for _, writeBatch := range writeBatchesSplitted {
		unprocessedRequests := writeBatch

		for len(unprocessedRequests) > 0 {
			output, err := client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
				RequestItems: unprocessedRequests,
			})
			if err != nil {
				return err
			}
			unprocessedRequests = output.UnprocessedItems
		}
	}
	return nil
}
