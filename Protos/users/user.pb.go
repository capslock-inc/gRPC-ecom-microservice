// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: Protos/users/user.proto

package usermodel

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

type Emptyparam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Emptyparam) Reset() {
	*x = Emptyparam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protos_users_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emptyparam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emptyparam) ProtoMessage() {}

func (x *Emptyparam) ProtoReflect() protoreflect.Message {
	mi := &file_Protos_users_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emptyparam.ProtoReflect.Descriptor instead.
func (*Emptyparam) Descriptor() ([]byte, []int) {
	return file_Protos_users_user_proto_rawDescGZIP(), []int{0}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protos_users_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_Protos_users_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_Protos_users_user_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Userid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Userid) Reset() {
	*x = Userid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protos_users_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Userid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Userid) ProtoMessage() {}

func (x *Userid) ProtoReflect() protoreflect.Message {
	mi := &file_Protos_users_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Userid.ProtoReflect.Descriptor instead.
func (*Userid) Descriptor() ([]byte, []int) {
	return file_Protos_users_user_proto_rawDescGZIP(), []int{2}
}

func (x *Userid) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protos_users_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_Protos_users_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_Protos_users_user_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Userlist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*User `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Userlist) Reset() {
	*x = Userlist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protos_users_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Userlist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Userlist) ProtoMessage() {}

func (x *Userlist) ProtoReflect() protoreflect.Message {
	mi := &file_Protos_users_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Userlist.ProtoReflect.Descriptor instead.
func (*Userlist) Descriptor() ([]byte, []int) {
	return file_Protos_users_user_proto_rawDescGZIP(), []int{4}
}

func (x *Userlist) GetList() []*User {
	if x != nil {
		return x.List
	}
	return nil
}

var File_Protos_users_user_proto protoreflect.FileDescriptor

var file_Protos_users_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x75, 0x73, 0x65, 0x72, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x0c, 0x0a, 0x0a, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x22, 0x1e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x2f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xc7, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x30, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x79, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x11, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61,
	0x70, 0x73, 0x6c, 0x6f, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x64, 0x65, 0x6d, 0x6f, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Protos_users_user_proto_rawDescOnce sync.Once
	file_Protos_users_user_proto_rawDescData = file_Protos_users_user_proto_rawDesc
)

func file_Protos_users_user_proto_rawDescGZIP() []byte {
	file_Protos_users_user_proto_rawDescOnce.Do(func() {
		file_Protos_users_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_Protos_users_user_proto_rawDescData)
	})
	return file_Protos_users_user_proto_rawDescData
}

var file_Protos_users_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Protos_users_user_proto_goTypes = []interface{}{
	(*Emptyparam)(nil), // 0: usermodel.emptyparam
	(*Status)(nil),     // 1: usermodel.status
	(*Userid)(nil),     // 2: usermodel.userid
	(*User)(nil),       // 3: usermodel.user
	(*Userlist)(nil),   // 4: usermodel.userlist
}
var file_Protos_users_user_proto_depIdxs = []int32{
	3, // 0: usermodel.userlist.list:type_name -> usermodel.user
	0, // 1: usermodel.UserService.GetAllUser:input_type -> usermodel.emptyparam
	2, // 2: usermodel.UserService.GetUserById:input_type -> usermodel.userid
	3, // 3: usermodel.UserService.AddNewUser:input_type -> usermodel.user
	3, // 4: usermodel.UserService.UpdateUser:input_type -> usermodel.user
	2, // 5: usermodel.UserService.DeleteUser:input_type -> usermodel.userid
	0, // 6: usermodel.UserService.Healthy:input_type -> usermodel.emptyparam
	4, // 7: usermodel.UserService.GetAllUser:output_type -> usermodel.userlist
	3, // 8: usermodel.UserService.GetUserById:output_type -> usermodel.user
	1, // 9: usermodel.UserService.AddNewUser:output_type -> usermodel.status
	1, // 10: usermodel.UserService.UpdateUser:output_type -> usermodel.status
	1, // 11: usermodel.UserService.DeleteUser:output_type -> usermodel.status
	1, // 12: usermodel.UserService.Healthy:output_type -> usermodel.status
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Protos_users_user_proto_init() }
func file_Protos_users_user_proto_init() {
	if File_Protos_users_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Protos_users_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emptyparam); i {
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
		file_Protos_users_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_Protos_users_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Userid); i {
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
		file_Protos_users_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_Protos_users_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Userlist); i {
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
			RawDescriptor: file_Protos_users_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Protos_users_user_proto_goTypes,
		DependencyIndexes: file_Protos_users_user_proto_depIdxs,
		MessageInfos:      file_Protos_users_user_proto_msgTypes,
	}.Build()
	File_Protos_users_user_proto = out.File
	file_Protos_users_user_proto_rawDesc = nil
	file_Protos_users_user_proto_goTypes = nil
	file_Protos_users_user_proto_depIdxs = nil
}
