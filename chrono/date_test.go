package chrono_test

import (
	"encoding/json"
	"testing"

	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/stretchr/testify/require"
)

func TestMarshalDate(t *testing.T) {
	ti := chrono.MustParse("2023-11-29")
	byt, err := json.Marshal(ti)
	require.NoError(t, err)
	require.Equal(t, `"2023-11-29T00:00:00+01:00"`, string(byt))
}
