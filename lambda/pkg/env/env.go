package env

import (
	"os"
)

const (
	// FromEmail is the environment variable that specifies the 'from' email address
	FromEmail string = "FROM_EMAIL"

	// AWSRegion is the environment variable that specifies the AWS region
	AWSRegion string = "AWS_REGION"

	// LogLevel is the environment variable that specifies the log level
	LogLevel string = "LOG_LEVEL"
)

// GetEnv retrieves an environment variable or returns a default value
func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
