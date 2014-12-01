package logo

import (
	"fmt"
)

/*
	List of logging levels
*/
const (
	Silly    = iota
	Debug    = iota
	Verbose  = iota
	Info     = iota
	Warn     = iota
	Error    = iota
	Critical = iota
)

/*
	A Logger struct
*/
type Logger struct {
	Transports []Transport
}

/*
	Log something with a parameterizable logging level
*/
func (l *Logger) Log(level int, args ...interface{}) {
	var s string = ""  // The formated string that will be used for logging
	var sc string = "" // The formated string with color

	for _, v := range args {
		s += fmt.Sprintf("%v ", v)
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	} else {
		s = ""
	}

	for _, t := range l.Transports { // List all transports to apply color
		switch t.Type {
		case Console:
			if t.ConsoleColorTheme != nil {
				sc = ApplyConsoleColor(s, level, *t.ConsoleColorTheme)
			} else {
				sc = s
			}
		default:
			sc = s
		}
		t.Write(sc) // Call the transport Write() function
	}
}

/*
	These functions call Log with the correct logging level :
*/
func (l *Logger) Silly(args ...interface{}) {
	l.Log(Silly, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.Log(Debug, args...)
}

func (l *Logger) Verbose(args ...interface{}) {
	l.Log(Verbose, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.Log(Info, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.Log(Warn, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Log(Error, args...)
}

func (l *Logger) Critical(args ...interface{}) {
	l.Log(Critical, args...)
}
