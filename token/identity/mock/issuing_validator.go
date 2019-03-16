
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456127546920960>

//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/token/identity"
)

type IssuingValidator struct {
	ValidateStub        func(creator identity.PublicInfo, tokenType string) error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		creator   identity.PublicInfo
		tokenType string
	}
	validateReturns struct {
		result1 error
	}
	validateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *IssuingValidator) Validate(creator identity.PublicInfo, tokenType string) error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		creator   identity.PublicInfo
		tokenType string
	}{creator, tokenType})
	fake.recordInvocation("Validate", []interface{}{creator, tokenType})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub(creator, tokenType)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validateReturns.result1
}

func (fake *IssuingValidator) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *IssuingValidator) ValidateArgsForCall(i int) (identity.PublicInfo, string) {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return fake.validateArgsForCall[i].creator, fake.validateArgsForCall[i].tokenType
}

func (fake *IssuingValidator) ValidateReturns(result1 error) {
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *IssuingValidator) ValidateReturnsOnCall(i int, result1 error) {
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IssuingValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *IssuingValidator) recordInvocation(key string, args []interface{}) {
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

var _ identity.IssuingValidator = new(IssuingValidator)

