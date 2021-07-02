package app

import (
	"BobGen/internal/output"
	"BobGen/internal/test"
	"strings"
	"testing"
)

func TestApp(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		args           []string
		expectInOut    string
		expectInErr    string
		expectExitCode int
	}{
		"should show usage if no arguments are passed": {
			args:           []string{},
			expectInErr:    "missing path\nUsage of BobGen: bob <path> [opts]",
			expectExitCode: 1,
		},
	}

	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			var outs strings.Builder
			var errs strings.Builder

			if code := Cli(output.New(&outs, &errs)).Run(tc.args); code != tc.expectExitCode {
				t.Errorf("expect exit code %d but got %d", tc.expectExitCode, code)
			}

			test.AssertThatInformed(t, outs, tc.expectInOut)
			test.AssertThatAlerted(t, errs, tc.expectInErr)
		})
	}
}
