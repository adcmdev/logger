package main

import (
	"errors"

	"github.com/innovafour/logger"
)

func main() {
	logger.New(logger.DebugLevel)

	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warning("Warning message")
	logger.Error("Error message", errors.New("error message x2"))
	logger.Debug("Current log level is", logger.LevelToString(logger.CurrentLevel()))
}
