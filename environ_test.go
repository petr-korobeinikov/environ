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

	t.Run(`negative`, func(t *testing.T) {
		const e = "ENV_STRING_NOT_SET"

		_, err := E(e).AsString()

		assert.ErrorIs(t, err, ErrEnvVarNotSet)
	})

	t.Run(`default`, func(t *testing.T) {
		t.Run(`for unset var`, func(t *testing.T) {
			const e = "ENV_STRING_NOT_SET_USE_DEFAULT"

			actual, err := E(e).Default("Hello, environ!").AsString()

			assert.NoError(t, err)
			assert.Equal(t, "Hello, environ!", actual)
		})

		t.Run(`for set var ignore default`, func(t *testing.T) {
			const e = "ENV_STRING_SET_DO_NOT_USE_DEFAULT"

			os.Setenv(e, "Hello, environ!")
			defer os.Unsetenv(e)

			actual, err := E(e).Default("Hello, world!").AsString()

			assert.NoError(t, err)
			assert.Equal(t, "Hello, environ!", actual)
		})
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

	t.Run(`negative`, func(t *testing.T) {
		const e = "ENV_INT_NOT_SET"

		_, err := E(e).AsInt()

		assert.ErrorIs(t, err, ErrEnvVarNotSet)
	})

	t.Run(`octal`, func(t *testing.T) {
		const e = "ENV_INT_OCTAL"

		os.Setenv(e, "0755")
		defer os.Unsetenv(e)

		actual, err := E(e).AsInt()

		assert.NoError(t, err)
		assert.Equal(t, 0755, actual)
	})

	t.Run(`hexadecimal`, func(t *testing.T) {
		const e = "ENV_INT_HEXADECIMAL"

		os.Setenv(e, "0xDEADBEEF")
		defer os.Unsetenv(e)

		actual, err := E(e).AsInt()

		assert.NoError(t, err)
		assert.Equal(t, 0xDEADBEEF, actual)
	})

	t.Run(`default`, func(t *testing.T) {
		t.Run(`for unset var`, func(t *testing.T) {
			const e = "ENV_INT_NOT_SET_USE_DEFAULT"

			actual, err := E(e).Default(42).AsInt()

			assert.NoError(t, err)
			assert.Equal(t, 42, actual)
		})

		t.Run(`for set var ignore default`, func(t *testing.T) {
			const e = "ENV_INT_SET_DO_NOT_USE_DEFAULT"

			os.Setenv(e, "1337")
			defer os.Unsetenv(e)

			actual, err := E(e).Default("42").AsInt()

			assert.NoError(t, err)
			assert.Equal(t, 1337, actual)
		})
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

	t.Run(`negative`, func(t *testing.T) {
		const e = "ENV_FLOAT_NOT_SET"

		_, err := E(e).AsFloat()

		assert.ErrorIs(t, err, ErrEnvVarNotSet)
	})

	t.Run(`default`, func(t *testing.T) {
		t.Run(`for unset var`, func(t *testing.T) {
			const e = "ENV_FLOAT_NOT_SET_USE_DEFAULT"

			actual, err := E(e).Default(4.2).AsFloat()

			assert.NoError(t, err)
			assert.Equal(t, 4.2, actual)
		})

		t.Run(`for set var ignore default`, func(t *testing.T) {
			const e = "ENV_FLOAT_SET_DO_NOT_USE_DEFAULT"

			os.Setenv(e, "13.37")
			defer os.Unsetenv(e)

			actual, err := E(e).Default(4.2).AsFloat()

			assert.NoError(t, err)
			assert.Equal(t, 13.37, actual)
		})
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

	t.Run(`negative`, func(t *testing.T) {
		const e = "ENV_DURATION_NOT_SET"

		_, err := E(e).AsDuration()

		assert.ErrorIs(t, err, ErrEnvVarNotSet)
	})

	t.Run(`default`, func(t *testing.T) {
		t.Run(`for unset var`, func(t *testing.T) {
			const e = "ENV_DURATION_NOT_SET_USE_DEFAULT"

			actual, err := E(e).Default(42 * time.Second).AsDuration()

			assert.NoError(t, err)
			assert.Equal(t, 42*time.Second, actual)
		})

		t.Run(`for set var ignore default`, func(t *testing.T) {
			const e = "ENV_DURATION_SET_DO_NOT_USE_DEFAULT"

			os.Setenv(e, "42s")
			defer os.Unsetenv(e)

			actual, err := E(e).Default(1337 * time.Second).AsDuration()

			assert.NoError(t, err)
			assert.Equal(t, 42*time.Second, actual)
		})
	})
}
