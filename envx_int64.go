package envx

import (
	"fmt"
	"os"
	"strconv"
)

// GetInt64 returns the int64 value of the environment variable or an error if not set or invalid.
func GetInt64(key string) (int64, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return 0, NewMissingEnvVarError(key)
	}
	res, err := strconv.ParseInt(aux, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid int64 value for %s: %s", key, aux)
	}
	return res, nil
}

// GetInt64Or returns the int64 value of the environment variable or the default if not set.
// If the variable is set but cannot be parsed as an int64, it returns an error.
func GetInt64Or(key string, defaultValue int64) (int64, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue, nil
	}
	res, err := strconv.ParseInt(aux, 10, 64)
	if err != nil {
		return defaultValue, fmt.Errorf("invalid int64 value for %s: %s", key, aux)
	}
	return res, nil
}

// MustGetInt64 returns the int64 value of an environment variable, panicking if it's missing.
func MustGetInt64(key string) int64 {
	aux := os.Getenv(key)
	if aux == "" {
		panic(NewMissingEnvVarError(key))
	}
	res, err := strconv.ParseInt(aux, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid int64 value for %s: %s", key, aux))
	}
	return res
}

// MustGetInt64Or returns the int64 value of an environment variable or a default if it's missing.
// It panics if the variable is set but has an invalid int64 format.
func MustGetInt64Or(key string, defaultValue int64) int64 {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue
	}
	res, err := strconv.ParseInt(aux, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid int64 value for %s: %s", key, aux))
	}
	return res
}
