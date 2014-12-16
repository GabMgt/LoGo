package main

/*
	Example of a Logger with console colors
*/

import (
	logo "github.com/nurza/logo"
)

func main() {
	var l logo.Logger                 // Create a simple Logger
	t := l.AddTransport(logo.Console) // Add a transport: Console
	t.AddColor(logo.ConsoleColor)     // Add a color: Console color
	l.EnableAllLevels()               // Enable all logging levels

	l.Silly("Silly")
	l.Debug("Debug")
	l.Verbose("Verbose")
	l.Info("Info")
	l.Warn("Warn")
	l.Error("Error")
	l.Critical("Critical")
}
