package envx

import "os"

func GetStrOr(key string, defaultValue string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultValue
	}
	return env
}

// MustGet returns the string value of an environment variable, panicking if it's missing.
func MustGet(key string) string {
	aux := os.Getenv(key)
	if aux == "" {
		panic(NewMissingEnvVarError(key))
	}
	return aux
}
