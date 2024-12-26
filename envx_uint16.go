package envx

import (
	"fmt"
	"os"
	"strconv"
)

// GetUint16 returns the uint16 value of the environment variable or an error if not set or invalid.
func GetUint16(key string) (uint16, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return 0, NewMissingEnvVarError(key)
	}
	res, err := strconv.ParseUint(aux, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("invalid uint16 value for %s: %s", key, aux)
	}
	return uint16(res), nil
}

// GetUint16Or returns the uint16 value of the environment variable or the default if not set.
// If the variable is set but cannot be parsed as an uint16, it returns an error.
func GetUint16Or(key string, defaultValue uint16) (uint16, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue, nil
	}
	res, err := strconv.ParseUint(aux, 10, 16)
	if err != nil {
		return defaultValue, fmt.Errorf("invalid uint16 value for %s: %s", key, aux)
	}
	return uint16(res), nil
}

// MustGetUint16 returns the uint16 value of an environment variable, panicking if it's missing.
func MustGetUint16(key string) uint16 {
	aux := os.Getenv(key)
	if aux == "" {
		panic(NewMissingEnvVarError(key))
	}
	res, err := strconv.ParseUint(aux, 10, 16)
	if err != nil {
		panic(fmt.Sprintf("invalid uint16 value for %s: %s", key, aux))
	}
	return uint16(res)
}

// MustGetUint16Or returns the uint16 value of an environment variable or a default if it's missing.
// It panics if the variable is set but has an invalid uint16 format.
func MustGetUint16Or(key string, defaultValue uint16) uint16 {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue
	}
	res, err := strconv.ParseUint(aux, 10, 16)
	if err != nil {
		panic(fmt.Sprintf("invalid uint16 value for %s: %s", key, aux))
	}
	return uint16(res)
}
