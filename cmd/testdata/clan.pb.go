// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: cmd/testdata/clan.proto

package testdata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Warg_Breed int32

const (
	Warg_BREED_UNSPECIFIED Warg_Breed = 0
	Warg_BREED_PLAINS      Warg_Breed = 1
	Warg_BREED_MOUNTAIN    Warg_Breed = 2
)

// Enum value maps for Warg_Breed.
var (
	Warg_Breed_name = map[int32]string{
		0: "BREED_UNSPECIFIED",
		1: "BREED_PLAINS",
		2: "BREED_MOUNTAIN",
	}
	Warg_Breed_value = map[string]int32{
		"BREED_UNSPECIFIED": 0,
		"BREED_PLAINS":      1,
		"BREED_MOUNTAIN":    2,
	}
)

func (x Warg_Breed) Enum() *Warg_Breed {
	p := new(Warg_Breed)
	*p = x
	return p
}

func (x Warg_Breed) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Warg_Breed) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_testdata_clan_proto_enumTypes[0].Descriptor()
}

func (Warg_Breed) Type() protoreflect.EnumType {
	return &file_cmd_testdata_clan_proto_enumTypes[0]
}

func (x Warg_Breed) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Warg_Breed.Descriptor instead.
func (Warg_Breed) EnumDescriptor() ([]byte, []int) {
	return file_cmd_testdata_clan_proto_rawDescGZIP(), []int{1, 0}
}

//message Clan {
//  OrcType type = 1;
//  string name = 2;
//  optional float gold = 3;
//  repeated Member orc = 4;
//}
//
//enum OrcType {
//  MORDOR = 0;
//  URUK_HAI = 1;
//  ISENGARD = 2;
//}
//
//message Member {
//  Orc details = 1;
//  google.protobuf.Any stats = 3; // expecting this to be interface{}
//  oneof steed {
//    Warg warg = 10;
//    Caragor caragor = 11;
//  };
//}
//
type Orc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age   uint32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Title *string `protobuf:"bytes,3,opt,name=title,proto3,oneof" json:"title,omitempty"`
}

func (x *Orc) Reset() {
	*x = Orc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_testdata_clan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orc) ProtoMessage() {}

func (x *Orc) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_testdata_clan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orc.ProtoReflect.Descriptor instead.
func (*Orc) Descriptor() ([]byte, []int) {
	return file_cmd_testdata_clan_proto_rawDescGZIP(), []int{0}
}

func (x *Orc) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Orc) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Orc) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

type Warg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Breed   Warg_Breed `protobuf:"varint,2,opt,name=breed,proto3,enum=Warg_Breed" json:"breed,omitempty"`
	Saddled bool       `protobuf:"varint,3,opt,name=saddled,proto3" json:"saddled,omitempty"`
}

func (x *Warg) Reset() {
	*x = Warg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_testdata_clan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warg) ProtoMessage() {}

func (x *Warg) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_testdata_clan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Warg.ProtoReflect.Descriptor instead.
func (*Warg) Descriptor() ([]byte, []int) {
	return file_cmd_testdata_clan_proto_rawDescGZIP(), []int{1}
}

func (x *Warg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Warg) GetBreed() Warg_Breed {
	if x != nil {
		return x.Breed
	}
	return Warg_BREED_UNSPECIFIED
}

func (x *Warg) GetSaddled() bool {
	if x != nil {
		return x.Saddled
	}
	return false
}

type Caragor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Equipment *Caragor_Equipment `protobuf:"bytes,2,opt,name=equipment,proto3" json:"equipment,omitempty"`
	Armoured  bool               `protobuf:"varint,3,opt,name=armoured,proto3" json:"armoured,omitempty"`
}

func (x *Caragor) Reset() {
	*x = Caragor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_testdata_clan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Caragor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Caragor) ProtoMessage() {}

func (x *Caragor) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_testdata_clan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Caragor.ProtoReflect.Descriptor instead.
func (*Caragor) Descriptor() ([]byte, []int) {
	return file_cmd_testdata_clan_proto_rawDescGZIP(), []int{2}
}

func (x *Caragor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Caragor) GetEquipment() *Caragor_Equipment {
	if x != nil {
		return x.Equipment
	}
	return nil
}

func (x *Caragor) GetArmoured() bool {
	if x != nil {
		return x.Armoured
	}
	return false
}

type Caragor_Equipment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spears     int32 `protobuf:"varint,1,opt,name=spears,proto3" json:"spears,omitempty"`
	Bows       int32 `protobuf:"varint,2,opt,name=bows,proto3" json:"bows,omitempty"`
	Provisions int32 `protobuf:"varint,3,opt,name=provisions,proto3" json:"provisions,omitempty"`
}

func (x *Caragor_Equipment) Reset() {
	*x = Caragor_Equipment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_testdata_clan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Caragor_Equipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Caragor_Equipment) ProtoMessage() {}

func (x *Caragor_Equipment) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_testdata_clan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Caragor_Equipment.ProtoReflect.Descriptor instead.
func (*Caragor_Equipment) Descriptor() ([]byte, []int) {
	return file_cmd_testdata_clan_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Caragor_Equipment) GetSpears() int32 {
	if x != nil {
		return x.Spears
	}
	return 0
}

func (x *Caragor_Equipment) GetBows() int32 {
	if x != nil {
		return x.Bows
	}
	return 0
}

func (x *Caragor_Equipment) GetProvisions() int32 {
	if x != nil {
		return x.Provisions
	}
	return 0
}

var File_cmd_testdata_clan_proto protoreflect.FileDescriptor

var file_cmd_testdata_clan_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6d, 0x64, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63,
	0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x03, 0x4f, 0x72, 0x63,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x04,
	0x57, 0x61, 0x72, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x62, 0x72, 0x65, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x57, 0x61, 0x72, 0x67, 0x2e, 0x42,
	0x72, 0x65, 0x65, 0x64, 0x52, 0x05, 0x62, 0x72, 0x65, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x61, 0x64, 0x64, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x61,
	0x64, 0x64, 0x6c, 0x65, 0x64, 0x22, 0x44, 0x0a, 0x05, 0x42, 0x72, 0x65, 0x65, 0x64, 0x12, 0x15,
	0x0a, 0x11, 0x42, 0x52, 0x45, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x52, 0x45, 0x45, 0x44, 0x5f, 0x50,
	0x4c, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x52, 0x45, 0x45, 0x44,
	0x5f, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x10, 0x02, 0x22, 0xc4, 0x01, 0x0a, 0x07,
	0x43, 0x61, 0x72, 0x61, 0x67, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x43, 0x61, 0x72, 0x61, 0x67, 0x6f, 0x72, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x09, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x72, 0x6d, 0x6f, 0x75, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x61, 0x72, 0x6d, 0x6f, 0x75, 0x72, 0x65, 0x64, 0x1a, 0x57, 0x0a, 0x09, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x65, 0x61, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x70, 0x65, 0x61, 0x72, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x6f,
	0x77, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_testdata_clan_proto_rawDescOnce sync.Once
	file_cmd_testdata_clan_proto_rawDescData = file_cmd_testdata_clan_proto_rawDesc
)

func file_cmd_testdata_clan_proto_rawDescGZIP() []byte {
	file_cmd_testdata_clan_proto_rawDescOnce.Do(func() {
		file_cmd_testdata_clan_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_testdata_clan_proto_rawDescData)
	})
	return file_cmd_testdata_clan_proto_rawDescData
}

var file_cmd_testdata_clan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_testdata_clan_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cmd_testdata_clan_proto_goTypes = []interface{}{
	(Warg_Breed)(0),           // 0: Warg.Breed
	(*Orc)(nil),               // 1: Orc
	(*Warg)(nil),              // 2: Warg
	(*Caragor)(nil),           // 3: Caragor
	(*Caragor_Equipment)(nil), // 4: Caragor.Equipment
}
var file_cmd_testdata_clan_proto_depIdxs = []int32{
	0, // 0: Warg.breed:type_name -> Warg.Breed
	4, // 1: Caragor.equipment:type_name -> Caragor.Equipment
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cmd_testdata_clan_proto_init() }
func file_cmd_testdata_clan_proto_init() {
	if File_cmd_testdata_clan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_testdata_clan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orc); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_testdata_clan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Warg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_testdata_clan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Caragor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_testdata_clan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Caragor_Equipment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_cmd_testdata_clan_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_testdata_clan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_testdata_clan_proto_goTypes,
		DependencyIndexes: file_cmd_testdata_clan_proto_depIdxs,
		EnumInfos:         file_cmd_testdata_clan_proto_enumTypes,
		MessageInfos:      file_cmd_testdata_clan_proto_msgTypes,
	}.Build()
	File_cmd_testdata_clan_proto = out.File
	file_cmd_testdata_clan_proto_rawDesc = nil
	file_cmd_testdata_clan_proto_goTypes = nil
	file_cmd_testdata_clan_proto_depIdxs = nil
}
