package main

/*
	Example of the simplest usage of LoGo Logger
*/

import (
	logo "github.com/nurza/logo"
)

type SimpleStructure struct {
	A string
	B bool
	C int
	D float64
	F byte
}

func main() {
	var l logo.Logger            // Create a simple Logger
	l.AddTransport(logo.Console) // Add a transport: Console

	l.Log(logo.Info, "Hello World !") // This is the same as:
	l.Info("Hello World !")
}
