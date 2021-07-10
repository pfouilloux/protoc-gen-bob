package dispatch

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/google/go-cmp/cmp"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/model"
	"github.com/pfouilloux/protoc-gen-bob/internal/test"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"reflect"
	"sort"
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
		"should produce a builder given a file with types": {
			input: NewRequest(test.MustMarshal(t, &pluginpb.CodeGeneratorRequest{ProtoFile: []*descriptorpb.FileDescriptorProto{{
				Name: proto.String("all_types.proto"),
				MessageType: []*descriptor.DescriptorProto{
					{Name: proto.String("Msg"), Field: fieldsOfEachType(t)},
				},
				EnumType: []*descriptor.EnumDescriptorProto{{
					Name:  proto.String("MyEnum"),
					Value: []*descriptorpb.EnumValueDescriptorProto{{Name: proto.String("TEST"), Number: proto.Int32(1)}},
				}},
				Options: &descriptorpb.FileOptions{
					GoPackage: proto.String("."),
				},
			}}})),
			expect: &pluginpb.CodeGeneratorResponse{
				SupportedFeatures: proto.Uint64((uint64)(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)),
				File: []*pluginpb.CodeGeneratorResponse_File{{
					Name:    proto.String("all_types.builder.go"),
					Content: proto.String(string(test.MustReadGoldenFile(t, "all_types.builder.go"))),
				}},
			},
		},
		"should handle nested enums": {
			input: NewRequest(test.MustMarshal(t, &pluginpb.CodeGeneratorRequest{ProtoFile: []*descriptorpb.FileDescriptorProto{{
				Name: proto.String("enum.proto"),
				MessageType: []*descriptor.DescriptorProto{{
					Name: proto.String("Msg"),
					EnumType: []*descriptorpb.EnumDescriptorProto{
						{Name: proto.String("MyEnum"), Value: []*descriptorpb.EnumValueDescriptorProto{
							{Name: proto.String("TEST"), Number: proto.Int32(1)},
						}}},
					Field: []*descriptorpb.FieldDescriptorProto{
						{
							Number:   proto.Int32(1),
							Name:     proto.String("val"),
							Type:     protoTypePtr(descriptorpb.FieldDescriptorProto_TYPE_ENUM),
							TypeName: proto.String(".Msg.MyEnum"),
						},
					},
				}},
				Options: &descriptorpb.FileOptions{
					GoPackage: proto.String("."),
				},
			}}})),
			expect: &pluginpb.CodeGeneratorResponse{
				SupportedFeatures: proto.Uint64((uint64)(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)),
				File: []*pluginpb.CodeGeneratorResponse_File{{
					Name:    proto.String("enum.builder.go"),
					Content: proto.String(string(test.MustReadGoldenFile(t, "enum.builder.go"))),
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
		if tc.expect == nil {
			tc.expect = &pluginpb.CodeGeneratorResponse{}
		}

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

			if test.IsUpdateFlagSet() {
				updateFiles(t, &out)
			} else {
				assertOutputMatches(t, tc.expect, &out)
			}
		})
	}
}

func updateFiles(t *testing.T, out *pluginpb.CodeGeneratorResponse) {
	t.Helper()

	for _, file := range out.File {
		test.UpdateTestData(t, *file.Name, bytes.NewBufferString(*file.Content).Bytes())
	}
}

func assertOutputMatches(t *testing.T, expect, actual *pluginpb.CodeGeneratorResponse) {
	if cmp.Equal(expect, actual, cmp.Comparer(cmpResp)) {
		return
	}

	if expect != nil && actual.File != nil && !reflect.DeepEqual(expect.File, actual.File) {
		printFileDifferences(t, expect.File, actual.File)
	} else {
		t.Errorf("output mismatch:\nexpected: %+v\n but got: %+v", expect, actual)
	}
}

func printFileDifferences(t *testing.T, expect, actual []*pluginpb.CodeGeneratorResponse_File) {
	if len(expect) != len(actual) {
		t.Errorf("expected %d files but got %d", len(expect), len(actual))
		return
	}

	for i := 0; i < len(expect); i++ {
		if diff := cmp.Diff(actual[i].Content, actual[i].Content); diff != "" {
			t.Errorf("file %d diff: %s", i, diff)
		}
	}
}

func cmpResp(r1, r2 *pluginpb.CodeGeneratorResponse) bool {
	return (r1 == nil && r2 == nil) || (r1 != nil && r2 != nil &&
		reflect.DeepEqual(r1.File, r2.File) &&
		reflect.DeepEqual(r1.Error, r2.Error) &&
		reflect.DeepEqual(r1.SupportedFeatures, r2.SupportedFeatures))
}

func protoTypePtr(t descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto_Type {
	return &t
}

func fieldsOfEachType(t *testing.T) []*descriptorpb.FieldDescriptorProto {
	t.Helper()

	var values []int
	for _, val := range descriptorpb.FieldDescriptorProto_Type_value {
		values = append(values, (int)(val))
	}

	sort.Ints(values)

	var fields []*descriptorpb.FieldDescriptorProto
	for i, val := range values {
		name := fmt.Sprintf("field%d", i)
		switch descriptorpb.FieldDescriptorProto_Type(val) {
		case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
			fields = append(fields, field(t, (int32)(val), name, proto.String("MyEnum"), descriptorpb.FieldDescriptorProto_Type(val)))
		case descriptorpb.FieldDescriptorProto_TYPE_GROUP: // groups are deprecated https://developers.google.com/protocol-buffers/docs/proto#groups
		case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
			fields = append(fields, field(t, (int32)(val), name, proto.String("Msg"), descriptorpb.FieldDescriptorProto_Type(val)))
		default:
			fields = append(fields, field(t, (int32)(val), name, nil, descriptorpb.FieldDescriptorProto_Type(val)))
		}
	}

	return fields
}

func field(t *testing.T, num int32, name string, typeName *string, protoType descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto {
	t.Helper()

	return &descriptorpb.FieldDescriptorProto{
		Number:   proto.Int32(num),
		Name:     proto.String(name),
		TypeName: typeName,
		Type:     protoTypePtr(protoType),
	}
}
