package iostreams

import (
	"bytes"
	"errors"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/google/go-cmp/cmp"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/generate"
	"github.com/pfouilloux/protoc-gen-bob/internal/test"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"path/filepath"
	"reflect"
	"testing"
)

const ExpectedFeatures = (uint64)(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

func TestProtocRequest(t *testing.T) {
	t.Parallel()
	var tcs = map[string]struct {
		mutagen   Mutagen
		input     []byte
		expect    *pluginpb.CodeGeneratorResponse
		expectErr string
	}{
		"should produce an empty response from an empty request": {
			input:  mustMarshall(t, &pluginpb.CodeGeneratorRequest{}),
			expect: &pluginpb.CodeGeneratorResponse{SupportedFeatures: proto.Uint64(ExpectedFeatures)},
		},
		"should produce a builder given a simple file": {
			input: mustMarshall(t, &pluginpb.CodeGeneratorRequest{ProtoFile: []*descriptorpb.FileDescriptorProto{{
				Name: proto.String("tea.proto"),
				MessageType: []*descriptor.DescriptorProto{
					{Name: proto.String("kettle"), Field: []*descriptorpb.FieldDescriptorProto{{
						Number: proto.Int32(1),
						Name:   proto.String("temperature"),
						Type:   protoTypePtr(descriptorpb.FieldDescriptorProto_TYPE_INT32),
					}}},
				},
				Options: &descriptorpb.FileOptions{
					GoPackage: proto.String("."),
				},
			}}}),
			expect: &pluginpb.CodeGeneratorResponse{
				SupportedFeatures: proto.Uint64(ExpectedFeatures),
				File: []*pluginpb.CodeGeneratorResponse_File{{
					Name:    proto.String("tea.builder.go"),
					Content: proto.String(string(test.MustReadFile(t, filepath.Join("testdata", "simple_message.expect")))),
				}},
			},
		},
		"should fail to write response": {
			mutagen: func(io *Io) { io.output = test.ExplodingWriter{Err: "splat"} },
			input: mustMarshall(t, &pluginpb.CodeGeneratorRequest{ProtoFile: []*descriptorpb.FileDescriptorProto{{
				Name: proto.String("tea.proto"),
				Options: &descriptorpb.FileOptions{
					GoPackage: proto.String("."),
				},
			}}}),
			expectErr: "failed to write response: splat",
		},
		"should fail to read input": {
			mutagen: func(io *Io) { io.input = test.ExplodingReader{Err: "boom"} },
			input: mustMarshall(t, &pluginpb.CodeGeneratorRequest{ProtoFile: []*descriptorpb.FileDescriptorProto{{
				Name: proto.String("tea.proto"),
				Options: &descriptorpb.FileOptions{
					GoPackage: proto.String("."),
				},
			}}}),
			expectErr: "failed to read request:",
		},
		"should fail to unmarshal invalid byte array": {
			input:     []byte{42},
			expectErr: "failed to unmarshal incoming message:",
		},
		"should fail to initialise plugin with incomplete request": {
			input:     mustMarshall(t, &pluginpb.CodeGeneratorRequest{FileToGenerate: []string{"hello"}}),
			expectErr: "failed to initialise protoc plugin:",
		},
		"should fail to generate builders": {
			mutagen:   func(io *Io) { io.generate = func(spec ...generate.Spec) error { return errors.New("fail") } },
			input:     mustMarshall(t, &pluginpb.CodeGeneratorRequest{}),
			expectErr: "failed to generate builders:",
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

			var out pluginpb.CodeGeneratorResponse
			test.AssertNoError(t, proto.Unmarshal(buf.Bytes(), &out))

			if tc.expect != nil && !cmp.Equal(tc.expect, &out, cmp.Comparer(cmpResp)) {
				t.Errorf("output mismatch:\nexpected: %+v\n but got: %+v", tc.expect, &out)
			}
		})
	}
}

func protoTypePtr(t descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto_Type {
	return &t
}

func cmpResp(r1, r2 *pluginpb.CodeGeneratorResponse) bool {
	return reflect.DeepEqual(r1.File, r2.File) &&
		reflect.DeepEqual(r1.Error, r2.Error) &&
		reflect.DeepEqual(r1.SupportedFeatures, r2.SupportedFeatures)
}

func mustMarshall(t *testing.T, req *pluginpb.CodeGeneratorRequest) []byte {
	bytes, err := proto.Marshal(req)
	test.AssertNoError(t, err)

	return bytes
}
