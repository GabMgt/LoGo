package logo

import (
	"fmt"
)

/*
	This function is the simpliest form to write logs with LoGo.
	It writes variables in the console like fmt.Println()

	This function can be rewrite like this:
	func Log(args ...interface{}) {
		fmt.Println(args...)
	}

	But I want to leave the choice of the format to the users.
	Like replacing:	s += fmt.Sprintf("%v ", v)
	with:			s += fmt.Sprintf("%+v ", v)
	To add struct field names
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
