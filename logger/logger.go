package logger

import (
	"github.com/sirupsen/logrus"
)

type ILogger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	Critical(v ...interface{})
}

var Log ILogger

func Init() {
	Log = NewLogrusLogger(logrus.ErrorLevel)
}

type logrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger(level logrus.Level) ILogger {
	logger := logrus.New()
	logger.SetLevel(level)
	return &logrusLogger{log: logger}
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.log.Debugf("DEBUG: %v", args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.log.Infof("INFO: %v", args...)
}

func (l *logrusLogger) Warning(args ...interface{}) {
	l.log.Warnf("WARN: %v", args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.log.Errorf("ERROR: %v", args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.log.Fatalf("FATAL: %v", args...)
}

func (l *logrusLogger) Critical(args ...interface{}) {
	l.log.Panicf("PANIC: %v", args...)
}

func Debug(args ...interface{}) {
	Log.Debug(args)
}

func Info(args ...interface{}) {
	Log.Info(args)
}

func Warning(args ...interface{}) {
	Log.Warning(args)
}

func Error(args ...interface{}) {
	Log.Error(args)
}

func Fatal(args ...interface{}) {
	Log.Fatal(args)
}

func Critical(args ...interface{}) {
	Log.Critical(args)
}
