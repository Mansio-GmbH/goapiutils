package tenant

import (
	"strings"

	"github.com/mansio-gmbh/goapiutils/stringnormalisation"
)

type Tenant struct {
	ID string `json:"id" validate:"required"`
}

type WithTenant struct {
	Tenant Tenant `json:"tenant" validate:"required"`
}

func (t Tenant) TenantID() string {
	return strings.ToUpper(stringnormalisation.NormaliseWithoutLengthCheck(t.ID))
}

func Parse(tenant string) Tenant {
	return Tenant{
		ID: tenant,
	}
}
