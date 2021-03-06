
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:28</date>
//</624456089307451392>

//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/orderer/common/broadcast"
	"github.com/hyperledger/fabric/orderer/common/msgprocessor"
	cb "github.com/hyperledger/fabric/protos/common"
)

type ChannelSupport struct {
	ClassifyMsgStub        func(chdr *cb.ChannelHeader) msgprocessor.Classification
	classifyMsgMutex       sync.RWMutex
	classifyMsgArgsForCall []struct {
		chdr *cb.ChannelHeader
	}
	classifyMsgReturns struct {
		result1 msgprocessor.Classification
	}
	classifyMsgReturnsOnCall map[int]struct {
		result1 msgprocessor.Classification
	}
	ProcessNormalMsgStub        func(env *cb.Envelope) (configSeq uint64, err error)
	processNormalMsgMutex       sync.RWMutex
	processNormalMsgArgsForCall []struct {
		env *cb.Envelope
	}
	processNormalMsgReturns struct {
		result1 uint64
		result2 error
	}
	processNormalMsgReturnsOnCall map[int]struct {
		result1 uint64
		result2 error
	}
	ProcessConfigUpdateMsgStub        func(env *cb.Envelope) (config *cb.Envelope, configSeq uint64, err error)
	processConfigUpdateMsgMutex       sync.RWMutex
	processConfigUpdateMsgArgsForCall []struct {
		env *cb.Envelope
	}
	processConfigUpdateMsgReturns struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}
	processConfigUpdateMsgReturnsOnCall map[int]struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}
	ProcessConfigMsgStub        func(env *cb.Envelope) (*cb.Envelope, uint64, error)
	processConfigMsgMutex       sync.RWMutex
	processConfigMsgArgsForCall []struct {
		env *cb.Envelope
	}
	processConfigMsgReturns struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}
	processConfigMsgReturnsOnCall map[int]struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}
	OrderStub        func(env *cb.Envelope, configSeq uint64) error
	orderMutex       sync.RWMutex
	orderArgsForCall []struct {
		env       *cb.Envelope
		configSeq uint64
	}
	orderReturns struct {
		result1 error
	}
	orderReturnsOnCall map[int]struct {
		result1 error
	}
	ConfigureStub        func(config *cb.Envelope, configSeq uint64) error
	configureMutex       sync.RWMutex
	configureArgsForCall []struct {
		config    *cb.Envelope
		configSeq uint64
	}
	configureReturns struct {
		result1 error
	}
	configureReturnsOnCall map[int]struct {
		result1 error
	}
	WaitReadyStub        func() error
	waitReadyMutex       sync.RWMutex
	waitReadyArgsForCall []struct{}
	waitReadyReturns     struct {
		result1 error
	}
	waitReadyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ChannelSupport) ClassifyMsg(chdr *cb.ChannelHeader) msgprocessor.Classification {
	fake.classifyMsgMutex.Lock()
	ret, specificReturn := fake.classifyMsgReturnsOnCall[len(fake.classifyMsgArgsForCall)]
	fake.classifyMsgArgsForCall = append(fake.classifyMsgArgsForCall, struct {
		chdr *cb.ChannelHeader
	}{chdr})
	fake.recordInvocation("ClassifyMsg", []interface{}{chdr})
	fake.classifyMsgMutex.Unlock()
	if fake.ClassifyMsgStub != nil {
		return fake.ClassifyMsgStub(chdr)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.classifyMsgReturns.result1
}

func (fake *ChannelSupport) ClassifyMsgCallCount() int {
	fake.classifyMsgMutex.RLock()
	defer fake.classifyMsgMutex.RUnlock()
	return len(fake.classifyMsgArgsForCall)
}

func (fake *ChannelSupport) ClassifyMsgArgsForCall(i int) *cb.ChannelHeader {
	fake.classifyMsgMutex.RLock()
	defer fake.classifyMsgMutex.RUnlock()
	return fake.classifyMsgArgsForCall[i].chdr
}

func (fake *ChannelSupport) ClassifyMsgReturns(result1 msgprocessor.Classification) {
	fake.ClassifyMsgStub = nil
	fake.classifyMsgReturns = struct {
		result1 msgprocessor.Classification
	}{result1}
}

func (fake *ChannelSupport) ClassifyMsgReturnsOnCall(i int, result1 msgprocessor.Classification) {
	fake.ClassifyMsgStub = nil
	if fake.classifyMsgReturnsOnCall == nil {
		fake.classifyMsgReturnsOnCall = make(map[int]struct {
			result1 msgprocessor.Classification
		})
	}
	fake.classifyMsgReturnsOnCall[i] = struct {
		result1 msgprocessor.Classification
	}{result1}
}

func (fake *ChannelSupport) ProcessNormalMsg(env *cb.Envelope) (configSeq uint64, err error) {
	fake.processNormalMsgMutex.Lock()
	ret, specificReturn := fake.processNormalMsgReturnsOnCall[len(fake.processNormalMsgArgsForCall)]
	fake.processNormalMsgArgsForCall = append(fake.processNormalMsgArgsForCall, struct {
		env *cb.Envelope
	}{env})
	fake.recordInvocation("ProcessNormalMsg", []interface{}{env})
	fake.processNormalMsgMutex.Unlock()
	if fake.ProcessNormalMsgStub != nil {
		return fake.ProcessNormalMsgStub(env)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.processNormalMsgReturns.result1, fake.processNormalMsgReturns.result2
}

func (fake *ChannelSupport) ProcessNormalMsgCallCount() int {
	fake.processNormalMsgMutex.RLock()
	defer fake.processNormalMsgMutex.RUnlock()
	return len(fake.processNormalMsgArgsForCall)
}

func (fake *ChannelSupport) ProcessNormalMsgArgsForCall(i int) *cb.Envelope {
	fake.processNormalMsgMutex.RLock()
	defer fake.processNormalMsgMutex.RUnlock()
	return fake.processNormalMsgArgsForCall[i].env
}

func (fake *ChannelSupport) ProcessNormalMsgReturns(result1 uint64, result2 error) {
	fake.ProcessNormalMsgStub = nil
	fake.processNormalMsgReturns = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *ChannelSupport) ProcessNormalMsgReturnsOnCall(i int, result1 uint64, result2 error) {
	fake.ProcessNormalMsgStub = nil
	if fake.processNormalMsgReturnsOnCall == nil {
		fake.processNormalMsgReturnsOnCall = make(map[int]struct {
			result1 uint64
			result2 error
		})
	}
	fake.processNormalMsgReturnsOnCall[i] = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *ChannelSupport) ProcessConfigUpdateMsg(env *cb.Envelope) (config *cb.Envelope, configSeq uint64, err error) {
	fake.processConfigUpdateMsgMutex.Lock()
	ret, specificReturn := fake.processConfigUpdateMsgReturnsOnCall[len(fake.processConfigUpdateMsgArgsForCall)]
	fake.processConfigUpdateMsgArgsForCall = append(fake.processConfigUpdateMsgArgsForCall, struct {
		env *cb.Envelope
	}{env})
	fake.recordInvocation("ProcessConfigUpdateMsg", []interface{}{env})
	fake.processConfigUpdateMsgMutex.Unlock()
	if fake.ProcessConfigUpdateMsgStub != nil {
		return fake.ProcessConfigUpdateMsgStub(env)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.processConfigUpdateMsgReturns.result1, fake.processConfigUpdateMsgReturns.result2, fake.processConfigUpdateMsgReturns.result3
}

func (fake *ChannelSupport) ProcessConfigUpdateMsgCallCount() int {
	fake.processConfigUpdateMsgMutex.RLock()
	defer fake.processConfigUpdateMsgMutex.RUnlock()
	return len(fake.processConfigUpdateMsgArgsForCall)
}

func (fake *ChannelSupport) ProcessConfigUpdateMsgArgsForCall(i int) *cb.Envelope {
	fake.processConfigUpdateMsgMutex.RLock()
	defer fake.processConfigUpdateMsgMutex.RUnlock()
	return fake.processConfigUpdateMsgArgsForCall[i].env
}

func (fake *ChannelSupport) ProcessConfigUpdateMsgReturns(result1 *cb.Envelope, result2 uint64, result3 error) {
	fake.ProcessConfigUpdateMsgStub = nil
	fake.processConfigUpdateMsgReturns = struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}{result1, result2, result3}
}

func (fake *ChannelSupport) ProcessConfigUpdateMsgReturnsOnCall(i int, result1 *cb.Envelope, result2 uint64, result3 error) {
	fake.ProcessConfigUpdateMsgStub = nil
	if fake.processConfigUpdateMsgReturnsOnCall == nil {
		fake.processConfigUpdateMsgReturnsOnCall = make(map[int]struct {
			result1 *cb.Envelope
			result2 uint64
			result3 error
		})
	}
	fake.processConfigUpdateMsgReturnsOnCall[i] = struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}{result1, result2, result3}
}

func (fake *ChannelSupport) ProcessConfigMsg(env *cb.Envelope) (*cb.Envelope, uint64, error) {
	fake.processConfigMsgMutex.Lock()
	ret, specificReturn := fake.processConfigMsgReturnsOnCall[len(fake.processConfigMsgArgsForCall)]
	fake.processConfigMsgArgsForCall = append(fake.processConfigMsgArgsForCall, struct {
		env *cb.Envelope
	}{env})
	fake.recordInvocation("ProcessConfigMsg", []interface{}{env})
	fake.processConfigMsgMutex.Unlock()
	if fake.ProcessConfigMsgStub != nil {
		return fake.ProcessConfigMsgStub(env)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.processConfigMsgReturns.result1, fake.processConfigMsgReturns.result2, fake.processConfigMsgReturns.result3
}

func (fake *ChannelSupport) ProcessConfigMsgCallCount() int {
	fake.processConfigMsgMutex.RLock()
	defer fake.processConfigMsgMutex.RUnlock()
	return len(fake.processConfigMsgArgsForCall)
}

func (fake *ChannelSupport) ProcessConfigMsgArgsForCall(i int) *cb.Envelope {
	fake.processConfigMsgMutex.RLock()
	defer fake.processConfigMsgMutex.RUnlock()
	return fake.processConfigMsgArgsForCall[i].env
}

func (fake *ChannelSupport) ProcessConfigMsgReturns(result1 *cb.Envelope, result2 uint64, result3 error) {
	fake.ProcessConfigMsgStub = nil
	fake.processConfigMsgReturns = struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}{result1, result2, result3}
}

func (fake *ChannelSupport) ProcessConfigMsgReturnsOnCall(i int, result1 *cb.Envelope, result2 uint64, result3 error) {
	fake.ProcessConfigMsgStub = nil
	if fake.processConfigMsgReturnsOnCall == nil {
		fake.processConfigMsgReturnsOnCall = make(map[int]struct {
			result1 *cb.Envelope
			result2 uint64
			result3 error
		})
	}
	fake.processConfigMsgReturnsOnCall[i] = struct {
		result1 *cb.Envelope
		result2 uint64
		result3 error
	}{result1, result2, result3}
}

func (fake *ChannelSupport) Order(env *cb.Envelope, configSeq uint64) error {
	fake.orderMutex.Lock()
	ret, specificReturn := fake.orderReturnsOnCall[len(fake.orderArgsForCall)]
	fake.orderArgsForCall = append(fake.orderArgsForCall, struct {
		env       *cb.Envelope
		configSeq uint64
	}{env, configSeq})
	fake.recordInvocation("Order", []interface{}{env, configSeq})
	fake.orderMutex.Unlock()
	if fake.OrderStub != nil {
		return fake.OrderStub(env, configSeq)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.orderReturns.result1
}

func (fake *ChannelSupport) OrderCallCount() int {
	fake.orderMutex.RLock()
	defer fake.orderMutex.RUnlock()
	return len(fake.orderArgsForCall)
}

func (fake *ChannelSupport) OrderArgsForCall(i int) (*cb.Envelope, uint64) {
	fake.orderMutex.RLock()
	defer fake.orderMutex.RUnlock()
	return fake.orderArgsForCall[i].env, fake.orderArgsForCall[i].configSeq
}

func (fake *ChannelSupport) OrderReturns(result1 error) {
	fake.OrderStub = nil
	fake.orderReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChannelSupport) OrderReturnsOnCall(i int, result1 error) {
	fake.OrderStub = nil
	if fake.orderReturnsOnCall == nil {
		fake.orderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.orderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChannelSupport) Configure(config *cb.Envelope, configSeq uint64) error {
	fake.configureMutex.Lock()
	ret, specificReturn := fake.configureReturnsOnCall[len(fake.configureArgsForCall)]
	fake.configureArgsForCall = append(fake.configureArgsForCall, struct {
		config    *cb.Envelope
		configSeq uint64
	}{config, configSeq})
	fake.recordInvocation("Configure", []interface{}{config, configSeq})
	fake.configureMutex.Unlock()
	if fake.ConfigureStub != nil {
		return fake.ConfigureStub(config, configSeq)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.configureReturns.result1
}

func (fake *ChannelSupport) ConfigureCallCount() int {
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	return len(fake.configureArgsForCall)
}

func (fake *ChannelSupport) ConfigureArgsForCall(i int) (*cb.Envelope, uint64) {
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	return fake.configureArgsForCall[i].config, fake.configureArgsForCall[i].configSeq
}

func (fake *ChannelSupport) ConfigureReturns(result1 error) {
	fake.ConfigureStub = nil
	fake.configureReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChannelSupport) ConfigureReturnsOnCall(i int, result1 error) {
	fake.ConfigureStub = nil
	if fake.configureReturnsOnCall == nil {
		fake.configureReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.configureReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChannelSupport) WaitReady() error {
	fake.waitReadyMutex.Lock()
	ret, specificReturn := fake.waitReadyReturnsOnCall[len(fake.waitReadyArgsForCall)]
	fake.waitReadyArgsForCall = append(fake.waitReadyArgsForCall, struct{}{})
	fake.recordInvocation("WaitReady", []interface{}{})
	fake.waitReadyMutex.Unlock()
	if fake.WaitReadyStub != nil {
		return fake.WaitReadyStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.waitReadyReturns.result1
}

func (fake *ChannelSupport) WaitReadyCallCount() int {
	fake.waitReadyMutex.RLock()
	defer fake.waitReadyMutex.RUnlock()
	return len(fake.waitReadyArgsForCall)
}

func (fake *ChannelSupport) WaitReadyReturns(result1 error) {
	fake.WaitReadyStub = nil
	fake.waitReadyReturns = struct {
		result1 error
	}{result1}
}

func (fake *ChannelSupport) WaitReadyReturnsOnCall(i int, result1 error) {
	fake.WaitReadyStub = nil
	if fake.waitReadyReturnsOnCall == nil {
		fake.waitReadyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitReadyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ChannelSupport) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.classifyMsgMutex.RLock()
	defer fake.classifyMsgMutex.RUnlock()
	fake.processNormalMsgMutex.RLock()
	defer fake.processNormalMsgMutex.RUnlock()
	fake.processConfigUpdateMsgMutex.RLock()
	defer fake.processConfigUpdateMsgMutex.RUnlock()
	fake.processConfigMsgMutex.RLock()
	defer fake.processConfigMsgMutex.RUnlock()
	fake.orderMutex.RLock()
	defer fake.orderMutex.RUnlock()
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	fake.waitReadyMutex.RLock()
	defer fake.waitReadyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ChannelSupport) recordInvocation(key string, args []interface{}) {
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

var _ broadcast.ChannelSupport = new(ChannelSupport)

