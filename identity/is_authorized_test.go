package identity_test

import (
	"testing"

	"github.com/mansio-gmbh/goapiutils/identity"
	"github.com/stretchr/testify/require"
)

type TestTenant struct{}

func (TestTenant) TenantID() string { return "TESTTENANT" }

type TestNetwork struct{}

func (TestNetwork) NetworkID() string { return "TESTNETWORK" }

func TestIsAuthorized(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"tenant:TESTTENANT",
			"role:transport-planner:tenant:TESTTENANT",
		},
	}

	require.NoError(t, iden.IsAuthorized(TestTenant{}, identity.Roles{"transport-planner"}))
}

func TestGlobalAdministratorIsAuthorized(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"role:global-administrator",
		},
	}

	require.NoError(t, iden.IsAuthorized(TestTenant{}, identity.Roles{"transport-planner"}))
}

func TestIsUnauthorized(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"tenant:TESTTENANT",
		},
	}

	require.Error(t, iden.IsAuthorized(TestTenant{}, identity.Roles{"transport-planner"}))
}

func TestIsAuthorizedInNetwork(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"network:TESTNETWORK",
			"role:transport-planner:network:TESTNETWORK",
		},
	}

	require.NoError(t, iden.IsAuthorizedInNetwork(TestNetwork{}, identity.Roles{"transport-planner"}))
}

func TestIsAuthorizedInNetworkFailedOtherNetwork(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"network:OTHERNETWORK",
			"role:transport-planner:network:TESTNETWORK",
		},
	}

	require.Error(t, iden.IsAuthorizedInNetwork(TestNetwork{}, identity.Roles{"transport-planner"}))
}

func TestIsAuthorizedInNetworkFailedOtherRole(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"network:TESTNETWORK",
			"role:accountant:tenant:TESTNETWORK",
		},
	}

	require.Error(t, iden.IsAuthorizedInNetwork(TestNetwork{}, identity.Roles{"transport-planner"}))
}

func TestIsAuthorizedWithTenantInNetwork(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"tenant:TESTTENANT",
			"network:TESTNETWORK",
			"role:transport-planner:tenant:TESTTENANT",
			"role:transport-planner:network:TESTNETWORK",
		},
	}

	require.NoError(t, iden.IsAuthorizedWithTenantInNetwork(TestTenant{}, TestNetwork{}, identity.Roles{"transport-planner"}))
}

func TestIsAuthorizedWithTenantInNetworkFailsNotPartOfTenant(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"tenant:OTHERTENANT",
			"network:TESTNETWORK",
			"role:transport-planner:tenant:TESTTENANT",
			"role:transport-planner:network:TESTNETWORK",
		},
	}

	require.Error(t, iden.IsAuthorizedWithTenantInNetwork(TestTenant{}, TestNetwork{}, identity.Roles{"transport-planner"}))
}

func TestIsAuthorizedWithTenantInNetworkFailsNotPartOfNetwork(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"tenant:TESTTENANT",
			"network:OTHERNETWORK",
			"role:transport-planner:tenant:TESTTENANT",
			"role:transport-planner:network:TESTNETWORK",
		},
	}

	require.Error(t, iden.IsAuthorizedWithTenantInNetwork(TestTenant{}, TestNetwork{}, identity.Roles{"transport-planner"}))
}

func TestIsAuthorizedWithTenantInNetworkFailsNoRoleAssigned(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"tenant:TESTTENANT",
			"network:TESTNETWORK",
		},
	}

	require.Error(t, iden.IsAuthorizedWithTenantInNetwork(TestTenant{}, TestNetwork{}, identity.Roles{"transport-planner"}))
}
