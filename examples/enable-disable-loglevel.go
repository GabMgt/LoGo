package main

/*
	Example of the feature : enable and disable logging levels
*/

import (
	logo "github.com/nurza/logo"
)

func DisplayAllLevels(l *logo.Logger) {
	l.Silly("\tSilly")
	l.Debug("\tDebug")
	l.Verbose("\tVerbose")
	l.Info("\tInfo")
	l.Warn("\tWarn")
	l.Error("\tError")
	l.Critical("\tCritical")
}

func main() {
	var l logo.Logger                 // Create a simple Logger
	t := l.AddTransport(logo.Console) // Add a transport: Console
	t.AddColor(logo.ConsoleColor)     // Add a color: Console color

	logo.Log("EnableAllLevels:")
	l.EnableAllLevels() // Enable all logging levels
	DisplayAllLevels(&l)

	logo.Log("DisableAllLevels:")
	l.DisableAllLevels() // Disable all logging levels
	DisplayAllLevels(&l)

	logo.Log("Enable Silly:")
	l.EnableLevel(logo.Silly)
	DisplayAllLevels(&l)

	logo.Log("Enable Debug:")
	l.EnableLevel(logo.Debug)
	DisplayAllLevels(&l)

	logo.Log("Enable Verbose:")
	l.EnableLevel(logo.Verbose)
	DisplayAllLevels(&l)

	logo.Log("Enable Info:")
	l.EnableLevel(logo.Info)
	DisplayAllLevels(&l)

	logo.Log("Enable Warn:")
	l.EnableLevel(logo.Warn)
	DisplayAllLevels(&l)

	logo.Log("Enable Error:")
	l.EnableLevel(logo.Error)
	DisplayAllLevels(&l)

	logo.Log("Enable Critical:")
	l.EnableLevel(logo.Critical)
	DisplayAllLevels(&l)

}
