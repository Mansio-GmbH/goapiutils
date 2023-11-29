package chrono_test

import (
	"encoding/json"
	"testing"

	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/stretchr/testify/require"
)

func TestMarshalTime(t *testing.T) {
	ti := chrono.MustParse("2023-11-29T14:00:00+01:00")
	byt, err := json.Marshal(ti)
	require.NoError(t, err)
	require.Equal(t, `"2023-11-29T14:00:00+01:00"`, string(byt))
}
