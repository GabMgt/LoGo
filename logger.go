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
	Transports               []Transport
	Time                     string
	Prefix                   string
	Suffix                   string
	SillyL                   bool
	DebugL                   bool
	VerboseL                 bool
	InfoL                    bool
	WarnL                    bool
	ErrorL                   bool
	CriticalL                bool
	SillyAttachedFunction    func(string)
	DebugAttachedFunction    func(string)
	VerboseAttachedFunction  func(string)
	InfoAttachedFunction     func(string)
	WarnAttachedFunction     func(string)
	ErrorAttachedFunction    func(string)
	CriticalAttachedFunction func(string)
}

/*
	Log something with a parameterizable logging level
*/
func (l *Logger) Log(level int, args ...interface{}) {
	var s string = ""                 // The formated string that will be used for logging
	var sc string = ""                // The formated string with color
	var attachedFunction func(string) // Attached function

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
		sc = s

		// Add prefix
		if l.Prefix != "" {
			sc = l.Prefix + sc
		}

		// Add suffix
		if l.Suffix != "" {
			sc += l.Suffix
		}

		// Apply color
		if t.ConsoleColorTheme != nil {
			sc = ApplyConsoleColor(sc, level, *t.ConsoleColorTheme)
		} else if t.HTMLColorTheme != nil {
			sc = ApplyHTMLColor(sc, level, *t.HTMLColorTheme)
		}

		// Add time
		if l.Time != "" {
			sc = time.Now().Format(l.Time) + " " + sc
		}

		// Call the transport Write() function
		t.Write(&t, sc+"\n")

		// Call the attached function if exists
		switch level {
		case Silly:
			attachedFunction = l.SillyAttachedFunction
		case Debug:
			attachedFunction = l.DebugAttachedFunction
		case Verbose:
			attachedFunction = l.VerboseAttachedFunction
		case Info:
			attachedFunction = l.InfoAttachedFunction
		case Warn:
			attachedFunction = l.WarnAttachedFunction
		case Error:
			attachedFunction = l.ErrorAttachedFunction
		case Critical:
			attachedFunction = l.CriticalAttachedFunction
		}
		if attachedFunction != nil {
			attachedFunction(s)
		}
	}
}

/*
	Log something with a parameterizable logging level without line return
*/
func (l *Logger) LogNR(level int, args ...interface{}) {
	var s string = ""                 // The formated string that will be used for logging
	var sc string = ""                // The formated string with color
	var attachedFunction func(string) // Attached function

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
		sc = s

		// Add prefix
		if l.Prefix != "" {
			sc = l.Prefix + sc
		}

		// Add suffix
		if l.Suffix != "" {
			sc += l.Suffix
		}

		// Apply color
		if t.ConsoleColorTheme != nil {
			sc = ApplyConsoleColor(sc, level, *t.ConsoleColorTheme)
		} else if t.HTMLColorTheme != nil {
			sc = ApplyHTMLColor(sc, level, *t.HTMLColorTheme)
		}

		// Add time
		if l.Time != "" {
			sc = time.Now().Format(l.Time) + " " + sc
		}

		// Call the transport Write() function
		t.Write(&t, sc)

		// Call the attached function if exists
		switch level {
		case Silly:
			attachedFunction = l.SillyAttachedFunction
		case Debug:
			attachedFunction = l.DebugAttachedFunction
		case Verbose:
			attachedFunction = l.VerboseAttachedFunction
		case Info:
			attachedFunction = l.InfoAttachedFunction
		case Warn:
			attachedFunction = l.WarnAttachedFunction
		case Error:
			attachedFunction = l.ErrorAttachedFunction
		case Critical:
			attachedFunction = l.CriticalAttachedFunction
		}
		if attachedFunction != nil {
			attachedFunction(s)
		}
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
	These functions call Log with the correct logging level without line return:
*/
func (l *Logger) SillyNR(args ...interface{}) {
	l.LogNR(Silly, args...)
}

func (l *Logger) DebugNR(args ...interface{}) {
	l.LogNR(Debug, args...)
}

func (l *Logger) VerboseNR(args ...interface{}) {
	l.LogNR(Verbose, args...)
}

func (l *Logger) InfoNR(args ...interface{}) {
	l.LogNR(Info, args...)
}

func (l *Logger) WarnNR(args ...interface{}) {
	l.LogNR(Warn, args...)
}

func (l *Logger) ErrorNR(args ...interface{}) {
	l.LogNR(Error, args...)
}

func (l *Logger) CriticalNR(args ...interface{}) {
	l.LogNR(Critical, args...)
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

/*
	Set a suffix
*/
func (l *Logger) SetSuffix(s string) *Logger {
	l.Suffix = s
	return l
}

/*
	Attach a function
*/
func (l *Logger) AttachFunction(level int, function func(string)) {
	switch level {
	case Silly:
		l.SillyAttachedFunction = function
	case Debug:
		l.DebugAttachedFunction = function
	case Verbose:
		l.VerboseAttachedFunction = function
	case Info:
		l.InfoAttachedFunction = function
	case Warn:
		l.WarnAttachedFunction = function
	case Error:
		l.ErrorAttachedFunction = function
	case Critical:
		l.CriticalAttachedFunction = function
	}
}
