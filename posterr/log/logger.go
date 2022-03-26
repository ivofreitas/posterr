package log

import (
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	serviceName string = "strider"
)

var (
	entry *logrus.Logger
	once  sync.Once
)

func GetLogger() *logrus.Entry {

	once.Do(func() {
		entry = logrus.New()
		entry.SetNoLock()
		entry.SetFormatter(&logrus.JSONFormatter{})
	})

	return entry.WithFields(logrus.Fields{
		"system":  serviceName,
		"version": "1.0.0",
		"mutex":   &sync.Mutex{},
		"type":    "json",
	})
}
