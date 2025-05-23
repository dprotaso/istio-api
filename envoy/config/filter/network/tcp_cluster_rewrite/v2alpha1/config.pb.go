// Copyright 2018 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: envoy/config/filter/network/tcp_cluster_rewrite/v2alpha1/config.proto

// $title: TCP cluster rewrite filter configuration for Envoy.

package v2alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TcpClusterRewrite is the config for the TCP cluster rewrite filter.
type TcpClusterRewrite struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Specifies the regex pattern to be matched in the cluster name.
	ClusterPattern string `protobuf:"bytes,1,opt,name=cluster_pattern,json=clusterPattern,proto3" json:"cluster_pattern,omitempty"`
	// Specifies the replacement for the matched cluster pattern.
	ClusterReplacement string `protobuf:"bytes,2,opt,name=cluster_replacement,json=clusterReplacement,proto3" json:"cluster_replacement,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *TcpClusterRewrite) Reset() {
	*x = TcpClusterRewrite{}
	mi := &file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TcpClusterRewrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TcpClusterRewrite) ProtoMessage() {}

func (x *TcpClusterRewrite) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TcpClusterRewrite.ProtoReflect.Descriptor instead.
func (*TcpClusterRewrite) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_rawDescGZIP(), []int{0}
}

func (x *TcpClusterRewrite) GetClusterPattern() string {
	if x != nil {
		return x.ClusterPattern
	}
	return ""
}

func (x *TcpClusterRewrite) GetClusterReplacement() string {
	if x != nil {
		return x.ClusterReplacement
	}
	return ""
}

var File_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto protoreflect.FileDescriptor

const file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_rawDesc = "" +
	"\n" +
	"Eenvoy/config/filter/network/tcp_cluster_rewrite/v2alpha1/config.proto\x12>istio.envoy.config.filter.network.tcp_cluster_rewrite.v2alpha1\"m\n" +
	"\x11TcpClusterRewrite\x12'\n" +
	"\x0fcluster_pattern\x18\x01 \x01(\tR\x0eclusterPattern\x12/\n" +
	"\x13cluster_replacement\x18\x02 \x01(\tR\x12clusterReplacementBGZEistio.io/api/envoy/config/filter/network/tcp_cluster_rewrite/v2alpha1b\x06proto3"

var (
	file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_rawDescOnce sync.Once
	file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_rawDescData []byte
)

func file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_rawDesc), len(file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_rawDesc)))
	})
	return file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_rawDescData
}

var file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_goTypes = []any{
	(*TcpClusterRewrite)(nil), // 0: istio.envoy.config.filter.network.tcp_cluster_rewrite.v2alpha1.TcpClusterRewrite
}
var file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_init() }
func file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_init() {
	if File_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_rawDesc), len(file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto = out.File
	file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_goTypes = nil
	file_envoy_config_filter_network_tcp_cluster_rewrite_v2alpha1_config_proto_depIdxs = nil
}
