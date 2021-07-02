package output

import (
	"BobGen/internal/test"
	"strings"
	"testing"
)

func TestOutput(t *testing.T) {
	t.Parallel()

	var outs strings.Builder
	var errs strings.Builder

	o := New(&outs, &errs)
	o.Inform("Hello world!")
	o.Alert("Oh no!")

	test.AssertThatInformed(t, outs, "Hello world!")
	test.AssertThatAlerted(t, errs, "Oh no!")
}
