package output

import (
	"io"
	"log"
)

// Output represents a way for the program to communicate to the user
type Output interface {
	// Inform communicates some information to the user
	Inform(msg string)
	// Alert communicates that the program has entered a fail state
	Alert(msg string)
}

type output struct {
	outs, errs *log.Logger
}

func New(outs, errs io.Writer) Output {
	return &output{
		log.New(outs, "", 0),
		log.New(errs, "", 0),
	}
}

func (o *output) Inform(msg string) {
	o.outs.Print(msg)
}

func (o *output) Alert(msg string) {
	o.errs.Print(msg)
}
