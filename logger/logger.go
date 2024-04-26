package logger

import (
	"os"

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

var logrusLoggerInstance *logrusLogger

func Init(logLevel Loglevel) {
	Log = NewLogrusLogger(logLevel)
}

type logrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger(level Loglevel) ILogger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(toLogrusLevel(level))

	logrusLoggerInstance = &logrusLogger{log: logger}

	return logrusLoggerInstance
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *logrusLogger) Warning(args ...interface{}) {
	l.log.Warn(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

func (l *logrusLogger) Critical(args ...interface{}) {
	l.log.Panic(args...)
}

func SetLevel(level Loglevel) {
	logrusLoggerInstance.log.SetLevel(toLogrusLevel(level))
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

func LevelFromString(level string) Loglevel {
	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	case "panic":
		return PanicLevel
	default:
		return ErrorLevel
	}
}

func LevelToString(level Loglevel) string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	default:
		return "error"
	}
}

func LogByLogLevel(level Loglevel, args ...interface{}) {
	switch level {
	case DebugLevel:
		Log.Debug(args...)
	case InfoLevel:
		Log.Info(args...)
	case WarnLevel:
		Log.Warning(args...)
	case ErrorLevel:
		Log.Error(args...)
	case FatalLevel:
		Log.Fatal(args...)
	case PanicLevel:
		Log.Critical(args...)
	default:
		Log.Error(args...)
	}
}
