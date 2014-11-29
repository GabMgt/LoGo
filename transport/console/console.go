package console

/*
	This transport writes logs in the console (standard output)
*/

import (
	"fmt"
)

/*
	Write function of the transport
*/
func Write(s string) {
	fmt.Println(s)
}
