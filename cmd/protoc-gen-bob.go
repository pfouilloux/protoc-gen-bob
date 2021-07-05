package main

import (
	"fmt"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/generate"
	"github.com/pfouilloux/protoc-gen-bob/internal/core/model"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	var req pluginpb.CodeGeneratorRequest
	proto.Unmarshal(input, &req)

	opts := protogen.Options{}
	plugin, err := opts.New(&req)
	if err != nil {
		panic(err)
	}
	features := (uint64)(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	plugin.SupportedFeatures = features

	var tasks []generate.Task
	for _, file := range plugin.Files {
		var messages []model.Message
		for _, msg := range file.Proto.MessageType {
			var fields []model.Field
			for _, field := range msg.Field {
				fields = append(fields, model.NewField(strings.Title(field.GetName()), goType(field), field.GetProto3Optional()))
			}

			messages = append(messages, model.NewMessage(strings.Title(msg.GetName()), fields...))
		}

		filename := file.GeneratedFilenamePrefix + ".builder.go"
		tasks = append(tasks, generate.NewTask(
			file.GeneratedFilenamePrefix,
			model.NewFile((string)(file.GoPackageName), messages...),
			plugin.NewGeneratedFile(filename, ".")))
	}

	generate.Builders(tasks...) // TODO Handle error

	stdout := plugin.Response()
	out, err := proto.Marshal(stdout)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, string(out))
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
