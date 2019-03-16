
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:35</date>
//</624456118143291392>

//由Protoc Gen Go生成的代码。不要编辑。
//来源：Ledger/rwset/rwset.proto

package rwset //导入“github.com/hyperledger/fabric/protos/ledger/rwset”

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

//引用导入以禁止错误（如果未使用）。
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//这是一个编译时断言，以确保生成的文件
//与正在编译的proto包兼容。
//此行的编译错误可能意味着您的
//需要更新proto包。
const _ = proto.ProtoPackageIsVersion2 //请升级proto包

type TxReadWriteSet_DataModel int32

const (
	TxReadWriteSet_KV TxReadWriteSet_DataModel = 0
)

var TxReadWriteSet_DataModel_name = map[int32]string{
	0: "KV",
}
var TxReadWriteSet_DataModel_value = map[string]int32{
	"KV": 0,
}

func (x TxReadWriteSet_DataModel) String() string {
	return proto.EnumName(TxReadWriteSet_DataModel_name, int32(x))
}
func (TxReadWriteSet_DataModel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rwset_28a2ad70bffb57df, []int{0, 0}
}

//txreadwriteset封装事务的读写集
//data model指定数据模型的枚举值
//nsrwset字段指定链码特定读写集的列表（每个链码一个）
type TxReadWriteSet struct {
	DataModel            TxReadWriteSet_DataModel `protobuf:"varint,1,opt,name=data_model,json=dataModel,proto3,enum=rwset.TxReadWriteSet_DataModel" json:"data_model,omitempty"`
	NsRwset              []*NsReadWriteSet        `protobuf:"bytes,2,rep,name=ns_rwset,json=nsRwset,proto3" json:"ns_rwset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TxReadWriteSet) Reset()         { *m = TxReadWriteSet{} }
func (m *TxReadWriteSet) String() string { return proto.CompactTextString(m) }
func (*TxReadWriteSet) ProtoMessage()    {}
func (*TxReadWriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_rwset_28a2ad70bffb57df, []int{0}
}
func (m *TxReadWriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxReadWriteSet.Unmarshal(m, b)
}
func (m *TxReadWriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxReadWriteSet.Marshal(b, m, deterministic)
}
func (dst *TxReadWriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxReadWriteSet.Merge(dst, src)
}
func (m *TxReadWriteSet) XXX_Size() int {
	return xxx_messageInfo_TxReadWriteSet.Size(m)
}
func (m *TxReadWriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_TxReadWriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_TxReadWriteSet proto.InternalMessageInfo

func (m *TxReadWriteSet) GetDataModel() TxReadWriteSet_DataModel {
	if m != nil {
		return m.DataModel
	}
	return TxReadWriteSet_KV
}

func (m *TxReadWriteSet) GetNsRwset() []*NsReadWriteSet {
	if m != nil {
		return m.NsRwset
	}
	return nil
}

//NsreadWriteset封装链码的读写集
type NsReadWriteSet struct {
	Namespace             string                          `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Rwset                 []byte                          `protobuf:"bytes,2,opt,name=rwset,proto3" json:"rwset,omitempty"`
	CollectionHashedRwset []*CollectionHashedReadWriteSet `protobuf:"bytes,3,rep,name=collection_hashed_rwset,json=collectionHashedRwset,proto3" json:"collection_hashed_rwset,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                        `json:"-"`
	XXX_unrecognized      []byte                          `json:"-"`
	XXX_sizecache         int32                           `json:"-"`
}

func (m *NsReadWriteSet) Reset()         { *m = NsReadWriteSet{} }
func (m *NsReadWriteSet) String() string { return proto.CompactTextString(m) }
func (*NsReadWriteSet) ProtoMessage()    {}
func (*NsReadWriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_rwset_28a2ad70bffb57df, []int{1}
}
func (m *NsReadWriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NsReadWriteSet.Unmarshal(m, b)
}
func (m *NsReadWriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NsReadWriteSet.Marshal(b, m, deterministic)
}
func (dst *NsReadWriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NsReadWriteSet.Merge(dst, src)
}
func (m *NsReadWriteSet) XXX_Size() int {
	return xxx_messageInfo_NsReadWriteSet.Size(m)
}
func (m *NsReadWriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_NsReadWriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_NsReadWriteSet proto.InternalMessageInfo

func (m *NsReadWriteSet) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NsReadWriteSet) GetRwset() []byte {
	if m != nil {
		return m.Rwset
	}
	return nil
}

func (m *NsReadWriteSet) GetCollectionHashedRwset() []*CollectionHashedReadWriteSet {
	if m != nil {
		return m.CollectionHashedRwset
	}
	return nil
}

//CollectionHashedReadWriteSet封装集合的专用读写集的哈希表示形式
type CollectionHashedReadWriteSet struct {
	CollectionName       string   `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	HashedRwset          []byte   `protobuf:"bytes,2,opt,name=hashed_rwset,json=hashedRwset,proto3" json:"hashed_rwset,omitempty"`
	PvtRwsetHash         []byte   `protobuf:"bytes,3,opt,name=pvt_rwset_hash,json=pvtRwsetHash,proto3" json:"pvt_rwset_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionHashedReadWriteSet) Reset()         { *m = CollectionHashedReadWriteSet{} }
func (m *CollectionHashedReadWriteSet) String() string { return proto.CompactTextString(m) }
func (*CollectionHashedReadWriteSet) ProtoMessage()    {}
func (*CollectionHashedReadWriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_rwset_28a2ad70bffb57df, []int{2}
}
func (m *CollectionHashedReadWriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionHashedReadWriteSet.Unmarshal(m, b)
}
func (m *CollectionHashedReadWriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionHashedReadWriteSet.Marshal(b, m, deterministic)
}
func (dst *CollectionHashedReadWriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionHashedReadWriteSet.Merge(dst, src)
}
func (m *CollectionHashedReadWriteSet) XXX_Size() int {
	return xxx_messageInfo_CollectionHashedReadWriteSet.Size(m)
}
func (m *CollectionHashedReadWriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionHashedReadWriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionHashedReadWriteSet proto.InternalMessageInfo

func (m *CollectionHashedReadWriteSet) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *CollectionHashedReadWriteSet) GetHashedRwset() []byte {
	if m != nil {
		return m.HashedRwset
	}
	return nil
}

func (m *CollectionHashedReadWriteSet) GetPvtRwsetHash() []byte {
	if m != nil {
		return m.PvtRwsetHash
	}
	return nil
}

//txpvtreadwriteset封装事务的专用读写集
type TxPvtReadWriteSet struct {
	DataModel            TxReadWriteSet_DataModel `protobuf:"varint,1,opt,name=data_model,json=dataModel,proto3,enum=rwset.TxReadWriteSet_DataModel" json:"data_model,omitempty"`
	NsPvtRwset           []*NsPvtReadWriteSet     `protobuf:"bytes,2,rep,name=ns_pvt_rwset,json=nsPvtRwset,proto3" json:"ns_pvt_rwset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TxPvtReadWriteSet) Reset()         { *m = TxPvtReadWriteSet{} }
func (m *TxPvtReadWriteSet) String() string { return proto.CompactTextString(m) }
func (*TxPvtReadWriteSet) ProtoMessage()    {}
func (*TxPvtReadWriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_rwset_28a2ad70bffb57df, []int{3}
}
func (m *TxPvtReadWriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxPvtReadWriteSet.Unmarshal(m, b)
}
func (m *TxPvtReadWriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxPvtReadWriteSet.Marshal(b, m, deterministic)
}
func (dst *TxPvtReadWriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxPvtReadWriteSet.Merge(dst, src)
}
func (m *TxPvtReadWriteSet) XXX_Size() int {
	return xxx_messageInfo_TxPvtReadWriteSet.Size(m)
}
func (m *TxPvtReadWriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_TxPvtReadWriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_TxPvtReadWriteSet proto.InternalMessageInfo

func (m *TxPvtReadWriteSet) GetDataModel() TxReadWriteSet_DataModel {
	if m != nil {
		return m.DataModel
	}
	return TxReadWriteSet_KV
}

func (m *TxPvtReadWriteSet) GetNsPvtRwset() []*NsPvtReadWriteSet {
	if m != nil {
		return m.NsPvtRwset
	}
	return nil
}

//nspvtreadwriteset封装链码的专用读写集
type NsPvtReadWriteSet struct {
	Namespace            string                       `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	CollectionPvtRwset   []*CollectionPvtReadWriteSet `protobuf:"bytes,2,rep,name=collection_pvt_rwset,json=collectionPvtRwset,proto3" json:"collection_pvt_rwset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *NsPvtReadWriteSet) Reset()         { *m = NsPvtReadWriteSet{} }
func (m *NsPvtReadWriteSet) String() string { return proto.CompactTextString(m) }
func (*NsPvtReadWriteSet) ProtoMessage()    {}
func (*NsPvtReadWriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_rwset_28a2ad70bffb57df, []int{4}
}
func (m *NsPvtReadWriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NsPvtReadWriteSet.Unmarshal(m, b)
}
func (m *NsPvtReadWriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NsPvtReadWriteSet.Marshal(b, m, deterministic)
}
func (dst *NsPvtReadWriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NsPvtReadWriteSet.Merge(dst, src)
}
func (m *NsPvtReadWriteSet) XXX_Size() int {
	return xxx_messageInfo_NsPvtReadWriteSet.Size(m)
}
func (m *NsPvtReadWriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_NsPvtReadWriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_NsPvtReadWriteSet proto.InternalMessageInfo

func (m *NsPvtReadWriteSet) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NsPvtReadWriteSet) GetCollectionPvtRwset() []*CollectionPvtReadWriteSet {
	if m != nil {
		return m.CollectionPvtRwset
	}
	return nil
}

//CollectionPvTreadWriteSet封装集合的专用读写集
type CollectionPvtReadWriteSet struct {
	CollectionName       string   `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	Rwset                []byte   `protobuf:"bytes,2,opt,name=rwset,proto3" json:"rwset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionPvtReadWriteSet) Reset()         { *m = CollectionPvtReadWriteSet{} }
func (m *CollectionPvtReadWriteSet) String() string { return proto.CompactTextString(m) }
func (*CollectionPvtReadWriteSet) ProtoMessage()    {}
func (*CollectionPvtReadWriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_rwset_28a2ad70bffb57df, []int{5}
}
func (m *CollectionPvtReadWriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionPvtReadWriteSet.Unmarshal(m, b)
}
func (m *CollectionPvtReadWriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionPvtReadWriteSet.Marshal(b, m, deterministic)
}
func (dst *CollectionPvtReadWriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionPvtReadWriteSet.Merge(dst, src)
}
func (m *CollectionPvtReadWriteSet) XXX_Size() int {
	return xxx_messageInfo_CollectionPvtReadWriteSet.Size(m)
}
func (m *CollectionPvtReadWriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionPvtReadWriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionPvtReadWriteSet proto.InternalMessageInfo

func (m *CollectionPvtReadWriteSet) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *CollectionPvtReadWriteSet) GetRwset() []byte {
	if m != nil {
		return m.Rwset
	}
	return nil
}

func init() {
	proto.RegisterType((*TxReadWriteSet)(nil), "rwset.TxReadWriteSet")
	proto.RegisterType((*NsReadWriteSet)(nil), "rwset.NsReadWriteSet")
	proto.RegisterType((*CollectionHashedReadWriteSet)(nil), "rwset.CollectionHashedReadWriteSet")
	proto.RegisterType((*TxPvtReadWriteSet)(nil), "rwset.TxPvtReadWriteSet")
	proto.RegisterType((*NsPvtReadWriteSet)(nil), "rwset.NsPvtReadWriteSet")
	proto.RegisterType((*CollectionPvtReadWriteSet)(nil), "rwset.CollectionPvtReadWriteSet")
	proto.RegisterEnum("rwset.TxReadWriteSet_DataModel", TxReadWriteSet_DataModel_name, TxReadWriteSet_DataModel_value)
}

func init() { proto.RegisterFile("ledger/rwset/rwset.proto", fileDescriptor_rwset_28a2ad70bffb57df) }

var fileDescriptor_rwset_28a2ad70bffb57df = []byte{
//gzip文件描述符或协议的421字节
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x71, 0xab, 0x16, 0xf2, 0x36, 0x0a, 0xd4, 0xb4, 0x22, 0x48, 0x95, 0x28, 0x05, 0x89,
	0x8a, 0x43, 0x02, 0xe5, 0xc6, 0x81, 0x03, 0x70, 0x40, 0x42, 0x54, 0xc8, 0x54, 0x4c, 0xea, 0x0e,
	0x91, 0x9b, 0x78, 0x4d, 0xa4, 0xfc, 0x53, 0xec, 0x75, 0xdd, 0x07, 0xd8, 0x79, 0xbb, 0xed, 0xbc,
	0x6f, 0x3a, 0xd5, 0x4e, 0xd3, 0x24, 0xdd, 0xbf, 0xc3, 0x2e, 0x51, 0xfc, 0xfa, 0x79, 0xf2, 0xfc,
	0xec, 0x37, 0x2f, 0x98, 0x21, 0xf3, 0x96, 0x2c, 0xb3, 0xb3, 0x13, 0xce, 0x84, 0x7a, 0x5a, 0x69,
	0x96, 0x88, 0x04, 0xb7, 0xe4, 0x62, 0x74, 0x89, 0xc0, 0x98, 0xad, 0x09, 0xa3, 0xde, 0x41, 0x16,
	0x08, 0xf6, 0x8f, 0x09, 0xfc, 0x0d, 0xc0, 0xa3, 0x82, 0x3a, 0x51, 0xe2, 0xb1, 0xd0, 0x44, 0x43,
	0x34, 0x36, 0x26, 0x6f, 0x2c, 0xe5, 0xad, 0x4a, 0xad, 0x9f, 0x54, 0xd0, 0x3f, 0x1b, 0x19, 0xd1,
	0xbc, 0xed, 0x2b, 0xfe, 0x04, 0xcf, 0x62, 0xee, 0x48, 0xbd, 0xd9, 0x18, 0x36, 0xc7, 0x9d, 0x49,
	0x3f, 0x77, 0x4f, 0x79, 0xd9, 0x4d, 0x9e, 0xc6, 0x9c, 0x48, 0x88, 0x97, 0xa0, 0x15, 0x5f, 0xc2,
	0x6d, 0x68, 0xfc, 0xfe, 0xff, 0xe2, 0xc9, 0xe8, 0x0a, 0x81, 0x51, 0x35, 0xe0, 0x01, 0x68, 0x31,
	0x8d, 0x18, 0x4f, 0xa9, 0xcb, 0x24, 0x98, 0x46, 0x76, 0x05, 0xdc, 0x83, 0xd6, 0x36, 0x14, 0x8d,
	0x75, 0xa2, 0x16, 0xf8, 0x10, 0x5e, 0xb9, 0x49, 0x18, 0x32, 0x57, 0x04, 0x49, 0xec, 0xf8, 0x94,
	0xfb, 0xcc, 0xcb, 0xe1, 0x9a, 0x12, 0xee, 0x5d, 0x0e, 0xf7, 0xa3, 0x50, 0xfd, 0x92, 0xa2, 0x0a,
	0x6a, 0xdf, 0xad, 0xef, 0x4a, 0xf0, 0x0b, 0x04, 0x83, 0xbb, 0x7c, 0xf8, 0x03, 0x3c, 0x2f, 0xa5,
	0x6f, 0x58, 0x73, 0x6e, 0x63, 0x57, 0x9e, 0xd2, 0x88, 0xe1, 0xb7, 0xa0, 0x57, 0xd8, 0xd4, 0x19,
	0x3a, 0xfe, 0x2e, 0x0c, 0xbf, 0x07, 0x23, 0x5d, 0x09, 0xb5, 0x2f, 0x0f, 0x62, 0x36, 0xa5, 0x48,
	0x4f, 0x57, 0x42, 0x2a, 0x36, 0xf9, 0xa3, 0x73, 0x04, 0xdd, 0xd9, 0xfa, 0xef, 0x4a, 0x3c, 0x6a,
	0x4f, 0xbf, 0x82, 0x1e, 0x73, 0xa7, 0x88, 0xcf, 0xfb, 0x6a, 0x16, 0x7d, 0xad, 0xe5, 0x11, 0x88,
	0x65, 0x49, 0x5e, 0xd2, 0x19, 0x82, 0xee, 0x9e, 0xe2, 0x9e, 0x5e, 0x12, 0xe8, 0x95, 0xee, 0xad,
	0x9e, 0x3b, 0xdc, 0x6b, 0x59, 0x3d, 0x1f, 0xbb, 0x95, 0x2d, 0xc9, 0x31, 0x87, 0xd7, 0xb7, 0x1a,
	0x1e, 0xde, 0xa8, 0x1b, 0xff, 0xb2, 0xef, 0x0e, 0x7c, 0x4c, 0xb2, 0xa5, 0xe5, 0x9f, 0xa6, 0x2c,
	0x53, 0x23, 0x67, 0x1d, 0xd1, 0x45, 0x16, 0xb8, 0x6a, 0xda, 0xb8, 0x95, 0x17, 0xa5, 0x7a, 0xfe,
	0x79, 0x19, 0x08, 0xff, 0x78, 0x61, 0xb9, 0x49, 0x64, 0x97, 0x2c, 0xb6, 0xb2, 0xd8, 0xca, 0x62,
	0x97, 0x47, 0x77, 0xd1, 0x96, 0xc5, 0x2f, 0xd7, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x8a, 0xfa,
	0x65, 0xd1, 0x03, 0x00, 0x00,
}

