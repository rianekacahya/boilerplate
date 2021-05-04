package gologger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func New() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)

	return log
}
