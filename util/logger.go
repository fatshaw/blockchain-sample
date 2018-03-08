package util

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Entry

func init() {
	Logger = logrus.WithFields(logrus.Fields{"appname": "blockchain"})
	Logger.Logger.Formatter = &logrus.TextFormatter{FullTimestamp:true}
}
