// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: models.proto

package ccipocr3pb

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

// CCIPMsg is a gRPC adapter to [github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3.CCIPMsg].
type CCIPMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SequenceNumber      uint64         `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	ChainFeeLimit       *BigInt        `protobuf:"bytes,2,opt,name=chain_fee_limit,json=chainFeeLimit,proto3" json:"chain_fee_limit,omitempty"`
	Nonce               uint64         `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ChainFeePrice       uint64         `protobuf:"varint,4,opt,name=chain_fee_price,json=chainFeePrice,proto3" json:"chain_fee_price,omitempty"`
	MessageId           []byte         `protobuf:"bytes,5,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"` // Hash [32]byte
	SourceChainSelector uint64         `protobuf:"varint,6,opt,name=source_chain_selector,json=sourceChainSelector,proto3" json:"source_chain_selector,omitempty"`
	Sender              string         `protobuf:"bytes,7,opt,name=sender,proto3" json:"sender,omitempty"`     // Address
	Receiver            string         `protobuf:"bytes,8,opt,name=receiver,proto3" json:"receiver,omitempty"` // Address
	Strict              bool           `protobuf:"varint,9,opt,name=strict,proto3" json:"strict,omitempty"`
	FeeToken            string         `protobuf:"bytes,10,opt,name=fee_token,json=feeToken,proto3" json:"fee_token,omitempty"` // Address
	FeeTokenAmount      *BigInt        `protobuf:"bytes,11,opt,name=fee_token_amount,json=feeTokenAmount,proto3" json:"fee_token_amount,omitempty"`
	Data                []byte         `protobuf:"bytes,12,opt,name=data,proto3" json:"data,omitempty"`
	TokenAmounts        []*TokenAmount `protobuf:"bytes,13,rep,name=token_amounts,json=tokenAmounts,proto3" json:"token_amounts,omitempty"`
	SourceTokenData     [][]byte       `protobuf:"bytes,14,rep,name=source_token_data,json=sourceTokenData,proto3" json:"source_token_data,omitempty"`
}

func (x *CCIPMsg) Reset() {
	*x = CCIPMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCIPMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCIPMsg) ProtoMessage() {}

func (x *CCIPMsg) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCIPMsg.ProtoReflect.Descriptor instead.
func (*CCIPMsg) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{0}
}

func (x *CCIPMsg) GetSequenceNumber() uint64 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

func (x *CCIPMsg) GetChainFeeLimit() *BigInt {
	if x != nil {
		return x.ChainFeeLimit
	}
	return nil
}

func (x *CCIPMsg) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *CCIPMsg) GetChainFeePrice() uint64 {
	if x != nil {
		return x.ChainFeePrice
	}
	return 0
}

func (x *CCIPMsg) GetMessageId() []byte {
	if x != nil {
		return x.MessageId
	}
	return nil
}

func (x *CCIPMsg) GetSourceChainSelector() uint64 {
	if x != nil {
		return x.SourceChainSelector
	}
	return 0
}

func (x *CCIPMsg) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *CCIPMsg) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *CCIPMsg) GetStrict() bool {
	if x != nil {
		return x.Strict
	}
	return false
}

func (x *CCIPMsg) GetFeeToken() string {
	if x != nil {
		return x.FeeToken
	}
	return ""
}

func (x *CCIPMsg) GetFeeTokenAmount() *BigInt {
	if x != nil {
		return x.FeeTokenAmount
	}
	return nil
}

func (x *CCIPMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CCIPMsg) GetTokenAmounts() []*TokenAmount {
	if x != nil {
		return x.TokenAmounts
	}
	return nil
}

func (x *CCIPMsg) GetSourceTokenData() [][]byte {
	if x != nil {
		return x.SourceTokenData
	}
	return nil
}

// BigInt represents a [big.Int].
type BigInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BigInt) Reset() {
	*x = BigInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigInt) ProtoMessage() {}

func (x *BigInt) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigInt.ProtoReflect.Descriptor instead.
func (*BigInt) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{1}
}

func (x *BigInt) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// TokenAmount is a helper type that defines a token and an amount.
type TokenAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string  `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // Token address
	Amount *BigInt `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TokenAmount) Reset() {
	*x = TokenAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenAmount) ProtoMessage() {}

func (x *TokenAmount) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenAmount.ProtoReflect.Descriptor instead.
func (*TokenAmount) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{2}
}

func (x *TokenAmount) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TokenAmount) GetAmount() *BigInt {
	if x != nil {
		return x.Amount
	}
	return nil
}

// CommitPluginReport is a gRPC adapter to [github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3.CommitPluginReport].
type CommitPluginReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PriceUpdates *PriceUpdates      `protobuf:"bytes,1,opt,name=price_updates,json=priceUpdates,proto3" json:"price_updates,omitempty"`
	MerkleRoots  []*MerkleRootChain `protobuf:"bytes,2,rep,name=merkle_roots,json=merkleRoots,proto3" json:"merkle_roots,omitempty"`
}

func (x *CommitPluginReport) Reset() {
	*x = CommitPluginReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitPluginReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitPluginReport) ProtoMessage() {}

func (x *CommitPluginReport) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitPluginReport.ProtoReflect.Descriptor instead.
func (*CommitPluginReport) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{3}
}

func (x *CommitPluginReport) GetPriceUpdates() *PriceUpdates {
	if x != nil {
		return x.PriceUpdates
	}
	return nil
}

func (x *CommitPluginReport) GetMerkleRoots() []*MerkleRootChain {
	if x != nil {
		return x.MerkleRoots
	}
	return nil
}

// PriceUpdates is a gRPC adapter to [github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3.PriceUpdates].
type PriceUpdates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenPriceUpdates []*TokenAmount   `protobuf:"bytes,1,rep,name=token_price_updates,json=tokenPriceUpdates,proto3" json:"token_price_updates,omitempty"`
	GasPriceUpdates   []*GasPriceChain `protobuf:"bytes,2,rep,name=gas_price_updates,json=gasPriceUpdates,proto3" json:"gas_price_updates,omitempty"`
}

func (x *PriceUpdates) Reset() {
	*x = PriceUpdates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceUpdates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceUpdates) ProtoMessage() {}

func (x *PriceUpdates) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceUpdates.ProtoReflect.Descriptor instead.
func (*PriceUpdates) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{4}
}

func (x *PriceUpdates) GetTokenPriceUpdates() []*TokenAmount {
	if x != nil {
		return x.TokenPriceUpdates
	}
	return nil
}

func (x *PriceUpdates) GetGasPriceUpdates() []*GasPriceChain {
	if x != nil {
		return x.GasPriceUpdates
	}
	return nil
}

// GasPriceChain is a gRPC adapter to [github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3.GasPriceChain].
type GasPriceChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainSelector uint64  `protobuf:"varint,1,opt,name=chain_selector,json=chainSelector,proto3" json:"chain_selector,omitempty"`
	Price         *BigInt `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GasPriceChain) Reset() {
	*x = GasPriceChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GasPriceChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GasPriceChain) ProtoMessage() {}

func (x *GasPriceChain) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GasPriceChain.ProtoReflect.Descriptor instead.
func (*GasPriceChain) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{5}
}

func (x *GasPriceChain) GetChainSelector() uint64 {
	if x != nil {
		return x.ChainSelector
	}
	return 0
}

func (x *GasPriceChain) GetPrice() *BigInt {
	if x != nil {
		return x.Price
	}
	return nil
}

// MerkleRootChain is a gRPC adapter to [github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3.MerkleRootChain].
type MerkleRootChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainSelector uint64       `protobuf:"varint,1,opt,name=chain_selector,json=chainSelector,proto3" json:"chain_selector,omitempty"`
	MerkleRoot    []byte       `protobuf:"bytes,2,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	SeqNumRange   *SeqNumRange `protobuf:"bytes,3,opt,name=seq_num_range,json=seqNumRange,proto3" json:"seq_num_range,omitempty"`
}

func (x *MerkleRootChain) Reset() {
	*x = MerkleRootChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleRootChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleRootChain) ProtoMessage() {}

func (x *MerkleRootChain) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleRootChain.ProtoReflect.Descriptor instead.
func (*MerkleRootChain) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{6}
}

func (x *MerkleRootChain) GetChainSelector() uint64 {
	if x != nil {
		return x.ChainSelector
	}
	return 0
}

func (x *MerkleRootChain) GetMerkleRoot() []byte {
	if x != nil {
		return x.MerkleRoot
	}
	return nil
}

func (x *MerkleRootChain) GetSeqNumRange() *SeqNumRange {
	if x != nil {
		return x.SeqNumRange
	}
	return nil
}

// SeqNumRange defines an inclusive range of sequence numbers.
type SeqNumRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start uint64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   uint64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *SeqNumRange) Reset() {
	*x = SeqNumRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeqNumRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeqNumRange) ProtoMessage() {}

func (x *SeqNumRange) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeqNumRange.ProtoReflect.Descriptor instead.
func (*SeqNumRange) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{7}
}

func (x *SeqNumRange) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *SeqNumRange) GetEnd() uint64 {
	if x != nil {
		return x.End
	}
	return 0
}

var File_models_proto protoreflect.FileDescriptor

var file_models_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62,
	0x2e, 0x63, 0x63, 0x69, 0x70, 0x6f, 0x63, 0x72, 0x33, 0x22, 0xd1, 0x04, 0x0a, 0x07, 0x43, 0x43,
	0x49, 0x50, 0x4d, 0x73, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x49,
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x6f,
	0x63, 0x72, 0x33, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x46, 0x65, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x46,
	0x65, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x65, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x4b, 0x0a, 0x10, 0x66, 0x65, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62,
	0x2e, 0x63, 0x63, 0x69, 0x70, 0x6f, 0x63, 0x72, 0x33, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74,
	0x52, 0x0e, 0x66, 0x65, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x4b, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63,
	0x63, 0x69, 0x70, 0x6f, 0x63, 0x72, 0x33, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a,
	0x06, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5e, 0x0a,
	0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x39, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x6f, 0x63, 0x72, 0x33, 0x2e, 0x42,
	0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb1, 0x01,
	0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x4c, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63,
	0x63, 0x69, 0x70, 0x6f, 0x63, 0x72, 0x33, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x12, 0x4d, 0x0a, 0x0c, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70,
	0x6f, 0x63, 0x72, 0x33, 0x2e, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x52, 0x0b, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74,
	0x73, 0x22, 0xbc, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x12, 0x56, 0x0a, 0x13, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x6f, 0x63, 0x72, 0x33, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x11, 0x67, 0x61,
	0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x6f, 0x63, 0x72,
	0x33, 0x2e, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52,
	0x0f, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x22, 0x6f, 0x0a, 0x0d, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x6f,
	0x63, 0x72, 0x33, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x22, 0xa5, 0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x4a, 0x0a,
	0x0d, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x6f, 0x63, 0x72, 0x33,
	0x2e, 0x53, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0b, 0x73, 0x65,
	0x71, 0x4e, 0x75, 0x6d, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x35, 0x0a, 0x0b, 0x53, 0x65, 0x71,
	0x4e, 0x75, 0x6d, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74, 0x2f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x63, 0x69, 0x70, 0x6f, 0x63, 0x72, 0x33, 0x3b, 0x63,
	0x63, 0x69, 0x70, 0x6f, 0x63, 0x72, 0x33, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_models_proto_rawDescOnce sync.Once
	file_models_proto_rawDescData = file_models_proto_rawDesc
)

func file_models_proto_rawDescGZIP() []byte {
	file_models_proto_rawDescOnce.Do(func() {
		file_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_proto_rawDescData)
	})
	return file_models_proto_rawDescData
}

var file_models_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_models_proto_goTypes = []interface{}{
	(*CCIPMsg)(nil),            // 0: loop.internal.pb.ccipocr3.CCIPMsg
	(*BigInt)(nil),             // 1: loop.internal.pb.ccipocr3.BigInt
	(*TokenAmount)(nil),        // 2: loop.internal.pb.ccipocr3.TokenAmount
	(*CommitPluginReport)(nil), // 3: loop.internal.pb.ccipocr3.CommitPluginReport
	(*PriceUpdates)(nil),       // 4: loop.internal.pb.ccipocr3.PriceUpdates
	(*GasPriceChain)(nil),      // 5: loop.internal.pb.ccipocr3.GasPriceChain
	(*MerkleRootChain)(nil),    // 6: loop.internal.pb.ccipocr3.MerkleRootChain
	(*SeqNumRange)(nil),        // 7: loop.internal.pb.ccipocr3.SeqNumRange
}
var file_models_proto_depIdxs = []int32{
	1,  // 0: loop.internal.pb.ccipocr3.CCIPMsg.chain_fee_limit:type_name -> loop.internal.pb.ccipocr3.BigInt
	1,  // 1: loop.internal.pb.ccipocr3.CCIPMsg.fee_token_amount:type_name -> loop.internal.pb.ccipocr3.BigInt
	2,  // 2: loop.internal.pb.ccipocr3.CCIPMsg.token_amounts:type_name -> loop.internal.pb.ccipocr3.TokenAmount
	1,  // 3: loop.internal.pb.ccipocr3.TokenAmount.amount:type_name -> loop.internal.pb.ccipocr3.BigInt
	4,  // 4: loop.internal.pb.ccipocr3.CommitPluginReport.price_updates:type_name -> loop.internal.pb.ccipocr3.PriceUpdates
	6,  // 5: loop.internal.pb.ccipocr3.CommitPluginReport.merkle_roots:type_name -> loop.internal.pb.ccipocr3.MerkleRootChain
	2,  // 6: loop.internal.pb.ccipocr3.PriceUpdates.token_price_updates:type_name -> loop.internal.pb.ccipocr3.TokenAmount
	5,  // 7: loop.internal.pb.ccipocr3.PriceUpdates.gas_price_updates:type_name -> loop.internal.pb.ccipocr3.GasPriceChain
	1,  // 8: loop.internal.pb.ccipocr3.GasPriceChain.price:type_name -> loop.internal.pb.ccipocr3.BigInt
	7,  // 9: loop.internal.pb.ccipocr3.MerkleRootChain.seq_num_range:type_name -> loop.internal.pb.ccipocr3.SeqNumRange
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_models_proto_init() }
func file_models_proto_init() {
	if File_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCIPMsg); i {
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
		file_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigInt); i {
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
		file_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenAmount); i {
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
		file_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitPluginReport); i {
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
		file_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceUpdates); i {
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
		file_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GasPriceChain); i {
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
		file_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleRootChain); i {
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
		file_models_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeqNumRange); i {
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
			RawDescriptor: file_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_proto_goTypes,
		DependencyIndexes: file_models_proto_depIdxs,
		MessageInfos:      file_models_proto_msgTypes,
	}.Build()
	File_models_proto = out.File
	file_models_proto_rawDesc = nil
	file_models_proto_goTypes = nil
	file_models_proto_depIdxs = nil
}