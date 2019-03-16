
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:18</date>
//</624456047364411392>

//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import discovery "github.com/hyperledger/fabric/discovery/client"
import mock "github.com/stretchr/testify/mock"
import protosdiscovery "github.com/hyperledger/fabric/protos/discovery"

//ServiceResponse是为ServiceResponse类型自动生成的模拟类型
type ServiceResponse struct {
	mock.Mock
}

//forchannel提供了一个具有给定字段的模拟函数：a0
func (_m *ServiceResponse) ForChannel(_a0 string) discovery.ChannelResponse {
	ret := _m.Called(_a0)

	var r0 discovery.ChannelResponse
	if rf, ok := ret.Get(0).(func(string) discovery.ChannelResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(discovery.ChannelResponse)
		}
	}

	return r0
}

//ForLocal为给定字段提供模拟函数：
func (_m *ServiceResponse) ForLocal() discovery.LocalResponse {
	ret := _m.Called()

	var r0 discovery.LocalResponse
	if rf, ok := ret.Get(0).(func() discovery.LocalResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(discovery.LocalResponse)
		}
	}

	return r0
}

//raw提供了一个具有给定字段的模拟函数：
func (_m *ServiceResponse) Raw() *protosdiscovery.Response {
	ret := _m.Called()

	var r0 *protosdiscovery.Response
	if rf, ok := ret.Get(0).(func() *protosdiscovery.Response); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protosdiscovery.Response)
		}
	}

	return r0
}

