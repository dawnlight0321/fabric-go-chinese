
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:26</date>
//</624456082261020672>

/*
版权所有IBM Corp.2017保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

   http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/


package mgmt

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/msp"
	mspproto "github.com/hyperledger/fabric/protos/msp"
	"github.com/pkg/errors"
)

//DeserializerManager是
//访问本地和通道反序列化程序
type DeserializersManager interface {

//反序列化接收SerializedEntity字节并返回未解析的窗体
//或失败时出错。
	Deserialize(raw []byte) (*mspproto.SerializedIdentity, error)

//getlocalmspidentifier返回本地msp标识符
	GetLocalMSPIdentifier() string

//GetLocalDeserializer返回本地标识反序列化程序
	GetLocalDeserializer() msp.IdentityDeserializer

//getChannelDeserializers返回通道反序列化器的映射
	GetChannelDeserializers() map[string]msp.IdentityDeserializer
}

//反序列化程序管理器返回反序列化程序管理器的新实例
func NewDeserializersManager() DeserializersManager {
	return &mspDeserializersManager{}
}

type mspDeserializersManager struct{}

func (m *mspDeserializersManager) Deserialize(raw []byte) (*mspproto.SerializedIdentity, error) {
	sId := &mspproto.SerializedIdentity{}
	err := proto.Unmarshal(raw, sId)
	if err != nil {
		return nil, errors.Wrap(err, "could not deserialize a SerializedIdentity")
	}
	return sId, nil
}

func (m *mspDeserializersManager) GetLocalMSPIdentifier() string {
	id, _ := GetLocalMSP().GetIdentifier()
	return id
}

func (m *mspDeserializersManager) GetLocalDeserializer() msp.IdentityDeserializer {
	return GetLocalMSP()
}

func (m *mspDeserializersManager) GetChannelDeserializers() map[string]msp.IdentityDeserializer {
	return GetDeserializers()
}

