package main

import (
	"github.com/sirupsen/logrus"
)

func main() {

	// create a dedicated logger
	logger := logrus.New()

	// produce JSON format
	logger.SetFormatter(&logrus.JSONFormatter{})

	// log info and high severity messages
	logger.SetLevel(logrus.InfoLevel)
	logger.Debug("Running the program ...")

	logger.WithFields(logrus.Fields{
		"Env": "DEV",
		"Service": "MyAppSrv",
	}).Info("This is a test info message")

	logger.WithFields(logrus.Fields{
		"Env": "DEV",
	}).Error("Sample error message")

	logger.Info("Sample Info message")
}