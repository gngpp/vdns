package vlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const (
	off = iota
	trace
	debug
	info
	warn
	err
	fatal
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

type LevelType string

var Level = LevelEnum{
	INFO:  "info",
	DEBUG: "debug",
	OFF:   "off",
	TRACE: "trace",
	WARN:  "warn",
	ERROR: "error",
	FATAL: "fatal",
}

type Color string

//goland:noinspection ALL
const (
	DEBUG_COLOR Color = "\033[0;36m"
	INFO_COLOR  Color = "\033[0;32m"
	WARN_COLOR  Color = "\033[0;33m"
	ERROR_COLOR Color = "\033[0;31m"
	TRACE_COLOR Color = "\033[0;34m"
	FATAL_COLOR Color = "\033[0;35m"
)

var syncOutput = false

var mu sync.Mutex

var writer *io.Writer

// New creates a logger.
func New(out io.Writer) *Logger {
	return NewFlag(out, log.LstdFlags|log.Lshortfile)
}

func NewFlag(out io.Writer, flag int) *Logger {
	ret := &Logger{level: logLevel, logger: log.New(out, "", flag)}

	loggers = append(loggers, ret)

	return ret
}

// Default create default logger
func Default() *Logger {
	if writer != nil {
		return New(*writer)
	}
	return New(os.Stdout)
}

// SetLevel sets the logging level of all loggers. (global)
func SetLevel(level LevelType) {
	mu.Lock()

	defer mu.Unlock()
	logLevel = getLevel(level)

	for _, l := range loggers {
		l.SetLevel(level)
	}
}

// SetOutput sets the output destination for the standard logger. (global)
func SetOutput(w io.Writer) {
	mu.Lock()
	defer mu.Unlock()
	writer = &w
	if syncOutput {
		log.SetOutput(w)
	}
	for _, logger := range loggers {
		logger.SetOutput(w)
	}
}

// SetSyncOutput synchronize output settings to the default log library
func SetSyncOutput(b bool) {
	mu.Lock()
	defer mu.Unlock()
	syncOutput = b
}

// all loggers.
var loggers []*Logger

// the global default logging level, it will be used for creating logger.
var logLevel = info

// Logger represents a simple logger with level.
// The underlying logger is the standard Go logging "log".
type Logger struct {
	level  int
	logger *log.Logger
}

// getLevel gets logging level int value corresponding to the specified level.
func getLevel(level LevelType) int {

	switch level {
	case Level.OFF:
		return off
	case Level.TRACE:
		return trace
	case Level.DEBUG:
		return debug
	case Level.INFO:
		return info
	case Level.WARN:
		return warn
	case Level.ERROR:
		return err
	case Level.FATAL:
		return fatal
	default:
		return info
	}
}

// SetOutput sets the output destination for the standard logger.
func (l *Logger) SetOutput(w io.Writer) {
	l.logger.SetOutput(w)
}

// SetLevel sets the logging level of a logger.
func (l *Logger) SetLevel(level LevelType) {
	l.level = getLevel(level)
}

// IsTraceEnabled determines whether the trace level is enabled.
func (l *Logger) IsTraceEnabled() bool {
	return l.level <= trace
}

// IsDebugEnabled determines whether the debug level is enabled.
func (l *Logger) IsDebugEnabled() bool {
	return l.level <= debug
}

// IsWarnEnabled determines whether the debug level is enabled.
func (l *Logger) IsWarnEnabled() bool {
	return l.level <= warn
}

// Trace prints trace level msg.
func (l *Logger) Trace(v ...interface{}) {
	if trace < l.level {
		return
	}

	l.logger.SetPrefix(string(TRACE_COLOR) + "[TRACE] ")
	err := l.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Tracef prints trace level msg with format.
func (l *Logger) Tracef(format string, v ...interface{}) {
	if trace < l.level {
		return
	}
	l.logger.SetPrefix(string(TRACE_COLOR) + "[TRACE] ")
	err := l.logger.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Debug debug prints debug level msg.
func (l *Logger) Debug(v ...interface{}) {
	if debug < l.level {
		return
	}

	l.logger.SetPrefix(string(DEBUG_COLOR) + "[DEBUG] ")
	err := l.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Debugf prints debug level msg with format.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if debug < l.level {
		return
	}

	l.logger.SetPrefix(string(DEBUG_COLOR) + "[DEBUG] ")
	err := l.logger.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Info prints info level msg.
func (l *Logger) Info(v ...interface{}) {
	if info < l.level {
		return
	}

	l.logger.SetPrefix(string(INFO_COLOR) + "[INFO] ")
	err := l.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Infof prints info level msg with format.
func (l *Logger) Infof(format string, v ...interface{}) {
	if info < l.level {
		return
	}

	l.logger.SetPrefix(string(INFO_COLOR) + "[INFO] ")
	err := l.logger.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Warn prints warning level msg.
func (l *Logger) Warn(v ...interface{}) {
	if warn < l.level {
		return
	}

	l.logger.SetPrefix(string(WARN_COLOR) + "[WARN] ")
	err := l.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Warnf prints warning level msg with format.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if warn < l.level {
		return
	}

	l.logger.SetPrefix(string(WARN_COLOR) + "[WARN] ")
	err := l.logger.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Error err prints err level msg.
func (l *Logger) Error(v ...interface{}) {
	if err < l.level {
		return
	}

	l.logger.SetPrefix(string(ERROR_COLOR) + "[ERROR] ")
	err := l.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Errorf prints err level msg with format.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if err < l.level {
		return
	}

	l.logger.SetPrefix(string(ERROR_COLOR) + "[ERROR] ")
	err := l.logger.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Fatal prints fatal level msg and exit process with code 1.
func (l *Logger) Fatal(v ...interface{}) {
	if fatal < l.level {
		return
	}

	l.logger.SetPrefix(string(FATAL_COLOR) + "[FATAL] ")
	err := l.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Exit(1)
}

// Fatalf prints fatal level msg with format and exit process with code 1.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if fatal < l.level {
		return
	}

	l.logger.SetPrefix(string(FATAL_COLOR) + "[FATAL] ")
	err := l.logger.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Exit(1)
}
