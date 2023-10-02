package identity

import (
	"errors"
	"log"
	"strings"
)

type Tenant interface {
	TenantID() string
}

type Network interface {
	NetworkID() string
}

type CognitoIdentity struct {
	Username string         `json:"username"`
	Groups   []string       `json:"groups"`
	Sub      string         `json:"sub"`
	Claims   map[string]any `json:"claims"`
}

type Roles []string

var (
	ErrUnauthorized = errors.New("unauthorized")
)

func FromStringClaims(stringClaims map[string]string) (*CognitoIdentity, error) {
	claims := make(map[string]any)
	for key, value := range stringClaims {
		claims[key] = value
	}
	return FromClaims(claims)
}

func FromClaims(claims map[string]any) (*CognitoIdentity, error) {
	identity := &CognitoIdentity{}
	identity.Sub, _ = claims["sub"].(string)
	identity.Username, _ = claims["username"].(string)
	groups := claims["cognito:groups"].(string)
	identity.Groups = strings.Split(groups[1:len(groups)-1], " ")
	identity.Claims = claims
	return identity, nil
}

func FromJWTAuthorizer(authorizer map[string]any) (*CognitoIdentity, error) {
	jwt := authorizer["jwt"]
	if jwt == nil {
		log.Print("authorizer does not contain 'jwt'")
		return nil, ErrUnauthorized
	}
	jwtAsMap := jwt.(map[string]any)
	claims := jwtAsMap["claims"]
	if claims == nil {
		log.Print("authorizer does not contain 'jwt.claims'")
		return nil, ErrUnauthorized
	}
	return FromClaims(claims.(map[string]any))
}
