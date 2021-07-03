package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestProtocGenBobE2E(t *testing.T) {
	t.Parallel()

	mustInstall(t)

	testcases := map[string]struct {
		proto       string
		expectInfo  string
		expectAlert string
		expect      string
		expectFail  bool
	}{
		"should create orcs.builder.go file": {
			proto:  "orcs.proto",
			expect: "orcs.builder.go",
		},
	}

	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			tmp := prepareTempDir(t)
			defer func() { rmDir(t, tmp) }()

			bobDir := fmt.Sprintf("--bob_out=%s", tmp)
			goDir := fmt.Sprintf("--go_out=%s", tmp)
			input := filepath.Join("testdata", tc.proto)

			cmd := exec.Command("protoc", bobDir, goDir, input)

			var outs strings.Builder
			cmd.Stdout = &outs

			var errs strings.Builder
			cmd.Stderr = &errs

			err := cmd.Run()
			if !tc.expectFail && err != nil {
				t.Errorf("unexpected error: %v", err)
				t.Errorf("stdOut: %s", outs.String())
				t.Errorf("stdErr: %s", errs.String())
			} else if tc.expectFail && err == nil {
				t.Errorf("expected to exit with an error but got none")
			}

			if tc.expect != "" {
				expectFile := mustReadFile(t, filepath.Join("testdata", tc.expect))
				actualFile := mustReadFile(t, filepath.Join(tmp, tc.expect))
				if diff := cmp.Diff(expectFile, actualFile); diff != "" {
					t.Errorf("generated code mismatch:\n%v", diff)
				}
			}

			if diff := cmp.Diff(outs.String(), tc.expectInfo); diff != "" {
				t.Errorf("inform mismatch:\n%v", diff)
			}

			if diff := cmp.Diff(errs.String(), tc.expectAlert); diff != "" {
				t.Errorf("alert mismatch:\n%v", diff)
			}
		})
	}
}

func mustInstall(t *testing.T) {
	t.Helper()

	cmd := exec.Command("go", "install", "./protoc-gen-bob.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to install protoc-gen-bob: %v", err)
	}

}

func prepareTempDir(t *testing.T) string {
	t.Helper()

	tmpDir, err := ioutil.TempDir(os.TempDir(), "bob_test_*")
	if err != nil {
		t.Fatalf("failed to create tmp dir: %v", err)
	}

	return tmpDir
}

func rmDir(t *testing.T, dir string) {
	t.Helper()

	if err := os.RemoveAll(dir); err != nil {
		println("failed to clean up dir ", dir, err)
	}
}

func mustReadFile(t *testing.T, path string) []byte {
	t.Helper()

	file, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read output: %v", err)
	}

	return file
}
