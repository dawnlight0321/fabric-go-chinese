
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:16</date>
//</624456040624164864>

//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"
)

type ACLProvider struct {
	CheckACLStub        func(resName string, channelID string, idinfo interface{}) error
	checkACLMutex       sync.RWMutex
	checkACLArgsForCall []struct {
		resName   string
		channelID string
		idinfo    interface{}
	}
	checkACLReturns struct {
		result1 error
	}
	checkACLReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ACLProvider) CheckACL(resName string, channelID string, idinfo interface{}) error {
	fake.checkACLMutex.Lock()
	ret, specificReturn := fake.checkACLReturnsOnCall[len(fake.checkACLArgsForCall)]
	fake.checkACLArgsForCall = append(fake.checkACLArgsForCall, struct {
		resName   string
		channelID string
		idinfo    interface{}
	}{resName, channelID, idinfo})
	fake.recordInvocation("CheckACL", []interface{}{resName, channelID, idinfo})
	fake.checkACLMutex.Unlock()
	if fake.CheckACLStub != nil {
		return fake.CheckACLStub(resName, channelID, idinfo)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.checkACLReturns.result1
}

func (fake *ACLProvider) CheckACLCallCount() int {
	fake.checkACLMutex.RLock()
	defer fake.checkACLMutex.RUnlock()
	return len(fake.checkACLArgsForCall)
}

func (fake *ACLProvider) CheckACLArgsForCall(i int) (string, string, interface{}) {
	fake.checkACLMutex.RLock()
	defer fake.checkACLMutex.RUnlock()
	return fake.checkACLArgsForCall[i].resName, fake.checkACLArgsForCall[i].channelID, fake.checkACLArgsForCall[i].idinfo
}

func (fake *ACLProvider) CheckACLReturns(result1 error) {
	fake.CheckACLStub = nil
	fake.checkACLReturns = struct {
		result1 error
	}{result1}
}

func (fake *ACLProvider) CheckACLReturnsOnCall(i int, result1 error) {
	fake.CheckACLStub = nil
	if fake.checkACLReturnsOnCall == nil {
		fake.checkACLReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkACLReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ACLProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkACLMutex.RLock()
	defer fake.checkACLMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ACLProvider) recordInvocation(key string, args []interface{}) {
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

