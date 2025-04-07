package types

import (
	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/mansio-gmbh/goapiutils/ct"
)

type HandoverStationDoc struct {
	BaseDoc
	Name                     string          `json:"name"`
	TenantID                 string          `json:"tenantID"`
	Location                 *ct.Location    `json:"location"`
	ExpectedHandoverDuration chrono.Duration `json:"expectedHandoverDuration"`
	MaxTravelDistance        int             `json:"maxTravelDistance"`
	MaxTravelDuration        int             `json:"maxTravelDuration"`
}
