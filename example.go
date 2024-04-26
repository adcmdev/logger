package main

func main() {
	Init(DebugLevel)

	Log.Debug("Debug message")
	Log.Info("Info message")
	Log.Warning("Warning message")
	Log.Error("Error message")
	// Log.Critical("Critical message")
	// Log.Fatal("Fatal message")
}
