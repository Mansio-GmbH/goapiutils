package identity

import "errors"

type WithCognitoIdentity struct {
	Identity *CognitoIdentity `json:"identity" validate:"required"`
}

func (w WithCognitoIdentity) EnsureAnyRole(anyRequiredRoles Roles) error {
	if w.Identity == nil {
		return errors.New("unauthorized. identity context incomplete")
	}
	return w.Identity.EnsureAnyRole(anyRequiredRoles)
}

func (w WithCognitoIdentity) IsAuthorized(selectedTenant Tenant, anyRequiredRoles Roles) error {
	if w.Identity == nil {
		return errors.New("unauthorized. identity context incomplete")
	}
	return w.Identity.IsAuthorized(selectedTenant, anyRequiredRoles)
}

func (w WithCognitoIdentity) IsAuthorizedInNetwork(selectedNetwork Network, anyRequiredRoles Roles) error {
	if w.Identity == nil {
		return errors.New("unauthorized. identity context incomplete")
	}
	return w.Identity.IsAuthorizedInNetwork(selectedNetwork, anyRequiredRoles)
}

func (w WithCognitoIdentity) IsAuthorizedWithTenantInNetwork(selectedTenant Tenant, selectedNetwork Network, anyRequiredRoles Roles) error {
	if w.Identity == nil {
		return errors.New("unauthorized. identity context incomplete")
	}
	return w.Identity.IsAuthorizedWithTenantInNetwork(selectedTenant, selectedNetwork, anyRequiredRoles)
}
