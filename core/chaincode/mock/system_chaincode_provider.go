
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:03</date>
//</624455984370159616>

//伪造者生成的代码。不要编辑。
package mock

import (
	sync "sync"
)

type SystemCCProvider struct {
	IsSysCCStub        func(string) bool
	isSysCCMutex       sync.RWMutex
	isSysCCArgsForCall []struct {
		arg1 string
	}
	isSysCCReturns struct {
		result1 bool
	}
	isSysCCReturnsOnCall map[int]struct {
		result1 bool
	}
	IsSysCCAndNotInvokableCC2CCStub        func(string) bool
	isSysCCAndNotInvokableCC2CCMutex       sync.RWMutex
	isSysCCAndNotInvokableCC2CCArgsForCall []struct {
		arg1 string
	}
	isSysCCAndNotInvokableCC2CCReturns struct {
		result1 bool
	}
	isSysCCAndNotInvokableCC2CCReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SystemCCProvider) IsSysCC(arg1 string) bool {
	fake.isSysCCMutex.Lock()
	ret, specificReturn := fake.isSysCCReturnsOnCall[len(fake.isSysCCArgsForCall)]
	fake.isSysCCArgsForCall = append(fake.isSysCCArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("IsSysCC", []interface{}{arg1})
	fake.isSysCCMutex.Unlock()
	if fake.IsSysCCStub != nil {
		return fake.IsSysCCStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isSysCCReturns
	return fakeReturns.result1
}

func (fake *SystemCCProvider) IsSysCCCallCount() int {
	fake.isSysCCMutex.RLock()
	defer fake.isSysCCMutex.RUnlock()
	return len(fake.isSysCCArgsForCall)
}

func (fake *SystemCCProvider) IsSysCCCalls(stub func(string) bool) {
	fake.isSysCCMutex.Lock()
	defer fake.isSysCCMutex.Unlock()
	fake.IsSysCCStub = stub
}

func (fake *SystemCCProvider) IsSysCCArgsForCall(i int) string {
	fake.isSysCCMutex.RLock()
	defer fake.isSysCCMutex.RUnlock()
	argsForCall := fake.isSysCCArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SystemCCProvider) IsSysCCReturns(result1 bool) {
	fake.isSysCCMutex.Lock()
	defer fake.isSysCCMutex.Unlock()
	fake.IsSysCCStub = nil
	fake.isSysCCReturns = struct {
		result1 bool
	}{result1}
}

func (fake *SystemCCProvider) IsSysCCReturnsOnCall(i int, result1 bool) {
	fake.isSysCCMutex.Lock()
	defer fake.isSysCCMutex.Unlock()
	fake.IsSysCCStub = nil
	if fake.isSysCCReturnsOnCall == nil {
		fake.isSysCCReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isSysCCReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *SystemCCProvider) IsSysCCAndNotInvokableCC2CC(arg1 string) bool {
	fake.isSysCCAndNotInvokableCC2CCMutex.Lock()
	ret, specificReturn := fake.isSysCCAndNotInvokableCC2CCReturnsOnCall[len(fake.isSysCCAndNotInvokableCC2CCArgsForCall)]
	fake.isSysCCAndNotInvokableCC2CCArgsForCall = append(fake.isSysCCAndNotInvokableCC2CCArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("IsSysCCAndNotInvokableCC2CC", []interface{}{arg1})
	fake.isSysCCAndNotInvokableCC2CCMutex.Unlock()
	if fake.IsSysCCAndNotInvokableCC2CCStub != nil {
		return fake.IsSysCCAndNotInvokableCC2CCStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isSysCCAndNotInvokableCC2CCReturns
	return fakeReturns.result1
}

func (fake *SystemCCProvider) IsSysCCAndNotInvokableCC2CCCallCount() int {
	fake.isSysCCAndNotInvokableCC2CCMutex.RLock()
	defer fake.isSysCCAndNotInvokableCC2CCMutex.RUnlock()
	return len(fake.isSysCCAndNotInvokableCC2CCArgsForCall)
}

func (fake *SystemCCProvider) IsSysCCAndNotInvokableCC2CCCalls(stub func(string) bool) {
	fake.isSysCCAndNotInvokableCC2CCMutex.Lock()
	defer fake.isSysCCAndNotInvokableCC2CCMutex.Unlock()
	fake.IsSysCCAndNotInvokableCC2CCStub = stub
}

func (fake *SystemCCProvider) IsSysCCAndNotInvokableCC2CCArgsForCall(i int) string {
	fake.isSysCCAndNotInvokableCC2CCMutex.RLock()
	defer fake.isSysCCAndNotInvokableCC2CCMutex.RUnlock()
	argsForCall := fake.isSysCCAndNotInvokableCC2CCArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SystemCCProvider) IsSysCCAndNotInvokableCC2CCReturns(result1 bool) {
	fake.isSysCCAndNotInvokableCC2CCMutex.Lock()
	defer fake.isSysCCAndNotInvokableCC2CCMutex.Unlock()
	fake.IsSysCCAndNotInvokableCC2CCStub = nil
	fake.isSysCCAndNotInvokableCC2CCReturns = struct {
		result1 bool
	}{result1}
}

func (fake *SystemCCProvider) IsSysCCAndNotInvokableCC2CCReturnsOnCall(i int, result1 bool) {
	fake.isSysCCAndNotInvokableCC2CCMutex.Lock()
	defer fake.isSysCCAndNotInvokableCC2CCMutex.Unlock()
	fake.IsSysCCAndNotInvokableCC2CCStub = nil
	if fake.isSysCCAndNotInvokableCC2CCReturnsOnCall == nil {
		fake.isSysCCAndNotInvokableCC2CCReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isSysCCAndNotInvokableCC2CCReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *SystemCCProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isSysCCMutex.RLock()
	defer fake.isSysCCMutex.RUnlock()
	fake.isSysCCAndNotInvokableCC2CCMutex.RLock()
	defer fake.isSysCCAndNotInvokableCC2CCMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SystemCCProvider) recordInvocation(key string, args []interface{}) {
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

