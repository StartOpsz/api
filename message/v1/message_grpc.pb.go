// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.14.0
// source: message/v1/message.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Message_SendMailMsg_FullMethodName                       = "/message.v1.Message/SendMailMsg"
	Message_SendLoginVerifyCodeMessage_FullMethodName        = "/message.v1.Message/SendLoginVerifyCodeMessage"
	Message_GetLoginVerifyCodeMessage_FullMethodName         = "/message.v1.Message/GetLoginVerifyCodeMessage"
	Message_SendModifyPasswdVerifyCodeMessage_FullMethodName = "/message.v1.Message/SendModifyPasswdVerifyCodeMessage"
	Message_GetModifyPasswdVerifyCodeMessage_FullMethodName  = "/message.v1.Message/GetModifyPasswdVerifyCodeMessage"
	Message_SendRegisterVerifyCodeMessage_FullMethodName     = "/message.v1.Message/SendRegisterVerifyCodeMessage"
	Message_GetRegisterVerifyCodeMessage_FullMethodName      = "/message.v1.Message/GetRegisterVerifyCodeMessage"
	Message_SendAlertNotifyMessage_FullMethodName            = "/message.v1.Message/SendAlertNotifyMessage"
	Message_SendKubernetesEventNotifyMessage_FullMethodName  = "/message.v1.Message/SendKubernetesEventNotifyMessage"
	Message_SendOrgDingTalkWorkMsg_FullMethodName            = "/message.v1.Message/SendOrgDingTalkWorkMsg"
	Message_CreateMessageGroup_FullMethodName                = "/message.v1.Message/CreateMessageGroup"
	Message_ListMessageGroup_FullMethodName                  = "/message.v1.Message/ListMessageGroup"
	Message_DelMessageGroup_FullMethodName                   = "/message.v1.Message/DelMessageGroup"
	Message_AddMessageGroupMember_FullMethodName             = "/message.v1.Message/AddMessageGroupMember"
	Message_GetMessageGroupMember_FullMethodName             = "/message.v1.Message/GetMessageGroupMember"
	Message_DelMessageGroupMember_FullMethodName             = "/message.v1.Message/DelMessageGroupMember"
	Message_SendMessageGroupContent_FullMethodName           = "/message.v1.Message/SendMessageGroupContent"
	Message_Translate_FullMethodName                         = "/message.v1.Message/Translate"
)

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClient interface {
	// Email
	SendMailMsg(ctx context.Context, in *SendMailMsgReq, opts ...grpc.CallOption) (*SendMailMsgReply, error)
	// 发送短信验证码消息 - 登录验证码
	SendLoginVerifyCodeMessage(ctx context.Context, in *SendLoginVerifyCodeMessageReq, opts ...grpc.CallOption) (*SendLoginVerifyCodeMessageReply, error)
	GetLoginVerifyCodeMessage(ctx context.Context, in *GetLoginVerifyCodeMessageReq, opts ...grpc.CallOption) (*GetLoginVerifyCodeMessageReply, error)
	// 发送短信验证码消息 - 修改密码验证码
	SendModifyPasswdVerifyCodeMessage(ctx context.Context, in *SendModifyPasswdVerifyCodeMessageReq, opts ...grpc.CallOption) (*SendModifyPasswdVerifyCodeMessageReply, error)
	GetModifyPasswdVerifyCodeMessage(ctx context.Context, in *GetModifyPasswdVerifyCodeMessageReq, opts ...grpc.CallOption) (*GetModifyPasswdVerifyCodeMessageReply, error)
	// 发送短信验证码消息 - 注册用户验证码
	SendRegisterVerifyCodeMessage(ctx context.Context, in *SendRegisterVerifyCodeMessageReq, opts ...grpc.CallOption) (*SendRegisterVerifyCodeMessageReply, error)
	GetRegisterVerifyCodeMessage(ctx context.Context, in *GetRegisterVerifyCodeMessageReq, opts ...grpc.CallOption) (*GetRegisterVerifyCodeMessageReply, error)
	// 发送通知消息 - 已放弃
	SendAlertNotifyMessage(ctx context.Context, in *SendAlertNotifyMessageReq, opts ...grpc.CallOption) (*SendAlertNotifyMessageReply, error)
	SendKubernetesEventNotifyMessage(ctx context.Context, in *SendKubernetesEventNotifyMessageReq, opts ...grpc.CallOption) (*SendKubernetesEventNotifyMessageReply, error)
	// 发送钉钉个人消息
	SendOrgDingTalkWorkMsg(ctx context.Context, in *SendOrgDingTalkWorkMsgReq, opts ...grpc.CallOption) (*SendOrgDingTalkWorkMsgReply, error)
	// 消息组
	CreateMessageGroup(ctx context.Context, in *CreateMessageGroupReq, opts ...grpc.CallOption) (*CreateMessageGroupReply, error)
	ListMessageGroup(ctx context.Context, in *ListMessageGroupReq, opts ...grpc.CallOption) (*ListMessageGroupReply, error)
	DelMessageGroup(ctx context.Context, in *DelMessageGroupReq, opts ...grpc.CallOption) (*DelMessageGroupReply, error)
	AddMessageGroupMember(ctx context.Context, in *AddMessageGroupMemberReq, opts ...grpc.CallOption) (*AddMessageGroupMemberReply, error)
	GetMessageGroupMember(ctx context.Context, in *GetMessageGroupMemberReq, opts ...grpc.CallOption) (*GetMessageGroupMemberReply, error)
	DelMessageGroupMember(ctx context.Context, in *DelMessageGroupMemberReq, opts ...grpc.CallOption) (*DelMessageGroupMemberReply, error)
	// 发送 “内容” 到消息组内
	SendMessageGroupContent(ctx context.Context, in *SendMessageGroupContentReq, opts ...grpc.CallOption) (*SendMessageGroupContentReply, error)
	// 翻译
	Translate(ctx context.Context, in *TranslateReq, opts ...grpc.CallOption) (*TranslateReply, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) SendMailMsg(ctx context.Context, in *SendMailMsgReq, opts ...grpc.CallOption) (*SendMailMsgReply, error) {
	out := new(SendMailMsgReply)
	err := c.cc.Invoke(ctx, Message_SendMailMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SendLoginVerifyCodeMessage(ctx context.Context, in *SendLoginVerifyCodeMessageReq, opts ...grpc.CallOption) (*SendLoginVerifyCodeMessageReply, error) {
	out := new(SendLoginVerifyCodeMessageReply)
	err := c.cc.Invoke(ctx, Message_SendLoginVerifyCodeMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) GetLoginVerifyCodeMessage(ctx context.Context, in *GetLoginVerifyCodeMessageReq, opts ...grpc.CallOption) (*GetLoginVerifyCodeMessageReply, error) {
	out := new(GetLoginVerifyCodeMessageReply)
	err := c.cc.Invoke(ctx, Message_GetLoginVerifyCodeMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SendModifyPasswdVerifyCodeMessage(ctx context.Context, in *SendModifyPasswdVerifyCodeMessageReq, opts ...grpc.CallOption) (*SendModifyPasswdVerifyCodeMessageReply, error) {
	out := new(SendModifyPasswdVerifyCodeMessageReply)
	err := c.cc.Invoke(ctx, Message_SendModifyPasswdVerifyCodeMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) GetModifyPasswdVerifyCodeMessage(ctx context.Context, in *GetModifyPasswdVerifyCodeMessageReq, opts ...grpc.CallOption) (*GetModifyPasswdVerifyCodeMessageReply, error) {
	out := new(GetModifyPasswdVerifyCodeMessageReply)
	err := c.cc.Invoke(ctx, Message_GetModifyPasswdVerifyCodeMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SendRegisterVerifyCodeMessage(ctx context.Context, in *SendRegisterVerifyCodeMessageReq, opts ...grpc.CallOption) (*SendRegisterVerifyCodeMessageReply, error) {
	out := new(SendRegisterVerifyCodeMessageReply)
	err := c.cc.Invoke(ctx, Message_SendRegisterVerifyCodeMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) GetRegisterVerifyCodeMessage(ctx context.Context, in *GetRegisterVerifyCodeMessageReq, opts ...grpc.CallOption) (*GetRegisterVerifyCodeMessageReply, error) {
	out := new(GetRegisterVerifyCodeMessageReply)
	err := c.cc.Invoke(ctx, Message_GetRegisterVerifyCodeMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SendAlertNotifyMessage(ctx context.Context, in *SendAlertNotifyMessageReq, opts ...grpc.CallOption) (*SendAlertNotifyMessageReply, error) {
	out := new(SendAlertNotifyMessageReply)
	err := c.cc.Invoke(ctx, Message_SendAlertNotifyMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SendKubernetesEventNotifyMessage(ctx context.Context, in *SendKubernetesEventNotifyMessageReq, opts ...grpc.CallOption) (*SendKubernetesEventNotifyMessageReply, error) {
	out := new(SendKubernetesEventNotifyMessageReply)
	err := c.cc.Invoke(ctx, Message_SendKubernetesEventNotifyMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SendOrgDingTalkWorkMsg(ctx context.Context, in *SendOrgDingTalkWorkMsgReq, opts ...grpc.CallOption) (*SendOrgDingTalkWorkMsgReply, error) {
	out := new(SendOrgDingTalkWorkMsgReply)
	err := c.cc.Invoke(ctx, Message_SendOrgDingTalkWorkMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) CreateMessageGroup(ctx context.Context, in *CreateMessageGroupReq, opts ...grpc.CallOption) (*CreateMessageGroupReply, error) {
	out := new(CreateMessageGroupReply)
	err := c.cc.Invoke(ctx, Message_CreateMessageGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) ListMessageGroup(ctx context.Context, in *ListMessageGroupReq, opts ...grpc.CallOption) (*ListMessageGroupReply, error) {
	out := new(ListMessageGroupReply)
	err := c.cc.Invoke(ctx, Message_ListMessageGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) DelMessageGroup(ctx context.Context, in *DelMessageGroupReq, opts ...grpc.CallOption) (*DelMessageGroupReply, error) {
	out := new(DelMessageGroupReply)
	err := c.cc.Invoke(ctx, Message_DelMessageGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) AddMessageGroupMember(ctx context.Context, in *AddMessageGroupMemberReq, opts ...grpc.CallOption) (*AddMessageGroupMemberReply, error) {
	out := new(AddMessageGroupMemberReply)
	err := c.cc.Invoke(ctx, Message_AddMessageGroupMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) GetMessageGroupMember(ctx context.Context, in *GetMessageGroupMemberReq, opts ...grpc.CallOption) (*GetMessageGroupMemberReply, error) {
	out := new(GetMessageGroupMemberReply)
	err := c.cc.Invoke(ctx, Message_GetMessageGroupMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) DelMessageGroupMember(ctx context.Context, in *DelMessageGroupMemberReq, opts ...grpc.CallOption) (*DelMessageGroupMemberReply, error) {
	out := new(DelMessageGroupMemberReply)
	err := c.cc.Invoke(ctx, Message_DelMessageGroupMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) SendMessageGroupContent(ctx context.Context, in *SendMessageGroupContentReq, opts ...grpc.CallOption) (*SendMessageGroupContentReply, error) {
	out := new(SendMessageGroupContentReply)
	err := c.cc.Invoke(ctx, Message_SendMessageGroupContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) Translate(ctx context.Context, in *TranslateReq, opts ...grpc.CallOption) (*TranslateReply, error) {
	out := new(TranslateReply)
	err := c.cc.Invoke(ctx, Message_Translate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
// All implementations must embed UnimplementedMessageServer
// for forward compatibility
type MessageServer interface {
	// Email
	SendMailMsg(context.Context, *SendMailMsgReq) (*SendMailMsgReply, error)
	// 发送短信验证码消息 - 登录验证码
	SendLoginVerifyCodeMessage(context.Context, *SendLoginVerifyCodeMessageReq) (*SendLoginVerifyCodeMessageReply, error)
	GetLoginVerifyCodeMessage(context.Context, *GetLoginVerifyCodeMessageReq) (*GetLoginVerifyCodeMessageReply, error)
	// 发送短信验证码消息 - 修改密码验证码
	SendModifyPasswdVerifyCodeMessage(context.Context, *SendModifyPasswdVerifyCodeMessageReq) (*SendModifyPasswdVerifyCodeMessageReply, error)
	GetModifyPasswdVerifyCodeMessage(context.Context, *GetModifyPasswdVerifyCodeMessageReq) (*GetModifyPasswdVerifyCodeMessageReply, error)
	// 发送短信验证码消息 - 注册用户验证码
	SendRegisterVerifyCodeMessage(context.Context, *SendRegisterVerifyCodeMessageReq) (*SendRegisterVerifyCodeMessageReply, error)
	GetRegisterVerifyCodeMessage(context.Context, *GetRegisterVerifyCodeMessageReq) (*GetRegisterVerifyCodeMessageReply, error)
	// 发送通知消息 - 已放弃
	SendAlertNotifyMessage(context.Context, *SendAlertNotifyMessageReq) (*SendAlertNotifyMessageReply, error)
	SendKubernetesEventNotifyMessage(context.Context, *SendKubernetesEventNotifyMessageReq) (*SendKubernetesEventNotifyMessageReply, error)
	// 发送钉钉个人消息
	SendOrgDingTalkWorkMsg(context.Context, *SendOrgDingTalkWorkMsgReq) (*SendOrgDingTalkWorkMsgReply, error)
	// 消息组
	CreateMessageGroup(context.Context, *CreateMessageGroupReq) (*CreateMessageGroupReply, error)
	ListMessageGroup(context.Context, *ListMessageGroupReq) (*ListMessageGroupReply, error)
	DelMessageGroup(context.Context, *DelMessageGroupReq) (*DelMessageGroupReply, error)
	AddMessageGroupMember(context.Context, *AddMessageGroupMemberReq) (*AddMessageGroupMemberReply, error)
	GetMessageGroupMember(context.Context, *GetMessageGroupMemberReq) (*GetMessageGroupMemberReply, error)
	DelMessageGroupMember(context.Context, *DelMessageGroupMemberReq) (*DelMessageGroupMemberReply, error)
	// 发送 “内容” 到消息组内
	SendMessageGroupContent(context.Context, *SendMessageGroupContentReq) (*SendMessageGroupContentReply, error)
	// 翻译
	Translate(context.Context, *TranslateReq) (*TranslateReply, error)
	mustEmbedUnimplementedMessageServer()
}

// UnimplementedMessageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (UnimplementedMessageServer) SendMailMsg(context.Context, *SendMailMsgReq) (*SendMailMsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMailMsg not implemented")
}
func (UnimplementedMessageServer) SendLoginVerifyCodeMessage(context.Context, *SendLoginVerifyCodeMessageReq) (*SendLoginVerifyCodeMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLoginVerifyCodeMessage not implemented")
}
func (UnimplementedMessageServer) GetLoginVerifyCodeMessage(context.Context, *GetLoginVerifyCodeMessageReq) (*GetLoginVerifyCodeMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginVerifyCodeMessage not implemented")
}
func (UnimplementedMessageServer) SendModifyPasswdVerifyCodeMessage(context.Context, *SendModifyPasswdVerifyCodeMessageReq) (*SendModifyPasswdVerifyCodeMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendModifyPasswdVerifyCodeMessage not implemented")
}
func (UnimplementedMessageServer) GetModifyPasswdVerifyCodeMessage(context.Context, *GetModifyPasswdVerifyCodeMessageReq) (*GetModifyPasswdVerifyCodeMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModifyPasswdVerifyCodeMessage not implemented")
}
func (UnimplementedMessageServer) SendRegisterVerifyCodeMessage(context.Context, *SendRegisterVerifyCodeMessageReq) (*SendRegisterVerifyCodeMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRegisterVerifyCodeMessage not implemented")
}
func (UnimplementedMessageServer) GetRegisterVerifyCodeMessage(context.Context, *GetRegisterVerifyCodeMessageReq) (*GetRegisterVerifyCodeMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisterVerifyCodeMessage not implemented")
}
func (UnimplementedMessageServer) SendAlertNotifyMessage(context.Context, *SendAlertNotifyMessageReq) (*SendAlertNotifyMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAlertNotifyMessage not implemented")
}
func (UnimplementedMessageServer) SendKubernetesEventNotifyMessage(context.Context, *SendKubernetesEventNotifyMessageReq) (*SendKubernetesEventNotifyMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendKubernetesEventNotifyMessage not implemented")
}
func (UnimplementedMessageServer) SendOrgDingTalkWorkMsg(context.Context, *SendOrgDingTalkWorkMsgReq) (*SendOrgDingTalkWorkMsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOrgDingTalkWorkMsg not implemented")
}
func (UnimplementedMessageServer) CreateMessageGroup(context.Context, *CreateMessageGroupReq) (*CreateMessageGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessageGroup not implemented")
}
func (UnimplementedMessageServer) ListMessageGroup(context.Context, *ListMessageGroupReq) (*ListMessageGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessageGroup not implemented")
}
func (UnimplementedMessageServer) DelMessageGroup(context.Context, *DelMessageGroupReq) (*DelMessageGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelMessageGroup not implemented")
}
func (UnimplementedMessageServer) AddMessageGroupMember(context.Context, *AddMessageGroupMemberReq) (*AddMessageGroupMemberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMessageGroupMember not implemented")
}
func (UnimplementedMessageServer) GetMessageGroupMember(context.Context, *GetMessageGroupMemberReq) (*GetMessageGroupMemberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageGroupMember not implemented")
}
func (UnimplementedMessageServer) DelMessageGroupMember(context.Context, *DelMessageGroupMemberReq) (*DelMessageGroupMemberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelMessageGroupMember not implemented")
}
func (UnimplementedMessageServer) SendMessageGroupContent(context.Context, *SendMessageGroupContentReq) (*SendMessageGroupContentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageGroupContent not implemented")
}
func (UnimplementedMessageServer) Translate(context.Context, *TranslateReq) (*TranslateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}
func (UnimplementedMessageServer) mustEmbedUnimplementedMessageServer() {}

// UnsafeMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServer will
// result in compilation errors.
type UnsafeMessageServer interface {
	mustEmbedUnimplementedMessageServer()
}

func RegisterMessageServer(s grpc.ServiceRegistrar, srv MessageServer) {
	s.RegisterService(&Message_ServiceDesc, srv)
}

func _Message_SendMailMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMailMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendMailMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendMailMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendMailMsg(ctx, req.(*SendMailMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SendLoginVerifyCodeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendLoginVerifyCodeMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendLoginVerifyCodeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendLoginVerifyCodeMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendLoginVerifyCodeMessage(ctx, req.(*SendLoginVerifyCodeMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_GetLoginVerifyCodeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginVerifyCodeMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).GetLoginVerifyCodeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_GetLoginVerifyCodeMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).GetLoginVerifyCodeMessage(ctx, req.(*GetLoginVerifyCodeMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SendModifyPasswdVerifyCodeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendModifyPasswdVerifyCodeMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendModifyPasswdVerifyCodeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendModifyPasswdVerifyCodeMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendModifyPasswdVerifyCodeMessage(ctx, req.(*SendModifyPasswdVerifyCodeMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_GetModifyPasswdVerifyCodeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModifyPasswdVerifyCodeMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).GetModifyPasswdVerifyCodeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_GetModifyPasswdVerifyCodeMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).GetModifyPasswdVerifyCodeMessage(ctx, req.(*GetModifyPasswdVerifyCodeMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SendRegisterVerifyCodeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRegisterVerifyCodeMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendRegisterVerifyCodeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendRegisterVerifyCodeMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendRegisterVerifyCodeMessage(ctx, req.(*SendRegisterVerifyCodeMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_GetRegisterVerifyCodeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisterVerifyCodeMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).GetRegisterVerifyCodeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_GetRegisterVerifyCodeMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).GetRegisterVerifyCodeMessage(ctx, req.(*GetRegisterVerifyCodeMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SendAlertNotifyMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendAlertNotifyMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendAlertNotifyMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendAlertNotifyMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendAlertNotifyMessage(ctx, req.(*SendAlertNotifyMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SendKubernetesEventNotifyMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendKubernetesEventNotifyMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendKubernetesEventNotifyMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendKubernetesEventNotifyMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendKubernetesEventNotifyMessage(ctx, req.(*SendKubernetesEventNotifyMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SendOrgDingTalkWorkMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOrgDingTalkWorkMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendOrgDingTalkWorkMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendOrgDingTalkWorkMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendOrgDingTalkWorkMsg(ctx, req.(*SendOrgDingTalkWorkMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_CreateMessageGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).CreateMessageGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_CreateMessageGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).CreateMessageGroup(ctx, req.(*CreateMessageGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_ListMessageGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessageGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).ListMessageGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_ListMessageGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).ListMessageGroup(ctx, req.(*ListMessageGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_DelMessageGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelMessageGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).DelMessageGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_DelMessageGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).DelMessageGroup(ctx, req.(*DelMessageGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_AddMessageGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMessageGroupMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).AddMessageGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_AddMessageGroupMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).AddMessageGroupMember(ctx, req.(*AddMessageGroupMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_GetMessageGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageGroupMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).GetMessageGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_GetMessageGroupMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).GetMessageGroupMember(ctx, req.(*GetMessageGroupMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_DelMessageGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelMessageGroupMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).DelMessageGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_DelMessageGroupMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).DelMessageGroupMember(ctx, req.(*DelMessageGroupMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_SendMessageGroupContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageGroupContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendMessageGroupContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_SendMessageGroupContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendMessageGroupContent(ctx, req.(*SendMessageGroupContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_Translate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Translate(ctx, req.(*TranslateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Message_ServiceDesc is the grpc.ServiceDesc for Message service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Message_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.v1.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMailMsg",
			Handler:    _Message_SendMailMsg_Handler,
		},
		{
			MethodName: "SendLoginVerifyCodeMessage",
			Handler:    _Message_SendLoginVerifyCodeMessage_Handler,
		},
		{
			MethodName: "GetLoginVerifyCodeMessage",
			Handler:    _Message_GetLoginVerifyCodeMessage_Handler,
		},
		{
			MethodName: "SendModifyPasswdVerifyCodeMessage",
			Handler:    _Message_SendModifyPasswdVerifyCodeMessage_Handler,
		},
		{
			MethodName: "GetModifyPasswdVerifyCodeMessage",
			Handler:    _Message_GetModifyPasswdVerifyCodeMessage_Handler,
		},
		{
			MethodName: "SendRegisterVerifyCodeMessage",
			Handler:    _Message_SendRegisterVerifyCodeMessage_Handler,
		},
		{
			MethodName: "GetRegisterVerifyCodeMessage",
			Handler:    _Message_GetRegisterVerifyCodeMessage_Handler,
		},
		{
			MethodName: "SendAlertNotifyMessage",
			Handler:    _Message_SendAlertNotifyMessage_Handler,
		},
		{
			MethodName: "SendKubernetesEventNotifyMessage",
			Handler:    _Message_SendKubernetesEventNotifyMessage_Handler,
		},
		{
			MethodName: "SendOrgDingTalkWorkMsg",
			Handler:    _Message_SendOrgDingTalkWorkMsg_Handler,
		},
		{
			MethodName: "CreateMessageGroup",
			Handler:    _Message_CreateMessageGroup_Handler,
		},
		{
			MethodName: "ListMessageGroup",
			Handler:    _Message_ListMessageGroup_Handler,
		},
		{
			MethodName: "DelMessageGroup",
			Handler:    _Message_DelMessageGroup_Handler,
		},
		{
			MethodName: "AddMessageGroupMember",
			Handler:    _Message_AddMessageGroupMember_Handler,
		},
		{
			MethodName: "GetMessageGroupMember",
			Handler:    _Message_GetMessageGroupMember_Handler,
		},
		{
			MethodName: "DelMessageGroupMember",
			Handler:    _Message_DelMessageGroupMember_Handler,
		},
		{
			MethodName: "SendMessageGroupContent",
			Handler:    _Message_SendMessageGroupContent_Handler,
		},
		{
			MethodName: "Translate",
			Handler:    _Message_Translate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message/v1/message.proto",
}
