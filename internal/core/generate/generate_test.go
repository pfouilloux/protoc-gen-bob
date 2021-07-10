package generate

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/model"
	"github.com/pfouilloux/protoc-gen-bob/internal/test"
	"strings"
	"testing"
)

func TestGenerateBuilder(t *testing.T) {
	t.Parallel()

	tcs := map[string]struct {
		input  model.File
		expect string
	}{
		"should generate a builder with only a package declaration": {
			input:  model.NewFile("testdata"),
			expect: "package_decl.builder.go",
		},
		"should generate a builder with an empty message": {
			input:  model.NewFile("testdata", model.NewMessage("Tea")),
			expect: "empty_message.builder.go",
		},
		"should generate a builder with three fields": {
			input: model.NewFile("testdata", model.NewMessage("Kettle",
				model.NewField("Material", "string", false),
				model.NewField("Colour", "string", true),
				model.NewField("Capacity", "uint32", false),
			)),
			expect: "message_with_fields.builder.go",
		},
		"should skip passing a value for a boolean field": {
			input: model.NewFile("testdata", model.NewMessage("Cup",
				model.NewField("Full", "bool", false),
				model.NewField("Hot", "bool", true),
			)),
			expect: "message_with_booleans.builder.go",
		},
	}

	for name, tc := range tcs {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var buf strings.Builder
			err := Generate(&buf, tc.input)
			test.AssertNoError(t, err)

			if test.IsUpdateFlagSet() {
				test.UpdateTestData(t, tc.expect, bytes.NewBufferString(buf.String()).Bytes())
			} else {
				expect := string(test.MustReadGoldenFile(t, tc.expect))
				if diff := cmp.Diff(expect, buf.String()); diff != "" {
					t.Errorf("mismatch:\n%s", diff)
				}
			}

		})

	}
}

func TestWriterErrorHandling(t *testing.T) {
	t.Parallel()

	err := Generate(test.ExplodingWriter{Err: "boom"}, model.NewFile("oops"))
	test.AssertSameError(t, "failed create a builder:", err)
}
