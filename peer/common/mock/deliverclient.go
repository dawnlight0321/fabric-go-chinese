
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:33</date>
//</624456110438354944>

//伪造者生成的代码。不要编辑。
package mock

import (
	"context"
	"sync"

	"github.com/hyperledger/fabric/peer/common/api"
	"google.golang.org/grpc"
)

type DeliverClient struct {
	DeliverStub        func(ctx context.Context, opts ...grpc.CallOption) (api.DeliverService, error)
	deliverMutex       sync.RWMutex
	deliverArgsForCall []struct {
		ctx  context.Context
		opts []grpc.CallOption
	}
	deliverReturns struct {
		result1 api.DeliverService
		result2 error
	}
	deliverReturnsOnCall map[int]struct {
		result1 api.DeliverService
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DeliverClient) Deliver(ctx context.Context, opts ...grpc.CallOption) (api.DeliverService, error) {
	fake.deliverMutex.Lock()
	ret, specificReturn := fake.deliverReturnsOnCall[len(fake.deliverArgsForCall)]
	fake.deliverArgsForCall = append(fake.deliverArgsForCall, struct {
		ctx  context.Context
		opts []grpc.CallOption
	}{ctx, opts})
	fake.recordInvocation("Deliver", []interface{}{ctx, opts})
	fake.deliverMutex.Unlock()
	if fake.DeliverStub != nil {
		return fake.DeliverStub(ctx, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deliverReturns.result1, fake.deliverReturns.result2
}

func (fake *DeliverClient) DeliverCallCount() int {
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	return len(fake.deliverArgsForCall)
}

func (fake *DeliverClient) DeliverArgsForCall(i int) (context.Context, []grpc.CallOption) {
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	return fake.deliverArgsForCall[i].ctx, fake.deliverArgsForCall[i].opts
}

func (fake *DeliverClient) DeliverReturns(result1 api.DeliverService, result2 error) {
	fake.DeliverStub = nil
	fake.deliverReturns = struct {
		result1 api.DeliverService
		result2 error
	}{result1, result2}
}

func (fake *DeliverClient) DeliverReturnsOnCall(i int, result1 api.DeliverService, result2 error) {
	fake.DeliverStub = nil
	if fake.deliverReturnsOnCall == nil {
		fake.deliverReturnsOnCall = make(map[int]struct {
			result1 api.DeliverService
			result2 error
		})
	}
	fake.deliverReturnsOnCall[i] = struct {
		result1 api.DeliverService
		result2 error
	}{result1, result2}
}

func (fake *DeliverClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DeliverClient) recordInvocation(key string, args []interface{}) {
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

var _ api.DeliverClient = new(DeliverClient)

