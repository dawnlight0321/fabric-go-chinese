
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:22</date>
//</624456064376508416>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package api

import (
	"github.com/hyperledger/fabric/gossip/common"
)

func init() {
//这只是为了满足代码覆盖工具
//错过任何方法
	switch true {

	}
}

//去：生成mokery-dir。-name securityadvisor-case下划线-输出../mocks/

//SecurityAdvisor定义外部辅助对象
//提供安全和身份相关功能
type SecurityAdvisor interface {
//orgByPeerIdentity返回orgIdentityType
//一个给定的对等身份。
//如果出现任何错误，则返回nil。
//此方法不验证对等标识。
//这个验证应该在执行流程中适当地完成。
	OrgByPeerIdentity(PeerIdentityType) OrgIdentityType
}

//channelnotifier由八卦组件实现，用于对等端
//通知JoinChannel事件的八卦组件的层
type ChannelNotifier interface {
	JoinChannel(joinMsg JoinChannelMessage, chainID common.ChainID)
}

//JoinChannelMessage是断言创建或突变的消息
//一个频道的成员名单，是八卦消息
//在同行中
type JoinChannelMessage interface {

//SequenceNumber返回配置块的序列号
//JoinChannelMessage源于
	SequenceNumber() uint64

//成员返回频道的组织
	Members() []OrgIdentityType

//anchor peers of返回给定组织的锚定对等方
	AnchorPeersOf(org OrgIdentityType) []AnchorPeer
}

//anchor peer是锚定对等的证书和终结点（主机：端口）
type AnchorPeer struct {
Host string //主机是远程对等机的主机名/IP地址
Port int    //端口是远程对等机正在侦听的端口
}

//OrgIdentityType定义组织的标识
type OrgIdentityType []byte

