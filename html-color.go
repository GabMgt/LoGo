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
	HTMLColorDefaultBackground   = "defbg"
	HTMLColorBlackBackground     = "blabg"
	HTMLColorRedBackground       = "redbg"
	HTMLColorGreenBackground     = "grebg"
	HTMLColorYellowBackground    = "yelbg"
	HTMLColorBlueBackground      = "blubg"
	HTMLColorMagentaBackground   = "magbg"
	HTMLColorCyanBackground      = "cyabg"
	HTMLColorLightGrayBackground = "lgrbg"
	HTMLColorDarkGrayBackground  = "dgrbg"
	HTMLColorWhiteBackground     = "whibg"
)
const (
	HTMLColorDefaultBackgroundCSS   = "LightSteelBlue"
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
	HTMLColorBold       = "bold"
	HTMLColorBolder     = "bolder"
	HTMLColorUnderlined = "underln"
	HTMLColorItalic     = "italic"
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
	List of available theme for HTML logging (At the moment there is only one, but more are coming)
*/
var (
	defaultHTMLColorTheme = HTMLColorTheme{
		HTMLColorAspect{HTMLColorDarkGrayText, HTMLColorDefaultBackground, []string{}},       // Silly
		HTMLColorAspect{HTMLColorDefaultText, HTMLColorDefaultBackground, []string{}},        // Debug
		HTMLColorAspect{HTMLColorGreenText, HTMLColorDefaultBackground, []string{}},          // Verbose
		HTMLColorAspect{HTMLColorBlueText, HTMLColorDefaultBackground, []string{}},           // Info
		HTMLColorAspect{HTMLColorYellowText, HTMLColorDefaultBackground, []string{}},         // Warn
		HTMLColorAspect{HTMLColorRedText, HTMLColorDefaultBackground, []string{}},            // Error
		HTMLColorAspect{HTMLColorWhiteText, HTMLColorRedBackground, []string{HTMLColorBold}}, // Critical
	}
)

/*
	HTML Header for CSS
*/
var (
	HTMLColorHeader string = `
<style>
*{background-color: ` + HTMLColorDefaultBackgroundCSS + `;}
.` + HTMLColorDefaultText + `{color: ` + HTMLColorDefaultTextCSS + `;}
.` + HTMLColorBlackText + `{color: ` + HTMLColorBlackTextCSS + `;}
.` + HTMLColorRedText + `{color: ` + HTMLColorRedTextCSS + `;}
.` + HTMLColorGreenText + `{color: ` + HTMLColorGreenTextCSS + `;}
.` + HTMLColorYellowText + `{color: ` + HTMLColorYellowTextCSS + `;}
.` + HTMLColorBlueText + `{color: ` + HTMLColorBlueTextCSS + `;}
.` + HTMLColorMagentaText + `{color: ` + HTMLColorMagentaTextCSS + `;}
.` + HTMLColorCyanText + `{color: ` + HTMLColorCyanTextCSS + `;}
.` + HTMLColorLightGrayText + `{color: ` + HTMLColorLightGrayTextCSS + `;}
.` + HTMLColorDarkGrayText + `{color: ` + HTMLColorDarkGrayTextCSS + `;}
.` + HTMLColorWhiteText + `{color: ` + HTMLColorWhiteTextCSS + `;}

.` + HTMLColorDefaultBackground + `{background-color: ` + HTMLColorDefaultBackgroundCSS + `;display:inline;}
.` + HTMLColorBlackBackground + `{background-color: ` + HTMLColorBlackBackgroundCSS + `;display:inline;}
.` + HTMLColorRedBackground + `{background-color: ` + HTMLColorRedBackgroundCSS + `;display:inline;}
.` + HTMLColorGreenBackground + `{background-color: ` + HTMLColorGreenBackgroundCSS + `;display:inline;}
.` + HTMLColorYellowBackground + `{background-color: ` + HTMLColorYellowBackgroundCSS + `;display:inline;}
.` + HTMLColorBlueBackground + `{background-color: ` + HTMLColorBlueBackgroundCSS + `;display:inline;}
.` + HTMLColorMagentaBackground + `{background-color: ` + HTMLColorMagentaBackgroundCSS + `;display:inline;}
.` + HTMLColorCyanBackground + `{background-color: ` + HTMLColorCyanBackgroundCSS + `;display:inline;}
.` + HTMLColorLightGrayBackground + `{background-color: ` + HTMLColorLightGrayBackgroundCSS + `;display:inline;}
.` + HTMLColorDarkGrayBackground + `{background-color: ` + HTMLColorDarkGrayBackgroundCSS + `;display:inline;}
.` + HTMLColorWhiteBackground + `{background-color: ` + HTMLColorWhiteBackgroundCSS + `;display:inline;}

.` + HTMLColorBold + `{font-weight: bold;}
.` + HTMLColorBolder + `{font-weight: bolder;}
.` + HTMLColorUnderlined + `{text-decoration: underline;}
.` + HTMLColorItalic + `{font-style: italic;}
</style>
`
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

	sr += "<div><p class=\""
	sr += ct.TextColor + " "               // Add text color
	sr += ct.BackgroundColor + " "         // Add background color
	sr += strings.Join(ct.TextEffect, " ") // Add effects
	sr += "\">"
	sr += s // Add the log
	sr += "</p></div>\n"

	return sr // Return result
}
