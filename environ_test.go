package environ_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/pkorobeinikov/environ"
)

func TestEnviron_AsString(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		const e = "ENV_STRING"

		os.Setenv(e, "foo_value")
		defer os.Unsetenv(e)

		actual, err := E(e).AsString()

		assert.NoError(t, err)
		assert.Equal(t, "foo_value", actual)
	})
}

func TestEnviron_AsInt(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		const e = "ENV_INT"

		os.Setenv(e, "42")
		defer os.Unsetenv(e)

		actual, err := E(e).AsInt()

		assert.NoError(t, err)
		assert.Equal(t, 42, actual)
	})
}

func TestEnviron_AsFloat(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		const e = "ENV_FLOAT"

		os.Setenv(e, "4.2")
		defer os.Unsetenv(e)

		actual, err := E(e).AsFloat()

		assert.NoError(t, err)
		assert.Equal(t, 4.2, actual)
	})
}

func TestEnviron_AsDuration(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		const e = "ENV_DURATION"

		os.Setenv(e, "42s")
		defer os.Unsetenv(e)

		actual, err := E(e).AsDuration()

		assert.NoError(t, err)
		assert.Equal(t, 42*time.Second, actual)
	})
}