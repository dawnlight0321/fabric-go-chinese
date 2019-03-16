
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:36</date>
//</624456125550432256>

//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/protos/common"
	ab "github.com/hyperledger/fabric/protos/orderer"
	"github.com/hyperledger/fabric/token/client"
)

type Broadcast struct {
	SendStub        func(m *common.Envelope) error
	sendMutex       sync.RWMutex
	sendArgsForCall []struct {
		m *common.Envelope
	}
	sendReturns struct {
		result1 error
	}
	sendReturnsOnCall map[int]struct {
		result1 error
	}
	RecvStub        func() (*ab.BroadcastResponse, error)
	recvMutex       sync.RWMutex
	recvArgsForCall []struct{}
	recvReturns     struct {
		result1 *ab.BroadcastResponse
		result2 error
	}
	recvReturnsOnCall map[int]struct {
		result1 *ab.BroadcastResponse
		result2 error
	}
	CloseSendStub        func() error
	closeSendMutex       sync.RWMutex
	closeSendArgsForCall []struct{}
	closeSendReturns     struct {
		result1 error
	}
	closeSendReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Broadcast) Send(m *common.Envelope) error {
	fake.sendMutex.Lock()
	ret, specificReturn := fake.sendReturnsOnCall[len(fake.sendArgsForCall)]
	fake.sendArgsForCall = append(fake.sendArgsForCall, struct {
		m *common.Envelope
	}{m})
	fake.recordInvocation("Send", []interface{}{m})
	fake.sendMutex.Unlock()
	if fake.SendStub != nil {
		return fake.SendStub(m)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sendReturns.result1
}

func (fake *Broadcast) SendCallCount() int {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return len(fake.sendArgsForCall)
}

func (fake *Broadcast) SendArgsForCall(i int) *common.Envelope {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return fake.sendArgsForCall[i].m
}

func (fake *Broadcast) SendReturns(result1 error) {
	fake.SendStub = nil
	fake.sendReturns = struct {
		result1 error
	}{result1}
}

func (fake *Broadcast) SendReturnsOnCall(i int, result1 error) {
	fake.SendStub = nil
	if fake.sendReturnsOnCall == nil {
		fake.sendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Broadcast) Recv() (*ab.BroadcastResponse, error) {
	fake.recvMutex.Lock()
	ret, specificReturn := fake.recvReturnsOnCall[len(fake.recvArgsForCall)]
	fake.recvArgsForCall = append(fake.recvArgsForCall, struct{}{})
	fake.recordInvocation("Recv", []interface{}{})
	fake.recvMutex.Unlock()
	if fake.RecvStub != nil {
		return fake.RecvStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.recvReturns.result1, fake.recvReturns.result2
}

func (fake *Broadcast) RecvCallCount() int {
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	return len(fake.recvArgsForCall)
}

func (fake *Broadcast) RecvReturns(result1 *ab.BroadcastResponse, result2 error) {
	fake.RecvStub = nil
	fake.recvReturns = struct {
		result1 *ab.BroadcastResponse
		result2 error
	}{result1, result2}
}

func (fake *Broadcast) RecvReturnsOnCall(i int, result1 *ab.BroadcastResponse, result2 error) {
	fake.RecvStub = nil
	if fake.recvReturnsOnCall == nil {
		fake.recvReturnsOnCall = make(map[int]struct {
			result1 *ab.BroadcastResponse
			result2 error
		})
	}
	fake.recvReturnsOnCall[i] = struct {
		result1 *ab.BroadcastResponse
		result2 error
	}{result1, result2}
}

func (fake *Broadcast) CloseSend() error {
	fake.closeSendMutex.Lock()
	ret, specificReturn := fake.closeSendReturnsOnCall[len(fake.closeSendArgsForCall)]
	fake.closeSendArgsForCall = append(fake.closeSendArgsForCall, struct{}{})
	fake.recordInvocation("CloseSend", []interface{}{})
	fake.closeSendMutex.Unlock()
	if fake.CloseSendStub != nil {
		return fake.CloseSendStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeSendReturns.result1
}

func (fake *Broadcast) CloseSendCallCount() int {
	fake.closeSendMutex.RLock()
	defer fake.closeSendMutex.RUnlock()
	return len(fake.closeSendArgsForCall)
}

func (fake *Broadcast) CloseSendReturns(result1 error) {
	fake.CloseSendStub = nil
	fake.closeSendReturns = struct {
		result1 error
	}{result1}
}

func (fake *Broadcast) CloseSendReturnsOnCall(i int, result1 error) {
	fake.CloseSendStub = nil
	if fake.closeSendReturnsOnCall == nil {
		fake.closeSendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeSendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Broadcast) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	fake.closeSendMutex.RLock()
	defer fake.closeSendMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Broadcast) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ client.Broadcast = new(Broadcast)
