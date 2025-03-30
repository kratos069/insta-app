// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: post.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Post struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId        int64                  `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	ContentUrl    string                 `protobuf:"bytes,3,opt,name=content_url,json=contentUrl,proto3" json:"content_url,omitempty"`
	Caption       string                 `protobuf:"bytes,4,opt,name=caption,proto3" json:"caption,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Post) Reset() {
	*x = Post{}
	mi := &file_post_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Post) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *Post) GetContentUrl() string {
	if x != nil {
		return x.ContentUrl
	}
	return ""
}

func (x *Post) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *Post) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_post_proto protoreflect.FileDescriptor

const file_post_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"post.proto\x12\x02pb\x1a\x1fgoogle/protobuf/timestamp.proto\"\xae\x01\n" +
	"\x04Post\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\x12\x17\n" +
	"\apost_id\x18\x02 \x01(\x03R\x06postId\x12\x1f\n" +
	"\vcontent_url\x18\x03 \x01(\tR\n" +
	"contentUrl\x12\x18\n" +
	"\acaption\x18\x04 \x01(\tR\acaption\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAtB\x19Z\x17github.com/insta-app/pbb\x06proto3"

var (
	file_post_proto_rawDescOnce sync.Once
	file_post_proto_rawDescData []byte
)

func file_post_proto_rawDescGZIP() []byte {
	file_post_proto_rawDescOnce.Do(func() {
		file_post_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_post_proto_rawDesc), len(file_post_proto_rawDesc)))
	})
	return file_post_proto_rawDescData
}

var file_post_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_post_proto_goTypes = []any{
	(*Post)(nil),                  // 0: pb.Post
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_post_proto_depIdxs = []int32{
	1, // 0: pb.Post.created_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_post_proto_init() }
func file_post_proto_init() {
	if File_post_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_post_proto_rawDesc), len(file_post_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_post_proto_goTypes,
		DependencyIndexes: file_post_proto_depIdxs,
		MessageInfos:      file_post_proto_msgTypes,
	}.Build()
	File_post_proto = out.File
	file_post_proto_goTypes = nil
	file_post_proto_depIdxs = nil
}
