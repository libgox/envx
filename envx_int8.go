package envx

import (
	"fmt"
	"os"
	"strconv"
)

// GetInt8 returns the int8 value of the environment variable or an error if not set or invalid.
func GetInt8(key string) (int8, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return 0, NewMissingEnvVarError(key)
	}
	res, err := strconv.ParseInt(aux, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("invalid int8 value for %s: %s", key, aux)
	}
	return int8(res), nil
}

// GetInt8Or returns the int8 value of the environment variable or the default if not set.
// If the variable is set but cannot be parsed as an int8, it returns an error.
func GetInt8Or(key string, defaultValue int8) (int8, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue, nil
	}
	res, err := strconv.ParseInt(aux, 10, 8)
	if err != nil {
		return defaultValue, fmt.Errorf("invalid int8 value for %s: %s", key, aux)
	}
	return int8(res), nil
}

// MustGetInt8 returns the int8 value of an environment variable, panicking if it's missing.
func MustGetInt8(key string) int8 {
	aux := os.Getenv(key)
	if aux == "" {
		panic(NewMissingEnvVarError(key))
	}
	res, err := strconv.ParseInt(aux, 10, 8)
	if err != nil {
		panic(fmt.Sprintf("invalid int8 value for %s: %s", key, aux))
	}
	return int8(res)
}

// MustGetInt8Or returns the int8 value of an environment variable or a default if it's missing.
// It panics if the variable is set but has an invalid int8 format.
func MustGetInt8Or(key string, defaultValue int8) int8 {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue
	}
	res, err := strconv.ParseInt(aux, 10, 8)
	if err != nil {
		panic(fmt.Sprintf("invalid int8 value for %s: %s", key, aux))
	}
	return int8(res)
}
