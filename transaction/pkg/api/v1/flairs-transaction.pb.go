// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: flairs-transaction.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TransactionType int32

const (
	Transaction_EXPENSE TransactionType = 0
	Transaction_INCOME  TransactionType = 1
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "EXPENSE",
		1: "INCOME",
	}
	TransactionType_value = map[string]int32{
		"EXPENSE": 0,
		"INCOME":  1,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_flairs_transaction_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_flairs_transaction_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_flairs_transaction_proto_rawDescGZIP(), []int{0, 0}
}

type TransactionClass int32

const (
	Transaction_Fund     TransactionClass = 0
	Transaction_Transfer TransactionClass = 1
)

// Enum value maps for TransactionClass.
var (
	TransactionClass_name = map[int32]string{
		0: "Fund",
		1: "Transfer",
	}
	TransactionClass_value = map[string]int32{
		"Fund":     0,
		"Transfer": 1,
	}
)

func (x TransactionClass) Enum() *TransactionClass {
	p := new(TransactionClass)
	*p = x
	return p
}

func (x TransactionClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionClass) Descriptor() protoreflect.EnumDescriptor {
	return file_flairs_transaction_proto_enumTypes[1].Descriptor()
}

func (TransactionClass) Type() protoreflect.EnumType {
	return &file_flairs_transaction_proto_enumTypes[1]
}

func (x TransactionClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionClass.Descriptor instead.
func (TransactionClass) EnumDescriptor() ([]byte, []int) {
	return file_flairs_transaction_proto_rawDescGZIP(), []int{0, 1}
}

type NewTransactionReqClass int32

const (
	NewTransactionReq_Fund     NewTransactionReqClass = 0
	NewTransactionReq_Transfer NewTransactionReqClass = 1
)

// Enum value maps for NewTransactionReqClass.
var (
	NewTransactionReqClass_name = map[int32]string{
		0: "Fund",
		1: "Transfer",
	}
	NewTransactionReqClass_value = map[string]int32{
		"Fund":     0,
		"Transfer": 1,
	}
)

func (x NewTransactionReqClass) Enum() *NewTransactionReqClass {
	p := new(NewTransactionReqClass)
	*p = x
	return p
}

func (x NewTransactionReqClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NewTransactionReqClass) Descriptor() protoreflect.EnumDescriptor {
	return file_flairs_transaction_proto_enumTypes[2].Descriptor()
}

func (NewTransactionReqClass) Type() protoreflect.EnumType {
	return &file_flairs_transaction_proto_enumTypes[2]
}

func (x NewTransactionReqClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NewTransactionReqClass.Descriptor instead.
func (NewTransactionReqClass) EnumDescriptor() ([]byte, []int) {
	return file_flairs_transaction_proto_rawDescGZIP(), []int{1, 0}
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                string          `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Amount            string          `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Memo              string          `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
	Message           string          `protobuf:"bytes,12,opt,name=message,proto3" json:"message,omitempty"`
	WalletId          string          `protobuf:"bytes,8,opt,name=walletId,proto3" json:"walletId,omitempty"`
	Status            string          `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Currency          string          `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	TxRef             string          `protobuf:"bytes,3,opt,name=txRef,proto3" json:"txRef,omitempty"`
	FlwRef            string          `protobuf:"bytes,7,opt,name=flwRef,proto3" json:"flwRef,omitempty"`
	Customer          string          `protobuf:"bytes,10,opt,name=customer,proto3" json:"customer,omitempty"`
	CustomerId        string          `protobuf:"bytes,11,opt,name=customerId,proto3" json:"customerId,omitempty"`
	PaymentType       string          `protobuf:"bytes,13,opt,name=paymentType,proto3" json:"paymentType,omitempty"`
	CreatedAt         string          `protobuf:"bytes,14,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	DeletedAt         string          `protobuf:"bytes,17,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	UpdatedAt         string          `protobuf:"bytes,9,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	CardLastFourDigit string          `protobuf:"bytes,15,opt,name=cardLastFourDigit,proto3" json:"cardLastFourDigit,omitempty"`
	CardType          string          `protobuf:"bytes,16,opt,name=cardType,proto3" json:"cardType,omitempty"`
	ThirdPartyID      string          `protobuf:"bytes,20,opt,name=thirdPartyID,proto3" json:"thirdPartyID,omitempty"`
	Source            string          `protobuf:"bytes,21,opt,name=source,proto3" json:"source,omitempty"`
	TransType         TransactionType `protobuf:"varint,19,opt,name=transType,proto3,enum=v1.TransactionType" json:"transType,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flairs_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_flairs_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_flairs_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Transaction) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Transaction) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Transaction) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Transaction) GetWalletId() string {
	if x != nil {
		return x.WalletId
	}
	return ""
}

func (x *Transaction) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Transaction) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Transaction) GetTxRef() string {
	if x != nil {
		return x.TxRef
	}
	return ""
}

func (x *Transaction) GetFlwRef() string {
	if x != nil {
		return x.FlwRef
	}
	return ""
}

func (x *Transaction) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

func (x *Transaction) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Transaction) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

func (x *Transaction) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Transaction) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Transaction) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Transaction) GetCardLastFourDigit() string {
	if x != nil {
		return x.CardLastFourDigit
	}
	return ""
}

func (x *Transaction) GetCardType() string {
	if x != nil {
		return x.CardType
	}
	return ""
}

func (x *Transaction) GetThirdPartyID() string {
	if x != nil {
		return x.ThirdPartyID
	}
	return ""
}

func (x *Transaction) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Transaction) GetTransType() TransactionType {
	if x != nil {
		return x.TransType
	}
	return Transaction_EXPENSE
}

type NewTransactionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThirdPartyID    string                 `protobuf:"bytes,4,opt,name=thirdPartyID,proto3" json:"thirdPartyID,omitempty"`
	Source          string                 `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
	UserMemo        string                 `protobuf:"bytes,2,opt,name=userMemo,proto3" json:"userMemo,omitempty"`
	InnerMemo       string                 `protobuf:"bytes,9,opt,name=innerMemo,proto3" json:"innerMemo,omitempty"`
	FromID          string                 `protobuf:"bytes,1,opt,name=fromID,proto3" json:"fromID,omitempty"`
	ToID            string                 `protobuf:"bytes,7,opt,name=toID,proto3" json:"toID,omitempty"`
	UserId          string                 `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Amount          float64                `protobuf:"fixed64,6,opt,name=amount,proto3" json:"amount,omitempty"`
	TransactionType NewTransactionReqClass `protobuf:"varint,5,opt,name=transactionType,proto3,enum=v1.NewTransactionReqClass" json:"transactionType,omitempty"`
}

func (x *NewTransactionReq) Reset() {
	*x = NewTransactionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flairs_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTransactionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTransactionReq) ProtoMessage() {}

func (x *NewTransactionReq) ProtoReflect() protoreflect.Message {
	mi := &file_flairs_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTransactionReq.ProtoReflect.Descriptor instead.
func (*NewTransactionReq) Descriptor() ([]byte, []int) {
	return file_flairs_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *NewTransactionReq) GetThirdPartyID() string {
	if x != nil {
		return x.ThirdPartyID
	}
	return ""
}

func (x *NewTransactionReq) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *NewTransactionReq) GetUserMemo() string {
	if x != nil {
		return x.UserMemo
	}
	return ""
}

func (x *NewTransactionReq) GetInnerMemo() string {
	if x != nil {
		return x.InnerMemo
	}
	return ""
}

func (x *NewTransactionReq) GetFromID() string {
	if x != nil {
		return x.FromID
	}
	return ""
}

func (x *NewTransactionReq) GetToID() string {
	if x != nil {
		return x.ToID
	}
	return ""
}

func (x *NewTransactionReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *NewTransactionReq) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *NewTransactionReq) GetTransactionType() NewTransactionReqClass {
	if x != nil {
		return x.TransactionType
	}
	return NewTransactionReq_Fund
}

type NewTransactionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,4,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *NewTransactionRes) Reset() {
	*x = NewTransactionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flairs_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTransactionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTransactionRes) ProtoMessage() {}

func (x *NewTransactionRes) ProtoReflect() protoreflect.Message {
	mi := &file_flairs_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTransactionRes.ProtoReflect.Descriptor instead.
func (*NewTransactionRes) Descriptor() ([]byte, []int) {
	return file_flairs_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *NewTransactionRes) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetMyTransactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetMyTransactionsRequest) Reset() {
	*x = GetMyTransactionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flairs_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyTransactionsRequest) ProtoMessage() {}

func (x *GetMyTransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flairs_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyTransactionsRequest.ProtoReflect.Descriptor instead.
func (*GetMyTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_flairs_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *GetMyTransactionsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type TransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *TransactionsResponse) Reset() {
	*x = TransactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flairs_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsResponse) ProtoMessage() {}

func (x *TransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flairs_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsResponse.ProtoReflect.Descriptor instead.
func (*TransactionsResponse) Descriptor() ([]byte, []int) {
	return file_flairs_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionsResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_flairs_transaction_proto protoreflect.FileDescriptor

var file_flairs_transaction_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x6c, 0x61, 0x69, 0x72, 0x73, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x05, 0x0a, 0x0b,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x78, 0x52, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x78, 0x52, 0x65, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6c, 0x77, 0x52, 0x65,
	0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6c, 0x77, 0x52, 0x65, 0x66, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x61, 0x72, 0x64, 0x4c,
	0x61, 0x73, 0x74, 0x46, 0x6f, 0x75, 0x72, 0x44, 0x69, 0x67, 0x69, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x63, 0x61, 0x72, 0x64, 0x4c, 0x61, 0x73, 0x74, 0x46, 0x6f, 0x75, 0x72,
	0x44, 0x69, 0x67, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50,
	0x45, 0x4e, 0x53, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x45,
	0x10, 0x01, 0x22, 0x1f, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x46,
	0x75, 0x6e, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x10, 0x01, 0x22, 0xcd, 0x02, 0x0a, 0x11, 0x6e, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x6d,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x6d,
	0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x6d, 0x6f, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x6d, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x49, 0x44, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x0f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x1f, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x46,
	0x75, 0x6e, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x10, 0x01, 0x22, 0x23, 0x0a, 0x11, 0x6e, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x32, 0x0a, 0x18, 0x67, 0x65, 0x74, 0x4d,
	0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x14,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xe8, 0x01, 0x0a, 0x18, 0x46, 0x6c,
	0x61, 0x69, 0x72, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x6e, 0x65, 0x77,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x76, 0x31,
	0x2e, 0x6e, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x6d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x76, 0x31, 0x2e,
	0x67, 0x65, 0x74, 0x4d, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x7d, 0x42, 0xb6, 0x02, 0x92, 0x41, 0xb2, 0x02, 0x12, 0x6a, 0x0a, 0x0e, 0x46,
	0x6c, 0x61, 0x69, 0x72, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x53, 0x0a,
	0x1c, 0x66, 0x6c, 0x61, 0x69, 0x72, 0x73, 0x20, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x20, 0x62,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x68,
	0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x70,
	0x6c, 0x75, 0x73, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x1a, 0x14, 0x64, 0x61,
	0x6d, 0x6d, 0x79, 0x64, 0x61, 0x72, 0x6d, 0x79, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63,
	0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f,
	0x73, 0x74, 0x3a, 0x39, 0x30, 0x30, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52,
	0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x5a, 0x4f, 0x0a, 0x4d,
	0x0a, 0x0c, 0x66, 0x6c, 0x61, 0x69, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3d,
	0x08, 0x02, 0x12, 0x28, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64,
	0x20, 0x69, 0x6e, 0x74, 0x6f, 0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x0d, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flairs_transaction_proto_rawDescOnce sync.Once
	file_flairs_transaction_proto_rawDescData = file_flairs_transaction_proto_rawDesc
)

func file_flairs_transaction_proto_rawDescGZIP() []byte {
	file_flairs_transaction_proto_rawDescOnce.Do(func() {
		file_flairs_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_flairs_transaction_proto_rawDescData)
	})
	return file_flairs_transaction_proto_rawDescData
}

var file_flairs_transaction_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_flairs_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_flairs_transaction_proto_goTypes = []interface{}{
	(TransactionType)(0),             // 0: v1.Transaction.type
	(TransactionClass)(0),            // 1: v1.Transaction.class
	(NewTransactionReqClass)(0),      // 2: v1.newTransactionReq.class
	(*Transaction)(nil),              // 3: v1.Transaction
	(*NewTransactionReq)(nil),        // 4: v1.newTransactionReq
	(*NewTransactionRes)(nil),        // 5: v1.newTransactionRes
	(*GetMyTransactionsRequest)(nil), // 6: v1.getMyTransactionsRequest
	(*TransactionsResponse)(nil),     // 7: v1.TransactionsResponse
}
var file_flairs_transaction_proto_depIdxs = []int32{
	0, // 0: v1.Transaction.transType:type_name -> v1.Transaction.type
	2, // 1: v1.newTransactionReq.transactionType:type_name -> v1.newTransactionReq.class
	3, // 2: v1.TransactionsResponse.transactions:type_name -> v1.Transaction
	4, // 3: v1.FlairsTransactionService.AddnewTransaction:input_type -> v1.newTransactionReq
	6, // 4: v1.FlairsTransactionService.GetMyTransactions:input_type -> v1.getMyTransactionsRequest
	5, // 5: v1.FlairsTransactionService.AddnewTransaction:output_type -> v1.newTransactionRes
	7, // 6: v1.FlairsTransactionService.GetMyTransactions:output_type -> v1.TransactionsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_flairs_transaction_proto_init() }
func file_flairs_transaction_proto_init() {
	if File_flairs_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flairs_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_flairs_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTransactionReq); i {
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
		file_flairs_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTransactionRes); i {
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
		file_flairs_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyTransactionsRequest); i {
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
		file_flairs_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionsResponse); i {
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
			RawDescriptor: file_flairs_transaction_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flairs_transaction_proto_goTypes,
		DependencyIndexes: file_flairs_transaction_proto_depIdxs,
		EnumInfos:         file_flairs_transaction_proto_enumTypes,
		MessageInfos:      file_flairs_transaction_proto_msgTypes,
	}.Build()
	File_flairs_transaction_proto = out.File
	file_flairs_transaction_proto_rawDesc = nil
	file_flairs_transaction_proto_goTypes = nil
	file_flairs_transaction_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FlairsTransactionServiceClient is the client API for FlairsTransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlairsTransactionServiceClient interface {
	// Create a new transaction
	AddnewTransaction(ctx context.Context, in *NewTransactionReq, opts ...grpc.CallOption) (*NewTransactionRes, error)
	// Get wallets beloging to a particular user
	GetMyTransactions(ctx context.Context, in *GetMyTransactionsRequest, opts ...grpc.CallOption) (*TransactionsResponse, error)
}

type flairsTransactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlairsTransactionServiceClient(cc grpc.ClientConnInterface) FlairsTransactionServiceClient {
	return &flairsTransactionServiceClient{cc}
}

func (c *flairsTransactionServiceClient) AddnewTransaction(ctx context.Context, in *NewTransactionReq, opts ...grpc.CallOption) (*NewTransactionRes, error) {
	out := new(NewTransactionRes)
	err := c.cc.Invoke(ctx, "/v1.FlairsTransactionService/AddnewTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flairsTransactionServiceClient) GetMyTransactions(ctx context.Context, in *GetMyTransactionsRequest, opts ...grpc.CallOption) (*TransactionsResponse, error) {
	out := new(TransactionsResponse)
	err := c.cc.Invoke(ctx, "/v1.FlairsTransactionService/GetMyTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlairsTransactionServiceServer is the server API for FlairsTransactionService service.
type FlairsTransactionServiceServer interface {
	// Create a new transaction
	AddnewTransaction(context.Context, *NewTransactionReq) (*NewTransactionRes, error)
	// Get wallets beloging to a particular user
	GetMyTransactions(context.Context, *GetMyTransactionsRequest) (*TransactionsResponse, error)
}

// UnimplementedFlairsTransactionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFlairsTransactionServiceServer struct {
}

func (*UnimplementedFlairsTransactionServiceServer) AddnewTransaction(context.Context, *NewTransactionReq) (*NewTransactionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddnewTransaction not implemented")
}
func (*UnimplementedFlairsTransactionServiceServer) GetMyTransactions(context.Context, *GetMyTransactionsRequest) (*TransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyTransactions not implemented")
}

func RegisterFlairsTransactionServiceServer(s *grpc.Server, srv FlairsTransactionServiceServer) {
	s.RegisterService(&_FlairsTransactionService_serviceDesc, srv)
}

func _FlairsTransactionService_AddnewTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTransactionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlairsTransactionServiceServer).AddnewTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.FlairsTransactionService/AddnewTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlairsTransactionServiceServer).AddnewTransaction(ctx, req.(*NewTransactionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlairsTransactionService_GetMyTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlairsTransactionServiceServer).GetMyTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.FlairsTransactionService/GetMyTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlairsTransactionServiceServer).GetMyTransactions(ctx, req.(*GetMyTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FlairsTransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.FlairsTransactionService",
	HandlerType: (*FlairsTransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddnewTransaction",
			Handler:    _FlairsTransactionService_AddnewTransaction_Handler,
		},
		{
			MethodName: "GetMyTransactions",
			Handler:    _FlairsTransactionService_GetMyTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flairs-transaction.proto",
}
