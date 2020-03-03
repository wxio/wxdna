// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.0
// 	protoc        v3.11.4
// source: sequence.proto

package sequence

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Sequence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Participant []string        `protobuf:"bytes,1,rep,name=participant,proto3" json:"participant,omitempty"`
	SimpleCall  []*SimpleReqRsp `protobuf:"bytes,2,rep,name=simpleCall,proto3" json:"simpleCall,omitempty"`
}

func (x *Sequence) Reset() {
	*x = Sequence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sequence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sequence) ProtoMessage() {}

func (x *Sequence) ProtoReflect() protoreflect.Message {
	mi := &file_sequence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sequence.ProtoReflect.Descriptor instead.
func (*Sequence) Descriptor() ([]byte, []int) {
	return file_sequence_proto_rawDescGZIP(), []int{0}
}

func (x *Sequence) GetParticipant() []string {
	if x != nil {
		return x.Participant
	}
	return nil
}

func (x *Sequence) GetSimpleCall() []*SimpleReqRsp {
	if x != nil {
		return x.SimpleCall
	}
	return nil
}

type SimpleReqRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *SimpleReqRsp) Reset() {
	*x = SimpleReqRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleReqRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleReqRsp) ProtoMessage() {}

func (x *SimpleReqRsp) ProtoReflect() protoreflect.Message {
	mi := &file_sequence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleReqRsp.ProtoReflect.Descriptor instead.
func (*SimpleReqRsp) Descriptor() ([]byte, []int) {
	return file_sequence_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleReqRsp) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SimpleReqRsp) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SimpleReqRsp) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

var file_sequence_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*Sequence)(nil),
		Field:         50014,
		Name:          "wxio.dna.extensions.sequence.sequence",
		Tag:           "bytes,50014,opt,name=sequence",
		Filename:      "sequence.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional wxio.dna.extensions.sequence.Sequence sequence = 50014;
	E_Sequence = &file_sequence_proto_extTypes[0]
)

var File_sequence_proto protoreflect.FileDescriptor

var file_sequence_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1c, 0x77, 0x78, 0x69, 0x6f, 0x2e, 0x64, 0x6e, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x78, 0x0a, 0x08, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x4a,
	0x0a, 0x0a, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x78, 0x69, 0x6f, 0x2e, 0x64, 0x6e, 0x61, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x52, 0x73, 0x70, 0x52, 0x0a,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x22, 0x46, 0x0a, 0x0c, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x3a, 0x63, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xde, 0x86,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x77, 0x78, 0x69, 0x6f, 0x2e, 0x64, 0x6e, 0x61,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x78, 0x69, 0x6f, 0x2f, 0x67, 0x6f, 0x64, 0x6e, 0x61,
	0x2f, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sequence_proto_rawDescOnce sync.Once
	file_sequence_proto_rawDescData = file_sequence_proto_rawDesc
)

func file_sequence_proto_rawDescGZIP() []byte {
	file_sequence_proto_rawDescOnce.Do(func() {
		file_sequence_proto_rawDescData = protoimpl.X.CompressGZIP(file_sequence_proto_rawDescData)
	})
	return file_sequence_proto_rawDescData
}

var file_sequence_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sequence_proto_goTypes = []interface{}{
	(*Sequence)(nil),                // 0: wxio.dna.extensions.sequence.Sequence
	(*SimpleReqRsp)(nil),            // 1: wxio.dna.extensions.sequence.SimpleReqRsp
	(*descriptor.FieldOptions)(nil), // 2: google.protobuf.FieldOptions
}
var file_sequence_proto_depIdxs = []int32{
	1, // 0: wxio.dna.extensions.sequence.Sequence.simpleCall:type_name -> wxio.dna.extensions.sequence.SimpleReqRsp
	2, // 1: wxio.dna.extensions.sequence.sequence:extendee -> google.protobuf.FieldOptions
	0, // 2: wxio.dna.extensions.sequence.sequence:type_name -> wxio.dna.extensions.sequence.Sequence
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sequence_proto_init() }
func file_sequence_proto_init() {
	if File_sequence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sequence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sequence); i {
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
		file_sequence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleReqRsp); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sequence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_sequence_proto_goTypes,
		DependencyIndexes: file_sequence_proto_depIdxs,
		MessageInfos:      file_sequence_proto_msgTypes,
		ExtensionInfos:    file_sequence_proto_extTypes,
	}.Build()
	File_sequence_proto = out.File
	file_sequence_proto_rawDesc = nil
	file_sequence_proto_goTypes = nil
	file_sequence_proto_depIdxs = nil
}