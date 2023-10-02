package stringnormalisation_test

import (
	"testing"

	"github.com/mansio-gmbh/goapiutils/stringnormalisation"
	"github.com/stretchr/testify/require"
)

func TestNormalisation(t *testing.T) {
	tests := []struct {
		In  string
		Out string
		err error
	}{
		{In: "Mansio", Out: "MANSIO"},
		{In: "Fast Carrot Global Traiding Company Ltd.", Out: "FASTCARROTGLOBALTRAIDINGCOMPANYLTD"},
		{In: "TDY", Out: "", err: stringnormalisation.ErrInputToShort},
		{In: "#*,.()", Out: "", err: stringnormalisation.ErrOutputToShort},
	}

	for _, test := range tests {
		out, err := stringnormalisation.Normalise(test.In)
		require.ErrorIs(t, err, test.err)
		require.Equal(t, test.Out, out)
	}
}
