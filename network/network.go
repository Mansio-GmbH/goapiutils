package network

import (
	"regexp"
	"strings"
)

type (
	Network struct {
		ID string `json:"id" validate:"required"`
	}
)

func (n Network) NetworkID() string {
	regex := regexp.MustCompile("[^a-zA-Z0-9-_~]")
	return strings.ToUpper(regex.ReplaceAllString(n.ID, ""))
}

func Parse(network string) Network {
	return Network{
		ID: network,
	}
}
