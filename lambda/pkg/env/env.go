package env

import (
	"os"
)

const Email string = "EMAIL"
const LogLevel string = "LOG_LEVEL"

// GetEnv retrieves an environment variable or returns a default value
func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
