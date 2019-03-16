
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:04</date>
//</624455989956972544>

//伪造者生成的代码。不要编辑。
package mock

import (
	tar "archive/tar"
	sync "sync"
)

type PackageWriter struct {
	WriteStub        func(string, []byte, *tar.Writer) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 string
		arg2 []byte
		arg3 *tar.Writer
	}
	writeReturns struct {
		result1 error
	}
	writeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PackageWriter) Write(arg1 string, arg2 []byte, arg3 *tar.Writer) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 string
		arg2 []byte
		arg3 *tar.Writer
	}{arg1, arg2Copy, arg3})
	fake.recordInvocation("Write", []interface{}{arg1, arg2Copy, arg3})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.writeReturns
	return fakeReturns.result1
}

func (fake *PackageWriter) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *PackageWriter) WriteCalls(stub func(string, []byte, *tar.Writer) error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = stub
}

func (fake *PackageWriter) WriteArgsForCall(i int) (string, []byte, *tar.Writer) {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	argsForCall := fake.writeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *PackageWriter) WriteReturns(result1 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *PackageWriter) WriteReturnsOnCall(i int, result1 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PackageWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PackageWriter) recordInvocation(key string, args []interface{}) {
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

