// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: user.proto

package user_pb

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

// 定义GetUserList 方法请求体
type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pn 当前页
	Pn uint32 `protobuf:"varint,1,opt,name=pn,proto3" json:"pn,omitempty"`
	// pSize 每页条数
	PSize uint32 `protobuf:"varint,2,opt,name=pSize,proto3" json:"pSize,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *PageInfo) GetPn() uint32 {
	if x != nil {
		return x.Pn
	}
	return 0
}

func (x *PageInfo) GetPSize() uint32 {
	if x != nil {
		return x.PSize
	}
	return 0
}

// UserInfoResponse 单条用户的信息
type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PassWord string `protobuf:"bytes,2,opt,name=passWord,proto3" json:"passWord,omitempty"`
	Mobile   string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	NickName string `protobuf:"bytes,4,opt,name=nickName,proto3" json:"nickName,omitempty"`
	BirthDay uint64 `protobuf:"varint,5,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	Gender   string `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Role     int32  `protobuf:"varint,7,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfoResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfoResponse) GetPassWord() string {
	if x != nil {
		return x.PassWord
	}
	return ""
}

func (x *UserInfoResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserInfoResponse) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserInfoResponse) GetBirthDay() uint64 {
	if x != nil {
		return x.BirthDay
	}
	return 0
}

func (x *UserInfoResponse) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserInfoResponse) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

// 定义GetUserList 方法响应体
type UserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总条数
	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 用户列表
	Data []*UserInfoResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *UserListResponse) Reset() {
	*x = UserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResponse) ProtoMessage() {}

func (x *UserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResponse.ProtoReflect.Descriptor instead.
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserListResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UserListResponse) GetData() []*UserInfoResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

// GetUserByMobile 方法请求体
type MobileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// mobile 手机号
	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *MobileRequest) Reset() {
	*x = MobileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MobileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MobileRequest) ProtoMessage() {}

func (x *MobileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MobileRequest.ProtoReflect.Descriptor instead.
func (*MobileRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *MobileRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type PasswordCheckInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 明文密码 123456
	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	// 密文密码 是调用 rpc GetUserByMobile之后得到密文密码
	EncryptedPassword string `protobuf:"bytes,2,opt,name=encryptedPassword,proto3" json:"encryptedPassword,omitempty"`
}

func (x *PasswordCheckInfo) Reset() {
	*x = PasswordCheckInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordCheckInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordCheckInfo) ProtoMessage() {}

func (x *PasswordCheckInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordCheckInfo.ProtoReflect.Descriptor instead.
func (*PasswordCheckInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *PasswordCheckInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PasswordCheckInfo) GetEncryptedPassword() string {
	if x != nil {
		return x.EncryptedPassword
	}
	return ""
}

// CheckResponse 检查密码返回体
type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// true 成功 false 失败
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *CheckResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// CreateUserInfo 创建用户请求体
type CreateUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 昵称
	NickName string `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	// 手机号
	Mobile string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	// 明文密码 123456
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserInfo) Reset() {
	*x = CreateUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserInfo) ProtoMessage() {}

func (x *CreateUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserInfo.ProtoReflect.Descriptor instead.
func (*CreateUserInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *CreateUserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *CreateUserInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *CreateUserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id 主键id主要是修改的条件
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 昵称
	NickName string `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	// 性别
	Gender string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	// 生日
	BirthDay uint64 `protobuf:"varint,4,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
}

func (x *UpdateUserInfo) Reset() {
	*x = UpdateUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfo) ProtoMessage() {}

func (x *UpdateUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfo.ProtoReflect.Descriptor instead.
func (*UpdateUserInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateUserInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UpdateUserInfo) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UpdateUserInfo) GetBirthDay() uint64 {
	if x != nil {
		return x.BirthDay
	}
	return 0
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x30, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x70,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x22, 0x4f, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x25,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x0d, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x5d,
	0x0a, 0x11, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x2c, 0x0a, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x29, 0x0a,
	0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x60, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x70, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x32, 0xf7, 0x01, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x09, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x50, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x2e, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x11, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_proto_goTypes = []any{
	(*Empty)(nil),             // 0: Empty
	(*PageInfo)(nil),          // 1: PageInfo
	(*UserInfoResponse)(nil),  // 2: UserInfoResponse
	(*UserListResponse)(nil),  // 3: UserListResponse
	(*MobileRequest)(nil),     // 4: MobileRequest
	(*PasswordCheckInfo)(nil), // 5: PasswordCheckInfo
	(*CheckResponse)(nil),     // 6: CheckResponse
	(*CreateUserInfo)(nil),    // 7: CreateUserInfo
	(*UpdateUserInfo)(nil),    // 8: UpdateUserInfo
}
var file_user_proto_depIdxs = []int32{
	2, // 0: UserListResponse.data:type_name -> UserInfoResponse
	1, // 1: User.GetUserList:input_type -> PageInfo
	4, // 2: User.GetUserByMobile:input_type -> MobileRequest
	5, // 3: User.CheckPassWord:input_type -> PasswordCheckInfo
	7, // 4: User.CreateUser:input_type -> CreateUserInfo
	8, // 5: User.UpdateUser:input_type -> UpdateUserInfo
	3, // 6: User.GetUserList:output_type -> UserListResponse
	2, // 7: User.GetUserByMobile:output_type -> UserInfoResponse
	6, // 8: User.CheckPassWord:output_type -> CheckResponse
	2, // 9: User.CreateUser:output_type -> UserInfoResponse
	0, // 10: User.UpdateUser:output_type -> Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PageInfo); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UserInfoResponse); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UserListResponse); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MobileRequest); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PasswordCheckInfo); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CheckResponse); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserInfo); i {
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
		file_user_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateUserInfo); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
