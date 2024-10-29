// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.17.3
// source: statement.proto

package statement

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

type StatementPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page         int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize     int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	PayinOrderID int64 `protobuf:"varint,3,opt,name=payinOrderID,proto3" json:"payinOrderID,omitempty"`
}

func (x *StatementPageRequest) Reset() {
	*x = StatementPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatementPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatementPageRequest) ProtoMessage() {}

func (x *StatementPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_statement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatementPageRequest.ProtoReflect.Descriptor instead.
func (*StatementPageRequest) Descriptor() ([]byte, []int) {
	return file_statement_proto_rawDescGZIP(), []int{0}
}

func (x *StatementPageRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *StatementPageRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *StatementPageRequest) GetPayinOrderID() int64 {
	if x != nil {
		return x.PayinOrderID
	}
	return 0
}

type StatementItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId         int64  `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	AccountId       int64  `protobuf:"varint,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
	AccountType     int32  `protobuf:"varint,3,opt,name=accountType,proto3" json:"accountType,omitempty"`
	FinanceCode     int32  `protobuf:"varint,4,opt,name=financeCode,proto3" json:"financeCode,omitempty"`
	Currency        string `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount          string `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Balance         string `protobuf:"bytes,7,opt,name=balance,proto3" json:"balance,omitempty"`
	PreviousBalance string `protobuf:"bytes,8,opt,name=previousBalance,proto3" json:"previousBalance,omitempty"`
	CreateTime      int64  `protobuf:"varint,9,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *StatementItem) Reset() {
	*x = StatementItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatementItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatementItem) ProtoMessage() {}

func (x *StatementItem) ProtoReflect() protoreflect.Message {
	mi := &file_statement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatementItem.ProtoReflect.Descriptor instead.
func (*StatementItem) Descriptor() ([]byte, []int) {
	return file_statement_proto_rawDescGZIP(), []int{1}
}

func (x *StatementItem) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *StatementItem) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *StatementItem) GetAccountType() int32 {
	if x != nil {
		return x.AccountType
	}
	return 0
}

func (x *StatementItem) GetFinanceCode() int32 {
	if x != nil {
		return x.FinanceCode
	}
	return 0
}

func (x *StatementItem) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *StatementItem) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *StatementItem) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *StatementItem) GetPreviousBalance() string {
	if x != nil {
		return x.PreviousBalance
	}
	return ""
}

func (x *StatementItem) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type StatementPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg        string           `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Items      []*StatementItem `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items,omitempty"`
	Pagination *Pagination      `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *StatementPageResponse) Reset() {
	*x = StatementPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatementPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatementPageResponse) ProtoMessage() {}

func (x *StatementPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_statement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatementPageResponse.ProtoReflect.Descriptor instead.
func (*StatementPageResponse) Descriptor() ([]byte, []int) {
	return file_statement_proto_rawDescGZIP(), []int{2}
}

func (x *StatementPageResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StatementPageResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *StatementPageResponse) GetItems() []*StatementItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *StatementPageResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize    int32 `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Page        int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	CurrentPage int32 `protobuf:"varint,3,opt,name=currentPage,proto3" json:"currentPage,omitempty"`
	Total       int32 `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_statement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_statement_proto_rawDescGZIP(), []int{3}
}

func (x *Pagination) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Pagination) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *Pagination) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_statement_proto protoreflect.FileDescriptor

var file_statement_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x6a, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61,
	0x79, 0x69, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0xa3,
	0x02, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x34, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x74, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x69, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x09, 0x50, 0x61,
	0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x25, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_statement_proto_rawDescOnce sync.Once
	file_statement_proto_rawDescData = file_statement_proto_rawDesc
)

func file_statement_proto_rawDescGZIP() []byte {
	file_statement_proto_rawDescOnce.Do(func() {
		file_statement_proto_rawDescData = protoimpl.X.CompressGZIP(file_statement_proto_rawDescData)
	})
	return file_statement_proto_rawDescData
}

var file_statement_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_statement_proto_goTypes = []interface{}{
	(*StatementPageRequest)(nil),  // 0: order.statement.StatementPageRequest
	(*StatementItem)(nil),         // 1: order.statement.StatementItem
	(*StatementPageResponse)(nil), // 2: order.statement.StatementPageResponse
	(*Pagination)(nil),            // 3: order.statement.Pagination
}
var file_statement_proto_depIdxs = []int32{
	1, // 0: order.statement.StatementPageResponse.Items:type_name -> order.statement.StatementItem
	3, // 1: order.statement.StatementPageResponse.pagination:type_name -> order.statement.Pagination
	0, // 2: order.statement.statement.PageQuery:input_type -> order.statement.StatementPageRequest
	2, // 3: order.statement.statement.PageQuery:output_type -> order.statement.StatementPageResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_statement_proto_init() }
func file_statement_proto_init() {
	if File_statement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_statement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatementPageRequest); i {
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
		file_statement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatementItem); i {
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
		file_statement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatementPageResponse); i {
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
		file_statement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
			RawDescriptor: file_statement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_statement_proto_goTypes,
		DependencyIndexes: file_statement_proto_depIdxs,
		MessageInfos:      file_statement_proto_msgTypes,
	}.Build()
	File_statement_proto = out.File
	file_statement_proto_rawDesc = nil
	file_statement_proto_goTypes = nil
	file_statement_proto_depIdxs = nil
}
