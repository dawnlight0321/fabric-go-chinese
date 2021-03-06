
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:28</date>
//</624456091966640128>

//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import orderer "github.com/hyperledger/fabric/protos/orderer"

//Dispatcher是为Dispatcher类型自动生成的模拟类型
type Dispatcher struct {
	mock.Mock
}

//DispatchStep提供了一个具有给定字段的模拟函数：ctx，request
func (_m *Dispatcher) DispatchStep(ctx context.Context, request *orderer.StepRequest) (*orderer.StepResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *orderer.StepResponse
	if rf, ok := ret.Get(0).(func(context.Context, *orderer.StepRequest) *orderer.StepResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orderer.StepResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *orderer.StepRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//DispatchSubmit提供了一个具有给定字段的模拟函数：ctx，request
func (_m *Dispatcher) DispatchSubmit(ctx context.Context, request *orderer.SubmitRequest) (*orderer.SubmitResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *orderer.SubmitResponse
	if rf, ok := ret.Get(0).(func(context.Context, *orderer.SubmitRequest) *orderer.SubmitResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orderer.SubmitResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *orderer.SubmitRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

