
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:28</date>
//</624456091710787584>

//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import mock "github.com/stretchr/testify/mock"

//ChannelLister是ChannelLister类型的自动生成的模拟类型
type ChannelLister struct {
	mock.Mock
}

//通道为给定字段提供模拟功能：
func (_m *ChannelLister) Channels() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

//close为给定字段提供模拟函数：
func (_m *ChannelLister) Close() {
	_m.Called()
}
