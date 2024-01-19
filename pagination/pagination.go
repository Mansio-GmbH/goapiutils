package pagination

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mansio-gmbh/goapiutils/lastevaluatedkey"
)

const DEFAULT_LIMIT = 100
const MAX_LIMIT = 500

type (
	Pagination struct {
		Key   *string `json:"key"`
		Limit *int64  `json:"limit" validate:"omitempty,gt=0,lte=500"`
	}

	PaginatedResponse struct {
		Items            []map[string]any `json:"items"`
		LastEvaluatedKey *string          `json:"lastEvaluatedKey"`
	}

	PaginatedResponseTyped[T any] struct {
		Items            []T     `json:"items"`
		LastEvaluatedKey *string `json:"lastEvaluatedKey"`
	}

	WithPagination struct {
		Pagination *Pagination `json:"pagination"`
	}
)

func (p Pagination) LimitCount() *int32 {
	if p.Limit == nil {
		return aws.Int32(100)
	}
	if *p.Limit > MAX_LIMIT {
		return aws.Int32(MAX_LIMIT)
	}
	return aws.Int32(int32(*p.Limit))
}

func (p Pagination) ExclusiveStartKey() map[string]types.AttributeValue {
	lek, _ := lastevaluatedkey.Decode(p.Key)
	return lek
}

func (w WithPagination) PaginationOrDefault() Pagination {
	if w.Pagination == nil {
		return Pagination{
			Limit: aws.Int64(100),
		}
	}
	return *w.Pagination

}
