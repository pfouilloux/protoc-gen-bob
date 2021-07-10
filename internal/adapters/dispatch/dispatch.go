package dispatch

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/generate"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/model"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"strings"
)

type Request struct{ bytes []byte }

func NewRequest(bytes []byte) Request { return Request{bytes: bytes} }

type Response struct{ bytes []byte }

func (r *Response) Bytes() []byte { return r.bytes }

const FEATURES = (uint64)(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
const EXT = ".builder.go" // TODO Make this configurable?

type Dispatcher interface {
	Dispatch(request Request) (Response, error)
}

type dispatcher struct {
	generate generate.Generator
}

func New() Dispatcher {
	return &dispatcher{generate: generate.Generate}
}

func (d *dispatcher) Dispatch(input Request) (Response, error) {
	var req pluginpb.CodeGeneratorRequest
	if err := proto.Unmarshal(input.bytes, &req); err != nil {
		return Response{}, fmt.Errorf("failed to unmarshal incoming message: %v", err)
	}

	plugin, err := initPlugin(&req)
	if err != nil {
		return Response{}, fmt.Errorf("failed to initialise protoc plugin: %v", err)
	}

	for i, file := range plugin.Files {
		writer := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+EXT, file.GoImportPath)
		if err := d.generate(writer, plan(file)); err != nil {
			return Response{}, fmt.Errorf("failed to generate builder %d: %v", i, err)
		}
	}

	resp, err := proto.Marshal(plugin.Response())
	if err != nil {
		return Response{}, fmt.Errorf("failed to marshal response")
	}

	return Response{resp}, nil
}

func initPlugin(req *pluginpb.CodeGeneratorRequest) (*protogen.Plugin, error) {
	plugin, err := protogen.Options{}.New(req)
	if err != nil {
		return nil, err
	}
	plugin.SupportedFeatures = FEATURES

	return plugin, nil
}

func plan(desc *protogen.File) model.File {
	var messages []model.Message
	for _, msg := range desc.Proto.MessageType {
		messages = append(messages, planMessage(msg, "")...)
	}

	return model.NewFile((string)(desc.GoPackageName), messages...)
}

func planMessage(desc *descriptor.DescriptorProto, parentName string) []model.Message {
	name := desc.GetName()
	if parentName != "" {
		name = parentName + "_" + name
	}

	var messages []model.Message
	fields := make([]model.Field, len(desc.GetField()))
	for f, field := range desc.Field {
		fields[f] = model.NewField(
			strings.Title(field.GetName()),
			goType(field),
			field.GetProto3Optional(),
			field.GetType() == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE,
		)
	}
	messages = append(messages, model.NewMessage(name, fields...))

	for _, nested := range desc.NestedType {
		messages = append(messages, planMessage(nested, name)...)
	}

	return messages
}

func goType(field *descriptorpb.FieldDescriptorProto) string {
	var out string
	switch field.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		out = "bool"
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		out = "string"
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		out = "float32"
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		out = "float64"
	case descriptorpb.FieldDescriptorProto_TYPE_INT32, descriptorpb.FieldDescriptorProto_TYPE_SINT32, descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
		out = "int32"
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32, descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
		out = "uint32"
	case descriptorpb.FieldDescriptorProto_TYPE_INT64, descriptorpb.FieldDescriptorProto_TYPE_SINT64, descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
		out = "int64"
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64, descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		out = "uint64"
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		out = "[]byte"
	case descriptorpb.FieldDescriptorProto_TYPE_GROUP:
		panic("unsupported!") // TODO error
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		out = formatNestedType(field.GetTypeName())
	}

	switch field.GetLabel() {
	case descriptorpb.FieldDescriptorProto_LABEL_REPEATED:
		return "[]" + out
	}

	return out
}

func formatNestedType(name string) string {
	return strings.ReplaceAll(strings.TrimPrefix(name, "."), ".", "_")
}
