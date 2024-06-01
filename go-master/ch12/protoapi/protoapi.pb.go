// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: protoapi.proto

package protoapi

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

// 난수 관련 메세지
type RandomParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seed  int64 `protobuf:"varint,1,opt,name=Seed,proto3" json:"Seed,omitempty"`
	Place int64 `protobuf:"varint,2,opt,name=Place,proto3" json:"Place,omitempty"`
}

func (x *RandomParams) Reset() {
	*x = RandomParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomParams) ProtoMessage() {}

func (x *RandomParams) ProtoReflect() protoreflect.Message {
	mi := &file_protoapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomParams.ProtoReflect.Descriptor instead.
func (*RandomParams) Descriptor() ([]byte, []int) {
	return file_protoapi_proto_rawDescGZIP(), []int{0}
}

func (x *RandomParams) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *RandomParams) GetPlace() int64 {
	if x != nil {
		return x.Place
	}
	return 0
}

type RandomInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *RandomInt) Reset() {
	*x = RandomInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomInt) ProtoMessage() {}

func (x *RandomInt) ProtoReflect() protoreflect.Message {
	mi := &file_protoapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomInt.ProtoReflect.Descriptor instead.
func (*RandomInt) Descriptor() ([]byte, []int) {
	return file_protoapi_proto_rawDescGZIP(), []int{1}
}

func (x *RandomInt) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// 날짜와 시간 관련 메세지
type DateTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *DateTime) Reset() {
	*x = DateTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateTime) ProtoMessage() {}

func (x *DateTime) ProtoReflect() protoreflect.Message {
	mi := &file_protoapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateTime.ProtoReflect.Descriptor instead.
func (*DateTime) Descriptor() ([]byte, []int) {
	return file_protoapi_proto_rawDescGZIP(), []int{2}
}

func (x *DateTime) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type RequestDataTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *RequestDataTime) Reset() {
	*x = RequestDataTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDataTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDataTime) ProtoMessage() {}

func (x *RequestDataTime) ProtoReflect() protoreflect.Message {
	mi := &file_protoapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDataTime.ProtoReflect.Descriptor instead.
func (*RequestDataTime) Descriptor() ([]byte, []int) {
	return file_protoapi_proto_rawDescGZIP(), []int{3}
}

func (x *RequestDataTime) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// 랜덤 패스워드 관련 메시지
type RequestPass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seed   int64 `protobuf:"varint,1,opt,name=Seed,proto3" json:"Seed,omitempty"`
	Length int64 `protobuf:"varint,8,opt,name=Length,proto3" json:"Length,omitempty"`
}

func (x *RequestPass) Reset() {
	*x = RequestPass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPass) ProtoMessage() {}

func (x *RequestPass) ProtoReflect() protoreflect.Message {
	mi := &file_protoapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPass.ProtoReflect.Descriptor instead.
func (*RequestPass) Descriptor() ([]byte, []int) {
	return file_protoapi_proto_rawDescGZIP(), []int{4}
}

func (x *RequestPass) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *RequestPass) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type RandomPass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *RandomPass) Reset() {
	*x = RandomPass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomPass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomPass) ProtoMessage() {}

func (x *RandomPass) ProtoReflect() protoreflect.Message {
	mi := &file_protoapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomPass.ProtoReflect.Descriptor instead.
func (*RandomPass) Descriptor() ([]byte, []int) {
	return file_protoapi_proto_rawDescGZIP(), []int{5}
}

func (x *RandomPass) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_protoapi_proto protoreflect.FileDescriptor

var file_protoapi_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x38, 0x0a, 0x0c, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x53, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x53, 0x65, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x52, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x0a,
	0x08, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x27, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x65, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x65, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x22, 0x28, 0x0a, 0x0a, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x50, 0x61, 0x73,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x84, 0x01,
	0x0a, 0x06, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x26, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x09, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x26, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x0d, 0x2e,
	0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0a, 0x2e, 0x52,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x12, 0x0c, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x1a, 0x0b, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x50, 0x61, 0x73, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoapi_proto_rawDescOnce sync.Once
	file_protoapi_proto_rawDescData = file_protoapi_proto_rawDesc
)

func file_protoapi_proto_rawDescGZIP() []byte {
	file_protoapi_proto_rawDescOnce.Do(func() {
		file_protoapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoapi_proto_rawDescData)
	})
	return file_protoapi_proto_rawDescData
}

var file_protoapi_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protoapi_proto_goTypes = []interface{}{
	(*RandomParams)(nil),    // 0: RandomParams
	(*RandomInt)(nil),       // 1: RandomInt
	(*DateTime)(nil),        // 2: DateTime
	(*RequestDataTime)(nil), // 3: RequestDataTime
	(*RequestPass)(nil),     // 4: RequestPass
	(*RandomPass)(nil),      // 5: RandomPass
}
var file_protoapi_proto_depIdxs = []int32{
	3, // 0: Random.GetData:input_type -> RequestDataTime
	0, // 1: Random.GetRandom:input_type -> RandomParams
	4, // 2: Random.GetRandomPass:input_type -> RequestPass
	2, // 3: Random.GetData:output_type -> DateTime
	1, // 4: Random.GetRandom:output_type -> RandomInt
	5, // 5: Random.GetRandomPass:output_type -> RandomPass
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoapi_proto_init() }
func file_protoapi_proto_init() {
	if File_protoapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomParams); i {
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
		file_protoapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomInt); i {
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
		file_protoapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateTime); i {
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
		file_protoapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDataTime); i {
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
		file_protoapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPass); i {
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
		file_protoapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomPass); i {
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
			RawDescriptor: file_protoapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoapi_proto_goTypes,
		DependencyIndexes: file_protoapi_proto_depIdxs,
		MessageInfos:      file_protoapi_proto_msgTypes,
	}.Build()
	File_protoapi_proto = out.File
	file_protoapi_proto_rawDesc = nil
	file_protoapi_proto_goTypes = nil
	file_protoapi_proto_depIdxs = nil
}
