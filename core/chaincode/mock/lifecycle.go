
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:03</date>
//</624455983715848192>

//伪造者生成的代码。不要编辑。
package mock

import (
	sync "sync"

	ccprovider "github.com/hyperledger/fabric/core/common/ccprovider"
	ledger "github.com/hyperledger/fabric/core/ledger"
)

type Lifecycle struct {
	ChaincodeContainerInfoStub        func(string, string) (*ccprovider.ChaincodeContainerInfo, error)
	chaincodeContainerInfoMutex       sync.RWMutex
	chaincodeContainerInfoArgsForCall []struct {
		arg1 string
		arg2 string
	}
	chaincodeContainerInfoReturns struct {
		result1 *ccprovider.ChaincodeContainerInfo
		result2 error
	}
	chaincodeContainerInfoReturnsOnCall map[int]struct {
		result1 *ccprovider.ChaincodeContainerInfo
		result2 error
	}
	ChaincodeDefinitionStub        func(string, ledger.QueryExecutor) (ccprovider.ChaincodeDefinition, error)
	chaincodeDefinitionMutex       sync.RWMutex
	chaincodeDefinitionArgsForCall []struct {
		arg1 string
		arg2 ledger.QueryExecutor
	}
	chaincodeDefinitionReturns struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}
	chaincodeDefinitionReturnsOnCall map[int]struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Lifecycle) ChaincodeContainerInfo(arg1 string, arg2 string) (*ccprovider.ChaincodeContainerInfo, error) {
	fake.chaincodeContainerInfoMutex.Lock()
	ret, specificReturn := fake.chaincodeContainerInfoReturnsOnCall[len(fake.chaincodeContainerInfoArgsForCall)]
	fake.chaincodeContainerInfoArgsForCall = append(fake.chaincodeContainerInfoArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ChaincodeContainerInfo", []interface{}{arg1, arg2})
	fake.chaincodeContainerInfoMutex.Unlock()
	if fake.ChaincodeContainerInfoStub != nil {
		return fake.ChaincodeContainerInfoStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.chaincodeContainerInfoReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Lifecycle) ChaincodeContainerInfoCallCount() int {
	fake.chaincodeContainerInfoMutex.RLock()
	defer fake.chaincodeContainerInfoMutex.RUnlock()
	return len(fake.chaincodeContainerInfoArgsForCall)
}

func (fake *Lifecycle) ChaincodeContainerInfoCalls(stub func(string, string) (*ccprovider.ChaincodeContainerInfo, error)) {
	fake.chaincodeContainerInfoMutex.Lock()
	defer fake.chaincodeContainerInfoMutex.Unlock()
	fake.ChaincodeContainerInfoStub = stub
}

func (fake *Lifecycle) ChaincodeContainerInfoArgsForCall(i int) (string, string) {
	fake.chaincodeContainerInfoMutex.RLock()
	defer fake.chaincodeContainerInfoMutex.RUnlock()
	argsForCall := fake.chaincodeContainerInfoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Lifecycle) ChaincodeContainerInfoReturns(result1 *ccprovider.ChaincodeContainerInfo, result2 error) {
	fake.chaincodeContainerInfoMutex.Lock()
	defer fake.chaincodeContainerInfoMutex.Unlock()
	fake.ChaincodeContainerInfoStub = nil
	fake.chaincodeContainerInfoReturns = struct {
		result1 *ccprovider.ChaincodeContainerInfo
		result2 error
	}{result1, result2}
}

func (fake *Lifecycle) ChaincodeContainerInfoReturnsOnCall(i int, result1 *ccprovider.ChaincodeContainerInfo, result2 error) {
	fake.chaincodeContainerInfoMutex.Lock()
	defer fake.chaincodeContainerInfoMutex.Unlock()
	fake.ChaincodeContainerInfoStub = nil
	if fake.chaincodeContainerInfoReturnsOnCall == nil {
		fake.chaincodeContainerInfoReturnsOnCall = make(map[int]struct {
			result1 *ccprovider.ChaincodeContainerInfo
			result2 error
		})
	}
	fake.chaincodeContainerInfoReturnsOnCall[i] = struct {
		result1 *ccprovider.ChaincodeContainerInfo
		result2 error
	}{result1, result2}
}

func (fake *Lifecycle) ChaincodeDefinition(arg1 string, arg2 ledger.QueryExecutor) (ccprovider.ChaincodeDefinition, error) {
	fake.chaincodeDefinitionMutex.Lock()
	ret, specificReturn := fake.chaincodeDefinitionReturnsOnCall[len(fake.chaincodeDefinitionArgsForCall)]
	fake.chaincodeDefinitionArgsForCall = append(fake.chaincodeDefinitionArgsForCall, struct {
		arg1 string
		arg2 ledger.QueryExecutor
	}{arg1, arg2})
	fake.recordInvocation("ChaincodeDefinition", []interface{}{arg1, arg2})
	fake.chaincodeDefinitionMutex.Unlock()
	if fake.ChaincodeDefinitionStub != nil {
		return fake.ChaincodeDefinitionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.chaincodeDefinitionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Lifecycle) ChaincodeDefinitionCallCount() int {
	fake.chaincodeDefinitionMutex.RLock()
	defer fake.chaincodeDefinitionMutex.RUnlock()
	return len(fake.chaincodeDefinitionArgsForCall)
}

func (fake *Lifecycle) ChaincodeDefinitionCalls(stub func(string, ledger.QueryExecutor) (ccprovider.ChaincodeDefinition, error)) {
	fake.chaincodeDefinitionMutex.Lock()
	defer fake.chaincodeDefinitionMutex.Unlock()
	fake.ChaincodeDefinitionStub = stub
}

func (fake *Lifecycle) ChaincodeDefinitionArgsForCall(i int) (string, ledger.QueryExecutor) {
	fake.chaincodeDefinitionMutex.RLock()
	defer fake.chaincodeDefinitionMutex.RUnlock()
	argsForCall := fake.chaincodeDefinitionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Lifecycle) ChaincodeDefinitionReturns(result1 ccprovider.ChaincodeDefinition, result2 error) {
	fake.chaincodeDefinitionMutex.Lock()
	defer fake.chaincodeDefinitionMutex.Unlock()
	fake.ChaincodeDefinitionStub = nil
	fake.chaincodeDefinitionReturns = struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *Lifecycle) ChaincodeDefinitionReturnsOnCall(i int, result1 ccprovider.ChaincodeDefinition, result2 error) {
	fake.chaincodeDefinitionMutex.Lock()
	defer fake.chaincodeDefinitionMutex.Unlock()
	fake.ChaincodeDefinitionStub = nil
	if fake.chaincodeDefinitionReturnsOnCall == nil {
		fake.chaincodeDefinitionReturnsOnCall = make(map[int]struct {
			result1 ccprovider.ChaincodeDefinition
			result2 error
		})
	}
	fake.chaincodeDefinitionReturnsOnCall[i] = struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *Lifecycle) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.chaincodeContainerInfoMutex.RLock()
	defer fake.chaincodeContainerInfoMutex.RUnlock()
	fake.chaincodeDefinitionMutex.RLock()
	defer fake.chaincodeDefinitionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Lifecycle) recordInvocation(key string, args []interface{}) {
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

