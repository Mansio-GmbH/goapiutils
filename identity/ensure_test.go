package identity_test

import (
	"testing"

	"github.com/mansio-gmbh/goapiutils/identity"
	"github.com/stretchr/testify/require"
)

func TestNotBelongToTenant(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"tenant:OTHERTENANT",
		},
	}

	require.Error(t, iden.EnsureBelongToTenant(TestTenant{}))
}

func TestEnsureBelongToNetwork(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"network:TESTNETWORK",
		},
	}

	require.NoError(t, iden.EnsureBelongToNetwork(TestNetwork{}))
}

func TestGlobalAdminstratorsBelongToNetwork(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"role:global-administrator",
		},
	}

	require.NoError(t, iden.EnsureBelongToNetwork(TestNetwork{}))
}

func TestNotEnsureBelongToNetwork(t *testing.T) {
	iden := identity.CognitoIdentity{
		Username: "florian.wehling@mansio-logistics.com",
		Groups: []string{
			"network:OTHERNETWORK",
		},
	}

	require.Error(t, iden.EnsureBelongToNetwork(TestNetwork{}))
}
