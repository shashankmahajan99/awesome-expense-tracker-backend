// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: api.proto

package AwesomeExpenseTrackerApi

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x0e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5a, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x4b, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x79, 0x32, 0xe5, 0x04, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a,
	0x22, 0x0b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x6b, 0x0a,
	0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x23, 0x2e,
	0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x77, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x2a, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x12, 0x6f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x3a, 0x01, 0x2a, 0x1a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x1e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x35, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x32, 0xdb, 0x04, 0x0a, 0x11, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x74, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x79, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x1a,
	0x10, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x76, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70,
	0x65, 0x6e, 0x73, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6e, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x32, 0x7f, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x12, 0x74, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61,
	0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22,
	0x08, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x32, 0x9e, 0x03, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x7b, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x70,
	0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x87, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x61,
	0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x1a, 0x12, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x87, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x32, 0xec, 0x01, 0x0a, 0x08, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x69, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x75, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x1a, 0x09,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x96, 0x01, 0x0a, 0x12, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1e, 0x2e, 0x2f,
	0x61, 0x70, 0x69, 0x3b, 0x41, 0x77, 0x65, 0x73, 0x6f, 0x6d, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x41, 0x70, 0x69, 0xa2, 0x02, 0x03, 0x41,
	0x58, 0x58, 0xaa, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),                         // 0: google.protobuf.Empty
	(*LoginUserRequest)(nil),                      // 1: apidefinitions.LoginUserRequest
	(*RegisterUserRequest)(nil),                   // 2: apidefinitions.RegisterUserRequest
	(*DeleteUserRequest)(nil),                     // 3: apidefinitions.DeleteUserRequest
	(*UpdateUserRequest)(nil),                     // 4: apidefinitions.UpdateUserRequest
	(*AuthenticateWithGoogleCallbackRequest)(nil), // 5: apidefinitions.AuthenticateWithGoogleCallbackRequest
	(*CreateExpenseRequest)(nil),                  // 6: apidefinitions.CreateExpenseRequest
	(*UpdateExpenseRequest)(nil),                  // 7: apidefinitions.UpdateExpenseRequest
	(*DeleteExpenseRequest)(nil),                  // 8: apidefinitions.DeleteExpenseRequest
	(*ListExpensesRequest)(nil),                   // 9: apidefinitions.ListExpensesRequest
	(*GetExpenseRequest)(nil),                     // 10: apidefinitions.GetExpenseRequest
	(*GenerateReportRequest)(nil),                 // 11: apidefinitions.GenerateReportRequest
	(*GetUserProfileRequest)(nil),                 // 12: apidefinitions.GetUserProfileRequest
	(*UpdateUserProfileRequest)(nil),              // 13: apidefinitions.UpdateUserProfileRequest
	(*CreateUserProfileRequest)(nil),              // 14: apidefinitions.CreateUserProfileRequest
	(*GetSettingsRequest)(nil),                    // 15: apidefinitions.GetSettingsRequest
	(*UpdateSettingsRequest)(nil),                 // 16: apidefinitions.UpdateSettingsRequest
	(*OAuth2Token)(nil),                           // 17: apidefinitions.OAuth2Token
	(*DeleteUserResponse)(nil),                    // 18: apidefinitions.DeleteUserResponse
	(*UpdateUserResponse)(nil),                    // 19: apidefinitions.UpdateUserResponse
	(*CreateExpenseResponse)(nil),                 // 20: apidefinitions.CreateExpenseResponse
	(*UpdateExpenseResponse)(nil),                 // 21: apidefinitions.UpdateExpenseResponse
	(*DeleteExpenseResponse)(nil),                 // 22: apidefinitions.DeleteExpenseResponse
	(*ListExpensesResponse)(nil),                  // 23: apidefinitions.ListExpensesResponse
	(*GetExpenseResponse)(nil),                    // 24: apidefinitions.GetExpenseResponse
	(*GenerateReportResponse)(nil),                // 25: apidefinitions.GenerateReportResponse
	(*GetUserProfileResponse)(nil),                // 26: apidefinitions.GetUserProfileResponse
	(*UpdateUserProfileResponse)(nil),             // 27: apidefinitions.UpdateUserProfileResponse
	(*CreateUserProfileResponse)(nil),             // 28: apidefinitions.CreateUserProfileResponse
	(*GetSettingsResponse)(nil),                   // 29: apidefinitions.GetSettingsResponse
	(*UpdateSettingsResponse)(nil),                // 30: apidefinitions.UpdateSettingsResponse
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: apidefinitions.HealthCheck.Healthy:input_type -> google.protobuf.Empty
	1,  // 1: apidefinitions.UserAuthentication.LoginUser:input_type -> apidefinitions.LoginUserRequest
	2,  // 2: apidefinitions.UserAuthentication.RegisterUser:input_type -> apidefinitions.RegisterUserRequest
	3,  // 3: apidefinitions.UserAuthentication.DeleteUser:input_type -> apidefinitions.DeleteUserRequest
	4,  // 4: apidefinitions.UserAuthentication.UpdateUser:input_type -> apidefinitions.UpdateUserRequest
	5,  // 5: apidefinitions.UserAuthentication.AuthenticateWithGoogleCallback:input_type -> apidefinitions.AuthenticateWithGoogleCallbackRequest
	6,  // 6: apidefinitions.ExpenseManagement.CreateExpense:input_type -> apidefinitions.CreateExpenseRequest
	7,  // 7: apidefinitions.ExpenseManagement.UpdateExpense:input_type -> apidefinitions.UpdateExpenseRequest
	8,  // 8: apidefinitions.ExpenseManagement.DeleteExpense:input_type -> apidefinitions.DeleteExpenseRequest
	9,  // 9: apidefinitions.ExpenseManagement.ListExpenses:input_type -> apidefinitions.ListExpensesRequest
	10, // 10: apidefinitions.ExpenseManagement.GetExpense:input_type -> apidefinitions.GetExpenseRequest
	11, // 11: apidefinitions.Reports.GenerateReport:input_type -> apidefinitions.GenerateReportRequest
	12, // 12: apidefinitions.UserProfile.GetUserProfile:input_type -> apidefinitions.GetUserProfileRequest
	13, // 13: apidefinitions.UserProfile.UpdateUserProfile:input_type -> apidefinitions.UpdateUserProfileRequest
	14, // 14: apidefinitions.UserProfile.CreateUserProfile:input_type -> apidefinitions.CreateUserProfileRequest
	15, // 15: apidefinitions.Settings.GetSettings:input_type -> apidefinitions.GetSettingsRequest
	16, // 16: apidefinitions.Settings.UpdateSettings:input_type -> apidefinitions.UpdateSettingsRequest
	0,  // 17: apidefinitions.HealthCheck.Healthy:output_type -> google.protobuf.Empty
	17, // 18: apidefinitions.UserAuthentication.LoginUser:output_type -> apidefinitions.OAuth2Token
	17, // 19: apidefinitions.UserAuthentication.RegisterUser:output_type -> apidefinitions.OAuth2Token
	18, // 20: apidefinitions.UserAuthentication.DeleteUser:output_type -> apidefinitions.DeleteUserResponse
	19, // 21: apidefinitions.UserAuthentication.UpdateUser:output_type -> apidefinitions.UpdateUserResponse
	17, // 22: apidefinitions.UserAuthentication.AuthenticateWithGoogleCallback:output_type -> apidefinitions.OAuth2Token
	20, // 23: apidefinitions.ExpenseManagement.CreateExpense:output_type -> apidefinitions.CreateExpenseResponse
	21, // 24: apidefinitions.ExpenseManagement.UpdateExpense:output_type -> apidefinitions.UpdateExpenseResponse
	22, // 25: apidefinitions.ExpenseManagement.DeleteExpense:output_type -> apidefinitions.DeleteExpenseResponse
	23, // 26: apidefinitions.ExpenseManagement.ListExpenses:output_type -> apidefinitions.ListExpensesResponse
	24, // 27: apidefinitions.ExpenseManagement.GetExpense:output_type -> apidefinitions.GetExpenseResponse
	25, // 28: apidefinitions.Reports.GenerateReport:output_type -> apidefinitions.GenerateReportResponse
	26, // 29: apidefinitions.UserProfile.GetUserProfile:output_type -> apidefinitions.GetUserProfileResponse
	27, // 30: apidefinitions.UserProfile.UpdateUserProfile:output_type -> apidefinitions.UpdateUserProfileResponse
	28, // 31: apidefinitions.UserProfile.CreateUserProfile:output_type -> apidefinitions.CreateUserProfileResponse
	29, // 32: apidefinitions.Settings.GetSettings:output_type -> apidefinitions.GetSettingsResponse
	30, // 33: apidefinitions.Settings.UpdateSettings:output_type -> apidefinitions.UpdateSettingsResponse
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   6,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
