package logging

import (
	log "github.com/sirupsen/logrus"
	"lambda/pkg/env"
	"os"
	"strings"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(envLogLevelToLevel(env.GetEnv(env.LogLevel, "debug")))
}

// envLogLevelToLevel converts the log level from the environment to a logrus log level
func envLogLevelToLevel(level string) log.Level {

	switch strings.ToLower(level) {
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

// Debug is a wrapper around logrus.Debug
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Debugf is a wrapper around logrus.Debugf
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Info is a wrapper around logrus.Info
func Info(args ...interface{}) {
	log.Info(args...)
}

// Infof is a wrapper around logrus.Info
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warn is a wrapper around logrus.Warn
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Warnf is a wrapper around logrus.Warnf
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Error is a wrapper around logrus.Error
func Error(args ...interface{}) {
	log.Error(args...)
}

// Errorf is a wrapper around logrus.Errorf
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatal is a wrapper around logrus.Fatal
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf is a wrapper around logrus.Fatalf
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
