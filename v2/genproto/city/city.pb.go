// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: proto/city/city.proto

package city

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_city_city_proto protoreflect.FileDescriptor

var file_proto_city_city_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x69, 0x74, 0x79, 0x1a, 0x1c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x44, 0x0a, 0x0b, 0x43,
	0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x13, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x69,
	0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_city_city_proto_goTypes = []interface{}{
	(*SearchRequest)(nil),  // 0: city.SearchRequest
	(*SearchResponse)(nil), // 1: city.SearchResponse
}
var file_proto_city_city_proto_depIdxs = []int32{
	0, // 0: city.CityService.Search:input_type -> city.SearchRequest
	1, // 1: city.CityService.Search:output_type -> city.SearchResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_city_city_proto_init() }
func file_proto_city_city_proto_init() {
	if File_proto_city_city_proto != nil {
		return
	}
	file_proto_city_type_search_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_city_city_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_city_city_proto_goTypes,
		DependencyIndexes: file_proto_city_city_proto_depIdxs,
	}.Build()
	File_proto_city_city_proto = out.File
	file_proto_city_city_proto_rawDesc = nil
	file_proto_city_city_proto_goTypes = nil
	file_proto_city_city_proto_depIdxs = nil
}
