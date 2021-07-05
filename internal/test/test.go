package test

import (
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"testing"
)

func AssertNoError(t *testing.T, err error) {
	t.Helper()

	AssertSameError(t, "", err)
}

func AssertSameError(t *testing.T, expected string, actual error) {
	t.Helper()

	switch {
	case expected == "" && actual != nil:
		t.Errorf("unexpected error: %v", actual)
	case expected != "" && actual == nil:
		fallthrough
	case actual != nil && expected != actual.Error():
		t.Errorf("error mismatch:\n%s", cmp.Diff(expected, actual.Error()))
	}
}

func MustReadFile(t *testing.T, path string) []byte {
	t.Helper()

	file, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read output: %v", err)
	}

	return file
}
