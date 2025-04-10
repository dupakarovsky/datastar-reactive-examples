package main

import "time"

// parseStrindDuration will parse a string to a time.Duration format, setting to zero if the
// parsing fails
func (app *application) parseStringDuration(d string) time.Duration {
	dur, err := time.ParseDuration(d)
	if err != nil {
		app.infoLog.Println("invalid duration. setting the duration to zero")
		dur = 0
	}
	return dur
}
