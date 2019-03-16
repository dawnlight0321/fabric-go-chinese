
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:36</date>
//</624456122538921984>

//由Protoc Gen Go生成的代码。不要编辑。
//来源：token/prover.proto

package token //导入“github.com/hyperledger/fabric/protos/token”

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

//引用导入以禁止错误（如果未使用）。
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//这是一个编译时断言，以确保生成的文件
//与正在编译的proto包兼容。
//此行的编译错误可能意味着您的
//需要更新proto包。
const _ = proto.ProtoPackageIsVersion2 //请升级proto包

//tokentoissue描述要在系统中颁发的令牌
type TokenToIssue struct {
//收件人是指要颁发的令牌的所有者
	Recipient []byte `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
//类型引用令牌类型
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
//数量是指要发行的代币单位的数量。
	Quantity             uint64   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenToIssue) Reset()         { *m = TokenToIssue{} }
func (m *TokenToIssue) String() string { return proto.CompactTextString(m) }
func (*TokenToIssue) ProtoMessage()    {}
func (*TokenToIssue) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{0}
}
func (m *TokenToIssue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenToIssue.Unmarshal(m, b)
}
func (m *TokenToIssue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenToIssue.Marshal(b, m, deterministic)
}
func (dst *TokenToIssue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenToIssue.Merge(dst, src)
}
func (m *TokenToIssue) XXX_Size() int {
	return xxx_messageInfo_TokenToIssue.Size(m)
}
func (m *TokenToIssue) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenToIssue.DiscardUnknown(m)
}

var xxx_messageInfo_TokenToIssue proto.InternalMessageInfo

func (m *TokenToIssue) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *TokenToIssue) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TokenToIssue) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

//RecipientTransfershare描述在令牌传输中收件人将收到多少信息
type RecipientTransferShare struct {
//接收者是指转让代币的潜在所有者。
	Recipient []byte `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
//数量是指要传输给接收者的令牌单元数。
	Quantity             uint64   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecipientTransferShare) Reset()         { *m = RecipientTransferShare{} }
func (m *RecipientTransferShare) String() string { return proto.CompactTextString(m) }
func (*RecipientTransferShare) ProtoMessage()    {}
func (*RecipientTransferShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{1}
}
func (m *RecipientTransferShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecipientTransferShare.Unmarshal(m, b)
}
func (m *RecipientTransferShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecipientTransferShare.Marshal(b, m, deterministic)
}
func (dst *RecipientTransferShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecipientTransferShare.Merge(dst, src)
}
func (m *RecipientTransferShare) XXX_Size() int {
	return xxx_messageInfo_RecipientTransferShare.Size(m)
}
func (m *RecipientTransferShare) XXX_DiscardUnknown() {
	xxx_messageInfo_RecipientTransferShare.DiscardUnknown(m)
}

var xxx_messageInfo_RecipientTransferShare proto.InternalMessageInfo

func (m *RecipientTransferShare) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *RecipientTransferShare) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

//tokenOutput用于指定ListRequest返回的令牌
type TokenOutput struct {
//ID用于唯一标识令牌
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
//类型是令牌的类型
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
//数量表示此类型令牌的编号
	Quantity             uint64   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenOutput) Reset()         { *m = TokenOutput{} }
func (m *TokenOutput) String() string { return proto.CompactTextString(m) }
func (*TokenOutput) ProtoMessage()    {}
func (*TokenOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{2}
}
func (m *TokenOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenOutput.Unmarshal(m, b)
}
func (m *TokenOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenOutput.Marshal(b, m, deterministic)
}
func (dst *TokenOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenOutput.Merge(dst, src)
}
func (m *TokenOutput) XXX_Size() int {
	return xxx_messageInfo_TokenOutput.Size(m)
}
func (m *TokenOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenOutput.DiscardUnknown(m)
}

var xxx_messageInfo_TokenOutput proto.InternalMessageInfo

func (m *TokenOutput) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TokenOutput) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TokenOutput) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

//未使用的令牌用于保存listRequest的输出
type UnspentTokens struct {
	Tokens               []*TokenOutput `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UnspentTokens) Reset()         { *m = UnspentTokens{} }
func (m *UnspentTokens) String() string { return proto.CompactTextString(m) }
func (*UnspentTokens) ProtoMessage()    {}
func (*UnspentTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{3}
}
func (m *UnspentTokens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnspentTokens.Unmarshal(m, b)
}
func (m *UnspentTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnspentTokens.Marshal(b, m, deterministic)
}
func (dst *UnspentTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnspentTokens.Merge(dst, src)
}
func (m *UnspentTokens) XXX_Size() int {
	return xxx_messageInfo_UnspentTokens.Size(m)
}
func (m *UnspentTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_UnspentTokens.DiscardUnknown(m)
}

var xxx_messageInfo_UnspentTokens proto.InternalMessageInfo

func (m *UnspentTokens) GetTokens() []*TokenOutput {
	if m != nil {
		return m.Tokens
	}
	return nil
}

//ListRequest用于请求未使用的令牌列表
type ListRequest struct {
	Credential           []byte   `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{4}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetCredential() []byte {
	if m != nil {
		return m.Credential
	}
	return nil
}

//importRequest用于请求创建导入
type ImportRequest struct {
//凭证包含有关请求操作的一方的信息
//此字段的内容取决于所使用的令牌管理器系统的特性。
	Credential []byte `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
//tokentoIssue包含有关要颁发的令牌的信息
	TokensToIssue        []*TokenToIssue `protobuf:"bytes,2,rep,name=tokens_to_issue,json=tokensToIssue,proto3" json:"tokens_to_issue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ImportRequest) Reset()         { *m = ImportRequest{} }
func (m *ImportRequest) String() string { return proto.CompactTextString(m) }
func (*ImportRequest) ProtoMessage()    {}
func (*ImportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{5}
}
func (m *ImportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportRequest.Unmarshal(m, b)
}
func (m *ImportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportRequest.Marshal(b, m, deterministic)
}
func (dst *ImportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportRequest.Merge(dst, src)
}
func (m *ImportRequest) XXX_Size() int {
	return xxx_messageInfo_ImportRequest.Size(m)
}
func (m *ImportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImportRequest proto.InternalMessageInfo

func (m *ImportRequest) GetCredential() []byte {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *ImportRequest) GetTokensToIssue() []*TokenToIssue {
	if m != nil {
		return m.TokensToIssue
	}
	return nil
}

//RequestTransfer用于请求创建传输
type TransferRequest struct {
	Credential           []byte                    `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	TokenIds             [][]byte                  `protobuf:"bytes,2,rep,name=token_ids,json=tokenIds,proto3" json:"token_ids,omitempty"`
	Shares               []*RecipientTransferShare `protobuf:"bytes,3,rep,name=shares,proto3" json:"shares,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TransferRequest) Reset()         { *m = TransferRequest{} }
func (m *TransferRequest) String() string { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()    {}
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{6}
}
func (m *TransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferRequest.Unmarshal(m, b)
}
func (m *TransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferRequest.Marshal(b, m, deterministic)
}
func (dst *TransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferRequest.Merge(dst, src)
}
func (m *TransferRequest) XXX_Size() int {
	return xxx_messageInfo_TransferRequest.Size(m)
}
func (m *TransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferRequest proto.InternalMessageInfo

func (m *TransferRequest) GetCredential() []byte {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *TransferRequest) GetTokenIds() [][]byte {
	if m != nil {
		return m.TokenIds
	}
	return nil
}

func (m *TransferRequest) GetShares() []*RecipientTransferShare {
	if m != nil {
		return m.Shares
	}
	return nil
}

//REDEEMREQUEST用于请求令牌赎回
type RedeemRequest struct {
//凭证包含请求操作的一方的信息
//此字段的内容取决于令牌管理器系统的特性
	Credential []byte `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
//令牌标识指定要兑换的令牌的标识
	TokenIds [][]byte `protobuf:"bytes,2,rep,name=token_ids,json=tokenIds,proto3" json:"token_ids,omitempty"`
//数量是指一个给定的代币需要赎回的单位数量。
	QuantityToRedeem     uint64   `protobuf:"varint,3,opt,name=quantity_to_redeem,json=quantityToRedeem,proto3" json:"quantity_to_redeem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedeemRequest) Reset()         { *m = RedeemRequest{} }
func (m *RedeemRequest) String() string { return proto.CompactTextString(m) }
func (*RedeemRequest) ProtoMessage()    {}
func (*RedeemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{7}
}
func (m *RedeemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedeemRequest.Unmarshal(m, b)
}
func (m *RedeemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedeemRequest.Marshal(b, m, deterministic)
}
func (dst *RedeemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedeemRequest.Merge(dst, src)
}
func (m *RedeemRequest) XXX_Size() int {
	return xxx_messageInfo_RedeemRequest.Size(m)
}
func (m *RedeemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RedeemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RedeemRequest proto.InternalMessageInfo

func (m *RedeemRequest) GetCredential() []byte {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *RedeemRequest) GetTokenIds() [][]byte {
	if m != nil {
		return m.TokenIds
	}
	return nil
}

func (m *RedeemRequest) GetQuantityToRedeem() uint64 {
	if m != nil {
		return m.QuantityToRedeem
	}
	return 0
}

//allowance定义了接收者可以代表其实际所有者转让的令牌的数量和内容。
type AllowanceRecipientShare struct {
//收件人是指允许使用由令牌ID标识的令牌中指定数量的实体。
	Recipient []byte `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
//Quantity是委托给收件人的令牌数
	Quantity             uint64   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllowanceRecipientShare) Reset()         { *m = AllowanceRecipientShare{} }
func (m *AllowanceRecipientShare) String() string { return proto.CompactTextString(m) }
func (*AllowanceRecipientShare) ProtoMessage()    {}
func (*AllowanceRecipientShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{8}
}
func (m *AllowanceRecipientShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllowanceRecipientShare.Unmarshal(m, b)
}
func (m *AllowanceRecipientShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllowanceRecipientShare.Marshal(b, m, deterministic)
}
func (dst *AllowanceRecipientShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowanceRecipientShare.Merge(dst, src)
}
func (m *AllowanceRecipientShare) XXX_Size() int {
	return xxx_messageInfo_AllowanceRecipientShare.Size(m)
}
func (m *AllowanceRecipientShare) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowanceRecipientShare.DiscardUnknown(m)
}

var xxx_messageInfo_AllowanceRecipientShare proto.InternalMessageInfo

func (m *AllowanceRecipientShare) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *AllowanceRecipientShare) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

//ApproveRequest用于请求从一个所有者到另一个所有者创建津贴。
type ApproveRequest struct {
//凭证是指请求创建者的公共凭证
	Credential []byte `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
//allowance描述请求创建者愿意委托的令牌
	AllowanceShares []*AllowanceRecipientShare `protobuf:"bytes,2,rep,name=allowance_shares,json=allowanceShares,proto3" json:"allowance_shares,omitempty"`
//tokenids是用于创建津贴的令牌标识符。
	TokenIds             [][]byte `protobuf:"bytes,3,rep,name=token_ids,json=tokenIds,proto3" json:"token_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApproveRequest) Reset()         { *m = ApproveRequest{} }
func (m *ApproveRequest) String() string { return proto.CompactTextString(m) }
func (*ApproveRequest) ProtoMessage()    {}
func (*ApproveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{9}
}
func (m *ApproveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApproveRequest.Unmarshal(m, b)
}
func (m *ApproveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApproveRequest.Marshal(b, m, deterministic)
}
func (dst *ApproveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApproveRequest.Merge(dst, src)
}
func (m *ApproveRequest) XXX_Size() int {
	return xxx_messageInfo_ApproveRequest.Size(m)
}
func (m *ApproveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApproveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApproveRequest proto.InternalMessageInfo

func (m *ApproveRequest) GetCredential() []byte {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *ApproveRequest) GetAllowanceShares() []*AllowanceRecipientShare {
	if m != nil {
		return m.AllowanceShares
	}
	return nil
}

func (m *ApproveRequest) GetTokenIds() [][]byte {
	if m != nil {
		return m.TokenIds
	}
	return nil
}

//ExpectationRequest用于根据令牌期望请求间接令牌导入或传输
type ExpectationRequest struct {
//凭证包含请求操作的一方的信息
//此字段的内容取决于令牌管理器系统的特性
	Credential []byte `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
//预期包含令牌导入或传输的预期输出
	Expectation *TokenExpectation `protobuf:"bytes,2,opt,name=expectation,proto3" json:"expectation,omitempty"`
//tokenids是用于满足期望的令牌标识符。
	TokenIds             [][]byte `protobuf:"bytes,3,rep,name=token_ids,json=tokenIds,proto3" json:"token_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpectationRequest) Reset()         { *m = ExpectationRequest{} }
func (m *ExpectationRequest) String() string { return proto.CompactTextString(m) }
func (*ExpectationRequest) ProtoMessage()    {}
func (*ExpectationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{10}
}
func (m *ExpectationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpectationRequest.Unmarshal(m, b)
}
func (m *ExpectationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpectationRequest.Marshal(b, m, deterministic)
}
func (dst *ExpectationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpectationRequest.Merge(dst, src)
}
func (m *ExpectationRequest) XXX_Size() int {
	return xxx_messageInfo_ExpectationRequest.Size(m)
}
func (m *ExpectationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpectationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExpectationRequest proto.InternalMessageInfo

func (m *ExpectationRequest) GetCredential() []byte {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *ExpectationRequest) GetExpectation() *TokenExpectation {
	if m != nil {
		return m.Expectation
	}
	return nil
}

func (m *ExpectationRequest) GetTokenIds() [][]byte {
	if m != nil {
		return m.TokenIds
	}
	return nil
}

//header是一个通用的重播预防和标识消息，要包含在签名的命令中
type Header struct {
//Timestamp是创建消息的本地时间
//由发送者
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
//channelid标识此消息绑定到的通道
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
//nonce是一个足够长的随机值
//用于确保请求具有足够的熵。
	Nonce []byte `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
//消息的创建者。
//通常，封送的msp.serialididEntity
	Creator              []byte   `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{11}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (dst *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(dst, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Header) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *Header) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *Header) GetCreator() []byte {
	if m != nil {
		return m.Creator
	}
	return nil
}

//命令描述客户端请求的操作类型。
type Command struct {
//header是此命令的头
	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
//有效载荷是这个命令的有效载荷。它可以假定以下值之一
//
//有效分配给有效负载的类型：
//*命令输入请求
//*命令传输请求
//*命令列表请求
//*命令REDEEMREQUEST
//*命令_ApproveRequest
//*命令传输请求
//*命令请求
	Payload              isCommand_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{12}
}
func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (dst *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(dst, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

type isCommand_Payload interface {
	isCommand_Payload()
}

type Command_ImportRequest struct {
	ImportRequest *ImportRequest `protobuf:"bytes,2,opt,name=import_request,json=importRequest,proto3,oneof"`
}

type Command_TransferRequest struct {
	TransferRequest *TransferRequest `protobuf:"bytes,3,opt,name=transfer_request,json=transferRequest,proto3,oneof"`
}

type Command_ListRequest struct {
	ListRequest *ListRequest `protobuf:"bytes,4,opt,name=list_request,json=listRequest,proto3,oneof"`
}

type Command_RedeemRequest struct {
	RedeemRequest *RedeemRequest `protobuf:"bytes,5,opt,name=redeem_request,json=redeemRequest,proto3,oneof"`
}

type Command_ApproveRequest struct {
	ApproveRequest *ApproveRequest `protobuf:"bytes,6,opt,name=approve_request,json=approveRequest,proto3,oneof"`
}

type Command_TransferFromRequest struct {
	TransferFromRequest *TransferRequest `protobuf:"bytes,7,opt,name=transfer_from_request,json=transferFromRequest,proto3,oneof"`
}

type Command_ExpectationRequest struct {
	ExpectationRequest *ExpectationRequest `protobuf:"bytes,8,opt,name=expectation_request,json=expectationRequest,proto3,oneof"`
}

func (*Command_ImportRequest) isCommand_Payload() {}

func (*Command_TransferRequest) isCommand_Payload() {}

func (*Command_ListRequest) isCommand_Payload() {}

func (*Command_RedeemRequest) isCommand_Payload() {}

func (*Command_ApproveRequest) isCommand_Payload() {}

func (*Command_TransferFromRequest) isCommand_Payload() {}

func (*Command_ExpectationRequest) isCommand_Payload() {}

func (m *Command) GetPayload() isCommand_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Command) GetImportRequest() *ImportRequest {
	if x, ok := m.GetPayload().(*Command_ImportRequest); ok {
		return x.ImportRequest
	}
	return nil
}

func (m *Command) GetTransferRequest() *TransferRequest {
	if x, ok := m.GetPayload().(*Command_TransferRequest); ok {
		return x.TransferRequest
	}
	return nil
}

func (m *Command) GetListRequest() *ListRequest {
	if x, ok := m.GetPayload().(*Command_ListRequest); ok {
		return x.ListRequest
	}
	return nil
}

func (m *Command) GetRedeemRequest() *RedeemRequest {
	if x, ok := m.GetPayload().(*Command_RedeemRequest); ok {
		return x.RedeemRequest
	}
	return nil
}

func (m *Command) GetApproveRequest() *ApproveRequest {
	if x, ok := m.GetPayload().(*Command_ApproveRequest); ok {
		return x.ApproveRequest
	}
	return nil
}

func (m *Command) GetTransferFromRequest() *TransferRequest {
	if x, ok := m.GetPayload().(*Command_TransferFromRequest); ok {
		return x.TransferFromRequest
	}
	return nil
}

func (m *Command) GetExpectationRequest() *ExpectationRequest {
	if x, ok := m.GetPayload().(*Command_ExpectationRequest); ok {
		return x.ExpectationRequest
	}
	return nil
}

//xxxoneoffuncs用于Proto包的内部使用。
func (*Command) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Command_OneofMarshaler, _Command_OneofUnmarshaler, _Command_OneofSizer, []interface{}{
		(*Command_ImportRequest)(nil),
		(*Command_TransferRequest)(nil),
		(*Command_ListRequest)(nil),
		(*Command_RedeemRequest)(nil),
		(*Command_ApproveRequest)(nil),
		(*Command_TransferFromRequest)(nil),
		(*Command_ExpectationRequest)(nil),
	}
}

func _Command_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Command)
//有效载荷
	switch x := m.Payload.(type) {
	case *Command_ImportRequest:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ImportRequest); err != nil {
			return err
		}
	case *Command_TransferRequest:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TransferRequest); err != nil {
			return err
		}
	case *Command_ListRequest:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListRequest); err != nil {
			return err
		}
	case *Command_RedeemRequest:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RedeemRequest); err != nil {
			return err
		}
	case *Command_ApproveRequest:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApproveRequest); err != nil {
			return err
		}
	case *Command_TransferFromRequest:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TransferFromRequest); err != nil {
			return err
		}
	case *Command_ExpectationRequest:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExpectationRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Command.Payload has unexpected type %T", x)
	}
	return nil
}

func _Command_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Command)
	switch tag {
case 2: //有效载荷.导入请求
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ImportRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &Command_ImportRequest{msg}
		return true, err
case 3: //有效载荷传送请求
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &Command_TransferRequest{msg}
		return true, err
case 4: //payload.list_请求
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &Command_ListRequest{msg}
		return true, err
case 5: //有效载荷.兑换请求
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RedeemRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &Command_RedeemRequest{msg}
		return true, err
case 6: //Payload.Approve_请求
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ApproveRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &Command_ApproveRequest{msg}
		return true, err
case 7: //Payload.Transfer_from_请求
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &Command_TransferFromRequest{msg}
		return true, err
case 8: //有效载荷期望请求
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExpectationRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &Command_ExpectationRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Command_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Command)
//
	switch x := m.Payload.(type) {
	case *Command_ImportRequest:
		s := proto.Size(x.ImportRequest)
n += 1 //标签和电线
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Command_TransferRequest:
		s := proto.Size(x.TransferRequest)
n += 1 //标签和电线
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Command_ListRequest:
		s := proto.Size(x.ListRequest)
n += 1 //标签和电线
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Command_RedeemRequest:
		s := proto.Size(x.RedeemRequest)
n += 1 //标签和电线
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Command_ApproveRequest:
		s := proto.Size(x.ApproveRequest)
n += 1 //标签和电线
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Command_TransferFromRequest:
		s := proto.Size(x.TransferFromRequest)
n += 1 //标签和电线
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Command_ExpectationRequest:
		s := proto.Size(x.ExpectationRequest)
n += 1 //标签和电线
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

//signedCommand是一个带有命令创建者签名的命令。
type SignedCommand struct {
//命令是命令消息的序列化版本
	Command []byte `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
//签名是对命令的签名
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedCommand) Reset()         { *m = SignedCommand{} }
func (m *SignedCommand) String() string { return proto.CompactTextString(m) }
func (*SignedCommand) ProtoMessage()    {}
func (*SignedCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{13}
}
func (m *SignedCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedCommand.Unmarshal(m, b)
}
func (m *SignedCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedCommand.Marshal(b, m, deterministic)
}
func (dst *SignedCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedCommand.Merge(dst, src)
}
func (m *SignedCommand) XXX_Size() int {
	return xxx_messageInfo_SignedCommand.Size(m)
}
func (m *SignedCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedCommand.DiscardUnknown(m)
}

var xxx_messageInfo_SignedCommand proto.InternalMessageInfo

func (m *SignedCommand) GetCommand() []byte {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *SignedCommand) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type CommandResponseHeader struct {
//timestamp是消息
//已按发件人的定义创建
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
//commandHash是根据signedCommand的命令和签名字段串联计算的哈希。
//如果没有不同的指定，则使用sha256
//哈希用于将响应与其请求链接起来，这两种方法都用于
//异步系统和出于安全原因（问责、不可抵赖）
	CommandHash []byte `protobuf:"bytes,2,opt,name=command_hash,json=commandHash,proto3" json:"command_hash,omitempty"`
//创建者是创建此邮件的一方的标识
	Creator              []byte   `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandResponseHeader) Reset()         { *m = CommandResponseHeader{} }
func (m *CommandResponseHeader) String() string { return proto.CompactTextString(m) }
func (*CommandResponseHeader) ProtoMessage()    {}
func (*CommandResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{14}
}
func (m *CommandResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandResponseHeader.Unmarshal(m, b)
}
func (m *CommandResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandResponseHeader.Marshal(b, m, deterministic)
}
func (dst *CommandResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandResponseHeader.Merge(dst, src)
}
func (m *CommandResponseHeader) XXX_Size() int {
	return xxx_messageInfo_CommandResponseHeader.Size(m)
}
func (m *CommandResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_CommandResponseHeader proto.InternalMessageInfo

func (m *CommandResponseHeader) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *CommandResponseHeader) GetCommandHash() []byte {
	if m != nil {
		return m.CommandHash
	}
	return nil
}

func (m *CommandResponseHeader) GetCreator() []byte {
	if m != nil {
		return m.Creator
	}
	return nil
}

//错误报告应用程序错误
type Error struct {
//与此响应关联的消息。
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
//可用于将元数据包含在此响应中的有效负载。
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{15}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

//命令和响应从验证程序返回给命令提交者。
type CommandResponse struct {
//响应的标题。
	Header *CommandResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
//响应的有效负载。
//
//有效分配给有效负载的类型：
//*命令响应错误
//*commandresponse_令牌事务
//*命令响应\u未使用的令牌
	Payload              isCommandResponse_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CommandResponse) Reset()         { *m = CommandResponse{} }
func (m *CommandResponse) String() string { return proto.CompactTextString(m) }
func (*CommandResponse) ProtoMessage()    {}
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{16}
}
func (m *CommandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandResponse.Unmarshal(m, b)
}
func (m *CommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandResponse.Marshal(b, m, deterministic)
}
func (dst *CommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandResponse.Merge(dst, src)
}
func (m *CommandResponse) XXX_Size() int {
	return xxx_messageInfo_CommandResponse.Size(m)
}
func (m *CommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommandResponse proto.InternalMessageInfo

func (m *CommandResponse) GetHeader() *CommandResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type isCommandResponse_Payload interface {
	isCommandResponse_Payload()
}

type CommandResponse_Err struct {
	Err *Error `protobuf:"bytes,2,opt,name=err,proto3,oneof"`
}

type CommandResponse_TokenTransaction struct {
	TokenTransaction *TokenTransaction `protobuf:"bytes,3,opt,name=token_transaction,json=tokenTransaction,proto3,oneof"`
}

type CommandResponse_UnspentTokens struct {
	UnspentTokens *UnspentTokens `protobuf:"bytes,4,opt,name=unspent_tokens,json=unspentTokens,proto3,oneof"`
}

func (*CommandResponse_Err) isCommandResponse_Payload() {}

func (*CommandResponse_TokenTransaction) isCommandResponse_Payload() {}

func (*CommandResponse_UnspentTokens) isCommandResponse_Payload() {}

func (m *CommandResponse) GetPayload() isCommandResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CommandResponse) GetErr() *Error {
	if x, ok := m.GetPayload().(*CommandResponse_Err); ok {
		return x.Err
	}
	return nil
}

func (m *CommandResponse) GetTokenTransaction() *TokenTransaction {
	if x, ok := m.GetPayload().(*CommandResponse_TokenTransaction); ok {
		return x.TokenTransaction
	}
	return nil
}

func (m *CommandResponse) GetUnspentTokens() *UnspentTokens {
	if x, ok := m.GetPayload().(*CommandResponse_UnspentTokens); ok {
		return x.UnspentTokens
	}
	return nil
}

//xxxoneoffuncs用于Proto包的内部使用。
func (*CommandResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CommandResponse_OneofMarshaler, _CommandResponse_OneofUnmarshaler, _CommandResponse_OneofSizer, []interface{}{
		(*CommandResponse_Err)(nil),
		(*CommandResponse_TokenTransaction)(nil),
		(*CommandResponse_UnspentTokens)(nil),
	}
}

func _CommandResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CommandResponse)
//有效载荷
	switch x := m.Payload.(type) {
	case *CommandResponse_Err:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Err); err != nil {
			return err
		}
	case *CommandResponse_TokenTransaction:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TokenTransaction); err != nil {
			return err
		}
	case *CommandResponse_UnspentTokens:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UnspentTokens); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CommandResponse.Payload has unexpected type %T", x)
	}
	return nil
}

func _CommandResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CommandResponse)
	switch tag {
case 2: //有效载荷
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Error)
		err := b.DecodeMessage(msg)
		m.Payload = &CommandResponse_Err{msg}
		return true, err
case 3: //payload.token_事务
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TokenTransaction)
		err := b.DecodeMessage(msg)
		m.Payload = &CommandResponse_TokenTransaction{msg}
		return true, err
case 4: //有效载荷.未使用的标记
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UnspentTokens)
		err := b.DecodeMessage(msg)
		m.Payload = &CommandResponse_UnspentTokens{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CommandResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CommandResponse)
//有效载荷
	switch x := m.Payload.(type) {
	case *CommandResponse_Err:
		s := proto.Size(x.Err)
n += 1 //标签和电线
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CommandResponse_TokenTransaction:
		s := proto.Size(x.TokenTransaction)
n += 1 //标签和电线
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CommandResponse_UnspentTokens:
		s := proto.Size(x.UnspentTokens)
n += 1 //标签和电线
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

//SignedCommandResponse是签名的命令响应
type SignedCommandResponse struct {
//Response是commandResponse消息的序列化版本
	Response []byte `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
//签名是对命令的签名
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedCommandResponse) Reset()         { *m = SignedCommandResponse{} }
func (m *SignedCommandResponse) String() string { return proto.CompactTextString(m) }
func (*SignedCommandResponse) ProtoMessage()    {}
func (*SignedCommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_prover_aacdf31643d38144, []int{17}
}
func (m *SignedCommandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedCommandResponse.Unmarshal(m, b)
}
func (m *SignedCommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedCommandResponse.Marshal(b, m, deterministic)
}
func (dst *SignedCommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedCommandResponse.Merge(dst, src)
}
func (m *SignedCommandResponse) XXX_Size() int {
	return xxx_messageInfo_SignedCommandResponse.Size(m)
}
func (m *SignedCommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedCommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignedCommandResponse proto.InternalMessageInfo

func (m *SignedCommandResponse) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *SignedCommandResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenToIssue)(nil), "protos.TokenToIssue")
	proto.RegisterType((*RecipientTransferShare)(nil), "protos.RecipientTransferShare")
	proto.RegisterType((*TokenOutput)(nil), "protos.TokenOutput")
	proto.RegisterType((*UnspentTokens)(nil), "protos.UnspentTokens")
	proto.RegisterType((*ListRequest)(nil), "protos.ListRequest")
	proto.RegisterType((*ImportRequest)(nil), "protos.ImportRequest")
	proto.RegisterType((*TransferRequest)(nil), "protos.TransferRequest")
	proto.RegisterType((*RedeemRequest)(nil), "protos.RedeemRequest")
	proto.RegisterType((*AllowanceRecipientShare)(nil), "protos.AllowanceRecipientShare")
	proto.RegisterType((*ApproveRequest)(nil), "protos.ApproveRequest")
	proto.RegisterType((*ExpectationRequest)(nil), "protos.ExpectationRequest")
	proto.RegisterType((*Header)(nil), "protos.Header")
	proto.RegisterType((*Command)(nil), "protos.Command")
	proto.RegisterType((*SignedCommand)(nil), "protos.SignedCommand")
	proto.RegisterType((*CommandResponseHeader)(nil), "protos.CommandResponseHeader")
	proto.RegisterType((*Error)(nil), "protos.Error")
	proto.RegisterType((*CommandResponse)(nil), "protos.CommandResponse")
	proto.RegisterType((*SignedCommandResponse)(nil), "protos.SignedCommandResponse")
}

//引用导入以禁止错误（如果未使用）。
var _ context.Context
var _ grpc.ClientConn

//这是一个编译时断言，以确保生成的文件
//与正在编译的GRPC包兼容。
const _ = grpc.SupportPackageIsVersion4

//prover client是prover服务的客户端API。
//
//有关CTX使用和关闭/结束流式RPC的语义，请参阅https://godoc.org/google.golang.org/grpc clientconn.newstream。
type ProverClient interface {
//processcommand处理传递的命令，确保正确的访问控制。
//返回的响应允许客户端了解
//操作已成功执行，如果没有，则响应
//报告失败的原因。
	ProcessCommand(ctx context.Context, in *SignedCommand, opts ...grpc.CallOption) (*SignedCommandResponse, error)
}

type proverClient struct {
	cc *grpc.ClientConn
}

func NewProverClient(cc *grpc.ClientConn) ProverClient {
	return &proverClient{cc}
}

func (c *proverClient) ProcessCommand(ctx context.Context, in *SignedCommand, opts ...grpc.CallOption) (*SignedCommandResponse, error) {
	out := new(SignedCommandResponse)
	err := c.cc.Invoke(ctx, "/protos.Prover/ProcessCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//prover server是prover服务的服务器API。
type ProverServer interface {
//processcommand处理传递的命令，确保正确的访问控制。
//返回的响应允许客户端了解
//操作已成功执行，如果没有，则响应
//报告失败的原因。
	ProcessCommand(context.Context, *SignedCommand) (*SignedCommandResponse, error)
}

func RegisterProverServer(s *grpc.Server, srv ProverServer) {
	s.RegisterService(&_Prover_serviceDesc, srv)
}

func _Prover_ProcessCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProverServer).ProcessCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Prover/ProcessCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProverServer).ProcessCommand(ctx, req.(*SignedCommand))
	}
	return interceptor(ctx, in, info, handler)
}

var _Prover_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Prover",
	HandlerType: (*ProverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessCommand",
			Handler:    _Prover_ProcessCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token/prover.proto",
}

func init() { proto.RegisterFile("token/prover.proto", fileDescriptor_prover_aacdf31643d38144) }

var fileDescriptor_prover_aacdf31643d38144 = []byte{
//gzip文件描述符或协议的1010字节
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6f, 0x1b, 0x45,
	0x10, 0xf7, 0xd9, 0xae, 0x13, 0x8f, 0xbf, 0xd2, 0x4d, 0xd3, 0x9c, 0x0c, 0x69, 0xd3, 0x43, 0x42,
	0x11, 0x1f, 0xb6, 0x14, 0x04, 0xaa, 0xa0, 0xaa, 0x48, 0xa1, 0x70, 0x41, 0x54, 0xb4, 0x1b, 0xf3,
	0x82, 0x90, 0xac, 0xcd, 0xdd, 0xc6, 0x5e, 0xe1, 0xbb, 0xbd, 0xee, 0xee, 0x01, 0xe1, 0x0f, 0xe0,
	0x0d, 0x24, 0x1e, 0x79, 0xe0, 0xff, 0xe4, 0x11, 0xdd, 0x7e, 0x9c, 0xef, 0xdc, 0xd0, 0x06, 0x95,
	0x27, 0x7b, 0x67, 0x66, 0x67, 0x7e, 0x3b, 0xfb, 0xfb, 0xcd, 0x1e, 0x20, 0xc5, 0x7f, 0xa0, 0xe9,
	0x34, 0x13, 0xfc, 0x47, 0x2a, 0x26, 0x99, 0xe0, 0x8a, 0xa3, 0x8e, 0xfe, 0x91, 0xe3, 0xbb, 0x0b,
	0xce, 0x17, 0x2b, 0x3a, 0xd5, 0xcb, 0xf3, 0xfc, 0x62, 0xaa, 0x58, 0x42, 0xa5, 0x22, 0x49, 0x66,
	0x02, 0xc7, 0xbe, 0xd9, 0x4c, 0x7f, 0xce, 0x68, 0xa4, 0x88, 0x62, 0x3c, 0x95, 0xd6, 0xb3, 0x6f,
	0x3c, 0x4a, 0x90, 0x54, 0x92, 0xa8, 0xf0, 0x18, 0x47, 0xf0, 0x3d, 0xf4, 0x67, 0x85, 0x6b, 0xc6,
	0x4f, 0xa5, 0xcc, 0x29, 0x7a, 0x13, 0xba, 0x82, 0x46, 0x2c, 0x63, 0x34, 0x55, 0xbe, 0x77, 0xe8,
	0x1d, 0xf5, 0xf1, 0xda, 0x80, 0x10, 0xb4, 0xd5, 0x65, 0x46, 0xfd, 0xe6, 0xa1, 0x77, 0xd4, 0xc5,
	0xfa, 0x3f, 0x1a, 0xc3, 0xf6, 0xf3, 0x9c, 0xa4, 0x8a, 0xa9, 0x4b, 0xbf, 0x75, 0xe8, 0x1d, 0xb5,
	0x71, 0xb9, 0x0e, 0x30, 0xdc, 0xc6, 0x6e, 0xf3, 0xac, 0xa8, 0x7d, 0x41, 0xc5, 0xd9, 0x92, 0x88,
	0x57, 0xd5, 0xa9, 0xe6, 0x6c, 0x6e, 0xe4, 0x7c, 0x02, 0x3d, 0x8d, 0xf8, 0x9b, 0x5c, 0x65, 0xb9,
	0x42, 0x43, 0x68, 0xb2, 0xd8, 0x66, 0x68, 0xb2, 0xf8, 0x3f, 0x43, 0x7c, 0x00, 0x83, 0x6f, 0x53,
	0x99, 0x15, 0x00, 0x8b, 0xac, 0x12, 0xbd, 0x0b, 0x1d, 0xdd, 0x2c, 0xe9, 0x7b, 0x87, 0xad, 0xa3,
	0xde, 0xf1, 0xae, 0xe9, 0x94, 0x9c, 0x54, 0xaa, 0x62, 0x1b, 0x12, 0xbc, 0x0f, 0xbd, 0xaf, 0x99,
	0x54, 0x98, 0x3e, 0xcf, 0xa9, 0x54, 0xe8, 0x0e, 0x40, 0x24, 0x68, 0x4c, 0x53, 0xc5, 0xc8, 0xca,
	0x82, 0xaa, 0x58, 0x82, 0x04, 0x06, 0xa7, 0x49, 0xc6, 0xc5, 0x75, 0x37, 0xa0, 0x07, 0x30, 0x32,
	0x95, 0xe6, 0x8a, 0xcf, 0x59, 0x71, 0x43, 0x7e, 0x53, 0xa3, 0xba, 0x55, 0x43, 0x65, 0x6f, 0x0f,
	0x0f, 0x4c, 0xb0, 0x5d, 0x06, 0xbf, 0x7a, 0x30, 0x72, 0x6d, 0xbf, 0x6e, 0xc5, 0x37, 0xa0, 0xab,
	0x93, 0xcc, 0x59, 0x2c, 0x75, 0xad, 0x3e, 0xde, 0xd6, 0x86, 0xd3, 0x58, 0xa2, 0x8f, 0xa0, 0x23,
	0x8b, 0xeb, 0x93, 0x7e, 0x4b, 0xa3, 0xb8, 0xe3, 0x50, 0x5c, 0x7d, 0xcb, 0xd8, 0x46, 0x07, 0xbf,
	0xc0, 0x00, 0xd3, 0x98, 0xd2, 0xe4, 0x7f, 0x41, 0xf1, 0x1e, 0x20, 0x77, 0x7d, 0x45, 0x5b, 0x84,
	0xce, 0x6c, 0x2f, 0x76, 0xc7, 0x79, 0x66, 0xdc, 0x54, 0x0c, 0xce, 0x60, 0xff, 0x64, 0xb5, 0xe2,
	0x3f, 0x91, 0x34, 0xa2, 0x25, 0xcc, 0xd7, 0x25, 0xe1, 0x9f, 0x1e, 0x0c, 0x4f, 0x32, 0xad, 0xd2,
	0xeb, 0x1e, 0xe9, 0x2b, 0xd8, 0x21, 0x0e, 0xc7, 0xdc, 0x76, 0xd1, 0xdc, 0xe5, 0x5d, 0xd7, 0xc5,
	0x7f, 0xc1, 0x89, 0x47, 0xe5, 0x46, 0xbd, 0x96, 0xf5, 0xf6, 0xb4, 0xea, 0xed, 0x09, 0x7e, 0xf3,
	0x00, 0x3d, 0x5e, 0x8f, 0x80, 0xeb, 0xe2, 0xfb, 0x18, 0x7a, 0x95, 0xc1, 0xa1, 0x4f, 0xdc, 0x3b,
	0xf6, 0x6b, 0x34, 0xab, 0x66, 0xad, 0x06, 0xbf, 0x1c, 0xcf, 0x1f, 0x1e, 0x74, 0x42, 0x4a, 0x62,
	0x2a, 0xd0, 0x7d, 0xe8, 0x96, 0x33, 0x4b, 0x43, 0xe8, 0x1d, 0x8f, 0x27, 0x66, 0xaa, 0x4d, 0xdc,
	0x54, 0x9b, 0xcc, 0x5c, 0x04, 0x5e, 0x07, 0xa3, 0x03, 0x80, 0x68, 0x49, 0xd2, 0x94, 0xae, 0xe6,
	0x2c, 0xb6, 0xe2, 0xee, 0x5a, 0xcb, 0x69, 0x8c, 0x6e, 0xc1, 0x8d, 0x94, 0xa7, 0x11, 0xd5, 0x2c,
	0xe8, 0x63, 0xb3, 0x40, 0x3e, 0x6c, 0x45, 0x82, 0x12, 0xc5, 0x85, 0xdf, 0xd6, 0x76, 0xb7, 0x0c,
	0xfe, 0x6a, 0xc3, 0xd6, 0x67, 0x3c, 0x49, 0x48, 0x1a, 0xa3, 0xb7, 0xa1, 0xb3, 0xd4, 0xf0, 0x2c,
	0xa2, 0xa1, 0x3b, 0xb3, 0x01, 0x8d, 0xad, 0x17, 0x3d, 0x84, 0x21, 0xd3, 0xe2, 0x9d, 0x0b, 0xd3,
	0x52, 0xdb, 0xa3, 0x3d, 0x17, 0x5f, 0x93, 0x76, 0xd8, 0xc0, 0x03, 0x56, 0xd3, 0xfa, 0xe7, 0xb0,
	0xa3, 0xac, 0x3a, 0xca, 0x0c, 0x2d, 0x9d, 0x61, 0xbf, 0xec, 0x72, 0x5d, 0xac, 0x61, 0x03, 0x8f,
	0xd4, 0x86, 0x7e, 0xef, 0x43, 0x7f, 0xc5, 0xe4, 0x1a, 0x43, 0x5b, 0x67, 0x28, 0x87, 0x54, 0x65,
	0x1a, 0x85, 0x0d, 0xdc, 0x5b, 0x55, 0x86, 0xd3, 0x43, 0x18, 0x1a, 0xa9, 0x94, 0x7b, 0x6f, 0xd4,
	0xf1, 0xd7, 0x24, 0x5a, 0xe0, 0x17, 0x35, 0xcd, 0x9e, 0xc0, 0x88, 0x18, 0xca, 0x97, 0x09, 0x3a,
	0x3a, 0xc1, 0xed, 0x92, 0xbf, 0x35, 0x45, 0x84, 0x0d, 0x3c, 0x24, 0x75, 0x8d, 0x3c, 0x81, 0xbd,
	0xb2, 0x05, 0x17, 0x82, 0xaf, 0x91, 0x6c, 0xbd, 0xaa, 0x0f, 0xbb, 0x6e, 0xdf, 0x17, 0x82, 0x27,
	0xeb, 0x74, 0xbb, 0x15, 0x16, 0x96, 0xc9, 0xb6, 0x2d, 0xb1, 0x6c, 0xb2, 0x17, 0xb5, 0x10, 0x36,
	0x30, 0xa2, 0x2f, 0x58, 0x1f, 0x75, 0x61, 0x2b, 0x23, 0x97, 0x2b, 0x4e, 0xe2, 0xe0, 0x4b, 0x18,
	0x9c, 0xb1, 0x45, 0x4a, 0x63, 0x47, 0x92, 0x82, 0x4a, 0xe6, 0xaf, 0x95, 0x8e, 0x5b, 0x16, 0x43,
	0x44, 0xb2, 0x45, 0x4a, 0x54, 0x2e, 0xcc, 0xab, 0xd3, 0xc7, 0x6b, 0x43, 0xf0, 0xbb, 0x07, 0x7b,
	0x36, 0x07, 0xa6, 0x32, 0xe3, 0xa9, 0xa4, 0xaf, 0xad, 0x85, 0x7b, 0xd0, 0xb7, 0xc5, 0xe7, 0x4b,
	0x22, 0x97, 0xb6, 0x68, 0xcf, 0xda, 0x42, 0x22, 0x97, 0x55, 0xe6, 0xb7, 0xea, 0xcc, 0xff, 0x04,
	0x6e, 0x3c, 0x16, 0x82, 0x8b, 0x22, 0x24, 0xa1, 0x52, 0x92, 0x05, 0xd5, 0xd5, 0xbb, 0xd8, 0x2d,
	0x0b, 0x8f, 0xed, 0x83, 0x4d, 0x5d, 0xb6, 0xe5, 0x6f, 0x0f, 0x46, 0x1b, 0xa7, 0x41, 0x1f, 0x6e,
	0xc8, 0xe7, 0xc0, 0xf5, 0xfd, 0xca, 0x63, 0x97, 0x6a, 0xba, 0x07, 0x2d, 0x2a, 0x84, 0x95, 0xd0,
	0xa0, 0xbc, 0xab, 0x02, 0x5a, 0xd8, 0xc0, 0x85, 0x0f, 0x7d, 0x0a, 0x37, 0xcd, 0x54, 0xa9, 0x7c,
	0xb6, 0x58, 0xc5, 0xdc, 0xb4, 0xef, 0xde, 0xda, 0x11, 0x36, 0xf0, 0x8e, 0xda, 0xb0, 0x15, 0x94,
	0xcf, 0xcd, 0xe3, 0x3e, 0xb7, 0x6f, 0x7a, 0xbb, 0x4e, 0xf9, 0xda, 0xd3, 0x5f, 0x50, 0x3e, 0xaf,
	0x1a, 0xaa, 0x8c, 0x78, 0x06, 0x7b, 0x35, 0x46, 0x94, 0xe7, 0x1f, 0xc3, 0xb6, 0xb0, 0xff, 0x2d,
	0x35, 0xca, 0xf5, 0xcb, 0xb9, 0x71, 0x8c, 0xa1, 0xf3, 0x54, 0x7f, 0xe7, 0xa1, 0x10, 0x86, 0x4f,
	0x05, 0x8f, 0xa8, 0x94, 0x8e, 0x6f, 0x25, 0xc2, 0x5a, 0xd1, 0xf1, 0xc1, 0x95, 0x66, 0x87, 0x25,
	0x68, 0x3c, 0x7a, 0x06, 0x6f, 0x71, 0xb1, 0x98, 0x2c, 0x2f, 0x33, 0x2a, 0x56, 0x34, 0x5e, 0x50,
	0x31, 0xb9, 0x20, 0xe7, 0x82, 0x45, 0x6e, 0xa3, 0xee, 0xc3, 0x77, 0xef, 0x2c, 0x98, 0x5a, 0xe6,
	0xe7, 0x93, 0x88, 0x27, 0xd3, 0x4a, 0xec, 0xd4, 0xc4, 0x9a, 0x2f, 0x4c, 0x39, 0xd5, 0xb1, 0xe7,
	0xe6, 0xf3, 0xf3, 0x83, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0x78, 0x90, 0x5b, 0x9b, 0x0a,
	0x00, 0x00,
}

