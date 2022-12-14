// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: command.proto

package api

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

type Command int32

const (
	Command_EMPTY     Command = 0
	Command_TERMINATE Command = 1
	Command_ADD_NODE  Command = 2
	Command_BARRIER   Command = 3
	Command_ACK       Command = 4
	Command_HEARTBEAT Command = 5
)

// Enum value maps for Command.
var (
	Command_name = map[int32]string{
		0: "EMPTY",
		1: "TERMINATE",
		2: "ADD_NODE",
		3: "BARRIER",
		4: "ACK",
		5: "HEARTBEAT",
	}
	Command_value = map[string]int32{
		"EMPTY":     0,
		"TERMINATE": 1,
		"ADD_NODE":  2,
		"BARRIER":   3,
		"ACK":       4,
		"HEARTBEAT": 5,
	}
)

func (x Command) Enum() *Command {
	p := new(Command)
	*p = x
	return p
}

func (x Command) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command) Descriptor() protoreflect.EnumDescriptor {
	return file_command_proto_enumTypes[0].Descriptor()
}

func (Command) Type() protoreflect.EnumType {
	return &file_command_proto_enumTypes[0]
}

func (x Command) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command.Descriptor instead.
func (Command) EnumDescriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

var File_command_proto protoreflect.FileDescriptor

var file_command_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0x56, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d,
	0x50, 0x54, 0x59, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41,
	0x54, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x44, 0x44, 0x5f, 0x4e, 0x4f, 0x44, 0x45,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x41, 0x52, 0x52, 0x49, 0x45, 0x52, 0x10, 0x03, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x43, 0x4b, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x45, 0x41, 0x52,
	0x54, 0x42, 0x45, 0x41, 0x54, 0x10, 0x05, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x63, 0x73, 0x79, 0x73, 0x2d, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2f, 0x70, 0x73, 0x2d, 0x6c, 0x69, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_proto_rawDescOnce sync.Once
	file_command_proto_rawDescData = file_command_proto_rawDesc
)

func file_command_proto_rawDescGZIP() []byte {
	file_command_proto_rawDescOnce.Do(func() {
		file_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_proto_rawDescData)
	})
	return file_command_proto_rawDescData
}

var file_command_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_command_proto_goTypes = []interface{}{
	(Command)(0), // 0: Command
}
var file_command_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_command_proto_init() }
func file_command_proto_init() {
	if File_command_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_command_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_command_proto_goTypes,
		DependencyIndexes: file_command_proto_depIdxs,
		EnumInfos:         file_command_proto_enumTypes,
	}.Build()
	File_command_proto = out.File
	file_command_proto_rawDesc = nil
	file_command_proto_goTypes = nil
	file_command_proto_depIdxs = nil
}
