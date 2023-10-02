package apigatewayresponse

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Option func(*events.APIGatewayProxyResponse)

func WithStatusCode(statuscode int) func(a *events.APIGatewayProxyResponse) {
	return func(a *events.APIGatewayProxyResponse) {
		a.StatusCode = statuscode
	}
}

func Make[ResponseT any](resp ResponseT, opts ...Option) (*events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	a := &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
		Headers: map[string]string{
			"content-type": "application/json",
		},
	}

	for _, opt := range opts {
		opt(a)
	}

	return a, nil
}

func MakeWithErr[ResponseT any](resp ResponseT, err error) (*events.APIGatewayProxyResponse, error) {
	if err != nil {
		return nil, err
	}
	return Make(resp)
}
