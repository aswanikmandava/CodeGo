package logger

import (
	"github.com/sirupsen/logrus"
)

// the function name must start with an uppercase letter to be accessible from another package
// for example, if the function name is getLogger(), it will not be accessible from another package
// GetLogger returns a logger instance with predefined settings
func GetLogger() *logrus.Entry {
	// create a dedicated logger
	logger := logrus.New()

	// produce JSON format
	logger.SetFormatter(&logrus.JSONFormatter{})

	// log info and high severity messages
	logger.SetLevel(logrus.InfoLevel)

	// set common context fields
	return logger.WithFields(logrus.Fields{
		"Env": "DEV",
		"Service": "MyAppSrv",
	})
}