package confield

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/spf13/cast"
)

// F is string to be config
type F string

const (
	sep          = "|"
	envvarPrefix = "$"
)

func (f *F) raw() string {
	return string(*f)
}

func (f *F) val() (string, error) {
	if f.raw() == "" {
		return "", errors.New("value not set")
	}
	pieces := strings.Split(f.raw(), sep)
	for _, p := range pieces {
		if strings.HasPrefix(p, envvarPrefix) {
			if env, ok := os.LookupEnv(strings.TrimLeft(p, envvarPrefix)); ok {
				return env, nil
			}
			continue
		}
		return p, nil
	}
	return "", errors.New("value not set")
}

// IsSet checks field has valid value
func (f *F) IsSet() bool {
	_, err := f.val()
	return err == nil
}

// String returns value as string
func (f *F) String() string {
	s, _ := f.val()
	return s
}

// Bool returns value as bool
func (f *F) Bool() bool {
	s, _ := f.val()
	return cast.ToBool(s)
}

// Int returns value as int
func (f *F) Int() int {
	s, _ := f.val()
	return cast.ToInt(s)
}

// Float64 returns value as float64
func (f *F) Float64() float64 {
	s, _ := f.val()
	return cast.ToFloat64(s)
}

// Time returns value as time.Time
func (f *F) Time() time.Time {
	s, _ := f.val()
	return cast.ToTime(s)
}

// Duration returns value as time.Duration
func (f *F) Duration() time.Duration {
	s, _ := f.val()
	return cast.ToDuration(s)
}

// StringE returns value as string
func (f *F) StringE() (string, error) {
	return f.val()
}

// BoolE returns value as bool
func (f *F) BoolE() (bool, error) {
	s, err := f.val()
	if err != nil {
		return false, err
	}
	return cast.ToBoolE(s)
}

// IntE returns value as int
func (f *F) IntE() (int, error) {
	s, err := f.val()
	if err != nil {
		return 0, err
	}
	return cast.ToIntE(s)
}

// Float64E returns value as float64
func (f *F) Float64E() (float64, error) {
	s, err := f.val()
	if err != nil {
		return 0, err
	}
	return cast.ToFloat64E(s)
}

// TimeE returns value as time.Time
func (f *F) TimeE() (time.Time, error) {
	s, err := f.val()
	if err != nil {
		return time.Time{}, err
	}
	return cast.ToTimeE(s)
}

// DurationE returns value as time.Duration
func (f *F) DurationE() (time.Duration, error) {
	s, err := f.val()
	if err != nil {
		return time.Duration(0), err
	}
	return cast.ToDurationE(s)
}
