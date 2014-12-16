package logo

/*
	This transport writes logs in a file
*/

import (
	"io/ioutil"
)

/*
	Write function of the transport
*/
func FileWrite(t *Transport, s string) {
	ioutil.WriteFile(t.Data[0].(string), []byte(s), 0664)
}
