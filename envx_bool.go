package envx

import (
	"fmt"
	"os"
	"strings"
)

// GetBool returns the boolean value of the environment variable or an error if not set or invalid.
func GetBool(key string) (bool, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return false, NewMissingEnvVarError(key)
	}
	return parseBool(key, aux)
}

// GetBoolOr returns the boolean value of the environment variable or the default if not set.
// If the variable is set but cannot be parsed as a boolean, it returns an error.
func GetBoolOr(key string, defaultValue bool) (bool, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue, nil
	}
	return parseBool(key, aux)
}

// MustGetBool returns the boolean value of an environment variable, panicking if it's missing.
func MustGetBool(key string) bool {
	aux := os.Getenv(key)
	if aux == "" {
		panic(NewMissingEnvVarError(key))
	}
	res, err := parseBool(key, aux)
	if err != nil {
		panic(err)
	}
	return res
}

// MustGetBoolOr returns the boolean value of an environment variable or a default if it's missing.
// It panics if the variable is set but has an invalid boolean format.
func MustGetBoolOr(key string, defaultValue bool) bool {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue
	}
	res, err := parseBool(key, aux)
	if err != nil {
		panic(err)
	}
	return res
}

func parseBool(key, value string) (bool, error) {
	value = strings.ToLower(value)
	if value == "true" {
		return true, nil
	}
	if value == "false" {
		return false, nil
	}
	return false, fmt.Errorf("invalid boolean value for %s: %s", key, value)
}
