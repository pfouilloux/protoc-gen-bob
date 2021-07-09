package iostreams

import (
	"bytes"
	"github.com/pfouilloux/protoc-gen-bob/internal/test"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"reflect"
	"testing"
)

func TestProtocRequest(t *testing.T) {
	t.Parallel()
	var tcs = map[string]struct {
		mutagen   func(*Io)
		input     []byte
		expect    []byte
		expectErr string
	}{
		"should produce an empty response from an empty request": {
			input: []byte{},
			expect: test.MustMarshal(t, &pluginpb.CodeGeneratorResponse{
				SupportedFeatures: proto.Uint64((uint64)(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)),
			}),
		},
		"should fail to write response": {
			mutagen:   func(io *Io) { io.output = test.ExplodingWriter{Err: "splat"} },
			expectErr: "failed to write response: splat",
		},
		"should fail to read input": {
			mutagen:   func(io *Io) { io.input = test.ExplodingReader{Err: "boom"} },
			expectErr: "failed to read request:",
		},
		"should fail to generate builders": {
			input:     []byte{42},
			expectErr: "failed to generate builder(s):",
		},
	}
	for name, tc := range tcs {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer
			io := New(bytes.NewReader(tc.input), &buf)
			if tc.mutagen != nil {
				tc.mutagen(io)
			}

			err := io.Handle()
			test.AssertSameError(t, tc.expectErr, err)

			if !reflect.DeepEqual(tc.expect, buf.Bytes()) {
				t.Errorf("mismatch:\nexpected: %v\n but got: %v", tc.expect, buf.Bytes())
			}
		})
	}
}
