package test

import (
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestBobGen(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		args        []string
		expectInfo  string
		expectAlert string
		expect      string
		expectFail  bool
	}{
		"should show usage if -h flag is set": {
			args:       []string{"-h"},
			expectInfo: "Usage of BobGen: bob <path> [opts]\n",
		},
	}

	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			tmp := prepareTempDir(t)
			defer func() { rmDir(t, tmp) }()

			args := append([]string{"run", "../cmd"}, tc.args...)
			cmd := exec.Command("go", args...)

			var outs strings.Builder
			cmd.Stdout = &outs

			var errs strings.Builder
			cmd.Stderr = &errs

			err := cmd.Run()
			if !tc.expectFail && err != nil {
				t.Errorf("unexpected error: %v", err)
			} else if tc.expectFail && err == nil {
				t.Errorf("expected to exit with an error but got none")
			}

			if tc.expect != "" {
				expectFile := mustReadFile(t, filepath.Join("testdata", "output", tc.expect))
				actualFile := mustReadFile(t, filepath.Join(tmp, tc.expect))
				cmp.Diff(expectFile, actualFile)
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

func prepareTempDir(t *testing.T) string {
	t.Helper()

	tmpDir, err := ioutil.TempDir(os.TempDir(), "bob_test_*")
	if err != nil {
		t.Fatalf("failed to create tmp dir: %v", err)
	}

	copyDir(t, "testdata", tmpDir)

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

func copyDir(t *testing.T, source, destination string) {
	t.Helper()

	if err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		var relPath = strings.Replace(path, source, "", 1)
		if relPath == "" {
			return nil
		}
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), info.Mode())
		} else {
			var data, err1 = ioutil.ReadFile(filepath.Join(source, relPath))
			if err1 != nil {
				return err1
			}
			return ioutil.WriteFile(filepath.Join(destination, relPath), data, info.Mode())
		}
	}); err != nil {
		t.Fatalf("failed to copy '%s' to '%s'", source, destination)
	}
}
