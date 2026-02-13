package network

import (
	"strings"

	"github.com/mansio-gmbh/goapiutils/stringnormalisation"
)

type (
	Network struct {
		ID string `json:"id"`
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

func (n Network) String() string {
	return n.NetworkID()
}

func (n Network) IsEmpty() bool {
	return n.ID == ""
}

func (n Network) IsZero() bool {
	return n.ID == ""
}

func (n Network) Equals(other Network) bool {
	return n.NetworkID() == other.NetworkID()
}

func (n Network) IsValid() bool {
	return n.ID != ""
}
