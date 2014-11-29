package logo

import (
// "./transport"
)

type Logger struct {
	// Transports []Transport
}

func (l *Logger) Log(level string, args ...interface{}) {

}

func (l *Logger) Warn(args ...interface{}) {
	l.Log("warn", args)
}
