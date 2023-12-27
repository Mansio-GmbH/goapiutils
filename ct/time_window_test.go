package ct_test

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/mansio-gmbh/goapiutils/ct"
	"github.com/stretchr/testify/require"
)

func TestTimeWindow(t *testing.T) {
	validate := validator.New()

	tw := ct.TimeWindow{
		StartsAt: chrono.NewTime(2023, chrono.December, 27, 12, 0, 0, 0, time.UTC),
		EndsAt:   chrono.NewTime(2023, chrono.December, 27, 14, 0, 0, 0, time.UTC),
	}

	err := validate.Struct(tw)
	require.NoError(t, err)
}
