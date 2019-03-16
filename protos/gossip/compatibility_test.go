
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:34</date>
//</624456115656069120>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package gossip

import (
	"encoding/hex"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

var digestMsg = &GossipMessage{
	Channel: []byte("mychannel"),
	Content: &GossipMessage_DataDig{
		DataDig: &DataDigest{
			Digests: [][]byte{
				{255},
				{255, 255},
				{255, 255, 255},
				{255, 255, 255, 255},
				[]byte("100"),
			},
		},
	},
}

var requestMsg = &GossipMessage{
	Channel: []byte("mychannel"),
	Content: &GossipMessage_DataReq{
		DataReq: &DataRequest{
			Digests: [][]byte{
				{255},
				{255, 255},
				{255, 255, 255},
				{255, 255, 255, 255},
				[]byte("100"),
			},
		},
	},
}

const (
	v12DataDigestBytes  = "12096d796368616e6e656c52171201ff1202ffff1203ffffff1204ffffffff1203313030"
	v12DataRequestBytes = "12096d796368616e6e656c5a171201ff1202ffff1203ffffff1204ffffffff1203313030"
)

func TestUnmarshalV12Digests(t *testing.T) {
//此测试确保数据摘要消息和数据请求的摘要
//源于结构v1.3的内容可以被v1.2成功解析
	for msgBytes, expectedMsg := range map[string]*GossipMessage{
		v12DataDigestBytes:  digestMsg,
		v12DataRequestBytes: requestMsg,
	} {
		var err error
		v13Envelope := &Envelope{}
		v13Envelope.Payload, err = hex.DecodeString(msgBytes)
		assert.NoError(t, err)
		sMsg := &SignedGossipMessage{
			Envelope: v13Envelope,
		}
		v13Digest, err := sMsg.ToGossipMessage()
		assert.NoError(t, err)
		assert.True(t, proto.Equal(expectedMsg, v13Digest.GossipMessage))
	}
}

