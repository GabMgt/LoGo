package logo

/*
	This transport writes logs in the console (standard output)
*/

import (
	"fmt"
)

/*
	Write function of the transport
*/
func ConsoleWrite(t *Transport, s string) {
	fmt.Print(s)
}
