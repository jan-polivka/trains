// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: trains.proto

package protos

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

type Train struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Drive string `protobuf:"bytes,2,opt,name=drive,proto3" json:"drive,omitempty"`
}

func (x *Train) Reset() {
	*x = Train{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trains_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Train) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Train) ProtoMessage() {}

func (x *Train) ProtoReflect() protoreflect.Message {
	mi := &file_trains_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Train.ProtoReflect.Descriptor instead.
func (*Train) Descriptor() ([]byte, []int) {
	return file_trains_proto_rawDescGZIP(), []int{0}
}

func (x *Train) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Train) GetDrive() string {
	if x != nil {
		return x.Drive
	}
	return ""
}

type Trains struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Train []*Train `protobuf:"bytes,1,rep,name=train,proto3" json:"train,omitempty"`
}

func (x *Trains) Reset() {
	*x = Trains{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trains_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trains) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trains) ProtoMessage() {}

func (x *Trains) ProtoReflect() protoreflect.Message {
	mi := &file_trains_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trains.ProtoReflect.Descriptor instead.
func (*Trains) Descriptor() ([]byte, []int) {
	return file_trains_proto_rawDescGZIP(), []int{1}
}

func (x *Trains) GetTrain() []*Train {
	if x != nil {
		return x.Train
	}
	return nil
}

type DispatchAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response int64 `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DispatchAck) Reset() {
	*x = DispatchAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trains_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchAck) ProtoMessage() {}

func (x *DispatchAck) ProtoReflect() protoreflect.Message {
	mi := &file_trains_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchAck.ProtoReflect.Descriptor instead.
func (*DispatchAck) Descriptor() ([]byte, []int) {
	return file_trains_proto_rawDescGZIP(), []int{2}
}

func (x *DispatchAck) GetResponse() int64 {
	if x != nil {
		return x.Response
	}
	return 0
}

var File_trains_proto protoreflect.FileDescriptor

var file_trains_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x31, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x72, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x64, 0x72, 0x69, 0x76, 0x65, 0x22, 0x2d, 0x0a, 0x06, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x52, 0x05, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x22, 0x29, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x48, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x0e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x73, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x1a, 0x13, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x73, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x6b, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6e, 0x2d,
	0x70, 0x6f, 0x6c, 0x69, 0x76, 0x6b, 0x61, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trains_proto_rawDescOnce sync.Once
	file_trains_proto_rawDescData = file_trains_proto_rawDesc
)

func file_trains_proto_rawDescGZIP() []byte {
	file_trains_proto_rawDescOnce.Do(func() {
		file_trains_proto_rawDescData = protoimpl.X.CompressGZIP(file_trains_proto_rawDescData)
	})
	return file_trains_proto_rawDescData
}

var file_trains_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_trains_proto_goTypes = []interface{}{
	(*Train)(nil),       // 0: trains.Train
	(*Trains)(nil),      // 1: trains.Trains
	(*DispatchAck)(nil), // 2: trains.DispatchAck
}
var file_trains_proto_depIdxs = []int32{
	0, // 0: trains.Trains.train:type_name -> trains.Train
	1, // 1: trains.DispatchService.DispatchTrains:input_type -> trains.Trains
	2, // 2: trains.DispatchService.DispatchTrains:output_type -> trains.DispatchAck
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trains_proto_init() }
func file_trains_proto_init() {
	if File_trains_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trains_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Train); i {
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
		file_trains_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trains); i {
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
		file_trains_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchAck); i {
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
			RawDescriptor: file_trains_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trains_proto_goTypes,
		DependencyIndexes: file_trains_proto_depIdxs,
		MessageInfos:      file_trains_proto_msgTypes,
	}.Build()
	File_trains_proto = out.File
	file_trains_proto_rawDesc = nil
	file_trains_proto_goTypes = nil
	file_trains_proto_depIdxs = nil
}
