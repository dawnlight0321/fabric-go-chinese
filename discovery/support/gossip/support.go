
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:18</date>
//</624456049088270336>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package gossip

import (
	"github.com/hyperledger/fabric/gossip/common"
	"github.com/hyperledger/fabric/gossip/discovery"
	gossip2 "github.com/hyperledger/fabric/gossip/gossip"
	"github.com/hyperledger/fabric/protos/gossip"
)

//DiscoverySupport实现用于服务发现的支持
//从流言蜚语中得到的
type DiscoverySupport struct {
	gossip2.Gossip
}

//新建DiscoverySupport创建新的DiscoverySupport
func NewDiscoverySupport(g gossip2.Gossip) *DiscoverySupport {
	return &DiscoverySupport{g}
}

//channel exists返回给定通道是否存在
func (s *DiscoverySupport) ChannelExists(channel string) bool {
	return s.SelfChannelInfo(common.ChainID(channel)) != nil
}

//peersofchannel返回被认为是活动的网络成员
//也订阅了给定的频道
func (s *DiscoverySupport) PeersOfChannel(chain common.ChainID) discovery.Members {
	msg := s.SelfChannelInfo(chain)
	if msg == nil {
		return nil
	}
	stateInf := msg.GetStateInfo()
	selfMember := discovery.NetworkMember{
		Properties: stateInf.Properties,
		PKIid:      stateInf.PkiId,
		Envelope:   msg.Envelope,
	}
	return append(s.Gossip.PeersOfChannel(chain), selfMember)
}

//对等方返回被认为是活动的网络成员
func (s *DiscoverySupport) Peers() discovery.Members {
	peers := s.Gossip.Peers()
	peers = append(peers, s.Gossip.SelfMembershipInfo())
//只返回具有外部端点的对等机，并清理信封。
	return discovery.Members(peers).Filter(discovery.HasExternalEndpoint).Map(sanitizeEnvelope)
}

func sanitizeEnvelope(member discovery.NetworkMember) discovery.NetworkMember {
//制作成员的本地副本
	returnedMember := member
	if returnedMember.Envelope == nil {
		return returnedMember
	}
	returnedMember.Envelope = &gossip.Envelope{
		Payload:   member.Envelope.Payload,
		Signature: member.Envelope.Signature,
	}
	return returnedMember
}

