package main

/*
	Example of a Logger with a Console transport
*/

import (
	logo "github.com/nurza/logo"
)

func main() {
	var l logo.Logger            // Create a simple Logger
	l.AddTransport(logo.Console) // Add a transport: Console
	l.EnableAllLevels()          // Enable all logging levels

	l.InfoF("Hello %s !\n%t \n%d \n%f \n%b \n%x ", "World", true, 42, 42.42, 42, 3735928559)
}
