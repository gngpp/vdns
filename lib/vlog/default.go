package vlog

import (
	"log"
	"os"
)

var logger = NewFlag(os.Stdout, log.LstdFlags)

// Debug debug prints debug level msg.
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Debugf prints debug level msg with format.
func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

// Info prints info level msg.
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Infof prints info level msg with format.
func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Warn prints warning level msg.
func Warn(v ...interface{}) {
	logger.Warn(v...)
}

// Warnf prints warning level msg with format.
func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

// Error err prints err level msg.
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Errorf prints err level msg with format.
func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

// Fatal prints fatal level msg and exit process with code 1.
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// Fatalf prints fatal level msg with format and exit process with code 1.
func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}
