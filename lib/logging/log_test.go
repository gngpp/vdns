package logging

import (
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func Test(t *testing.T) {
	log.SetFormatter(&log.JSONFormatter{
		// HideKeys:        true,
		TimestampFormat: time.RFC3339,
		FieldMap:        log.FieldMap{"name": "age"},
	})
	log.WithFields(log.Fields{
		"name": "dj",
		"age":  18,
	}).Info("info msg")
	log.Trace("trace msg")
	log.Debug("debug msg")
	log.Info("info msg")
	log.Warn("warn msg")
	log.Error("error msg")
	log.Fatal("fatal msg")
	log.Panic("panic msg")
}
