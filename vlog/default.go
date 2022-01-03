package vlog

var vlog = Default()

// Debug debug prints debug level msg.
func Debug(v ...interface{}) {
	vlog.Debug(v)
}

// Debugf prints debug level msg with format.
func Debugf(format string, v ...interface{}) {
	vlog.Debugf(format, v)
}

// Info prints info level msg.
func Info(v ...interface{}) {
	vlog.Info(v)
}

// Infof prints info level msg with format.
func Infof(format string, v ...interface{}) {
	vlog.Infof(format, v)
}

// Warn prints warning level msg.
func Warn(v ...interface{}) {
	vlog.Warn(v)
}

// Warnf prints warning level msg with format.
func Warnf(format string, v ...interface{}) {
	vlog.Warnf(format, v)
}

// Error err prints err level msg.
func Error(v ...interface{}) {
	vlog.Error(v)
}

// Errorf prints err level msg with format.
func Errorf(format string, v ...interface{}) {
	vlog.Errorf(format, v)
}

// Fatal prints fatal level msg and exit process with code 1.
func Fatal(v ...interface{}) {
	vlog.Fatal(v)
}

// Fatalf prints fatal level msg with format and exit process with code 1.
func Fatalf(format string, v ...interface{}) {
	vlog.Fatalf(format, v)
}
