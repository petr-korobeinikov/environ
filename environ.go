package environ

import (
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
	return e.raw, nil
}

func (e *Environ) AsInt() (int, error) {
	return strconv.Atoi(e.raw)
}

func (e *Environ) AsFloat() (float64, error) {
	return strconv.ParseFloat(e.raw, strconv.IntSize)
}

func (e *Environ) AsDuration() (time.Duration, error) {
	return time.ParseDuration(e.raw)
}

type Environ struct {
	raw   string
	found bool
}
