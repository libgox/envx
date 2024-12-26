package envx

import (
	"fmt"
	"os"
	"strconv"
)

// GetInt16 returns the int16 value of the environment variable or an error if not set or invalid.
func GetInt16(key string) (int16, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return 0, NewMissingEnvVarError(key)
	}
	res, err := strconv.ParseInt(aux, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("invalid int16 value for %s: %s", key, aux)
	}
	return int16(res), nil
}

// GetInt16Or returns the int16 value of the environment variable or the default if not set.
// If the variable is set but cannot be parsed as an int16, it returns an error.
func GetInt16Or(key string, defaultValue int16) (int16, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue, nil
	}
	res, err := strconv.ParseInt(aux, 10, 16)
	if err != nil {
		return defaultValue, fmt.Errorf("invalid int16 value for %s: %s", key, aux)
	}
	return int16(res), nil
}

// MustGetInt16 returns the int16 value of an environment variable, panicking if it's missing.
func MustGetInt16(key string) int16 {
	aux := os.Getenv(key)
	if aux == "" {
		panic(NewMissingEnvVarError(key))
	}
	res, err := strconv.ParseInt(aux, 10, 16)
	if err != nil {
		panic(fmt.Sprintf("invalid int16 value for %s: %s", key, aux))
	}
	return int16(res)
}

// MustGetInt16Or returns the int16 value of an environment variable or a default if it's missing.
// It panics if the variable is set but has an invalid int16 format.
func MustGetInt16Or(key string, defaultValue int16) int16 {
	aux := os.Getenv(key)
	if aux == "" {
		return defaultValue
	}
	res, err := strconv.ParseInt(aux, 10, 16)
	if err != nil {
		panic(fmt.Sprintf("invalid int16 value for %s: %s", key, aux))
	}
	return int16(res)
}
