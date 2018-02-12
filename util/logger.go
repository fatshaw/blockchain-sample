package util

import (
	log "github.com/sirupsen/logrus"
)

var Logger *log.Entry

func init() {
	Logger = log.WithFields(log.Fields{"appname": "blockchain"})
	Logger.Logger.Formatter = &log.TextFormatter{FullTimestamp:true}
}
