
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455977445363712>

//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import chaincode "github.com/hyperledger/fabric/common/chaincode"
import mock "github.com/stretchr/testify/mock"

//
type LifeCycleChangeListener struct {
	mock.Mock
}

//LifecycleChangeListener提供具有给定字段的模拟函数：channel、chaincodes
func (_m *LifeCycleChangeListener) LifeCycleChangeListener(channel string, chaincodes chaincode.MetadataSet) {
	_m.Called(channel, chaincodes)
}

