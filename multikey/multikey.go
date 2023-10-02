package multikey

import (
	"strings"
)

const (
	TENANT           = "TENANT"
	TOUR             = "TOUR"
	TRIP             = "TRIP"
	IMPORTEDTOUR     = "IMPORTEDTOUR"
	HANDOVERSTATION  = "HANDOVERSTATION"
	NAME             = "NAME"
	ADDRESS          = "ADDRESS"
	NETWORK          = "NETWORK"
	ID               = "ID"
	DATA             = "DATA"
	EXTERNALNUMBER   = "EXTERNALNUMBER"
	METADATA         = "METADATA"
	USER             = "USER"
	PLANNEDDEPARTURE = "PLANNEDDEPARTURE"
	PLANNEDARRIVAL   = "PLANNEDARRIVAL"
	MATCHSTATE       = "MATCHSTATE"
	EVENT            = "EVENT"
	COMPANY          = "COMPANY"
	TRAILER          = "TRAILER"
	LICENSEPLATE     = "LICENSEPLATE"
	CREATEDAT        = "CREATEDAT"
	VARIANT          = "VARIANT"
	QUERY            = "QUERY"
	SERVICECASE      = "SERVICECASE"
	STATE            = "STATE"
	CREATED_AT       = "CREATED_AT"
	TRUCK            = "TRUCK"
	DEPOT            = "DEPOT"
	DOCUMENTUPLOADER = "DOCUMENTUPLOADER"
	BUCKET           = "BUCKET"
	SECRET           = "SECRET"
	MATCHINGRESULT   = "MATCHINGRESULT"
)

func Key(part0 string, parts ...string) *string {
	s := strings.Join(append([]string{part0}, parts...), "#")
	return &s
}
