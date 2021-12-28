package environ

import (
	"errors"
	"os"
	"strconv"
	"strings"
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
	switch v := v.(type) {
	case string:
		e.defaultString = &v
	case int:
		e.defaultInt = &v
	case float64:
		e.defaultFloat = &v
	case time.Duration:
		e.defaultDuration = &v
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

	const (
		hexPrefix = "0x"
		octPrefix = "0"
	)

	var (
		base   int
		prefix string
	)

	switch {
	default:
		base = 10
	case strings.HasPrefix(strings.ToLower(e.raw), hexPrefix):
		base = 16
		prefix = hexPrefix
	case strings.HasPrefix(strings.ToLower(e.raw), octPrefix):
		base = 8
		prefix = octPrefix
	}

	i, err := strconv.ParseInt(strings.TrimPrefix(e.raw, prefix), base, strconv.IntSize)

	return int(i), err
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
