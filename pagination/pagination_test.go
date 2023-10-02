package pagination_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/mansio-gmbh/goapiutils/pagination"
	"github.com/stretchr/testify/require"
)

func TestValidation(t *testing.T) {
	page := pagination.Pagination{}

	validate := validator.New()
	err := validate.Struct(page)
	require.NoError(t, err)
}
