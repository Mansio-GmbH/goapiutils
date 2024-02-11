package ct

import "github.com/mansio-gmbh/goapiutils/chrono"

type TimeWindow struct {
	StartsAt chrono.Time `json:"startsAt" dynamodbav:"startsAt"`
	EndsAt   chrono.Time `json:"endsAt" dynamodbav:"endsAt"`
}
