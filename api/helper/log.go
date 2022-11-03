package helper

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Log struct {
	log *logrus.Logger
}

func NewLog(log *logrus.Logger) Log {
	return Log{
		log: log,
	}
}

func (log Log) SetUp(output string) {
	log.log.SetLevel(logrus.InfoLevel)
	file, error := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if error != nil {
		log.log.Fatal("Failed to log to file, using default stderr")
	}
	log.log.SetOutput(file)
}

func (log Log) Infof(format string, args ...interface{}) {
	log.log.Infof(format, args...)
}

func (log Log) Errorf(format string, args ...interface{}) {
	log.log.Errorf(format, args...)
}

func (l Log) Fatal(message string) {
	l.log.Fatal(message)
}

func (l Log) Panic(message string) {
	l.log.Panic(message)
}

func (l Log) Debug(message string) {
	l.log.Debug(message)
}

func (l Log) Warn(message string) {
	l.log.Warn(message)
}

func (l Log) Trace(message string) {
	l.log.Trace(message)
}
