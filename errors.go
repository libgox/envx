package envx

import "fmt"

// NoSuchEnvError is returned when a required environment variable is missing.
type NoSuchEnvError struct {
	Key string
}

func (e *NoSuchEnvError) Error() string {
	return fmt.Sprintf("no such environment variable: %s", e.Key)
}

func NewMissingEnvVarError(key string) error {
	return &NoSuchEnvError{Key: key}
}
