package logging

import (
	log "github.com/sirupsen/logrus"
	"lambda/pkg/env"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(envLogLevelToLevel(env.GetEnv(env.LogLevel, "error")))
}

// envLogLevelToLevel converts the log level from the environment to a logrus log level
func envLogLevelToLevel(string level) log.Level {
	switch level {
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn":
		return log.WarnLevel
	default:
		return log.ErrorLevel
	}
}
