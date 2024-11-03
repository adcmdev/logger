package main

import "github.com/adcmdev/logger"

func main() {
	logger.New(logger.DebugLevel)

	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warning("Warning message")
	logger.Error("Error message")

	logger.Debug("Current log level is ", logger.LevelToString(logger.CurrentLevel()))
}
