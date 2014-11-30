package logo

import ()

/*
	List of Colors
*/
const (
	ConsoleColor = iota
)

/*
	Attach a new Color to a Logger
*/
func (t *Transport) AddColor(colorType int) {
	switch colorType {
	case ConsoleColor:
		if t.Type == Console {
			t.ConsoleColorTheme = &defaultConsoleColorTheme
		}
	}
}
