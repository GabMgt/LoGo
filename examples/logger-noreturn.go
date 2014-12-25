package main

/*
	Example of a Logger with a Console transport
*/

import (
	".."
	// logo "github.com/nurza/logo"
)

func main() {
	var l logo.Logger            // Create a simple Logger
	l.AddTransport(logo.Console) // Add a transport: Console
	l.EnableAllLevels()          // Enable all logging levels

	l.Info("1")   // Line return
	l.Info("2")   // Same
	l.InfoNR("3") // No line return
	l.Info("4")

	/*
		Display:
		1
		2
		34
	*/
}
