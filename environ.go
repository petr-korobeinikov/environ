package environ

import (
	"errors"
	"os"
	"strconv"
	"time"
)

func E(name string) *Environ {
	s, found := os.LookupEnv(name)

	return &Environ{
		raw:   s,
		found: found,
	}
}

func (e *Environ) Default(interface{}) *Environ {
	return e
}

func (e *Environ) AsString() (string, error) {
	if !e.found {
		return "", ErrEnvVarNotSet
	}

	return e.raw, nil
}

func (e *Environ) AsInt() (int, error) {
	if !e.found {
		return 0, ErrEnvVarNotSet
	}

	return strconv.Atoi(e.raw)
}

func (e *Environ) AsFloat() (float64, error) {
	if !e.found {
		return 0, ErrEnvVarNotSet
	}

	return strconv.ParseFloat(e.raw, strconv.IntSize)
}

func (e *Environ) AsDuration() (time.Duration, error) {
	if !e.found {
		return 0, ErrEnvVarNotSet
	}

	return time.ParseDuration(e.raw)
}

type Environ struct {
	raw   string
	found bool
}

var (
	ErrEnvVarNotSet = errors.New("env var not set")
)
