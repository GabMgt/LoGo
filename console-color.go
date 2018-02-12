package logo

import (
	"strings"
)

/*
	List of Console Text Colors
*/
const (
	ConsoleColorDefaultText      = "\x1b[39m"
	ConsoleColorBlackText        = "\x1b[30m"
	ConsoleColorRedText          = "\x1b[31m"
	ConsoleColorGreenText        = "\x1b[32m"
	ConsoleColorYellowText       = "\x1b[33m"
	ConsoleColorBlueText         = "\x1b[34m"
	ConsoleColorMagentaText      = "\x1b[35m"
	ConsoleColorCyanText         = "\x1b[36m"
	ConsoleColorLightGrayText    = "\x1b[37m"
	ConsoleColorDarkGrayText     = "\x1b[90m"
	ConsoleColorLightRedText     = "\x1b[91m"
	ConsoleColorLightGreenText   = "\x1b[92m"
	ConsoleColorLightYellowText  = "\x1b[93m"
	ConsoleColorLightBlueText    = "\x1b[94m"
	ConsoleColorLightMagentaText = "\x1b[95m"
	ConsoleColorLightCyanText    = "\x1b[96m"
	ConsoleColorWhiteText        = "\x1b[97m"
)

/*
	List of Console Background Colors
*/
const (
	ConsoleColorDefaultBackground      = "\x1b[49m"
	ConsoleColorBlackBackground        = "\x1b[40m"
	ConsoleColorRedBackground          = "\x1b[41m"
	ConsoleColorGreenBackground        = "\x1b[42m"
	ConsoleColorYellowBackground       = "\x1b[43m"
	ConsoleColorBlueBackground         = "\x1b[44m"
	ConsoleColorMagentaBackground      = "\x1b[45m"
	ConsoleColorCyanBackground         = "\x1b[46m"
	ConsoleColorLightGrayBackground    = "\x1b[47m"
	ConsoleColorDarkGrayBackground     = "\x1b[100m"
	ConsoleColorLightRedBackground     = "\x1b[101m"
	ConsoleColorLightGreenBackground   = "\x1b[102m"
	ConsoleColorLightYellowBackground  = "\x1b[103m"
	ConsoleColorLightBlueBackground    = "\x1b[104m"
	ConsoleColorLightMagentaBackground = "\x1b[105m"
	ConsoleColorLightCyanBackground    = "\x1b[106m"
	ConsoleColorWhiteBackground        = "\x1b[107m"
)

/*
	List of Console Text Effects
*/
const (
	ConsoleColorBold       = "\x1b[1m"
	ConsoleColorDim        = "\x1b[2m"
	ConsoleColorUnderlined = "\x1b[4m"
	ConsoleColorBlink      = "\x1b[5m"
	ConsoleColorReverse    = "\x1b[7m"
	ConsoleColorHidden     = "\x1b[8m"
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
	darkConsoleColorTheme = ConsoleColorTheme{
		ConsoleColorAspect{ConsoleColorDarkGrayText, ConsoleColorDefaultBackground, []string{}},            // Silly
		ConsoleColorAspect{ConsoleColorDefaultText, ConsoleColorDefaultBackground, []string{}},             // Debug
		ConsoleColorAspect{ConsoleColorGreenText, ConsoleColorDefaultBackground, []string{}},               // Verbose
		ConsoleColorAspect{ConsoleColorLightBlueText, ConsoleColorDefaultBackground, []string{}},           // Info
		ConsoleColorAspect{ConsoleColorLightYellowText, ConsoleColorDefaultBackground, []string{}},         // Warn
		ConsoleColorAspect{ConsoleColorLightRedText, ConsoleColorDefaultBackground, []string{}},            // Error
		ConsoleColorAspect{ConsoleColorDefaultText, ConsoleColorRedBackground, []string{ConsoleColorBold}}, // Critical
	}
	lightConsoleColorTheme = ConsoleColorTheme{
		ConsoleColorAspect{ConsoleColorDarkGrayText, ConsoleColorDefaultBackground, []string{}},            // Silly
		ConsoleColorAspect{ConsoleColorDefaultText, ConsoleColorDefaultBackground, []string{}},             // Debug
		ConsoleColorAspect{ConsoleColorGreenText, ConsoleColorDefaultBackground, []string{}},               // Verbose
		ConsoleColorAspect{ConsoleColorBlueText, ConsoleColorDefaultBackground, []string{}},                // Info
		ConsoleColorAspect{ConsoleColorYellowText, ConsoleColorDefaultBackground, []string{}},              // Warn
		ConsoleColorAspect{ConsoleColorRedText, ConsoleColorDefaultBackground, []string{}},                 // Error
		ConsoleColorAspect{ConsoleColorDefaultText, ConsoleColorRedBackground, []string{ConsoleColorBold}}, // Critical
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
