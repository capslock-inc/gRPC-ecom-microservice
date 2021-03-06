// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: Protos/cart/cart.proto

package cartmodel

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
		mi := &file_Protos_cart_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emptyparam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emptyparam) ProtoMessage() {}

func (x *Emptyparam) ProtoReflect() protoreflect.Message {
	mi := &file_Protos_cart_cart_proto_msgTypes[0]
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
	return file_Protos_cart_cart_proto_rawDescGZIP(), []int{0}
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
		mi := &file_Protos_cart_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_Protos_cart_cart_proto_msgTypes[1]
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
	return file_Protos_cart_cart_proto_rawDescGZIP(), []int{1}
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
		mi := &file_Protos_cart_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Userid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Userid) ProtoMessage() {}

func (x *Userid) ProtoReflect() protoreflect.Message {
	mi := &file_Protos_cart_cart_proto_msgTypes[2]
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
	return file_Protos_cart_cart_proto_rawDescGZIP(), []int{2}
}

func (x *Userid) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Productid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Userid string `protobuf:"bytes,2,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *Productid) Reset() {
	*x = Productid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protos_cart_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Productid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Productid) ProtoMessage() {}

func (x *Productid) ProtoReflect() protoreflect.Message {
	mi := &file_Protos_cart_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Productid.ProtoReflect.Descriptor instead.
func (*Productid) Descriptor() ([]byte, []int) {
	return file_Protos_cart_cart_proto_rawDescGZIP(), []int{3}
}

func (x *Productid) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Productid) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Productidlist []string `protobuf:"bytes,2,rep,name=productidlist,proto3" json:"productidlist,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Protos_cart_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_Protos_cart_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_Protos_cart_cart_proto_rawDescGZIP(), []int{4}
}

func (x *Cart) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cart) GetProductidlist() []string {
	if x != nil {
		return x.Productidlist
	}
	return nil
}

var File_Protos_cart_cart_proto protoreflect.FileDescriptor

var file_Protos_cart_cart_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x61, 0x72, 0x74, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0x0c, 0x0a, 0x0a, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x22, 0x1e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x18, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64,
	0x22, 0x3c, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x64, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x64, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x9c,
	0x02, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x11, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64,
	0x1a, 0x0f, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x12, 0x35, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x64, 0x1a, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x64, 0x1a, 0x11, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x11, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x1a,
	0x11, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x39, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x64, 0x1a, 0x11, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x70, 0x73,
	0x6c, 0x6f, 0x63, 0x6b, 0x2d, 0x69, 0x6e, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65,
	0x6d, 0x6f, 0x3b, 0x63, 0x61, 0x72, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Protos_cart_cart_proto_rawDescOnce sync.Once
	file_Protos_cart_cart_proto_rawDescData = file_Protos_cart_cart_proto_rawDesc
)

func file_Protos_cart_cart_proto_rawDescGZIP() []byte {
	file_Protos_cart_cart_proto_rawDescOnce.Do(func() {
		file_Protos_cart_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_Protos_cart_cart_proto_rawDescData)
	})
	return file_Protos_cart_cart_proto_rawDescData
}

var file_Protos_cart_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Protos_cart_cart_proto_goTypes = []interface{}{
	(*Emptyparam)(nil), // 0: cartmodel.emptyparam
	(*Status)(nil),     // 1: cartmodel.status
	(*Userid)(nil),     // 2: cartmodel.userid
	(*Productid)(nil),  // 3: cartmodel.productid
	(*Cart)(nil),       // 4: cartmodel.cart
}
var file_Protos_cart_cart_proto_depIdxs = []int32{
	2, // 0: cartmodel.CartService.GetCartItem:input_type -> cartmodel.userid
	2, // 1: cartmodel.CartService.CreateNewCart:input_type -> cartmodel.userid
	3, // 2: cartmodel.CartService.AddToCart:input_type -> cartmodel.productid
	2, // 3: cartmodel.CartService.DeleteCart:input_type -> cartmodel.userid
	3, // 4: cartmodel.CartService.DeleteCartItem:input_type -> cartmodel.productid
	4, // 5: cartmodel.CartService.GetCartItem:output_type -> cartmodel.cart
	1, // 6: cartmodel.CartService.CreateNewCart:output_type -> cartmodel.status
	1, // 7: cartmodel.CartService.AddToCart:output_type -> cartmodel.status
	1, // 8: cartmodel.CartService.DeleteCart:output_type -> cartmodel.status
	1, // 9: cartmodel.CartService.DeleteCartItem:output_type -> cartmodel.status
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Protos_cart_cart_proto_init() }
func file_Protos_cart_cart_proto_init() {
	if File_Protos_cart_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Protos_cart_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Protos_cart_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_Protos_cart_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_Protos_cart_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Productid); i {
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
		file_Protos_cart_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
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
			RawDescriptor: file_Protos_cart_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Protos_cart_cart_proto_goTypes,
		DependencyIndexes: file_Protos_cart_cart_proto_depIdxs,
		MessageInfos:      file_Protos_cart_cart_proto_msgTypes,
	}.Build()
	File_Protos_cart_cart_proto = out.File
	file_Protos_cart_cart_proto_rawDesc = nil
	file_Protos_cart_cart_proto_goTypes = nil
	file_Protos_cart_cart_proto_depIdxs = nil
}
