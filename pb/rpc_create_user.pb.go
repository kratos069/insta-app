// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: rpc_create_user.proto

package pb

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

type CreateUserRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Username       string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password       string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	FullName       string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email          string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	ProfilePicture string                 `protobuf:"bytes,5,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Bio            string                 `protobuf:"bytes,6,opt,name=bio,proto3" json:"bio,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_rpc_create_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_rpc_create_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetProfilePicture() string {
	if x != nil {
		return x.ProfilePicture
	}
	return ""
}

func (x *CreateUserRequest) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_rpc_create_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_rpc_create_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_rpc_create_user_proto protoreflect.FileDescriptor

const file_rpc_create_user_proto_rawDesc = "" +
	"\n" +
	"\x15rpc_create_user.proto\x12\x02pb\x1a\n" +
	"user.proto\"\xb9\x01\n" +
	"\x11CreateUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12\x1b\n" +
	"\tfull_name\x18\x03 \x01(\tR\bfullName\x12\x14\n" +
	"\x05email\x18\x04 \x01(\tR\x05email\x12'\n" +
	"\x0fprofile_picture\x18\x05 \x01(\tR\x0eprofilePicture\x12\x10\n" +
	"\x03bio\x18\x06 \x01(\tR\x03bio\"2\n" +
	"\x12CreateUserResponse\x12\x1c\n" +
	"\x04user\x18\x01 \x01(\v2\b.pb.UserR\x04userB\x19Z\x17github.com/insta-app/pbb\x06proto3"

var (
	file_rpc_create_user_proto_rawDescOnce sync.Once
	file_rpc_create_user_proto_rawDescData []byte
)

func file_rpc_create_user_proto_rawDescGZIP() []byte {
	file_rpc_create_user_proto_rawDescOnce.Do(func() {
		file_rpc_create_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rpc_create_user_proto_rawDesc), len(file_rpc_create_user_proto_rawDesc)))
	})
	return file_rpc_create_user_proto_rawDescData
}

var file_rpc_create_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_create_user_proto_goTypes = []any{
	(*CreateUserRequest)(nil),  // 0: pb.CreateUserRequest
	(*CreateUserResponse)(nil), // 1: pb.CreateUserResponse
	(*User)(nil),               // 2: pb.User
}
var file_rpc_create_user_proto_depIdxs = []int32{
	2, // 0: pb.CreateUserResponse.user:type_name -> pb.User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_create_user_proto_init() }
func file_rpc_create_user_proto_init() {
	if File_rpc_create_user_proto != nil {
		return
	}
	file_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rpc_create_user_proto_rawDesc), len(file_rpc_create_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_create_user_proto_goTypes,
		DependencyIndexes: file_rpc_create_user_proto_depIdxs,
		MessageInfos:      file_rpc_create_user_proto_msgTypes,
	}.Build()
	File_rpc_create_user_proto = out.File
	file_rpc_create_user_proto_goTypes = nil
	file_rpc_create_user_proto_depIdxs = nil
}
