package identity

func (i CognitoIdentity) IsAuthorized(selectedTenant Tenant, anyRequiredRoles Roles) error {
	if err := i.EnsureBelongToTenant(selectedTenant); err != nil {
		return err
	}

	// scope role to tenant
	rolesWithPostfix := make([]string, len(anyRequiredRoles))
	for idx, role := range anyRequiredRoles {
		rolesWithPostfix[idx] = role + ":tenant:" + selectedTenant.TenantID()
	}

	if err := i.EnsureAnyRole(rolesWithPostfix); err != nil {
		return err
	}

	return nil
}

func (i CognitoIdentity) IsAuthorizedInNetwork(selectedNetwork Network, anyRequiredRoles Roles) error {
	if err := i.EnsureBelongToNetwork(selectedNetwork); err != nil {
		return err
	}

	// scope role to network
	rolesWithPostfix := make([]string, len(anyRequiredRoles))
	for idx, role := range anyRequiredRoles {
		rolesWithPostfix[idx] = role + ":network:" + selectedNetwork.NetworkID()
	}

	if err := i.EnsureAnyRole(rolesWithPostfix); err != nil {
		return err
	}

	return nil
}

func (i CognitoIdentity) IsAuthorizedWithTenantInNetwork(selectedTenant Tenant, selectedNetwork Network, anyRequiredRoles Roles) error {
	if err := i.IsAuthorized(selectedTenant, anyRequiredRoles); err != nil {
		return err
	}

	if err := i.IsAuthorizedInNetwork(selectedNetwork, anyRequiredRoles); err != nil {
		return err
	}

	return nil
}
