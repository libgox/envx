package envx

import "os"

func GetStrOr(key string, defaultValue string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultValue
	}
	return env
}
