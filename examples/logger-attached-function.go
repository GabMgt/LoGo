package main

/*
	Example of a Logger with attached function
*/

import (
	logo "github.com/nurza/logo"
	"os"
)

/*
	Exit the program when critical error
*/
func CriticalExit(s string) {
	println("Critical error => Exit program")
	os.Exit(1)
}

func main() {
	var l logo.Logger                 // Create a simple Logger
	t := l.AddTransport(logo.Console) // Add a transport: Console
	t.AddColor(logo.ConsoleColor)     // Add a color: Console color
	l.EnableAllLevels()               // Enable all logging levels

	l.AttachFunction(logo.Critical, CriticalExit) // Attach the function CriticalExit(string)

	l.Silly("Silly")
	l.Debug("Debug")
	l.Verbose("Verbose")
	l.Info("Info")
	l.Warn("Warn")
	l.Error("Error")
	l.Critical("Critical") // Critical log => CriticalExit(string) is called

}
