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

func (t Tenant) String() string {
	return t.TenantID()
}

func (t Tenant) IsEmpty() bool {
	return t.ID == ""
}

func (t Tenant) IsZero() bool {
	return t.ID == ""
}

func (t Tenant) Equals(other Tenant) bool {
	return t.TenantID() == other.TenantID()
}

func (t Tenant) IsValid() bool {
	return t.ID != ""
}
