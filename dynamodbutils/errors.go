package dynamodbutils

import (
	"errors"

	dyntypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func IsResourceNotFound(err error) bool {
	if err == nil {
		return false
	}
	{
		var rnf *dyntypes.ResourceNotFoundException
		if errors.As(err, &rnf) {
			return true
		}
	}
	return false
}
