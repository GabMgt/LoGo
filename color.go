package logo

import ()

/*
	List of Colors
*/
const (
	ConsoleColor = iota
	HTMLColor    = iota
)

/*
	Attach a new Color to a Logger
*/
func (t *Transport) AddColor(colorType int) {
	switch colorType {
	case ConsoleColor:
		t.ConsoleColorTheme = &defaultConsoleColorTheme
	case HTMLColor:
		t.HTMLColorTheme = &defaultHTMLColorTheme
		t.Write(t, HTMLColorHeader)
	}
}
