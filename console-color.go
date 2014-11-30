package logo

import (
	"strings"
)

/*
	List of Console Text Colors
*/
const (
	DefaultText      = "\x1b[39m"
	BlackText        = "\x1b[30m"
	RedText          = "\x1b[31m"
	GreenText        = "\x1b[32m"
	YellowText       = "\x1b[33m"
	BlueText         = "\x1b[34m"
	MagentaText      = "\x1b[35m"
	CyanText         = "\x1b[36m"
	LightGrayText    = "\x1b[37m"
	DarkGrayText     = "\x1b[90m"
	LightRedText     = "\x1b[91m"
	LightGreenText   = "\x1b[92m"
	LightYellowText  = "\x1b[93m"
	LightBlueText    = "\x1b[94m"
	LightMagentaText = "\x1b[95m"
	LightCyanText    = "\x1b[96m"
	WhiteText        = "\x1b[97m"
)

/*
	List of Console Background Colors
*/
const (
	DefaultBackground      = "\x1b[49m"
	BlackBackground        = "\x1b[40m"
	RedBackground          = "\x1b[41m"
	GreenBackground        = "\x1b[42m"
	YellowBackground       = "\x1b[43m"
	BlueBackground         = "\x1b[44m"
	MagentaBackground      = "\x1b[45m"
	CyanBackground         = "\x1b[46m"
	LightGrayBackground    = "\x1b[47m"
	DarkGrayBackground     = "\x1b[100m"
	LightRedBackground     = "\x1b[101m"
	LightGreenBackground   = "\x1b[102m"
	LightYellowBackground  = "\x1b[103m"
	LightBlueBackground    = "\x1b[104m"
	LightMagentaBackground = "\x1b[105m"
	LightCyanBackground    = "\x1b[106m"
	WhiteBackground        = "\x1b[107m"
)

/*
	List of Console Text Effects
*/
const (
	Bold       = "\x1b[1m"
	Dim        = "\x1b[2m"
	Underlined = "\x1b[4m"
	Blink      = "\x1b[5m"
	Reverse    = "\x1b[7m"
	Hidden     = "\x1b[8m"
)

/*
	A Color Aspect describes text and background colors and effects
*/
type ConsoleColorAspect struct {
	TextColor       string
	BackgroundColor string
	TextEffect      []string
}

/*
	A Color Theme is a list of Color Aspect (one for each logging level)
*/
type ConsoleColorTheme struct {
	Silly    ConsoleColorAspect
	Debug    ConsoleColorAspect
	Verbose  ConsoleColorAspect
	Info     ConsoleColorAspect
	Warn     ConsoleColorAspect
	Error    ConsoleColorAspect
	Critical ConsoleColorAspect
}

/*
	List of available theme for console logging (At the moment there is only one, but more are coming)
*/
var (
	defaultConsoleColorTheme = ConsoleColorTheme{
		ConsoleColorAspect{DarkGrayText, DefaultBackground, []string{}},    // Silly
		ConsoleColorAspect{DefaultText, DefaultBackground, []string{}},     // Debug
		ConsoleColorAspect{LightGreenText, DefaultBackground, []string{}},  // Verbose
		ConsoleColorAspect{LightBlueText, DefaultBackground, []string{}},   // Info
		ConsoleColorAspect{LightYellowText, DefaultBackground, []string{}}, // Warn
		ConsoleColorAspect{LightRedText, DefaultBackground, []string{}},    // Error
		ConsoleColorAspect{DefaultText, RedBackground, []string{Bold}},     // Critical
	}
)

/*
	Apply a color theme on a string with the corect logging level
*/
func ApplyConsoleColor(s string, level int, c ConsoleColorTheme) string {
	var sr string = ""        // String to return
	var ct ConsoleColorAspect // Console Color Aspect to use

	// Get the correct logging level
	switch level {
	case Silly:
		ct = c.Silly
	case Debug:
		ct = c.Debug
	case Verbose:
		ct = c.Verbose
	case Info:
		ct = c.Info
	case Warn:
		ct = c.Warn
	case Error:
		ct = c.Error
	case Critical:
		ct = c.Critical
	}

	sr += ct.TextColor                    // Add text color
	sr += ct.BackgroundColor              // Add background color
	sr += strings.Join(ct.TextEffect, "") // Add effects
	sr += s                               // Add the log
	sr += "\x1b[0m"                       // Reset all colors and effects

	return sr // Return result
}
