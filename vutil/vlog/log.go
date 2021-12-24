package vlog

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logging level.
const (
	Off = iota
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
)

//goland:noinspection ALL
const (
	INFO  = "info"
	DEBUG = "debug"
	OFF   = "off"
	TRACE = "trace"
	WARN  = "warn"
	ERROR = "error"
	FATAL = "fatal"
)

type LevelEnum struct {
	INFO  LevelType
	DEBUG LevelType
	OFF   LevelType
	TRACE LevelType
	WARN  LevelType
	ERROR LevelType
	FATAL LevelType
}

type VLog struct{}

type LevelType string

// Log utilities.
var Log = VLog{}

var Level = LevelEnum{
	INFO:  INFO,
	DEBUG: DEBUG,
	OFF:   OFF,
	TRACE: TRACE,
	WARN:  WARN,
	ERROR: ERROR,
	FATAL: FATAL,
}

// all loggers.
var loggers []*Logger

// the global default logging level, it will be used for creating logger.
var logLevel = Debug

// Logger represents a simple logger with level.
// The underlying logger is the standard Go logging "log".
type Logger struct {
	level  int
	logger *log.Logger
}

// NewLogger creates a logger.
func (*VLog) NewLogger(out io.Writer) *Logger {
	ret := &Logger{level: logLevel, logger: log.New(out, "", log.Ldate|log.Ltime|log.Lshortfile)}

	loggers = append(loggers, ret)

	return ret
}

// Default create default logger
func (v *VLog) Default() *Logger {
	return v.NewLogger(os.Stdout)
}

// SetLevel sets the logging level of all loggers.
func (*VLog) SetLevel(level LevelType) {
	logLevel = getLevel(level)

	for _, l := range loggers {
		l.SetLevel(level)
	}
}

// getLevel gets logging level int value corresponding to the specified level.
func getLevel(level LevelType) int {

	switch level {
	case OFF:
		return Off
	case TRACE:
		return Trace
	case DEBUG:
		return Debug
	case INFO:
		return Info
	case WARN:
		return Warn
	case ERROR:
		return Error
	case FATAL:
		return Fatal
	default:
		return Info
	}
}

// SetLevel sets the logging level of a logger.
func (l *Logger) SetLevel(level LevelType) {
	l.level = getLevel(level)
}

// IsTraceEnabled determines whether the trace level is enabled.
func (l *Logger) IsTraceEnabled() bool {
	return l.level <= Trace
}

// IsDebugEnabled determines whether the debug level is enabled.
func (l *Logger) IsDebugEnabled() bool {
	return l.level <= Debug
}

// IsWarnEnabled determines whether the debug level is enabled.
func (l *Logger) IsWarnEnabled() bool {
	return l.level <= Warn
}

// Trace prints trace level message.
func (l *Logger) Trace(v ...interface{}) {
	if Trace < l.level {
		return
	}

	l.logger.SetPrefix("TRACE ")
	l.logger.Output(2, fmt.Sprintln(v...))
}

// Tracef prints trace level message with format.
func (l *Logger) Tracef(format string, v ...interface{}) {
	if Trace < l.level {
		return
	}
	l.logger.SetPrefix("TRACE ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Debug prints debug level message.
func (l *Logger) Debug(v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.logger.SetPrefix("DEBUG ")
	l.logger.Output(2, fmt.Sprintln(v...))
}

// Debugf prints debug level message with format.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.logger.SetPrefix("DEBUG ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Info prints info level message.
func (l *Logger) Info(v ...interface{}) {
	if Info < l.level {
		return
	}

	l.logger.SetPrefix("INFO ")
	l.logger.Output(2, fmt.Sprintln(v...))
}

// Infof prints info level message with format.
func (l *Logger) Infof(format string, v ...interface{}) {
	if Info < l.level {
		return
	}

	l.logger.SetPrefix("INFO ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Warn prints warning level message.
func (l *Logger) Warn(v ...interface{}) {
	if Warn < l.level {
		return
	}

	l.logger.SetPrefix("WARN ")
	l.logger.Output(2, fmt.Sprintln(v...))
}

// Warnf prints warning level message with format.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if Warn < l.level {
		return
	}

	l.logger.SetPrefix("WARN ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Error prints error level message.
func (l *Logger) Error(v ...interface{}) {
	if Error < l.level {
		return
	}

	l.logger.SetPrefix("ERROR ")
	l.logger.Output(2, fmt.Sprintln(v...))
}

// Errorf prints error level message with format.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if Error < l.level {
		return
	}

	l.logger.SetPrefix("ERROR ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Fatal prints fatal level message and exit process with code 1.
func (l *Logger) Fatal(v ...interface{}) {
	if Fatal < l.level {
		return
	}

	l.logger.SetPrefix("FATAL ")
	l.logger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Fatalf prints fatal level message with format and exit process with code 1.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if Fatal < l.level {
		return
	}

	l.logger.SetPrefix("FATAL ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
