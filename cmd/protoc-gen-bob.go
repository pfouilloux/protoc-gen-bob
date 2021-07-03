package main

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
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

	for _, file := range plugin.Files {
		var buf bytes.Buffer

		pkg := fmt.Sprintf("package %s", file.GoPackageName)
		buf.Write([]byte(pkg))

		for _, msg := range file.Proto.MessageType {
			var fields []string
			var ctorArgs []string
			for _, field := range msg.Field {
				var ft string

				switch *field.Type {
				case descriptor.FieldDescriptorProto_TYPE_INT32:
					ft = "int32"
				default:
					panic("Not supported!")
				}

				fields = append(fields, fmt.Sprintf("%s %s", *field.Name, ft))
				ctorArgs = append(ctorArgs, fmt.Sprintf("%s: %s", strings.Title(*field.Name), *field.Name))
			}

			buf.WriteString(fmt.Sprintf(`
			// %[1]sBuilder provides a fluent api for building %[3]s.%[1]s
			type %[1]sBuilder struct {
				%[2]s *%[1]s
			}

			// New%[1]s starts building the %[3]s.%[1]s
			func New%[1]s(%[4]s) *%[1]sBuilder {
				return &%[1]sBuilder{&%[1]s{%[5]s}}
			}

			// Build builds the %[3]s.%[1]s
			func (%[6]s *%[1]sBuilder) Build() *%[1]s {
				return %[6]s.%[2]s
			}
			`, *msg.Name, strings.ToLower(*msg.Name), file.GoPackageName, strings.Join(fields, ","), strings.Join(ctorArgs, ","), "x")) // TODO what if x collides with something?
		}

		filename := file.GeneratedFilenamePrefix + ".builder.go"
		file := plugin.NewGeneratedFile(filename, ".")

		file.Write(buf.Bytes())
	}

	stdout := plugin.Response()
	out, err := proto.Marshal(stdout)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, string(out))
}
