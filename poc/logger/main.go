package main

import "github.com/nibrasmuhamed/go-modules/logger"

func main() {
	logger := logger.Init()

	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Debug("This is a debug message")
	logger.Error("This is an error message")
}
