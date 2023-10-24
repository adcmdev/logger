package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Loglevel int

const (
	PanicLevel Loglevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel // Note: zap doesn't have a trace level, so this will map to debug
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
	Log = NewZapLogger(logLevel)
}

type zapLogger struct {
	log *zap.SugaredLogger
}

func NewZapLogger(level Loglevel) ILogger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(toZapLevel(level))
	logger, _ := config.Build()
	sugar := logger.Sugar()

	return &zapLogger{log: sugar}
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *zapLogger) Warning(args ...interface{}) {
	l.log.Warn(args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

func (l *zapLogger) Critical(args ...interface{}) {
	l.log.DPanic(args...) // Using DPanic for critical which will panic in development and log error in production
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

func toZapLevel(level Loglevel) zapcore.Level {
	switch level {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	case PanicLevel:
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}
