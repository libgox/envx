package envx

import (
	"fmt"
	"os"
	"strconv"
)

// GetUint32 returns the uint32 value of the environment variable or an error if not set or invalid.
func GetUint32(key string) (uint32, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return 0, NewMissingEnvVarError(key)
	}
	res, err := strconv.ParseUint(aux, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid uint32 value for %s: %s", key, aux)
	}
	return uint32(res), nil
}

// GetUint32Or returns the uint32 value of the environment variable or the default if not set.
// If the variable is set but cannot be parsed as an uint32, it returns an error.
func GetUint32Or(key string, defaultValue uint32) (uint32, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue, nil
	}
	res, err := strconv.ParseUint(aux, 10, 32)
	if err != nil {
		return defaultValue, fmt.Errorf("invalid uint32 value for %s: %s", key, aux)
	}
	return uint32(res), nil
}

// MustGetUint32 returns the uint32 value of an environment variable, panicking if it's missing.
func MustGetUint32(key string) uint32 {
	aux := os.Getenv(key)
	if aux == "" {
		panic(NewMissingEnvVarError(key))
	}
	res, err := strconv.ParseUint(aux, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("invalid uint32 value for %s: %s", key, aux))
	}
	return uint32(res)
}

// MustGetUint32Or returns the uint32 value of an environment variable or a default if it's missing.
// It panics if the variable is set but has an invalid uint32 format.
func MustGetUint32Or(key string, defaultValue uint32) uint32 {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue
	}
	res, err := strconv.ParseUint(aux, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("invalid uint32 value for %s: %s", key, aux))
	}
	return uint32(res)
}
