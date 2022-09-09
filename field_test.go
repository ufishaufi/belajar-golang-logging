package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "imanulhaq").Info("Hello World")

	logger.WithField("username", "shaufi").
		WithField("name", "Muhammad Shaufi").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "shaufi",
		"name":     "Muhammad Shaufi",
	}).Info("Hello World")
}
