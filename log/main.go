package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.WithFields(logrus.Fields{"severity": "info"}).Info("233")
	logrus.WithFields(logrus.Fields{"c": "d"}).WithTime(time.Now()).Info("233")

	return
}
