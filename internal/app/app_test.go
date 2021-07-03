package app

import (
	"github.com/google/go-cmp/cmp"
	"github.com/pfouilloux/protoc-gen-bob/internal/display"
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
			expectInErr:    "missing path\nUsage of BobGen: bob <path> [opts]\n",
			expectExitCode: 1,
		},
	}

	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			var outs strings.Builder
			var errs strings.Builder

			if code := Cli(display.New(&outs, &errs)).Run(tc.args); code != tc.expectExitCode {
				t.Errorf("expect exit code %d but got %d", tc.expectExitCode, code)
			}

			if diff := cmp.Diff(outs.String(), tc.expectInOut); diff != "" {
				t.Errorf("inform mismatch:\n%v", diff)
			}

			if diff := cmp.Diff(errs.String(), tc.expectInErr); diff != "" {
				t.Errorf("inform mismatch:\n%v", diff)
			}
		})
	}
}
