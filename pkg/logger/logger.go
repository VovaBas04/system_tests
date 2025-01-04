package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{*logrus.New()}
}
