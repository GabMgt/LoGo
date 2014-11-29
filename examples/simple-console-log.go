package main

/*
	Example of the simplest usage of LoGo
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
	logo.Log(true)                                                 // bool
	logo.Log(42)                                                   // int
	logo.Log(42.42)                                                // float64
	logo.Log('A')                                                  // byte
	logo.Log(SimpleStructure{"Hello World", true, 42, 42.42, 'A'}) // Structure
	logo.Log("Hello", true, 42, "World")                           // Mixed

	logo.Log("Hello", "World", "!") // Is the same as :
	logo.Log("Hello World !")
}
