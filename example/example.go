package logger

import "github.com/adcmdev/logger"

func main() {
	logger.Init(logger.DebugLevel)

	logger.Log.Debug("Debug message")
	logger.Log.Info("Info message")
	logger.Log.Warning("Warning message")
	logger.Log.Error("Error message")
	// logger.Log.Critical("Critical message")
	// logger.Log.Fatal("Fatal message")
}
