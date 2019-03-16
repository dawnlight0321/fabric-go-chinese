
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:22</date>
//</624456065437667328>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package comm

import (
	"sync"

	"github.com/hyperledger/fabric/gossip/common"
)

//ChannelDisultiplexer是一个可以接收通道注册（addChannel）的结构。
//和出版物（解编），并将出版物广播给登记处。
//根据其谓词
type ChannelDeMultiplexer struct {
	channels []*channel
	lock     *sync.RWMutex
	closed   bool
}

//新建信道解复用器创建新信道解复用器
func NewChannelDemultiplexer() *ChannelDeMultiplexer {
	return &ChannelDeMultiplexer{
		channels: make([]*channel, 0),
		lock:     &sync.RWMutex{},
	}
}

type channel struct {
	pred common.MessageAcceptor
	ch   chan interface{}
}

func (m *ChannelDeMultiplexer) isClosed() bool {
	return m.closed
}

//关闭关闭此频道，使所有频道在之前注册
//也可以关闭。
func (m *ChannelDeMultiplexer) Close() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.closed = true
	for _, ch := range m.channels {
		close(ch.ch)
	}
	m.channels = nil
}

//addchannel用某个谓词注册一个通道
func (m *ChannelDeMultiplexer) AddChannel(predicate common.MessageAcceptor) chan interface{} {
	m.lock.Lock()
	defer m.lock.Unlock()
	ch := &channel{ch: make(chan interface{}, 10), pred: predicate}
	m.channels = append(m.channels, ch)
	return ch.ch
}

//解复用将消息广播到返回的所有通道
//通过addchannel调用并保存受尊重的谓词。
func (m *ChannelDeMultiplexer) DeMultiplex(msg interface{}) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if m.isClosed() {
		return
	}
	for _, ch := range m.channels {
		if ch.pred(msg) {
			ch.ch <- msg
		}
	}
}

