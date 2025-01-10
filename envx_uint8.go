package envx

import (
	"fmt"
	"os"
	"strconv"
)

// GetUint8 returns the uint8 value of the environment variable or an error if not set or invalid.
func GetUint8(key string) (uint8, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return 0, NewMissingEnvVarError(key)
	}
	res, err := strconv.ParseUint(aux, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("invalid uint8 value for %s: %s", key, aux)
	}
	return uint8(res), nil
}

// GetUint8Or returns the uint8 value of the environment variable or the default if not set.
// If the variable is set but cannot be parsed as an uint8, it returns an error.
func GetUint8Or(key string, defaultValue uint8) (uint8, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue, nil
	}
	res, err := strconv.ParseUint(aux, 10, 8)
	if err != nil {
		return defaultValue, fmt.Errorf("invalid uint8 value for %s: %s", key, aux)
	}
	return uint8(res), nil
}

// MustGetUint8 returns the uint8 value of an environment variable, panicking if it's missing.
func MustGetUint8(key string) uint8 {
	aux := os.Getenv(key)
	if aux == "" {
		panic(NewMissingEnvVarError(key))
	}
	res, err := strconv.ParseUint(aux, 10, 8)
	if err != nil {
		panic(fmt.Sprintf("invalid uint8 value for %s: %s", key, aux))
	}
	return uint8(res)
}

// MustGetUint8Or returns the uint8 value of an environment variable or a default if it's missing.
// It panics if the variable is set but has an invalid uint8 format.
func MustGetUint8Or(key string, defaultValue uint8) uint8 {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue
	}
	res, err := strconv.ParseUint(aux, 10, 8)
	if err != nil {
		panic(fmt.Sprintf("invalid uint8 value for %s: %s", key, aux))
	}
	return uint8(res)
}
