package helpers

import "github.com/sirupsen/logrus"

var logger *logrus.Entry

func InitLogger() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})
	logger = log.WithFields(logrus.Fields{
		"service": "gopher-translator",
	})
}

func GetLogger() *logrus.Entry {
	if logger == nil {
		InitLogger()
	}
	return logger
}
