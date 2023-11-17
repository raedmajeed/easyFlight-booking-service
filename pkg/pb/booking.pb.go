// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: booking.proto

package __

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

type Type int32

const (
	Type_BUSINESS Type = 0
	Type_ECONOMY  Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "BUSINESS",
		1: "ECONOMY",
	}
	Type_value = map[string]int32{
		"BUSINESS": 0,
		"ECONOMY":  1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_booking_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_booking_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0}
}

type SearchFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        Type   `protobuf:"varint,1,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	Adults      string `protobuf:"bytes,2,opt,name=adults,proto3" json:"adults,omitempty"`
	Children    string `protobuf:"bytes,3,opt,name=children,proto3" json:"children,omitempty"`
	FromAirport string `protobuf:"bytes,4,opt,name=from_airport,json=fromAirport,proto3" json:"from_airport,omitempty"`
	ToAirport   string `protobuf:"bytes,5,opt,name=to_airport,json=toAirport,proto3" json:"to_airport,omitempty"`
	DepartDate  string `protobuf:"bytes,6,opt,name=depart_date,json=departDate,proto3" json:"depart_date,omitempty"`
	ReturnDate  string `protobuf:"bytes,7,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	Page        string `protobuf:"bytes,8,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *SearchFlightRequest) Reset() {
	*x = SearchFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlightRequest) ProtoMessage() {}

func (x *SearchFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFlightRequest.ProtoReflect.Descriptor instead.
func (*SearchFlightRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0}
}

func (x *SearchFlightRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_BUSINESS
}

func (x *SearchFlightRequest) GetAdults() string {
	if x != nil {
		return x.Adults
	}
	return ""
}

func (x *SearchFlightRequest) GetChildren() string {
	if x != nil {
		return x.Children
	}
	return ""
}

func (x *SearchFlightRequest) GetFromAirport() string {
	if x != nil {
		return x.FromAirport
	}
	return ""
}

func (x *SearchFlightRequest) GetToAirport() string {
	if x != nil {
		return x.ToAirport
	}
	return ""
}

func (x *SearchFlightRequest) GetDepartDate() string {
	if x != nil {
		return x.DepartDate
	}
	return ""
}

func (x *SearchFlightRequest) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

func (x *SearchFlightRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

type SearchFlightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SearchFlightResponse) Reset() {
	*x = SearchFlightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFlightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFlightResponse) ProtoMessage() {}

func (x *SearchFlightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFlightResponse.ProtoReflect.Descriptor instead.
func (*SearchFlightResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{1}
}

func (x *SearchFlightResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

var File_booking_proto protoreflect.FileDescriptor

var file_booking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfc, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72,
	0x6f, 0x6d, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f,
	0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x6f, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x30,
	0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2a, 0x21, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x53, 0x49,
	0x4e, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x43, 0x4f, 0x4e, 0x4f, 0x4d,
	0x59, 0x10, 0x01, 0x32, 0x4e, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x43,
	0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booking_proto_rawDescOnce sync.Once
	file_booking_proto_rawDescData = file_booking_proto_rawDesc
)

func file_booking_proto_rawDescGZIP() []byte {
	file_booking_proto_rawDescOnce.Do(func() {
		file_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_proto_rawDescData)
	})
	return file_booking_proto_rawDescData
}

var file_booking_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_booking_proto_goTypes = []interface{}{
	(Type)(0),                    // 0: Type
	(*SearchFlightRequest)(nil),  // 1: SearchFlightRequest
	(*SearchFlightResponse)(nil), // 2: SearchFlightResponse
}
var file_booking_proto_depIdxs = []int32{
	0, // 0: SearchFlightRequest.type:type_name -> Type
	1, // 1: Booking.RegisterSearchFlight:input_type -> SearchFlightRequest
	2, // 2: Booking.RegisterSearchFlight:output_type -> SearchFlightResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_booking_proto_init() }
func file_booking_proto_init() {
	if File_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFlightRequest); i {
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
		file_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFlightResponse); i {
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
			RawDescriptor: file_booking_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_proto_goTypes,
		DependencyIndexes: file_booking_proto_depIdxs,
		EnumInfos:         file_booking_proto_enumTypes,
		MessageInfos:      file_booking_proto_msgTypes,
	}.Build()
	File_booking_proto = out.File
	file_booking_proto_rawDesc = nil
	file_booking_proto_goTypes = nil
	file_booking_proto_depIdxs = nil
}
