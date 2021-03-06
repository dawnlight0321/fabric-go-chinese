
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:35</date>
//</624456119674212352>

//由Protoc Gen Go生成的代码。不要编辑。
//来源：order/etcdraft/configuration.proto

package etcdraft //导入“github.com/hyperledger/fabric/protos/order/etcdraft”

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

//元数据已序列化并设置为ConsensusType的值。中的元数据
//当consensistype.type设置为“etcdraft”时的信道配置。
type Metadata struct {
	Consenters           []*Consenter `protobuf:"bytes,1,rep,name=consenters,proto3" json:"consenters,omitempty"`
	Options              *Options     `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_7287b6fe5795d8b9, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetConsenters() []*Consenter {
	if m != nil {
		return m.Consenters
	}
	return nil
}

func (m *Metadata) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

//同意者表示同意节点（即副本）。
type Consenter struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ClientTlsCert        []byte   `protobuf:"bytes,3,opt,name=client_tls_cert,json=clientTlsCert,proto3" json:"client_tls_cert,omitempty"`
	ServerTlsCert        []byte   `protobuf:"bytes,4,opt,name=server_tls_cert,json=serverTlsCert,proto3" json:"server_tls_cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consenter) Reset()         { *m = Consenter{} }
func (m *Consenter) String() string { return proto.CompactTextString(m) }
func (*Consenter) ProtoMessage()    {}
func (*Consenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_7287b6fe5795d8b9, []int{1}
}
func (m *Consenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consenter.Unmarshal(m, b)
}
func (m *Consenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consenter.Marshal(b, m, deterministic)
}
func (dst *Consenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consenter.Merge(dst, src)
}
func (m *Consenter) XXX_Size() int {
	return xxx_messageInfo_Consenter.Size(m)
}
func (m *Consenter) XXX_DiscardUnknown() {
	xxx_messageInfo_Consenter.DiscardUnknown(m)
}

var xxx_messageInfo_Consenter proto.InternalMessageInfo

func (m *Consenter) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Consenter) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Consenter) GetClientTlsCert() []byte {
	if m != nil {
		return m.ClientTlsCert
	}
	return nil
}

func (m *Consenter) GetServerTlsCert() []byte {
	if m != nil {
		return m.ServerTlsCert
	}
	return nil
}

//为所有ETCD/RAFT节点指定选项。这些可以在
//每通道基础。
type Options struct {
	TickInterval         uint64   `protobuf:"varint,1,opt,name=tick_interval,json=tickInterval,proto3" json:"tick_interval,omitempty"`
	ElectionTick         uint32   `protobuf:"varint,2,opt,name=election_tick,json=electionTick,proto3" json:"election_tick,omitempty"`
	HeartbeatTick        uint32   `protobuf:"varint,3,opt,name=heartbeat_tick,json=heartbeatTick,proto3" json:"heartbeat_tick,omitempty"`
	MaxInflightMsgs      uint32   `protobuf:"varint,4,opt,name=max_inflight_msgs,json=maxInflightMsgs,proto3" json:"max_inflight_msgs,omitempty"`
	MaxSizePerMsg        uint64   `protobuf:"varint,5,opt,name=max_size_per_msg,json=maxSizePerMsg,proto3" json:"max_size_per_msg,omitempty"`
	SnapshotInterval     uint64   `protobuf:"varint,6,opt,name=snapshot_interval,json=snapshotInterval,proto3" json:"snapshot_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_7287b6fe5795d8b9, []int{2}
}
func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (dst *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(dst, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetTickInterval() uint64 {
	if m != nil {
		return m.TickInterval
	}
	return 0
}

func (m *Options) GetElectionTick() uint32 {
	if m != nil {
		return m.ElectionTick
	}
	return 0
}

func (m *Options) GetHeartbeatTick() uint32 {
	if m != nil {
		return m.HeartbeatTick
	}
	return 0
}

func (m *Options) GetMaxInflightMsgs() uint32 {
	if m != nil {
		return m.MaxInflightMsgs
	}
	return 0
}

func (m *Options) GetMaxSizePerMsg() uint64 {
	if m != nil {
		return m.MaxSizePerMsg
	}
	return 0
}

func (m *Options) GetSnapshotInterval() uint64 {
	if m != nil {
		return m.SnapshotInterval
	}
	return 0
}

//raftmetadata存储raft osn在以下情况下使用的数据：
//相互协调，连载成
//阻止meta dta字段，并在failres和restart之后使用。
type RaftMetadata struct {
//在集群的OSN之间维护映射
//还有他们的救生筏ID。
	Consenters map[uint64]*Consenter `protobuf:"bytes,1,rep,name=consenters,proto3" json:"consenters,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
//携带将分配的筏板ID值
//到将加入该集群的下一个OSN。
	NextConsenterId uint64 `protobuf:"varint,2,opt,name=next_consenter_id,json=nextConsenterId,proto3" json:"next_consenter_id,omitempty"`
//筏组配置计数
	ConfChangeCounts uint64 `protobuf:"varint,3,opt,name=conf_change_counts,json=confChangeCounts,proto3" json:"conf_change_counts,omitempty"`
//当前区块ETCD/筏入口索引。
	RaftIndex            uint64   `protobuf:"varint,4,opt,name=raft_index,json=raftIndex,proto3" json:"raft_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RaftMetadata) Reset()         { *m = RaftMetadata{} }
func (m *RaftMetadata) String() string { return proto.CompactTextString(m) }
func (*RaftMetadata) ProtoMessage()    {}
func (*RaftMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_7287b6fe5795d8b9, []int{3}
}
func (m *RaftMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RaftMetadata.Unmarshal(m, b)
}
func (m *RaftMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RaftMetadata.Marshal(b, m, deterministic)
}
func (dst *RaftMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftMetadata.Merge(dst, src)
}
func (m *RaftMetadata) XXX_Size() int {
	return xxx_messageInfo_RaftMetadata.Size(m)
}
func (m *RaftMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RaftMetadata proto.InternalMessageInfo

func (m *RaftMetadata) GetConsenters() map[uint64]*Consenter {
	if m != nil {
		return m.Consenters
	}
	return nil
}

func (m *RaftMetadata) GetNextConsenterId() uint64 {
	if m != nil {
		return m.NextConsenterId
	}
	return 0
}

func (m *RaftMetadata) GetConfChangeCounts() uint64 {
	if m != nil {
		return m.ConfChangeCounts
	}
	return 0
}

func (m *RaftMetadata) GetRaftIndex() uint64 {
	if m != nil {
		return m.RaftIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*Metadata)(nil), "etcdraft.Metadata")
	proto.RegisterType((*Consenter)(nil), "etcdraft.Consenter")
	proto.RegisterType((*Options)(nil), "etcdraft.Options")
	proto.RegisterType((*RaftMetadata)(nil), "etcdraft.RaftMetadata")
	proto.RegisterMapType((map[uint64]*Consenter)(nil), "etcdraft.RaftMetadata.ConsentersEntry")
}

func init() {
	proto.RegisterFile("orderer/etcdraft/configuration.proto", fileDescriptor_configuration_7287b6fe5795d8b9)
}

var fileDescriptor_configuration_7287b6fe5795d8b9 = []byte{
//gzip文件描述符或协议的535字节
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x4d, 0x6b, 0x1b, 0x3d,
	0x14, 0x85, 0x99, 0x78, 0xf2, 0xa5, 0x78, 0x5e, 0xdb, 0x7a, 0x37, 0xa6, 0x50, 0x30, 0x6e, 0x9b,
	0xba, 0x49, 0x99, 0x81, 0x84, 0x42, 0xe9, 0xb2, 0xa6, 0x05, 0x2f, 0x4c, 0x8b, 0x9a, 0x55, 0x37,
	0x42, 0xd6, 0x5c, 0xcf, 0x08, 0x8f, 0xa5, 0x41, 0x92, 0x8d, 0x9d, 0x6d, 0x7f, 0x4a, 0xff, 0x66,
	0x17, 0x45, 0x9a, 0x0f, 0xbb, 0x21, 0x3b, 0x71, 0xce, 0x73, 0x2f, 0x47, 0xba, 0x57, 0xe8, 0xb5,
	0xd2, 0x29, 0x68, 0xd0, 0x09, 0x58, 0x9e, 0x6a, 0xb6, 0xb4, 0x09, 0x57, 0x72, 0x29, 0xb2, 0x8d,
	0x66, 0x56, 0x28, 0x19, 0x97, 0x5a, 0x59, 0x85, 0x2f, 0x1a, 0x77, 0x5c, 0xa0, 0x8b, 0x39, 0x58,
	0x96, 0x32, 0xcb, 0xf0, 0x3d, 0x42, 0x5c, 0x49, 0x03, 0xd2, 0x82, 0x36, 0xc3, 0x60, 0xd4, 0x99,
	0x5c, 0xdd, 0xfd, 0x1f, 0x37, 0x68, 0x3c, 0x6d, 0x3c, 0x72, 0x84, 0xe1, 0x5b, 0x74, 0xae, 0x4a,
	0xd7, 0xda, 0x0c, 0x4f, 0x46, 0xc1, 0xe4, 0xea, 0x6e, 0x70, 0xa8, 0xf8, 0x56, 0x19, 0xa4, 0x21,
	0xc6, 0xbf, 0x02, 0x74, 0xd9, 0xb6, 0xc1, 0x18, 0x85, 0xb9, 0x32, 0x76, 0x18, 0x8c, 0x82, 0xc9,
	0x25, 0xf1, 0x67, 0xa7, 0x95, 0x4a, 0x5b, 0xdf, 0x2b, 0x22, 0xfe, 0x8c, 0xaf, 0x51, 0x8f, 0x17,
	0x02, 0xa4, 0xa5, 0xb6, 0x30, 0x94, 0x83, 0xb6, 0xc3, 0xce, 0x28, 0x98, 0x74, 0x49, 0x54, 0xc9,
	0x0f, 0x85, 0x99, 0x42, 0xc5, 0x19, 0xd0, 0x5b, 0xd0, 0x07, 0x2e, 0xac, 0xb8, 0x4a, 0xae, 0xb9,
	0xf1, 0x9f, 0x00, 0x9d, 0xd7, 0xd1, 0xf0, 0x2b, 0x14, 0x59, 0xc1, 0x57, 0x54, 0xb8, 0x44, 0x5b,
	0x56, 0xf8, 0x30, 0x21, 0xe9, 0x3a, 0x71, 0x56, 0x6b, 0x0e, 0x82, 0x02, 0xb8, 0xab, 0xa0, 0xce,
	0xa8, 0xd3, 0x75, 0x1b, 0xf1, 0x41, 0xf0, 0x15, 0x7e, 0x83, 0xfe, 0xcb, 0x81, 0x69, 0xbb, 0x00,
	0x66, 0x2b, 0xaa, 0xe3, 0xa9, 0xa8, 0x55, 0x3d, 0x76, 0x83, 0x06, 0x6b, 0xb6, 0xa3, 0x42, 0x2e,
	0x0b, 0x91, 0xe5, 0x96, 0xae, 0x4d, 0x66, 0x7c, 0xcc, 0x88, 0xf4, 0xd6, 0x6c, 0x37, 0xab, 0xf5,
	0xb9, 0xc9, 0x0c, 0x7e, 0x8b, 0xfa, 0x8e, 0x35, 0xe2, 0x11, 0x68, 0x09, 0xda, 0xb1, 0xc3, 0x53,
	0x9f, 0x2f, 0x5a, 0xb3, 0xdd, 0x0f, 0xf1, 0x08, 0xdf, 0x41, 0xcf, 0x4d, 0x86, 0x6f, 0xd1, 0xc0,
	0x48, 0x56, 0x9a, 0x5c, 0xd9, 0xc3, 0x4d, 0xce, 0x3c, 0xd9, 0x6f, 0x8c, 0xe6, 0x36, 0xe3, 0xdf,
	0x27, 0xa8, 0x4b, 0xd8, 0xd2, 0xb6, 0x73, 0xff, 0xfa, 0xcc, 0xdc, 0xaf, 0x0f, 0x53, 0x3c, 0x66,
	0x0f, 0x4b, 0x60, 0xbe, 0x48, 0xab, 0xf7, 0xff, 0xac, 0xc2, 0x0d, 0x1a, 0x48, 0xd8, 0x59, 0xda,
	0x4a, 0x54, 0xa4, 0xfe, 0xa9, 0x42, 0xd2, 0x73, 0x46, 0x5b, 0x3b, 0x4b, 0xf1, 0x7b, 0x84, 0xdd,
	0x62, 0x52, 0x9e, 0x33, 0x99, 0x01, 0xe5, 0x6a, 0x23, 0xad, 0xf1, 0x2f, 0x16, 0x92, 0xbe, 0x73,
	0xa6, 0xde, 0x98, 0x7a, 0x1d, 0xbf, 0x44, 0xc8, 0x45, 0xa1, 0x42, 0xa6, 0xb0, 0xf3, 0xaf, 0x15,
	0x92, 0x4b, 0xa7, 0xcc, 0x9c, 0xf0, 0x82, 0xa0, 0xde, 0x93, 0x5c, 0xb8, 0x8f, 0x3a, 0x2b, 0xd8,
	0xd7, 0xd3, 0x74, 0x47, 0xfc, 0x0e, 0x9d, 0x6e, 0x59, 0xb1, 0x81, 0x7a, 0x4d, 0x9f, 0x5d, 0xec,
	0x8a, 0xf8, 0x74, 0xf2, 0x31, 0xf8, 0x9c, 0xa1, 0x58, 0xe9, 0x2c, 0xce, 0xf7, 0x25, 0xe8, 0x02,
	0xd2, 0x0c, 0x74, 0xbc, 0x64, 0x0b, 0x2d, 0x78, 0xf5, 0x85, 0x4c, 0x5c, 0x7f, 0xb4, 0xb6, 0xcd,
	0xcf, 0x0f, 0x99, 0xb0, 0xf9, 0x66, 0x11, 0x73, 0xb5, 0x4e, 0x8e, 0xca, 0x92, 0xaa, 0x2c, 0xa9,
	0xca, 0x92, 0xa7, 0xff, 0x73, 0x71, 0xe6, 0x8d, 0xfb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x56,
	0xd3, 0xd9, 0x86, 0xba, 0x03, 0x00, 0x00,
}

