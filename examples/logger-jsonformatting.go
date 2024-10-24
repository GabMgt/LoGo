package main

/*
	Example of the feature: enable json formatting
*/

import (
	logo "github.com/nurza/logo"
)

func main() {
	var l logo.Logger                                // Create a simple Logger
	t1 := l.AddTransport(logo.Console)               // Add a transport: Console
	t1.AddColor(logo.ConsoleColor)                   // Add a color: Console color
	t2 := l.AddTransport(logo.File, "test.log.json") // Add a transport: File
	t2.JSONFormatted = true                          // Enable JSON formatting
	l.EnableAllLevels()                              // Enable all logging levels
	l.SetPrefix("{ ")                                // Set Prefix
	l.SetSuffix(" }")                                // Set Suffix
	l.AddTime("2006-01-02 15:04:05")                 // Add time template

	l.Silly("Silly")
	l.Debug("Debug")
	l.Verbose("Verbose")
	l.Info("Info")
	l.Warn("Warn")
	l.Error("Error")
	l.Critical("Critical")
}
