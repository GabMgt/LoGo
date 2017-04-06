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
	Type              int
	Write             func(*Transport, string)
	ConsoleColorTheme *ConsoleColorTheme
	HTMLColorTheme    *HTMLColorTheme
	Data              []interface{}
	JSONFormatted     bool
}

/*
	JSON Formatted struct
*/
type JSONFormatted struct {
	Prefix      string `json:"prefix,omitempty"`
	Suffix      string `json:"suffix,omitempty"`
	Level       string `json:"level"`
	GoRoutineID string `json:"goroutineid,omitempty"`
	Time        string `json:"time,omitempty"`
	Text        string `json:"text"`
}

/*
	Attach a new Transport to a Logger
*/
func (l *Logger) AddTransport(t int, param ...interface{}) *Transport {
	var newTransport Transport

	switch t {
	case Console:
		newTransport.Type = Console
		newTransport.Write = ConsoleWrite
		l.Transports = append(l.Transports, newTransport)
	case File:
		newTransport.Type = File
		newTransport.Write = FileWrite
		newTransport.Data = []interface{}{param[0].(string)}
		l.Transports = append(l.Transports, newTransport)
	}

	return &(l.Transports[len(l.Transports)-1]) // Return the pointer of last Transport of the Logger
}
