package logo

import (
	"strings"
)

/*
	List of HTML Text Colors
*/
const (
	HTMLColorDefaultText   = "deft"
	HTMLColorBlackText     = "blat"
	HTMLColorRedText       = "redt"
	HTMLColorGreenText     = "gret"
	HTMLColorYellowText    = "yelt"
	HTMLColorBlueText      = "blut"
	HTMLColorMagentaText   = "magt"
	HTMLColorCyanText      = "cyat"
	HTMLColorLightGrayText = "lgrt"
	HTMLColorDarkGrayText  = "dgrt"
	HTMLColorWhiteText     = "whit"
)
const (
	HTMLColorDefaultTextCSS   = "black"
	HTMLColorBlackTextCSS     = "black"
	HTMLColorRedTextCSS       = "red"
	HTMLColorGreenTextCSS     = "green"
	HTMLColorYellowTextCSS    = "yellow"
	HTMLColorBlueTextCSS      = "blue"
	HTMLColorMagentaTextCSS   = "magenta"
	HTMLColorCyanTextCSS      = "cyan"
	HTMLColorLightGrayTextCSS = "lightgray"
	HTMLColorDarkGrayTextCSS  = "darkgray"
	HTMLColorWhiteTextCSS     = "white"
)

/*
	List of HTML Background Colors
*/
const (
	HTMLColorDefaultBackground      = "defbg"
	HTMLColorBlackBackground        = "blabg"
	HTMLColorRedBackground          = "redbg"
	HTMLColorGreenBackground        = "grebg"
	HTMLColorYellowBackground       = "yelbg"
	HTMLColorBlueBackground         = "blubg"
	HTMLColorMagentaBackground      = "magbg"
	HTMLColorCyanBackground         = "cyabg"
	HTMLColorLightGrayBackground    = "lgrabg"
	HTMLColorDarkGrayBackground     = "darbg"
	HTMLColorLightRedBackground     = "lrebg"
	HTMLColorLightGreenBackground   = "lgrebg"
	HTMLColorLightYellowBackground  = "lyebg"
	HTMLColorLightBlueBackground    = "lblbg"
	HTMLColorLightMagentaBackground = "lmabg"
	HTMLColorLightCyanBackground    = "lcybg"
	HTMLColorWhiteBackground        = "whibg"
)
const (
	HTMLColorDefaultBackgroundCSS   = "black"
	HTMLColorBlackBackgroundCSS     = "black"
	HTMLColorRedBackgroundCSS       = "red"
	HTMLColorGreenBackgroundCSS     = "green"
	HTMLColorYellowBackgroundCSS    = "yellow"
	HTMLColorBlueBackgroundCSS      = "blue"
	HTMLColorMagentaBackgroundCSS   = "magenta"
	HTMLColorCyanBackgroundCSS      = "cyan"
	HTMLColorLightGrayBackgroundCSS = "lightgray"
	HTMLColorDarkGrayBackgroundCSS  = "darkgray"
	HTMLColorWhiteBackgroundCSS     = "white"
)

/*
	List of HTML Text Effects
*/
const (
	Bold       = "bold"
	Underlined = "underln"
)

/*
	A Color Aspect describes text and background colors and effects
*/
type HTMLColorAspect struct {
	TextColor       string
	BackgroundColor string
	TextEffect      []string
}

/*
	A Color Theme is a list of Color Aspect (one for each logging level)
*/
type HTMLColorTheme struct {
	Silly    HTMLColorAspect
	Debug    HTMLColorAspect
	Verbose  HTMLColorAspect
	Info     HTMLColorAspect
	Warn     HTMLColorAspect
	Error    HTMLColorAspect
	Critical HTMLColorAspect
}

/*
	List of available theme for console logging (At the moment there is only one, but more are coming)
*/
var (
	defaultHTMLColorTheme = HTMLColorTheme{
		HTMLColorAspect{HTMLColorDarkGrayText, HTMLColorDefaultBackground, []string{}}, // Silly
		HTMLColorAspect{HTMLColorDefaultText, HTMLColorDefaultBackground, []string{}},  // Debug
		HTMLColorAspect{HTMLColorGreenText, HTMLColorDefaultBackground, []string{}},    // Verbose
		HTMLColorAspect{HTMLColorBlueText, HTMLColorDefaultBackground, []string{}},     // Info
		HTMLColorAspect{HTMLColorYellowText, HTMLColorDefaultBackground, []string{}},   // Warn
		HTMLColorAspect{HTMLColorRedText, HTMLColorDefaultBackground, []string{}},      // Error
		HTMLColorAspect{HTMLColorDefaultText, HTMLColorRedBackground, []string{Bold}},  // Critical
	}
)

/*
	Apply a color theme on a string with the corect logging level
*/
func ApplyHTMLColor(s string, level int, c HTMLColorTheme) string {
	var sr string = ""     // String to return
	var ct HTMLColorAspect // HTML Color Aspect to use

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

	sr += "<div class=\""
	sr += ct.TextColor + " "               // Add text color
	sr += ct.BackgroundColor + " "         // Add background color
	sr += strings.Join(ct.TextEffect, " ") // Add effects
	sr += "\">"
	sr += s // Add the log
	sr += "</div>"

	return sr // Return result
}
