package dispatch

import (
	"errors"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/google/go-cmp/cmp"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/model"
	"github.com/pfouilloux/protoc-gen-bob/internal/test"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"path/filepath"
	"reflect"
	"testing"
)

func TestProtocGen(t *testing.T) {
	t.Parallel()

	tcs := map[string]struct {
		mutagen   func(*dispatcher)
		input     Request
		expect    *pluginpb.CodeGeneratorResponse
		expectErr string
	}{
		"should produce a builder given a simple file": {
			input: NewRequest(test.MustMarshal(t, &pluginpb.CodeGeneratorRequest{ProtoFile: []*descriptorpb.FileDescriptorProto{{
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
			}}})),
			expect: &pluginpb.CodeGeneratorResponse{
				SupportedFeatures: proto.Uint64((uint64)(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)),
				File: []*pluginpb.CodeGeneratorResponse_File{{
					Name:    proto.String("tea.builder.go"),
					Content: proto.String(string(test.MustReadFile(t, filepath.Join("testdata", "simple_message.expect")))),
				}},
			},
		},
		"should fail to generate builders": {
			mutagen: func(d *dispatcher) {
				d.generate = func(_ io.Writer, _ model.File) error { return errors.New("boom") }
			},
			input: NewRequest(test.MustMarshal(t, &pluginpb.CodeGeneratorRequest{ProtoFile: []*descriptorpb.FileDescriptorProto{{
				Name: proto.String("tea.proto"),
				Options: &descriptorpb.FileOptions{
					GoPackage: proto.String("."),
				},
			}}})),
			expectErr: "failed to generate builder 0:",
		},
		"should fail to unmarshal invalid byte array": {
			input:     NewRequest([]byte{42}),
			expectErr: "failed to unmarshal incoming message:",
		},
		"should fail to initialise plugin with incomplete request": {
			input:     NewRequest(test.MustMarshal(t, &pluginpb.CodeGeneratorRequest{FileToGenerate: []string{"hello"}})),
			expectErr: "failed to initialise protoc plugin:",
		},
	}

	for name, tc := range tcs {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			d := New()
			if tc.mutagen != nil {
				tc.mutagen(d.(*dispatcher))
			}

			resp, err := d.Dispatch(tc.input)
			test.AssertSameError(t, tc.expectErr, err)

			var out pluginpb.CodeGeneratorResponse
			test.AssertNoError(t, proto.Unmarshal(resp.bytes, &out))

			if tc.expect != nil && !cmp.Equal(tc.expect, &out, cmp.Comparer(cmpResp)) {
				t.Errorf("output mismatch:\nexpected: %+v\n but got: %+v", tc.expect, &out)
			}
		})
	}
}

func cmpResp(r1, r2 *pluginpb.CodeGeneratorResponse) bool {
	return reflect.DeepEqual(r1.File, r2.File) &&
		reflect.DeepEqual(r1.Error, r2.Error) &&
		reflect.DeepEqual(r1.SupportedFeatures, r2.SupportedFeatures)
}

func protoTypePtr(t descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto_Type {
	return &t
}
