package generate

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/model"
	"github.com/pfouilloux/protoc-gen-bob/internal/test"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateBuilder(t *testing.T) {
	t.Parallel()

	tcs := map[string]struct {
		input  []model.File
		expect []string
	}{
		"should generate a builder with only a package declaration": {
			input:  []model.File{model.NewFile("testdata")},
			expect: []string{"package_decl.test"},
		},
		"should generate multiple builders": {
			input:  []model.File{model.NewFile("testdata"), model.NewFile("testdata")},
			expect: []string{"package_decl.test", "package_decl.test"},
		},
		"should generate a builder with an empty message": {
			input:  []model.File{model.NewFile("testdata", model.NewMessage("Kettle"))},
			expect: []string{"empty_message.test"},
		},
		"should generate a builder with three fields": {
			input: []model.File{model.NewFile("testdata", model.NewMessage("Kettle",
				model.NewField("Material", "string", false),
				model.NewField("Colour", "string", true),
				model.NewField("Capacity", "uint32", false),
			))},
			expect: []string{"message_with_fields.test"},
		},
	}

	for name, tc := range tcs {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if len(tc.input) != len(tc.expect) {
				t.Fatalf("$d inputs != $d expectations. These must be the same for comparisons to work. " +
					"Please pad out with nils or zero values if need be")
			}

			buffers := make([]strings.Builder, len(tc.input))
			tasks := make([]Task, len(tc.input))
			for i, file := range tc.input {
				buffers[i] = strings.Builder{}
				tasks[i] = NewTask(fmt.Sprint("test", i), file, &buffers[i])
			}

			err := Builders(tasks...)
			test.AssertNoError(t, err)

			for i, xf := range tc.expect {
				expect := string(test.MustReadFile(t, filepath.Join("testdata", xf)))
				actual := buffers[i].String()
				if diff := cmp.Diff(expect, actual); diff != "" {
					t.Errorf("item %d mismatch:\n%s", i, diff)
				}
			}
		})

	}
}

func TestWriterErrorHandling(t *testing.T) {
	t.Parallel()

	task := NewTask("trap", model.NewFile("bla"), ExplodingWriter{})

	err := Builders(task)
	test.AssertSameError(t, "failed create a builder for task 0 'trap': boom", err)
}

type ExplodingWriter struct{}

func (e ExplodingWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("boom")
}
