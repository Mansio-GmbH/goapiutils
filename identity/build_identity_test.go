package identity_test

import (
	"testing"

	"github.com/mansio-gmbh/goapiutils/identity"
	"github.com/stretchr/testify/require"
)

func TestBuildFromClaims(t *testing.T) {
	iden, err := identity.FromClaims(map[string]any{
		"sub":            "b15169e8-12b6-44e0-afac-9516ec1b0a92",
		"username":       "florian.wehling@mansio-logistics.com",
		"cognito:groups": "[tenant:TESTTEANT role:transport-planner]",
	})
	require.NoError(t, err)
	require.Equal(t, "b15169e8-12b6-44e0-afac-9516ec1b0a92", iden.Sub)
	require.Equal(t, "florian.wehling@mansio-logistics.com", iden.Username)
}

func TestBuildFromStringClaims(t *testing.T) {
	iden, err := identity.FromStringClaims(map[string]string{
		"sub":            "b15169e8-12b6-44e0-afac-9516ec1b0a92",
		"username":       "florian.wehling@mansio-logistics.com",
		"cognito:groups": "[tenant:TESTTEANT role:transport-planner]",
	})
	require.NoError(t, err)
	require.Equal(t, "b15169e8-12b6-44e0-afac-9516ec1b0a92", iden.Sub)
	require.Equal(t, "florian.wehling@mansio-logistics.com", iden.Username)
}

func TestBuildFromJWTAuthorizer(t *testing.T) {
	iden, err := identity.FromJWTAuthorizer(map[string]any{
		"jwt": map[string]any{
			"claims": map[string]any{
				"sub":            "b15169e8-12b6-44e0-afac-9516ec1b0a92",
				"username":       "florian.wehling@mansio-logistics.com",
				"cognito:groups": "[tenant:TESTTEANT role:transport-planner]",
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, "b15169e8-12b6-44e0-afac-9516ec1b0a92", iden.Sub)
	require.Equal(t, "florian.wehling@mansio-logistics.com", iden.Username)
}

func TestBuildFromNonJWTAuthorizer(t *testing.T) {
	iden, err := identity.FromJWTAuthorizer(map[string]any{})
	require.Error(t, err)
	require.Nil(t, iden)
}

func TestBuildFromJWTAuthorizerWithoutClaims(t *testing.T) {
	iden, err := identity.FromJWTAuthorizer(map[string]any{
		"jwt": map[string]any{},
	})
	require.Error(t, err)
	require.Nil(t, iden)
}
