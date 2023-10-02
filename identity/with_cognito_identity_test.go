package identity_test

import (
	"testing"

	"github.com/mansio-gmbh/goapiutils/identity"
	"github.com/stretchr/testify/require"
)

type Iden struct {
	identity.WithCognitoIdentity
}

func TestWithCognitoIdentity(t *testing.T) {
	iden := Iden{
		WithCognitoIdentity: identity.WithCognitoIdentity{
			Identity: &identity.CognitoIdentity{
				Username: "florian.wehling@mansio-logistics.com",
				Groups: []string{
					"tenant:TESTTENANT",
					"network:TESTNETWORK",
					"role:transport-planner:tenant:TESTTENANT",
					"role:transport-planner:network:TESTNETWORK",
				},
			},
		},
	}

	require.NoError(t, iden.EnsureAnyRole(identity.Roles{"transport-planner"}))
	require.NoError(t, iden.IsAuthorized(TestTenant{}, identity.Roles{"transport-planner"}))
	require.NoError(t, iden.IsAuthorizedInNetwork(TestNetwork{}, identity.Roles{"transport-planner"}))
	require.NoError(t, iden.IsAuthorizedWithTenantInNetwork(TestTenant{}, TestNetwork{}, identity.Roles{"transport-planner"}))
}

func TestWithCognitoIdentityIsNil(t *testing.T) {
	iden := Iden{WithCognitoIdentity: identity.WithCognitoIdentity{}}

	require.Error(t, iden.EnsureAnyRole(identity.Roles{"transport-planner"}))
	require.Error(t, iden.IsAuthorized(TestTenant{}, identity.Roles{"transport-planner"}))
	require.Error(t, iden.IsAuthorizedInNetwork(TestNetwork{}, identity.Roles{"transport-planner"}))
	require.Error(t, iden.IsAuthorizedWithTenantInNetwork(TestTenant{}, TestNetwork{}, identity.Roles{"transport-planner"}))
}
