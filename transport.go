package logo

import ()

/*
	List of transports
*/
const (
	Console = iota
	File    = iota
)

/*
	A Transport struct
*/
type Transport struct {
	Write func(string)
}

/*
	Attach a new Transport to a Logger
*/
func (l *Logger) AddTransport(t int) (lo *Logger) {
	var newTransport Transport
	switch t {
	case Console:
		newTransport.Write = ConsoleWrite
		l.Transports = append(l.Transports, newTransport)
	case File: // Todo
	default:
	}

	return l
}
