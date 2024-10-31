package envx

import (
	"fmt"
	"os"
	"time"
)

// GetDuration returns the duration value of the environment variable or an error if not set or invalid.
func GetDuration(key string) (time.Duration, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return 0, NewMissingEnvVarError(key)
	}
	res, err := time.ParseDuration(aux)
	if err != nil {
		return 0, fmt.Errorf("invalid duration value for %s: %s", key, aux)
	}
	return res, nil
}

// GetDurationOr returns the duration value of the environment variable or the default if not set.
// If the variable is set but cannot be parsed as a duration, it returns an error.
func GetDurationOr(key string, defaultValue time.Duration) (time.Duration, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue, nil
	}
	res, err := time.ParseDuration(aux)
	if err != nil {
		return defaultValue, fmt.Errorf("invalid duration value for %s: %s", key, aux)
	}
	return res, nil
}

// MustGetDuration returns the duration value of an environment variable, panicking if it's missing or invalid.
func MustGetDuration(key string) time.Duration {
	aux := os.Getenv(key)
	if aux == "" {
		panic(NewMissingEnvVarError(key))
	}
	res, err := time.ParseDuration(aux)
	if err != nil {
		panic(fmt.Sprintf("invalid duration value for %s: %s", key, aux))
	}
	return res
}

// MustGetDurationOr returns the duration value of an environment variable or a default if it's missing.
// It panics if the variable is set but has an invalid duration format.
func MustGetDurationOr(key string, defaultValue time.Duration) time.Duration {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue
	}
	res, err := time.ParseDuration(aux)
	if err != nil {
		panic(fmt.Sprintf("invalid duration value for %s: %s", key, aux))
	}
	return res
}
