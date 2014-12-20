package logo

import (
	"fmt"
	"time"
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
	Time       string
	Prefix     string
	SillyL     bool
	DebugL     bool
	VerboseL   bool
	InfoL      bool
	WarnL      bool
	ErrorL     bool
	CriticalL  bool
}

/*
	Log something with a parameterizable logging level
*/
func (l *Logger) Log(level int, args ...interface{}) {
	var s string = ""  // The formated string that will be used for logging
	var sc string = "" // The formated string with color

	// Check if the logging level is enabled
	switch level {
	case Silly:
		if !l.SillyL {
			return
		}
	case Debug:
		if !l.DebugL {
			return
		}
	case Verbose:
		if !l.VerboseL {
			return
		}
	case Info:
		if !l.InfoL {
			return
		}
	case Warn:
		if !l.WarnL {
			return
		}
	case Error:
		if !l.ErrorL {
			return
		}
	case Critical:
		if !l.CriticalL {
			return
		}
	}

	// Format the message
	for _, v := range args {
		s += fmt.Sprintf("%v ", v)
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	} else {
		s = ""
	}

	// List all transports to apply color, prefix and time
	for _, t := range l.Transports {
		// Apply color
		if t.ConsoleColorTheme != nil {
			sc = ApplyConsoleColor(s, level, *t.ConsoleColorTheme)
		} else if t.HTMLColorTheme != nil {
			sc = ApplyHTMLColor(s, level, *t.HTMLColorTheme)
		} else {
			sc = s
		}

		// Add prefix
		if l.Prefix != "" {
			if t.ConsoleColorTheme != nil {
				sc = ApplyConsoleColor(l.Prefix+sc, level, *t.ConsoleColorTheme)
			} else if t.HTMLColorTheme != nil {
				sc = ApplyHTMLColor(l.Prefix+sc, level, *t.HTMLColorTheme)
			} else {
				sc = l.Prefix + sc
			}
		}

		// Add time
		if l.Time != "" {
			sc = time.Now().Format(l.Time) + " " + sc
		}

		t.Write(&t, sc) // Call the transport Write() function
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

/*
	These functions disable or enable logging level :
*/
func (l *Logger) DisableAllLevels() *Logger {
	l.SillyL = false
	l.DebugL = false
	l.VerboseL = false
	l.InfoL = false
	l.WarnL = false
	l.ErrorL = false
	l.CriticalL = false
	return l
}

func (l *Logger) EnableAllLevels() *Logger {
	l.SillyL = true
	l.DebugL = true
	l.VerboseL = true
	l.InfoL = true
	l.WarnL = true
	l.ErrorL = true
	l.CriticalL = true
	return l
}

func (l *Logger) EnableLevel(level int) *Logger {
	switch level {
	case Silly:
		l.SillyL = true
	case Debug:
		l.DebugL = true
	case Verbose:
		l.VerboseL = true
	case Info:
		l.InfoL = true
	case Warn:
		l.WarnL = true
	case Error:
		l.ErrorL = true
	case Critical:
		l.CriticalL = true
	}
	return l
}

func (l *Logger) DisableLevel(level int) *Logger {
	switch level {
	case Silly:
		l.SillyL = false
	case Debug:
		l.DebugL = false
	case Verbose:
		l.VerboseL = false
	case Info:
		l.InfoL = false
	case Warn:
		l.WarnL = false
	case Error:
		l.ErrorL = false
	case Critical:
		l.CriticalL = false
	}
	return l
}

/*
	Set a prefix
*/
func (l *Logger) SetPrefix(s string) *Logger {
	l.Prefix = s
	return l
}
