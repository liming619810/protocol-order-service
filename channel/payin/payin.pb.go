// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.17.3
// source: payin.proto

package payin

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

type PayinQueryToExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	MerchantNo  int64 `protobuf:"varint,2,opt,name=merchantNo,proto3" json:"merchantNo,omitempty"`
	CreateBegin int64 `protobuf:"varint,3,opt,name=createBegin,proto3" json:"createBegin,omitempty"`
}

func (x *PayinQueryToExecuteRequest) Reset() {
	*x = PayinQueryToExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayinQueryToExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayinQueryToExecuteRequest) ProtoMessage() {}

func (x *PayinQueryToExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayinQueryToExecuteRequest.ProtoReflect.Descriptor instead.
func (*PayinQueryToExecuteRequest) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{0}
}

func (x *PayinQueryToExecuteRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PayinQueryToExecuteRequest) GetMerchantNo() int64 {
	if x != nil {
		return x.MerchantNo
	}
	return 0
}

func (x *PayinQueryToExecuteRequest) GetCreateBegin() int64 {
	if x != nil {
		return x.CreateBegin
	}
	return 0
}

type PayinExcuteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId           int64  `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	OutTradeNo        string `protobuf:"bytes,2,opt,name=outTradeNo,proto3" json:"outTradeNo,omitempty"`
	Status            uint32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	TargetOrg         string `protobuf:"bytes,4,opt,name=targetOrg,proto3" json:"targetOrg,omitempty"`
	PayAmount         uint32 `protobuf:"varint,5,opt,name=payAmount,proto3" json:"payAmount,omitempty"`
	PaymentMethod     string `protobuf:"bytes,6,opt,name=paymentMethod,proto3" json:"paymentMethod,omitempty"`
	PayCurrency       string `protobuf:"bytes,7,opt,name=payCurrency,proto3" json:"payCurrency,omitempty"`
	PaymentMethodType string `protobuf:"bytes,8,opt,name=paymentMethodType,proto3" json:"paymentMethodType,omitempty"`
	Country           string `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`
	Subject           string `protobuf:"bytes,10,opt,name=subject,proto3" json:"subject,omitempty"`
	CreateTime        int64  `protobuf:"varint,11,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *PayinExcuteInfo) Reset() {
	*x = PayinExcuteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayinExcuteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayinExcuteInfo) ProtoMessage() {}

func (x *PayinExcuteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayinExcuteInfo.ProtoReflect.Descriptor instead.
func (*PayinExcuteInfo) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{1}
}

func (x *PayinExcuteInfo) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *PayinExcuteInfo) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

func (x *PayinExcuteInfo) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PayinExcuteInfo) GetTargetOrg() string {
	if x != nil {
		return x.TargetOrg
	}
	return ""
}

func (x *PayinExcuteInfo) GetPayAmount() uint32 {
	if x != nil {
		return x.PayAmount
	}
	return 0
}

func (x *PayinExcuteInfo) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *PayinExcuteInfo) GetPayCurrency() string {
	if x != nil {
		return x.PayCurrency
	}
	return ""
}

func (x *PayinExcuteInfo) GetPaymentMethodType() string {
	if x != nil {
		return x.PaymentMethodType
	}
	return ""
}

func (x *PayinExcuteInfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *PayinExcuteInfo) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *PayinExcuteInfo) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type PayinQueryToExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *BaseResponse                      `protobuf:"bytes,1,opt,name=baseResponse,proto3" json:"baseResponse,omitempty"`
	Items        *PayinQueryToExecuteResponse_Items `protobuf:"bytes,2,opt,name=items,proto3" json:"items,omitempty"`
}

func (x *PayinQueryToExecuteResponse) Reset() {
	*x = PayinQueryToExecuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayinQueryToExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayinQueryToExecuteResponse) ProtoMessage() {}

func (x *PayinQueryToExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayinQueryToExecuteResponse.ProtoReflect.Descriptor instead.
func (*PayinQueryToExecuteResponse) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{2}
}

func (x *PayinQueryToExecuteResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *PayinQueryToExecuteResponse) GetItems() *PayinQueryToExecuteResponse_Items {
	if x != nil {
		return x.Items
	}
	return nil
}

type PayinQueryPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page          int32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32   `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderStatuses []int32 `protobuf:"varint,3,rep,packed,name=orderStatuses,proto3" json:"orderStatuses,omitempty"` // [0, 1, 2] 数组格式
	MerchantNo    string  `protobuf:"bytes,4,opt,name=merchantNo,proto3" json:"merchantNo,omitempty"`
	OutTradeNo    string  `protobuf:"bytes,5,opt,name=outTradeNo,proto3" json:"outTradeNo,omitempty"`
	UserId        int64   `protobuf:"varint,6,opt,name=userId,proto3" json:"userId,omitempty"`   // 商户下的用户id
	OrderId       int64   `protobuf:"varint,7,opt,name=orderId,proto3" json:"orderId,omitempty"` // 订单ID
	CreateBegin   int64   `protobuf:"varint,8,opt,name=createBegin,proto3" json:"createBegin,omitempty"`
	CreateEnd     int64   `protobuf:"varint,9,opt,name=createEnd,proto3" json:"createEnd,omitempty"`
	UpdateBegin   int64   `protobuf:"varint,10,opt,name=updateBegin,proto3" json:"updateBegin,omitempty"`
	UpdateEnd     int64   `protobuf:"varint,11,opt,name=updateEnd,proto3" json:"updateEnd,omitempty"`
}

func (x *PayinQueryPageRequest) Reset() {
	*x = PayinQueryPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayinQueryPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayinQueryPageRequest) ProtoMessage() {}

func (x *PayinQueryPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayinQueryPageRequest.ProtoReflect.Descriptor instead.
func (*PayinQueryPageRequest) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{3}
}

func (x *PayinQueryPageRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PayinQueryPageRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PayinQueryPageRequest) GetOrderStatuses() []int32 {
	if x != nil {
		return x.OrderStatuses
	}
	return nil
}

func (x *PayinQueryPageRequest) GetMerchantNo() string {
	if x != nil {
		return x.MerchantNo
	}
	return ""
}

func (x *PayinQueryPageRequest) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

func (x *PayinQueryPageRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PayinQueryPageRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *PayinQueryPageRequest) GetCreateBegin() int64 {
	if x != nil {
		return x.CreateBegin
	}
	return 0
}

func (x *PayinQueryPageRequest) GetCreateEnd() int64 {
	if x != nil {
		return x.CreateEnd
	}
	return 0
}

func (x *PayinQueryPageRequest) GetUpdateBegin() int64 {
	if x != nil {
		return x.UpdateBegin
	}
	return 0
}

func (x *PayinQueryPageRequest) GetUpdateEnd() int64 {
	if x != nil {
		return x.UpdateEnd
	}
	return 0
}

type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"` // 返回码
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`    // 错误信息
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{4}
}

func (x *BaseResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BaseResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PayinQueryToExecuteResponse_Items struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayinExcuteInfo []*PayinExcuteInfo `protobuf:"bytes,1,rep,name=payinExcuteInfo,proto3" json:"payinExcuteInfo,omitempty"`
}

func (x *PayinQueryToExecuteResponse_Items) Reset() {
	*x = PayinQueryToExecuteResponse_Items{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayinQueryToExecuteResponse_Items) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayinQueryToExecuteResponse_Items) ProtoMessage() {}

func (x *PayinQueryToExecuteResponse_Items) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayinQueryToExecuteResponse_Items.ProtoReflect.Descriptor instead.
func (*PayinQueryToExecuteResponse_Items) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{2, 0}
}

func (x *PayinQueryToExecuteResponse_Items) GetPayinExcuteInfo() []*PayinExcuteInfo {
	if x != nil {
		return x.PayinExcuteInfo
	}
	return nil
}

var File_payin_proto protoreflect.FileDescriptor

var file_payin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x22, 0x74, 0x0a, 0x1a, 0x50, 0x61,
	0x79, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4e, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e,
	0x22, 0xe9, 0x02, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x69, 0x6e, 0x45, 0x78, 0x63, 0x75, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x4f, 0x72, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x4f, 0x72, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x79, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x61, 0x79, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xf3, 0x01, 0x0a,
	0x1b, 0x50, 0x61, 0x79, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x69, 0x6e,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x79, 0x69, 0x6e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x1a, 0x4f, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x46, 0x0a, 0x0f, 0x70, 0x61,
	0x79, 0x69, 0x6e, 0x45, 0x78, 0x63, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x69,
	0x6e, 0x2e, 0x50, 0x61, 0x79, 0x69, 0x6e, 0x45, 0x78, 0x63, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0f, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x45, 0x78, 0x63, 0x75, 0x74, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0xdf, 0x02, 0x0a, 0x15, 0x50, 0x61, 0x79, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x4e, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x4e, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x4e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65,
	0x67, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x65,
	0x67, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x64, 0x22, 0x34, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x6e, 0x0a, 0x05, 0x70, 0x61,
	0x79, 0x69, 0x6e, 0x12, 0x65, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61,
	0x79, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x79, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x79,
	0x69, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e,
	0x2f, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payin_proto_rawDescOnce sync.Once
	file_payin_proto_rawDescData = file_payin_proto_rawDesc
)

func file_payin_proto_rawDescGZIP() []byte {
	file_payin_proto_rawDescOnce.Do(func() {
		file_payin_proto_rawDescData = protoimpl.X.CompressGZIP(file_payin_proto_rawDescData)
	})
	return file_payin_proto_rawDescData
}

var file_payin_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_payin_proto_goTypes = []interface{}{
	(*PayinQueryToExecuteRequest)(nil),        // 0: order.payin.PayinQueryToExecuteRequest
	(*PayinExcuteInfo)(nil),                   // 1: order.payin.PayinExcuteInfo
	(*PayinQueryToExecuteResponse)(nil),       // 2: order.payin.PayinQueryToExecuteResponse
	(*PayinQueryPageRequest)(nil),             // 3: order.payin.PayinQueryPageRequest
	(*BaseResponse)(nil),                      // 4: order.payin.BaseResponse
	(*PayinQueryToExecuteResponse_Items)(nil), // 5: order.payin.PayinQueryToExecuteResponse.Items
}
var file_payin_proto_depIdxs = []int32{
	4, // 0: order.payin.PayinQueryToExecuteResponse.baseResponse:type_name -> order.payin.BaseResponse
	5, // 1: order.payin.PayinQueryToExecuteResponse.items:type_name -> order.payin.PayinQueryToExecuteResponse.Items
	1, // 2: order.payin.PayinQueryToExecuteResponse.Items.payinExcuteInfo:type_name -> order.payin.PayinExcuteInfo
	0, // 3: order.payin.payin.QueryToExecute:input_type -> order.payin.PayinQueryToExecuteRequest
	2, // 4: order.payin.payin.QueryToExecute:output_type -> order.payin.PayinQueryToExecuteResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_payin_proto_init() }
func file_payin_proto_init() {
	if File_payin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayinQueryToExecuteRequest); i {
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
		file_payin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayinExcuteInfo); i {
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
		file_payin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayinQueryToExecuteResponse); i {
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
		file_payin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayinQueryPageRequest); i {
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
		file_payin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResponse); i {
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
		file_payin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayinQueryToExecuteResponse_Items); i {
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
			RawDescriptor: file_payin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payin_proto_goTypes,
		DependencyIndexes: file_payin_proto_depIdxs,
		MessageInfos:      file_payin_proto_msgTypes,
	}.Build()
	File_payin_proto = out.File
	file_payin_proto_rawDesc = nil
	file_payin_proto_goTypes = nil
	file_payin_proto_depIdxs = nil
}
