package display

import (
	"io"
	"log"
)

// Display represents a way for the program to communicate to the user
type Display interface {
	// Inform communicates some information to the user
	Inform(msg string)
	// Alert communicates that the program has entered a fail state
	Alert(msg string)
}

type output struct {
	outs, errs *log.Logger
}

// New instantiates a new Display
func New(outs, errs io.Writer) Display {
	return &output{
		log.New(outs, "", 0),
		log.New(errs, "", 0),
	}
}

// Inform implements Display.Inform
func (o *output) Inform(msg string) {
	o.outs.Print(msg)
}

// Alert implements Display.Alert
func (o *output) Alert(msg string) {
	o.errs.Print(msg)
}
