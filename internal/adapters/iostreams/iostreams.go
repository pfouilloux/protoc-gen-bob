package iostreams

import (
	"fmt"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/generate"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/model"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"io/ioutil"
	"strings"
)

const FEATURES = (uint64)(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
const EXT = ".builder.go"

type Io struct {
	input    io.Reader
	output   io.Writer
	generate func(...generate.Spec) error
}

type Mutagen func(*Io)

func New(input io.Reader, output io.Writer, mutagens ...Mutagen) *Io {
	io := &Io{input: input, output: output, generate: generate.Builders}

	for _, mutate := range mutagens {
		mutate(io)
	}

	return io
}

func (io *Io) Handle() error {
	in, err := ioutil.ReadAll(io.input)
	if err != nil {
		return fmt.Errorf("failed to read request: %v", err)
	}

	var req pluginpb.CodeGeneratorRequest
	if err := proto.Unmarshal(in, &req); err != nil {
		return fmt.Errorf("failed to unmarshal incoming message: %v", err)
	}

	// TODO Move this stuff into another package
	plugin, err := initPlugin(&req)
	if err != nil {
		return fmt.Errorf("failed to initialise protoc plugin: %v", err)
	}

	specs := prepareSpecs(plugin)
	if err := io.generate(specs...); err != nil {
		return fmt.Errorf("failed to generate builders: %v", err)
	}
	// TODO Stop moving stuff

	resp, err := proto.Marshal(plugin.Response())
	if err != nil {
		return fmt.Errorf("failed to marshal response")
	}

	_, err = io.output.Write(resp)
	if err != nil {
		return fmt.Errorf("failed to write response: %v", err)
	}

	return nil
}

func initPlugin(req *pluginpb.CodeGeneratorRequest) (*protogen.Plugin, error) {
	plugin, err := protogen.Options{}.New(req)
	if err != nil {
		return nil, err
	}
	plugin.SupportedFeatures = FEATURES

	return plugin, nil
}

func prepareSpecs(plugin *protogen.Plugin) []generate.Spec {
	tasks := make([]generate.Spec, len(plugin.Files))
	for i, file := range plugin.Files {
		writer := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+EXT, file.GoImportPath)
		tasks[i] = generate.NewSpec(file.GeneratedFilenamePrefix, plan(file), writer)
	}

	return tasks
}

func plan(desc *protogen.File) model.File {
	messages := make([]model.Message, len(desc.Proto.MessageType))
	for m, msg := range desc.Proto.MessageType {
		fields := make([]model.Field, len(msg.GetField()))
		for f, field := range msg.Field {
			fields[f] = model.NewField(strings.Title(field.GetName()), goType(field), field.GetProto3Optional())
		}
		messages[m] = model.NewMessage(strings.Title(msg.GetName()), fields...)
	}

	return model.NewFile((string)(desc.GoPackageName), messages...)
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
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		out = "[]byte"
	case descriptorpb.FieldDescriptorProto_TYPE_GROUP:
		panic("unsupported!")
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		panic("unsupported!")
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		panic("unsupported!")
	}

	switch field.GetLabel() {
	case descriptorpb.FieldDescriptorProto_LABEL_REPEATED:
		return "[]" + out
	}

	return out
}
