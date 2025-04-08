package types

import "github.com/mansio-gmbh/goapiutils/chrono"

// Matching Response represents a response from a matching schedule.
// Please ensure that the ID field is set to the ID of the matching schedule.
// The ApiResponse field should contain the actual response from the matching schedule. It will be returned as it is via the API.
type MatchingResponse struct {
	BaseDoc
	ApiResponse any         `json:"apiResponse"`
	CreatedAt   chrono.Time `json:"createdAt"`
}
