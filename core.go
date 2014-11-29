package logo

import (
	"fmt"
)

/*
	This function is the simpliest form to write logs with LoGo.

	It writes variables in the console like fmt.Println()
*/
func Log(args ...interface{}) {
	var s string = ""

	for _, v := range args {
		s += fmt.Sprintf("%v ", v)
	}
	if len(s) > 0 {
		fmt.Println(s[:len(s)-1])
	} else {
		fmt.Println()
	}
}
