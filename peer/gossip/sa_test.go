
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:33</date>
//</624456112166408192>

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


package gossip

import (
	"testing"

	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/peer/gossip/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMspSecurityAdvisor_OrgByPeerIdentity(t *testing.T) {
	dm := &mocks.DeserializersManager{
		LocalDeserializer: &mocks.IdentityDeserializer{Identity: []byte("Alice"), Msg: []byte("msg1"), Mock: mock.Mock{}},
		ChannelDeserializers: map[string]msp.IdentityDeserializer{
			"A": &mocks.IdentityDeserializer{Identity: []byte("Bob"), Msg: []byte("msg2"), Mock: mock.Mock{}},
		},
	}

	advisor := NewSecurityAdvisor(dm)
	assert.NotNil(t, advisor.OrgByPeerIdentity([]byte("Alice")))
	assert.NotNil(t, advisor.OrgByPeerIdentity([]byte("Bob")))
	assert.Nil(t, advisor.OrgByPeerIdentity([]byte("Charlie")))
	assert.Nil(t, advisor.OrgByPeerIdentity(nil))
}

