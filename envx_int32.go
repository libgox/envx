package envx

import (
	"fmt"
	"os"
	"strconv"
)

// GetInt32 returns the int32 value of the environment variable or an error if not set or invalid.
func GetInt32(key string) (int32, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return 0, NewMissingEnvVarError(key)
	}
	res, err := strconv.ParseInt(aux, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid int32 value for %s: %s", key, aux)
	}
	return int32(res), nil
}

// GetInt32Or returns the int32 value of the environment variable or the default if not set.
// If the variable is set but cannot be parsed as an int32, it returns an error.
func GetInt32Or(key string, defaultValue int32) (int32, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue, nil
	}
	res, err := strconv.ParseInt(aux, 10, 32)
	if err != nil {
		return defaultValue, fmt.Errorf("invalid int32 value for %s: %s", key, aux)
	}
	return int32(res), nil
}

// MustGetInt32 returns the int32 value of an environment variable, panicking if it's missing.
func MustGetInt32(key string) int32 {
	aux := os.Getenv(key)
	if aux == "" {
		panic(NewMissingEnvVarError(key))
	}
	res, err := strconv.ParseInt(aux, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("invalid int32 value for %s: %s", key, aux))
	}
	return int32(res)
}

// MustGetInt32Or returns the int32 value of an environment variable or a default if it's missing.
// It panics if the variable is set but has an invalid int32 format.
func MustGetInt32Or(key string, defaultValue int32) int32 {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue
	}
	res, err := strconv.ParseInt(aux, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("invalid int32 value for %s: %s", key, aux))
	}
	return int32(res)
}
