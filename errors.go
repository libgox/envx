package envx

import "fmt"

// MissingEnvVarError is returned when a required environment variable is missing.
type MissingEnvVarError struct {
	Key string
}

func (e *MissingEnvVarError) Error() string {
	return fmt.Sprintf("missing required environment variable: %s", e.Key)
}

func NewMissingEnvVarError(key string) error {
	return &MissingEnvVarError{Key: key}
}
