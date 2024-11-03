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
)

type ILogger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	Critical(v ...interface{})
}

var zapLoggerInstance *zapLogger

type zapLogger struct {
	log   *zap.SugaredLogger
	level zap.AtomicLevel
}

func New(level Loglevel) ILogger {
	atom := zap.NewAtomicLevelAt(toZapLevel(level))

	encoderConfig := zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
	}

	config := zap.Config{
		Level:            atom,
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ := config.Build()
	sugar := logger.Sugar()

	zapLoggerInstance = &zapLogger{log: sugar, level: atom}

	return zapLoggerInstance
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
	l.log.DPanic(args...)
}

func CurrentLevel() Loglevel {
	return Loglevel(zapLoggerInstance.level.Level())
}

func SetLevel(level Loglevel) {
	zapLoggerInstance.level.SetLevel(toZapLevel(level))
	LogByLogLevel(level, "Log level changed to ", LevelToString(level))
}

func Debug(args ...interface{}) {
	zapLoggerInstance.Debug(args...)
}

func Info(args ...interface{}) {
	zapLoggerInstance.Info(args...)
}

func Warning(args ...interface{}) {
	zapLoggerInstance.Warning(args...)
}

func Error(args ...interface{}) {
	zapLoggerInstance.Error(args...)
}

func Fatal(args ...interface{}) {
	zapLoggerInstance.Fatal(args...)
}

func Critical(args ...interface{}) {
	zapLoggerInstance.Critical(args...)
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
		zapLoggerInstance.Debug(args...)
	case InfoLevel:
		zapLoggerInstance.Info(args...)
	case WarnLevel:
		zapLoggerInstance.Warning(args...)
	case ErrorLevel:
		zapLoggerInstance.Error(args...)
	case FatalLevel:
		zapLoggerInstance.Fatal(args...)
	case PanicLevel:
		zapLoggerInstance.Critical(args...)
	default:
		zapLoggerInstance.Error(args...)
	}
}
