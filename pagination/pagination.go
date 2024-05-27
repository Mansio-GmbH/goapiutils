package pagination

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/elliotchance/pie/v2"
	"github.com/mansio-gmbh/goapiutils/lastevaluatedkey"
	"github.com/mansio-gmbh/goapiutils/must"
)

const DEFAULT_LIMIT = 500
const MAX_LIMIT = 1000

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
		return aws.Int32(DEFAULT_LIMIT)
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

func (w WithPagination) PaginationOrDefault(limit ...int) Pagination {
	if w.Pagination == nil {
		l := DEFAULT_LIMIT
		if len(limit) > 0 {
			l = limit[0]
		}
		return Pagination{
			Limit: aws.Int64(int64(l)),
		}
	}
	return *w.Pagination
}

func FilteredPaginatedResponse[T any](items []*T, lastEvaluatedKey map[string]types.AttributeValue, selectedFields []string) (PaginatedResponseTyped[*T], error) {
	objectsToFilter := make([]any, 0, len(items))
	must.WithoutError(json.Unmarshal(must.Must(json.Marshal(items)), &objectsToFilter))
	filterSlice(selectedFields, &objectsToFilter, "items")
	filteredItems := make([]*T, 0, len(objectsToFilter))
	must.WithoutError(json.Unmarshal(must.Must(json.Marshal(objectsToFilter)), &filteredItems))

	return PaginatedResponseTyped[*T]{
		Items:            filteredItems,
		LastEvaluatedKey: must.Must(lastevaluatedkey.Encode(lastEvaluatedKey)),
	}, nil
}

func filterMap(keys []string, data *map[string]any, prefix string) {
	for key, value := range *data {
		if !pie.Contains(keys, prefix+"/"+key) {
			delete(*data, key)
		} else {
			switch v := value.(type) {
			case map[string]any:
				filterMap(keys, &v, prefix+"/"+key)
			case []any:
				filterSlice(keys, &v, prefix+"/"+key)
			}
		}
	}
}

func filterSlice(keys []string, data *[]any, prefix string) {
	for i := 0; i < len(*data); i++ {
		switch v := (*data)[i].(type) {
		case map[string]any:
			filterMap(keys, &v, prefix)
		case []any:
			filterSlice(keys, &v, prefix)
		}
	}
}
