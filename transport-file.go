package logo

/*
	This transport writes logs in a file
*/

import (
	"os"
)

/*
	Write function of the transport

	Append a string at the end of the file
*/
func FileWrite(t *Transport, s string) {
	f, err := os.OpenFile(t.Data[0].(string), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(s); err != nil {
		panic(err)
	}
}
