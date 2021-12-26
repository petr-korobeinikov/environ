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

func (e *Environ) Default(v interface{}) *Environ {
	switch v.(type) {
	case string:
		s := v.(string)
		e.defaultString = &s
	case int:
		i := v.(int)
		e.defaultInt = &i
	case float64:
		f := v.(float64)
		e.defaultFloat = &f
	case time.Duration:
		d := v.(time.Duration)
		e.defaultDuration = &d
	}

	return e
}

func (e *Environ) AsString() (string, error) {
	if !e.found && e.defaultString != nil {
		return *e.defaultString, nil
	}

	if !e.found {
		return "", ErrEnvVarNotSet
	}

	return e.raw, nil
}

func (e *Environ) AsInt() (int, error) {
	if !e.found && e.defaultInt != nil {
		return *e.defaultInt, nil
	}

	if !e.found {
		return 0, ErrEnvVarNotSet
	}

	return strconv.Atoi(e.raw)
}

func (e *Environ) AsFloat() (float64, error) {
	if !e.found && e.defaultFloat != nil {
		return *e.defaultFloat, nil
	}

	if !e.found {
		return 0, ErrEnvVarNotSet
	}

	return strconv.ParseFloat(e.raw, strconv.IntSize)
}

func (e *Environ) AsDuration() (time.Duration, error) {
	if !e.found && e.defaultDuration != nil {
		return *e.defaultDuration, nil
	}

	if !e.found {
		return 0, ErrEnvVarNotSet
	}

	return time.ParseDuration(e.raw)
}

type Environ struct {
	raw   string
	found bool

	defaultString   *string
	defaultInt      *int
	defaultFloat    *float64
	defaultDuration *time.Duration
}

var (
	ErrEnvVarNotSet = errors.New("env var not set")
)
