package golog

import (
	"os"

	"github.com/go-logr/logr"
)

var log = logr.Discard()

func SetLogger(l logr.Logger) {
	_, log = l.WithCallStackHelper()
}

func Info(msg string, keysAndValues ...interface{}) {
	log.Info(msg, keysAndValues...)
}

func Error(err error, msg string, keysAndValues ...interface{}) {
	log.Error(err, msg, keysAndValues...)
}

func Fatal(err error, msg string, keysAndValues ...interface{}) {
	log.Error(err, msg, keysAndValues...)
	os.Exit(19)
}

func V(level int) logr.Logger {
	return log.V(level)
}

func WithValues(keysAndValues ...interface{}) logr.Logger {
	return log.WithValues(keysAndValues...)
}
