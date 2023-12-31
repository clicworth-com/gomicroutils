// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: proto/city/type/search.proto

package city

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

type SearchType int32

const (
	SearchType_SEARCH_TYPE_UNSPECIFIED SearchType = 0
	SearchType_SEARCH_TYPE_PROJECT     SearchType = 1
	SearchType_SEARCH_TYPE_LOCALITY    SearchType = 2
	SearchType_SEARCH_TYPE_DEVELOPER   SearchType = 3
	SearchType_SEARCH_TYPE_PINCODE     SearchType = 4
	SearchType_SEARCH_TYPE_RERAID      SearchType = 5
)

// Enum value maps for SearchType.
var (
	SearchType_name = map[int32]string{
		0: "SEARCH_TYPE_UNSPECIFIED",
		1: "SEARCH_TYPE_PROJECT",
		2: "SEARCH_TYPE_LOCALITY",
		3: "SEARCH_TYPE_DEVELOPER",
		4: "SEARCH_TYPE_PINCODE",
		5: "SEARCH_TYPE_RERAID",
	}
	SearchType_value = map[string]int32{
		"SEARCH_TYPE_UNSPECIFIED": 0,
		"SEARCH_TYPE_PROJECT":     1,
		"SEARCH_TYPE_LOCALITY":    2,
		"SEARCH_TYPE_DEVELOPER":   3,
		"SEARCH_TYPE_PINCODE":     4,
		"SEARCH_TYPE_RERAID":      5,
	}
)

func (x SearchType) Enum() *SearchType {
	p := new(SearchType)
	*p = x
	return p
}

func (x SearchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_city_type_search_proto_enumTypes[0].Descriptor()
}

func (SearchType) Type() protoreflect.EnumType {
	return &file_proto_city_type_search_proto_enumTypes[0]
}

func (x SearchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchType.Descriptor instead.
func (SearchType) EnumDescriptor() ([]byte, []int) {
	return file_proto_city_type_search_proto_rawDescGZIP(), []int{0}
}

type SearchResultSource int32

const (
	SearchResultSource_SEARCH_RESULT_SOURCE_UNSPECIFIED SearchResultSource = 0
	SearchResultSource_SEARCH_RESULT_SOURCE_CLICBRICS   SearchResultSource = 1
	SearchResultSource_SEARCH_RESULT_SOURCE_GOOGLE      SearchResultSource = 2
)

// Enum value maps for SearchResultSource.
var (
	SearchResultSource_name = map[int32]string{
		0: "SEARCH_RESULT_SOURCE_UNSPECIFIED",
		1: "SEARCH_RESULT_SOURCE_CLICBRICS",
		2: "SEARCH_RESULT_SOURCE_GOOGLE",
	}
	SearchResultSource_value = map[string]int32{
		"SEARCH_RESULT_SOURCE_UNSPECIFIED": 0,
		"SEARCH_RESULT_SOURCE_CLICBRICS":   1,
		"SEARCH_RESULT_SOURCE_GOOGLE":      2,
	}
)

func (x SearchResultSource) Enum() *SearchResultSource {
	p := new(SearchResultSource)
	*p = x
	return p
}

func (x SearchResultSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchResultSource) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_city_type_search_proto_enumTypes[1].Descriptor()
}

func (SearchResultSource) Type() protoreflect.EnumType {
	return &file_proto_city_type_search_proto_enumTypes[1]
}

func (x SearchResultSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchResultSource.Descriptor instead.
func (SearchResultSource) EnumDescriptor() ([]byte, []int) {
	return file_proto_city_type_search_proto_rawDescGZIP(), []int{1}
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input      string     `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	City       string     `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Type       SearchType `protobuf:"varint,3,opt,name=type,proto3,enum=city.SearchType" json:"type,omitempty"`
	MaxResults int32      `protobuf:"varint,4,opt,name=maxResults,proto3" json:"maxResults,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_city_type_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_city_type_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_city_type_search_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *SearchRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *SearchRequest) GetType() SearchType {
	if x != nil {
		return x.Type
	}
	return SearchType_SEARCH_TYPE_UNSPECIFIED
}

func (x *SearchRequest) GetMaxResults() int32 {
	if x != nil {
		return x.MaxResults
	}
	return 0
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    SearchType               `protobuf:"varint,1,opt,name=type,proto3,enum=city.SearchType" json:"type,omitempty"`
	Results []*SearchResponse_Result `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_city_type_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_city_type_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_proto_city_type_search_proto_rawDescGZIP(), []int{1}
}

func (x *SearchResponse) GetType() SearchType {
	if x != nil {
		return x.Type
	}
	return SearchType_SEARCH_TYPE_UNSPECIFIED
}

func (x *SearchResponse) GetResults() []*SearchResponse_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type SearchResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address      string             `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	ResultSource SearchResultSource `protobuf:"varint,4,opt,name=resultSource,proto3,enum=city.SearchResultSource" json:"resultSource,omitempty"`
}

func (x *SearchResponse_Result) Reset() {
	*x = SearchResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_city_type_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse_Result) ProtoMessage() {}

func (x *SearchResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_proto_city_type_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse_Result.ProtoReflect.Descriptor instead.
func (*SearchResponse_Result) Descriptor() ([]byte, []int) {
	return file_proto_city_type_search_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SearchResponse_Result) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SearchResponse_Result) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchResponse_Result) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SearchResponse_Result) GetResultSource() SearchResultSource {
	if x != nil {
		return x.ResultSource
	}
	return SearchResultSource_SEARCH_RESULT_SOURCE_UNSPECIFIED
}

var File_proto_city_type_search_proto protoreflect.FileDescriptor

var file_proto_city_type_search_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x22, 0x7f, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xf4, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x35,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x84, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3c,
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2a, 0xa8, 0x01, 0x0a,
	0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53,
	0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x41, 0x52,
	0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x10,
	0x01, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x53,
	0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c,
	0x4f, 0x50, 0x45, 0x52, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x4e, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x04, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x45, 0x52, 0x41, 0x49, 0x44, 0x10, 0x05, 0x2a, 0x7f, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a,
	0x20, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x43,
	0x42, 0x52, 0x49, 0x43, 0x53, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x45, 0x41, 0x52, 0x43,
	0x48, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f,
	0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_city_type_search_proto_rawDescOnce sync.Once
	file_proto_city_type_search_proto_rawDescData = file_proto_city_type_search_proto_rawDesc
)

func file_proto_city_type_search_proto_rawDescGZIP() []byte {
	file_proto_city_type_search_proto_rawDescOnce.Do(func() {
		file_proto_city_type_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_city_type_search_proto_rawDescData)
	})
	return file_proto_city_type_search_proto_rawDescData
}

var file_proto_city_type_search_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_city_type_search_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_city_type_search_proto_goTypes = []interface{}{
	(SearchType)(0),               // 0: city.SearchType
	(SearchResultSource)(0),       // 1: city.SearchResultSource
	(*SearchRequest)(nil),         // 2: city.SearchRequest
	(*SearchResponse)(nil),        // 3: city.SearchResponse
	(*SearchResponse_Result)(nil), // 4: city.SearchResponse.Result
}
var file_proto_city_type_search_proto_depIdxs = []int32{
	0, // 0: city.SearchRequest.type:type_name -> city.SearchType
	0, // 1: city.SearchResponse.type:type_name -> city.SearchType
	4, // 2: city.SearchResponse.results:type_name -> city.SearchResponse.Result
	1, // 3: city.SearchResponse.Result.resultSource:type_name -> city.SearchResultSource
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_city_type_search_proto_init() }
func file_proto_city_type_search_proto_init() {
	if File_proto_city_type_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_city_type_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_proto_city_type_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_proto_city_type_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse_Result); i {
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
			RawDescriptor: file_proto_city_type_search_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_city_type_search_proto_goTypes,
		DependencyIndexes: file_proto_city_type_search_proto_depIdxs,
		EnumInfos:         file_proto_city_type_search_proto_enumTypes,
		MessageInfos:      file_proto_city_type_search_proto_msgTypes,
	}.Build()
	File_proto_city_type_search_proto = out.File
	file_proto_city_type_search_proto_rawDesc = nil
	file_proto_city_type_search_proto_goTypes = nil
	file_proto_city_type_search_proto_depIdxs = nil
}
