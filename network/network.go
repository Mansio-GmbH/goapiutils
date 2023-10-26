package network

import (
	"strings"

	"github.com/mansio-gmbh/goapiutils/stringnormalisation"
)

type (
	Network struct {
		ID string `json:"id" validate:"required"`
	}
)

func (n Network) NetworkID() string {
	return strings.ToUpper(stringnormalisation.NormaliseWithoutLengthCheck(n.ID))
}

func Parse(network string) Network {
	return Network{
		ID: network,
	}
}
