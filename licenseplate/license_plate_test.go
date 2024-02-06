package licenseplate_test

import (
	"testing"

	"github.com/mansio-gmbh/goapiutils/licenseplate"
)

func TestLicensePlate_MarshalJSON(t *testing.T) {
	lp := licenseplate.LicensePlate{
		LicensePlate: []string{"ABC", "123"},
	}
	_, err := lp.MarshalJSON()
	if err != nil {
		t.Errorf("Failed to marshal LicensePlate: %v", err)
	}
}

func TestLicensePlate_UnmarshalJSON(t *testing.T) {
	lp := &licenseplate.LicensePlate{}
	err := lp.UnmarshalJSON([]byte(`"ABC-123"`))
	if err != nil {
		t.Errorf("Failed to unmarshal LicensePlate: %v", err)
	}
}
