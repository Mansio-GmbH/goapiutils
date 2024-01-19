package pagination

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mansio-gmbh/goapiutils/lastevaluatedkey"
)

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

func (p Pagination) QueryParameter() (map[string]types.AttributeValue, *int32) {
	var limit *int32 = aws.Int32(100)
	lek, _ := lastevaluatedkey.Decode(p.Key)
	if p.Limit != nil {
		limit = aws.Int32(int32(*p.Limit))
	}
	return lek, limit
}

func (w WithPagination) PaginationOrDefault() Pagination {
	if w.Pagination == nil {
		return Pagination{
			Limit: aws.Int64(100),
		}
	}
	return *w.Pagination

}
