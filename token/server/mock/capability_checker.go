
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456129149145088>

//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/token/server"
)

type CapabilityChecker struct {
	FabTokenStub        func(channelId string) (bool, error)
	fabTokenMutex       sync.RWMutex
	fabTokenArgsForCall []struct {
		channelId string
	}
	fabTokenReturns struct {
		result1 bool
		result2 error
	}
	fabTokenReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CapabilityChecker) FabToken(channelId string) (bool, error) {
	fake.fabTokenMutex.Lock()
	ret, specificReturn := fake.fabTokenReturnsOnCall[len(fake.fabTokenArgsForCall)]
	fake.fabTokenArgsForCall = append(fake.fabTokenArgsForCall, struct {
		channelId string
	}{channelId})
	fake.recordInvocation("FabToken", []interface{}{channelId})
	fake.fabTokenMutex.Unlock()
	if fake.FabTokenStub != nil {
		return fake.FabTokenStub(channelId)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fabTokenReturns.result1, fake.fabTokenReturns.result2
}

func (fake *CapabilityChecker) FabTokenCallCount() int {
	fake.fabTokenMutex.RLock()
	defer fake.fabTokenMutex.RUnlock()
	return len(fake.fabTokenArgsForCall)
}

func (fake *CapabilityChecker) FabTokenArgsForCall(i int) string {
	fake.fabTokenMutex.RLock()
	defer fake.fabTokenMutex.RUnlock()
	return fake.fabTokenArgsForCall[i].channelId
}

func (fake *CapabilityChecker) FabTokenReturns(result1 bool, result2 error) {
	fake.FabTokenStub = nil
	fake.fabTokenReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *CapabilityChecker) FabTokenReturnsOnCall(i int, result1 bool, result2 error) {
	fake.FabTokenStub = nil
	if fake.fabTokenReturnsOnCall == nil {
		fake.fabTokenReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.fabTokenReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *CapabilityChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fabTokenMutex.RLock()
	defer fake.fabTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CapabilityChecker) recordInvocation(key string, args []interface{}) {
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

var _ server.CapabilityChecker = new(CapabilityChecker)

