package logging

import (
	log "github.com/sirupsen/logrus"
	"lambda/pkg/env"
	"os"
	"strings"
)

const (
	tagKey = "tag"
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
func Debug(tag string, args ...interface{}) {
	if tag != "" {
		log.WithField(tagKey, tag).Debug(args...)
		return
	}

	log.Debug(args...)
}

// Debugf is a wrapper around logrus.Debugf
func Debugf(tag string, format string, args ...interface{}) {
	if tag != "" {
		log.WithField(tagKey, tag).Debugf(format, args...)
		return
	}

	log.Debugf(format, args...)
}

// Info is a wrapper around logrus.Info
func Info(tag string, args ...interface{}) {
	if tag != "" {
		log.WithField(tagKey, tag).Info(args...)
		return
	}

	log.Info(args...)
}

// Infof is a wrapper around logrus.Info
func Infof(tag string, format string, args ...interface{}) {
	if tag != "" {
		log.WithField(tagKey, tag).Infof(format, args...)
		return
	}

	log.Infof(format, args...)
}

// Warn is a wrapper around logrus.Warn
func Warn(tag string, args ...interface{}) {
	if tag != "" {
		log.WithField(tagKey, tag).Warn(args...)
		return
	}

	log.Warn(args...)
}

// Warnf is a wrapper around logrus.Warnf
func Warnf(tag string, format string, args ...interface{}) {
	if tag != "" {
		log.WithField(tagKey, tag).Warnf(format, args...)
		return
	}

	log.Warnf(format, args...)
}

// Error is a wrapper around logrus.Error
func Error(tag string, args ...interface{}) {
	if tag != "" {
		log.WithField(tagKey, tag).Error(args...)
		return
	}

	log.Error(args...)
}

// Errorf is a wrapper around logrus.Errorf
func Errorf(tag string, format string, args ...interface{}) {
	if tag != "" {
		log.WithField(tagKey, tag).Errorf(format, args...)
		return
	}

	log.Errorf(format, args...)
}

// Fatal is a wrapper around logrus.Fatal
func Fatal(tag string, args ...interface{}) {
	if tag != "" {
		log.WithField(tagKey, tag).Fatal(args...)
		return
	}

	log.Fatal(args...)
}

// Fatalf is a wrapper around logrus.Fatalf
func Fatalf(tag string, format string, args ...interface{}) {
	if tag != "" {
		log.WithField(tagKey, tag).Fatalf(format, args...)
		return
	}

	log.Fatalf(format, args...)
}
