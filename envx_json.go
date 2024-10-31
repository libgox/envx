package envx

import (
	"encoding/json"
	"fmt"
	"os"
)

// GetJson parses the JSON environment variable into the provided struct pointer.
// Returns an error if the variable is missing or contains invalid JSON.
func GetJson(key string, out any) error {
	value := os.Getenv(key)
	if value == "" {
		return NewMissingEnvVarError(key)
	}
	if err := json.Unmarshal([]byte(value), out); err != nil {
		return fmt.Errorf("invalid JSON format for %s: %s", key, value)
	}
	return nil
}
