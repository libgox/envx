package envx

import (
	"fmt"
	"os"
	"strconv"
)

// GetUint64 returns the uint64 value of the environment variable or an error if not set or invalid.
func GetUint64(key string) (uint64, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return 0, NewMissingEnvVarError(key)
	}
	res, err := strconv.ParseUint(aux, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid uint64 value for %s: %s", key, aux)
	}
	return uint64(res), nil
}

// GetUint64Or returns the uint64 value of the environment variable or the default if not set.
// If the variable is set but cannot be parsed as an uint64, it returns an error.
func GetUint64Or(key string, defaultValue uint64) (uint64, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue, nil
	}
	res, err := strconv.ParseUint(aux, 10, 64)
	if err != nil {
		return defaultValue, fmt.Errorf("invalid uint64 value for %s: %s", key, aux)
	}
	return uint64(res), nil
}

// MustGetUint64 returns the uint64 value of an environment variable, panicking if it's missing.
func MustGetUint64(key string) uint64 {
	aux := os.Getenv(key)
	if aux == "" {
		panic(NewMissingEnvVarError(key))
	}
	res, err := strconv.ParseUint(aux, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid uint64 value for %s: %s", key, aux))
	}
	return uint64(res)
}

// MustGetUint64Or returns the uint64 value of an environment variable or a default if it's missing.
// It panics if the variable is set but has an invalid uint64 format.
func MustGetUint64Or(key string, defaultValue uint64) uint64 {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue
	}
	res, err := strconv.ParseUint(aux, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid uint64 value for %s: %s", key, aux))
	}
	return uint64(res)
}
