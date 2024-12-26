package envx

import (
	"fmt"
	"os"
	"strconv"
)

// GetInt returns the integer value of the environment variable or an error if not set or invalid.
func GetInt(key string) (int, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return 0, NewMissingEnvVarError(key)
	}
	res, err := strconv.Atoi(aux)
	if err != nil {
		return 0, fmt.Errorf("invalid integer value for %s: %s", key, aux)
	}
	return res, nil
}

// GetIntOr returns the integer value of the environment variable or the default if not set.
// If the variable is set but cannot be parsed as an integer, it returns an error.
func GetIntOr(key string, defaultValue int) (int, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue, nil
	}
	res, err := strconv.Atoi(aux)
	if err != nil {
		return defaultValue, fmt.Errorf("invalid integer value for %s: %s", key, aux)
	}
	return res, nil
}

// MustGetInt returns the int value of an environment variable, panicking if it's missing.
func MustGetInt(key string) int {
	aux := os.Getenv(key)
	if aux == "" {
		panic(NewMissingEnvVarError(key))
	}
	res, err := strconv.Atoi(aux)
	if err != nil {
		panic(fmt.Sprintf("invalid int value for %s: %s", key, aux))
	}
	return res
}

// MustGetIntOr returns the integer value of an environment variable or a default if it's missing.
// It panics if the variable is set but has an invalid integer format.
func MustGetIntOr(key string, defaultValue int) int {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue
	}
	res, err := strconv.Atoi(aux)
	if err != nil {
		panic(fmt.Sprintf("invalid integer value for %s: %s", key, aux))
	}
	return res
}
