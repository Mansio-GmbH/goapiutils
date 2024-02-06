package multikey

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/mansio-gmbh/goapiutils/ct"
	"github.com/mansio-gmbh/goapiutils/must"
	"github.com/mansio-gmbh/goapiutils/network"
	"github.com/mansio-gmbh/goapiutils/tenant"
	"github.com/oklog/ulid/v2"
)

const (
	TENANT    = "TENANT"
	NETWORK   = "NETWORK"
	ID        = "ID"
	DATA      = "DATA"
	METADATA  = "METADATA"
	CREATEDAT = "CREATEDAT"
	STATE     = "STATE"
)

type ToString interface {
	ToString() string
}

func Key(part0 any, parts ...any) *string {
	p0 := anyToString(part0)
	ps := make([]string, len(parts))
	for i := 0; i < len(parts); i++ {
		ps[i] = anyToString(parts[i])
	}
	s := strings.Join(append([]string{p0}, ps...), "#")
	return &s
}

func KeyAV(part0 any, parts ...any) types.AttributeValue {
	k := Key(part0, parts...)
	return must.Must(attributevalue.Marshal(k))
}

func anyToString(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case *string:
		return *val
	case *chrono.Time:
		return val.UTC().Format()
	case chrono.Time:
		return val.UTC().Format()
	case *chrono.Date:
		return val.Format()
	case chrono.Date:
		return val.Format()
	case *time.Time:
		return val.UTC().Format(time.RFC3339)
	case time.Time:
		return val.UTC().Format(time.RFC3339)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	case tenant.Tenant:
		return *Key(TENANT, val.TenantID())
	case *tenant.Tenant:
		return *Key(TENANT, val.TenantID())
	case network.Network:
		return *Key(NETWORK, val.NetworkID())
	case *network.Network:
		return *Key(NETWORK, val.NetworkID())
	case ulid.ULID:
		return val.String()
	case *ulid.ULID:
		return val.String()
	case ToString:
		return val.ToString()
	case ct.LicensePlate:
		return val.String()
	case *ct.LicensePlate:
		return val.String()
	}
	log.Fatal("val type not handled")
	return ""
}

func DateOnly(t *time.Time) string {
	return t.Format("2006-01-02")
}

func DateOnlyVal(t time.Time) string {
	return DateOnly(&t)
}
