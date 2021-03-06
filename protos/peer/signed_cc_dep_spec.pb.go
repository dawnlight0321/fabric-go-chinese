
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:36</date>
//</624456121964302336>

//由Protoc Gen Go生成的代码。不要编辑。
//来源：peer/signed_cc_dep_spec.proto

package peer //导入“github.com/hyperledger/fabric/protos/peer”

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

//签名的haincodedeploymentspec携带CD和背书
type SignedChaincodeDeploymentSpec struct {
//这是chaincodedeploymentspec的字节
	ChaincodeDeploymentSpec []byte `protobuf:"bytes,1,opt,name=chaincode_deployment_spec,json=chaincodeDeploymentSpec,proto3" json:"chaincode_deployment_spec,omitempty"`
//这是在结构上相同的实例化策略
//背书政策。此策略在提交时由VSCC检查
//实例化的时间（所有对等方将获得与其相同的策略
//将是LSCC实例化记录的一部分，并且将是
//散列也一样）
	InstantiationPolicy []byte `protobuf:"bytes,2,opt,name=instantiation_policy,json=instantiationPolicy,proto3" json:"instantiation_policy,omitempty"`
//以上部署规范的认可，业主签字
//链码\部署\规范和背书。背书人。
	OwnerEndorsements    []*Endorsement `protobuf:"bytes,3,rep,name=owner_endorsements,json=ownerEndorsements,proto3" json:"owner_endorsements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SignedChaincodeDeploymentSpec) Reset()         { *m = SignedChaincodeDeploymentSpec{} }
func (m *SignedChaincodeDeploymentSpec) String() string { return proto.CompactTextString(m) }
func (*SignedChaincodeDeploymentSpec) ProtoMessage()    {}
func (*SignedChaincodeDeploymentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_signed_cc_dep_spec_80ee01f5d5123140, []int{0}
}
func (m *SignedChaincodeDeploymentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedChaincodeDeploymentSpec.Unmarshal(m, b)
}
func (m *SignedChaincodeDeploymentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedChaincodeDeploymentSpec.Marshal(b, m, deterministic)
}
func (dst *SignedChaincodeDeploymentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedChaincodeDeploymentSpec.Merge(dst, src)
}
func (m *SignedChaincodeDeploymentSpec) XXX_Size() int {
	return xxx_messageInfo_SignedChaincodeDeploymentSpec.Size(m)
}
func (m *SignedChaincodeDeploymentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedChaincodeDeploymentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_SignedChaincodeDeploymentSpec proto.InternalMessageInfo

func (m *SignedChaincodeDeploymentSpec) GetChaincodeDeploymentSpec() []byte {
	if m != nil {
		return m.ChaincodeDeploymentSpec
	}
	return nil
}

func (m *SignedChaincodeDeploymentSpec) GetInstantiationPolicy() []byte {
	if m != nil {
		return m.InstantiationPolicy
	}
	return nil
}

func (m *SignedChaincodeDeploymentSpec) GetOwnerEndorsements() []*Endorsement {
	if m != nil {
		return m.OwnerEndorsements
	}
	return nil
}

func init() {
	proto.RegisterType((*SignedChaincodeDeploymentSpec)(nil), "protos.SignedChaincodeDeploymentSpec")
}

func init() {
	proto.RegisterFile("peer/signed_cc_dep_spec.proto", fileDescriptor_signed_cc_dep_spec_80ee01f5d5123140)
}

var fileDescriptor_signed_cc_dep_spec_80ee01f5d5123140 = []byte{
//gzip文件描述符或协议的255字节
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xa9, 0x05, 0x0f, 0xab, 0x17, 0x53, 0xc1, 0x28, 0x16, 0x4a, 0x4f, 0xf5, 0x92, 0xa0,
	0xde, 0x3c, 0x56, 0x3d, 0x2b, 0xed, 0xcd, 0xcb, 0x92, 0x4c, 0xc6, 0x64, 0x21, 0xdd, 0x19, 0x66,
	0x56, 0x24, 0xaf, 0xe9, 0x13, 0x49, 0x76, 0x2d, 0xd6, 0x83, 0xa7, 0x85, 0xfd, 0xbe, 0x7f, 0x66,
	0xf8, 0xcd, 0x9c, 0x11, 0xa5, 0x54, 0xd7, 0x7a, 0x6c, 0x2c, 0x80, 0x6d, 0x90, 0xad, 0x32, 0x42,
	0xc1, 0x42, 0x81, 0xb2, 0xe3, 0xf8, 0xe8, 0xd5, 0x75, 0xd4, 0x58, 0x88, 0x49, 0xab, 0xde, 0x0a,
	0x2a, 0x93, 0x57, 0x4c, 0xd6, 0xf2, 0x6b, 0x62, 0xe6, 0xdb, 0x38, 0xe2, 0xb1, 0xab, 0x9c, 0x07,
	0x6a, 0xf0, 0x09, 0xb9, 0xa7, 0x61, 0x87, 0x3e, 0x6c, 0x19, 0x21, 0x7b, 0x30, 0x97, 0xb0, 0x47,
	0xe3, 0x8e, 0x1f, 0x16, 0x57, 0xe5, 0x93, 0xc5, 0x64, 0x75, 0xba, 0xb9, 0x80, 0x7f, 0xb2, 0xb7,
	0xe6, 0xdc, 0x79, 0x0d, 0x95, 0x0f, 0xae, 0x0a, 0x8e, 0xbc, 0x65, 0xea, 0x1d, 0x0c, 0xf9, 0x51,
	0x8c, 0xcd, 0xfe, 0xb0, 0xd7, 0x88, 0xb2, 0xb5, 0xc9, 0xe8, 0xd3, 0xa3, 0x58, 0xf4, 0x0d, 0x89,
	0xe2, 0x38, 0x4b, 0xf3, 0xe9, 0x62, 0xba, 0x3a, 0xb9, 0x9b, 0xa5, 0xa3, 0xb5, 0x78, 0xfe, 0x65,
	0x9b, 0xb3, 0xa8, 0x1f, 0xfc, 0xe8, 0xfa, 0xc5, 0x2c, 0x49, 0xda, 0xa2, 0x1b, 0x18, 0xa5, 0xc7,
	0xa6, 0x45, 0x29, 0xde, 0xab, 0x5a, 0x1c, 0xec, 0xf3, 0x63, 0x25, 0x6f, 0x37, 0xad, 0x0b, 0xdd,
	0x47, 0x5d, 0x00, 0xed, 0xca, 0x03, 0xb5, 0x4c, 0x6a, 0x99, 0xd4, 0x72, 0x54, 0xeb, 0xd4, 0xe5,
	0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x40, 0x4c, 0x9e, 0x73, 0x01, 0x00, 0x00,
}

