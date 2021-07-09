package test

import (
	"errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"io/ioutil"
	"strings"
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
		t.Errorf("expected '%s' but got nil", expected)
	case actual != nil && !strings.HasPrefix(actual.Error(), expected):
		t.Errorf("error mismatch:\n     expected: %s\nto start with: %s\n", actual.Error(), expected)
	}
}

func MustMarshal(t *testing.T, msg protoreflect.ProtoMessage) []byte {
	t.Helper()

	out, err := proto.Marshal(msg)
	AssertNoError(t, err)

	return out
}

func MustReadFile(t *testing.T, path string) []byte {
	t.Helper()

	file, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read output: %v", err)
	}

	return file
}

type ExplodingReader struct{ Err string }

func (e ExplodingReader) Read(_ []byte) (int, error) {
	return 0, errors.New(e.Err)
}

type ExplodingWriter struct{ Err string }

func (e ExplodingWriter) Write(_ []byte) (int, error) {
	return 0, errors.New(e.Err)
}
