package pagination

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
