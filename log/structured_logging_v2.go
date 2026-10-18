package main

import (
	"github.com/sirupsen/logrus"
)

func getLogger() *logrus.Entry {
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

func main() {
	logger := getLogger()
	logger.Info("This is a test info message")
	logger.Error("Sample error message")
	logger.Info("Sample Info message")
}