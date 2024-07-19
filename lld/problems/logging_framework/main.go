package main

import (
	"main/models"
)

func main() {
	// consolePrinter := models.ConsolePrinter{}
	consolePrinter := models.FilePrinter{}
	logger := models.NewLogger(&consolePrinter)
	logger.Debug("Hello ji saale kaise ho.")
	logger.Error("Hello ji saale kaise ho.")
	logger.Info("Hello ji saale kaise ho.")
	logger.Warn("Hello ji saale kaise ho.")
	
}
