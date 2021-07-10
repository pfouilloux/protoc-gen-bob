package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/pfouilloux/protoc-gen-bob/internal/test"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestProtocGenBobE2E(t *testing.T) {
	t.Parallel()

	mustInstall(t)

	testcases := map[string]struct {
		proto      string
		expect     string
		expectFail bool
	}{
		"should create clan.builder.go file": {
			proto:  "clan.proto",
			expect: "clan.builder.go",
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
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			test.AssertNoError(t, err)

			if tc.expect != "" {
				expectFile := string(test.MustReadGoldenFile(t, tc.expect))
				actualFile := string(test.MustReadFile(t, filepath.Join(tmp, tc.expect)))
				if diff := cmp.Diff(expectFile, actualFile); diff != "" {
					t.Errorf("generated code mismatch:\n%v", diff)
				}
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
