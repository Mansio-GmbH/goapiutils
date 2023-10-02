package identity

import (
	"fmt"
	"strings"
)

func (i CognitoIdentity) EnsureAnyRole(anyRequiredRoles Roles) error {
	accessByRolePermitted := false
	for _, role := range anyRequiredRoles {
		roleGroupIdentifier := fmt.Sprintf("role:%s", role)
		for _, group := range i.Groups {
			if strings.HasPrefix(group, roleGroupIdentifier) {
				accessByRolePermitted = true
				break
			} else if group == "role:global-administrator" {
				accessByRolePermitted = true
				break
			}
		}

		if accessByRolePermitted {
			break
		}
	}
	if !accessByRolePermitted {
		return fmt.Errorf("unauthorized. missing required role")
	}
	return nil
}

func (i CognitoIdentity) EnsureBelongToTenant(selectedTenant Tenant) error {
	tenantGroupIdentifier := fmt.Sprintf("tenant:%s", selectedTenant.TenantID())
	accessToTenatPermitted := false
	for _, group := range i.Groups {
		if group == tenantGroupIdentifier {
			accessToTenatPermitted = true
			break
		} else if group == "role:global-administrator" {
			accessToTenatPermitted = true
			break
		}
	}

	if !accessToTenatPermitted {
		return fmt.Errorf("unauthorized for selected tenant")
	}
	return nil
}

func (i CognitoIdentity) EnsureBelongToNetwork(selectedNetwork Network) error {
	networkGroupIdentifier := fmt.Sprintf("network:%s", selectedNetwork.NetworkID())
	accessToTenatPermitted := false
	for _, group := range i.Groups {
		if group == networkGroupIdentifier {
			accessToTenatPermitted = true
			break
		} else if group == "role:global-administrator" {
			accessToTenatPermitted = true
			break
		}
	}

	if !accessToTenatPermitted {
		return fmt.Errorf("unauthorized for selected network")
	}
	return nil
}
