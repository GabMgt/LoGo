package logo

import ()

/*
	Attach a new time to a Logger
*/
func (l *Logger) AddTime(layout string) {
	l.Time = layout
}

/*
	Remove time to a Logger
*/
func (l *Logger) RemoveTime() {
	l.Time = ""
}
