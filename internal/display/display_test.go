package display

import (
	"github.com/google/go-cmp/cmp"
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

	if diff := cmp.Diff(outs.String(), "Hello world!\n"); diff != "" {
		t.Errorf("inform mismatch:\n%v", diff)
	}

	if diff := cmp.Diff(errs.String(), "Oh no!\n"); diff != "" {
		t.Errorf("alert mismatch:\n%v", diff)
	}
}
