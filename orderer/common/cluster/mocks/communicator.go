
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:28</date>
//</624456091886948352>

//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import cluster "github.com/hyperledger/fabric/orderer/common/cluster"
import mock "github.com/stretchr/testify/mock"

//Communicator是为Communicator类型自动生成的模拟类型
type Communicator struct {
	mock.Mock
}

//配置为给定字段提供模拟函数：通道、成员
func (_m *Communicator) Configure(channel string, members []cluster.RemoteNode) {
	_m.Called(channel, members)
}

//远程提供了一个具有给定字段的模拟函数：channel，id
func (_m *Communicator) Remote(channel string, id uint64) (*cluster.RemoteContext, error) {
	ret := _m.Called(channel, id)

	var r0 *cluster.RemoteContext
	if rf, ok := ret.Get(0).(func(string, uint64) *cluster.RemoteContext); ok {
		r0 = rf(channel, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.RemoteContext)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint64) error); ok {
		r1 = rf(channel, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//Shutdown为给定字段提供模拟函数：
func (_m *Communicator) Shutdown() {
	_m.Called()
}

