package logger

import (
	"github.com/sirupsen/logrus"
)

type Loglevel int

const (
	PanicLevel Loglevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
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

func Init(logLevel Loglevel) {
	Log = NewLogrusLogger(logLevel)
}

type logrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger(level Loglevel) ILogger {
	logger := logrus.New()
	logger.SetLevel(toLogrusLevel(level))

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

func toLogrusLevel(level Loglevel) logrus.Level {
	switch level {
	case DebugLevel:
		return logrus.DebugLevel
	case InfoLevel:
		return logrus.InfoLevel
	case WarnLevel:
		return logrus.WarnLevel
	case ErrorLevel:
		return logrus.ErrorLevel
	case FatalLevel:
		return logrus.FatalLevel
	case PanicLevel:
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}
