package main

/*
	Example of a Logger with a Console transport
*/

import (
	logo "github.com/nurza/logo"
)

func main() {
	var l logo.Logger                     // Create a simple Logger
	l.AddTransport(logo.File, "test.log") // Add a transport: Console
	l.EnableAllLevels()                   // Enable all logging levels

	l.Log(logo.Info, "Hello World !") // This is the same as:
	l.Info("Hello World !")
}
