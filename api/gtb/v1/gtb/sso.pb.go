// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
// source: sso.proto

package gtb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LoginMethodType int32

const (
	LoginMethodType_LMT_UNKNOWN  LoginMethodType = 0
	LoginMethodType_LMT_EMAIL    LoginMethodType = 1
	LoginMethodType_LMT_PHONE    LoginMethodType = 2
	LoginMethodType_LMT_FACEBOOK LoginMethodType = 3
	LoginMethodType_LMT_TWITTER  LoginMethodType = 4
	LoginMethodType_LMT_GOOGLE   LoginMethodType = 5
	LoginMethodType_LMT_SNAPCHAT LoginMethodType = 6
	LoginMethodType_LMT_APPLE    LoginMethodType = 7
)

// Enum value maps for LoginMethodType.
var (
	LoginMethodType_name = map[int32]string{
		0: "LMT_UNKNOWN",
		1: "LMT_EMAIL",
		2: "LMT_PHONE",
		3: "LMT_FACEBOOK",
		4: "LMT_TWITTER",
		5: "LMT_GOOGLE",
		6: "LMT_SNAPCHAT",
		7: "LMT_APPLE",
	}
	LoginMethodType_value = map[string]int32{
		"LMT_UNKNOWN":  0,
		"LMT_EMAIL":    1,
		"LMT_PHONE":    2,
		"LMT_FACEBOOK": 3,
		"LMT_TWITTER":  4,
		"LMT_GOOGLE":   5,
		"LMT_SNAPCHAT": 6,
		"LMT_APPLE":    7,
	}
)

func (x LoginMethodType) Enum() *LoginMethodType {
	p := new(LoginMethodType)
	*p = x
	return p
}

func (x LoginMethodType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginMethodType) Descriptor() protoreflect.EnumDescriptor {
	return file_sso_proto_enumTypes[0].Descriptor()
}

func (LoginMethodType) Type() protoreflect.EnumType {
	return &file_sso_proto_enumTypes[0]
}

func (x LoginMethodType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginMethodType.Descriptor instead.
func (LoginMethodType) EnumDescriptor() ([]byte, []int) {
	return file_sso_proto_rawDescGZIP(), []int{0}
}

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account     string          `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password    string          `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // 密码登录
	Token       string          `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`       // 第三方登录填access token
	LoginMethod LoginMethodType `protobuf:"varint,4,opt,name=login_method,json=loginMethod,proto3,enum=sso.LoginMethodType" json:"login_method,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_sso_proto_rawDescGZIP(), []int{0}
}

func (x *SignInRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignInRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SignInRequest) GetLoginMethod() LoginMethodType {
	if x != nil {
		return x.LoginMethod
	}
	return LoginMethodType_LMT_UNKNOWN
}

type TokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ValidSecs uint32 `protobuf:"varint,2,opt,name=valid_secs,json=validSecs,proto3" json:"valid_secs,omitempty"`
}

func (x *TokenInfo) Reset() {
	*x = TokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfo) ProtoMessage() {}

func (x *TokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sso_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenInfo.ProtoReflect.Descriptor instead.
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return file_sso_proto_rawDescGZIP(), []int{1}
}

func (x *TokenInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TokenInfo) GetValidSecs() uint32 {
	if x != nil {
		return x.ValidSecs
	}
	return 0
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid          uint64          `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	AccessToken  *TokenInfo      `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`                            // 设置到channel的ChannelCredentials的key: "x-auth-token"
	RefreshToken *TokenInfo      `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`                         // 每次请求access_token自动延长有效期, 过期需要重新登录
	LoginMethod  LoginMethodType `protobuf:"varint,10,opt,name=login_method,json=loginMethod,proto3,enum=sso.LoginMethodType" json:"login_method,omitempty"` // 账号类型
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_sso_proto_rawDescGZIP(), []int{2}
}

func (x *SignInResponse) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SignInResponse) GetAccessToken() *TokenInfo {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

func (x *SignInResponse) GetRefreshToken() *TokenInfo {
	if x != nil {
		return x.RefreshToken
	}
	return nil
}

func (x *SignInResponse) GetLoginMethod() LoginMethodType {
	if x != nil {
		return x.LoginMethod
	}
	return LoginMethodType_LMT_UNKNOWN
}

var File_sso_proto protoreflect.FileDescriptor

var file_sso_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x73, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x73, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94,
	0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x0c,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x40, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x53, 0x65, 0x63, 0x73, 0x22, 0xc3, 0x01, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x33, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x73, 0x73, 0x6f,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2a, 0x94, 0x01,
	0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x4d, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x4d, 0x54, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x4d, 0x54, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x02,
	0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x4d, 0x54, 0x5f, 0x46, 0x41, 0x43, 0x45, 0x42, 0x4f, 0x4f, 0x4b,
	0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x4d, 0x54, 0x5f, 0x54, 0x57, 0x49, 0x54, 0x54, 0x45,
	0x52, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4d, 0x54, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c,
	0x45, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x4d, 0x54, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x43,
	0x48, 0x41, 0x54, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x4d, 0x54, 0x5f, 0x41, 0x50, 0x50,
	0x4c, 0x45, 0x10, 0x07, 0x32, 0x5d, 0x0a, 0x03, 0x53, 0x73, 0x6f, 0x12, 0x56, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x12, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x73, 0x6f, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x67, 0x74, 0x62, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x6e,
	0x3a, 0x01, 0x2a, 0x42, 0x33, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x74, 0x62, 0x2e, 0x76,
	0x31, 0x42, 0x0a, 0x53, 0x73, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a,
	0x17, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6e, 0x36,
	0x35, 0x35, 0x33, 0x35, 0x2f, 0x67, 0x74, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sso_proto_rawDescOnce sync.Once
	file_sso_proto_rawDescData = file_sso_proto_rawDesc
)

func file_sso_proto_rawDescGZIP() []byte {
	file_sso_proto_rawDescOnce.Do(func() {
		file_sso_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_proto_rawDescData)
	})
	return file_sso_proto_rawDescData
}

var file_sso_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sso_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sso_proto_goTypes = []interface{}{
	(LoginMethodType)(0),   // 0: sso.LoginMethodType
	(*SignInRequest)(nil),  // 1: sso.SignInRequest
	(*TokenInfo)(nil),      // 2: sso.TokenInfo
	(*SignInResponse)(nil), // 3: sso.SignInResponse
}
var file_sso_proto_depIdxs = []int32{
	0, // 0: sso.SignInRequest.login_method:type_name -> sso.LoginMethodType
	2, // 1: sso.SignInResponse.access_token:type_name -> sso.TokenInfo
	2, // 2: sso.SignInResponse.refresh_token:type_name -> sso.TokenInfo
	0, // 3: sso.SignInResponse.login_method:type_name -> sso.LoginMethodType
	1, // 4: sso.Sso.SignIn:input_type -> sso.SignInRequest
	3, // 5: sso.Sso.SignIn:output_type -> sso.SignInResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sso_proto_init() }
func file_sso_proto_init() {
	if File_sso_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sso_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
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
		file_sso_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenInfo); i {
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
		file_sso_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
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
			RawDescriptor: file_sso_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_proto_goTypes,
		DependencyIndexes: file_sso_proto_depIdxs,
		EnumInfos:         file_sso_proto_enumTypes,
		MessageInfos:      file_sso_proto_msgTypes,
	}.Build()
	File_sso_proto = out.File
	file_sso_proto_rawDesc = nil
	file_sso_proto_goTypes = nil
	file_sso_proto_depIdxs = nil
}
