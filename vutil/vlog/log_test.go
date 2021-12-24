package vlog

import (
	"testing"
)

func TestLog(t *testing.T) {
	logger := Log.Default()
	logger.Infof("info")
	logger.Infof("info value:%s", "test")

}

func TestSetLevel(t *testing.T) {
	Log.SetLevel("trace")
}

func TestTrace(t *testing.T) {
	logger := Log.Default()
	Log.SetLevel("trace")
	logger.Trace("trace")
	logger.SetLevel("off")
	logger.Trace("trace")
}

func TestTracef(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel("trace")
	logger.Tracef("tracef")
	logger.SetLevel("off")
	if logger.IsTraceEnabled() {
		logger.Tracef("tracef")
	}
}

func TestDebug(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel("debug")
	logger.Debug("debug")
	logger.SetLevel("off")
	logger.Debug("debug")
}

func TestDebugf(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel("debug")
	logger.Debugf("debugf")
	logger.SetLevel("off")
	logger.Debug("debug")
}

func TestInfo(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel("info")
	logger.Info("info")
	logger.SetLevel("off")
	logger.Info("info")
}

func TestInfof(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel("info")
	logger.Infof("infof")
	logger.SetLevel("off")
	logger.Infof("infof")
}

func TestWarn(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel("warn")
	logger.Warn("warn")
	logger.SetLevel("off")
	logger.Warn("warn")
}

func TestWarnf(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel("warn")
	logger.Warnf("warnf")
	logger.SetLevel("off")
	logger.Warnf("warnf")
}

func TestError(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel("error")
	logger.Error("error")
	logger.SetLevel("off")
	logger.Error("error")
}

func TestErrorf(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel(Level.OFF)
	logger.Errorf("errorf")
	logger.SetLevel("off")
	logger.Errorf("errorf")
}

func TestGetLevel(t *testing.T) {
	if getLevel("trace") != Trace {
		t.FailNow()

		return
	}

	if getLevel("debug") != Debug {
		t.FailNow()

		return
	}

	if getLevel("info") != Info {
		t.FailNow()

		return
	}

	if getLevel("warn") != Warn {
		t.FailNow()

		return
	}

	if getLevel("error") != Error {
		t.FailNow()

		return
	}
}

func TestLoggerSetLevel(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel("trace")

	if logger.level != Trace {
		t.FailNow()

		return
	}
}

func TestIsTraceEnabled(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel("trace")

	if !logger.IsTraceEnabled() {
		t.FailNow()

		return
	}
}

func TestIsDebugEnabled(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel("debug")

	if !logger.IsDebugEnabled() {
		t.FailNow()

		return
	}
}

func TestIsWarnEnabled(t *testing.T) {
	logger := Log.Default()
	logger.SetLevel("warn")
	if !logger.IsWarnEnabled() {
		t.FailNow()
		return
	}
}
