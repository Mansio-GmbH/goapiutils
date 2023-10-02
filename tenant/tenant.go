package tenant

import (
	"regexp"
	"strings"
)

type Tenant struct {
	ID string `json:"id" validate:"required"`
}

type WithTenant struct {
	Tenant Tenant `json:"tenant" validate:"required"`
}

func (t Tenant) TenantID() string {
	regex := regexp.MustCompile("[^a-zA-Z0-9-_~]")
	return strings.ToUpper(regex.ReplaceAllString(t.ID, ""))
}

func Parse(tenant string) Tenant {
	return Tenant{
		ID: tenant,
	}
}
