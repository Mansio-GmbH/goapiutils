package ct

import "github.com/mansio-gmbh/goapiutils/chrono"

type LoadingWindow struct {
	StartsAt chrono.Time `json:"startsAt" dynamodbav:"startsAt" validate:"required"`
	EndsAt   chrono.Time `json:"endsAt" dynamodbav:"endsAt" validate:"required,gtcsfield=StartsAt"`
}
