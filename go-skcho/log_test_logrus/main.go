package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {

	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     false,
		FullTimestamp:   true,
		ForceQuote:      false,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	log.Out = os.Stdout
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.SetReportCaller(true)
	log.SetLevel(logrus.InfoLevel)

}

func main() {

	log.Infof("main() start")

	task()

}

func task() {
	log.Warnf("task() start: id=%v", 12)
}
