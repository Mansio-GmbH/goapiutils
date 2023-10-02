package apigatewayresponse_test

import (
	"testing"

	"github.com/mansio-gmbh/api/lib/apigatewayresponse"
	"github.com/stretchr/testify/require"
)

func TestMake(t *testing.T) {
	type Response struct {
		Foo string `json:"foo"`
	}

	resp, err := apigatewayresponse.Make(Response{
		Foo: "foo",
	})
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, `{"foo":"foo"}`, resp.Body)
	require.Equal(t, "application/json", resp.Headers["content-type"])
}
