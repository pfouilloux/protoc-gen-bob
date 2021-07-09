package iostreams

import (
	"fmt"
	"github.com/pfouilloux/protoc-gen-bob/internal/adapters/dispatch"
	"io"
	"io/ioutil"
)

type Io struct {
	input      io.Reader
	output     io.Writer
	dispatcher dispatch.Dispatcher
}

func New(input io.Reader, output io.Writer) *Io {
	i := &Io{input: input, output: output, dispatcher: dispatch.New()}

	return i
}

func (io *Io) Handle() error {
	in, err := ioutil.ReadAll(io.input)
	if err != nil {
		return fmt.Errorf("failed to read request: %v", err)
	}

	resp, err := io.dispatcher.Dispatch(dispatch.NewRequest(in))
	if err != nil {
		return fmt.Errorf("failed to generate builder(s): %v", err)
	}

	_, err = io.output.Write(resp.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write response: %v", err)
	}

	return nil
}
