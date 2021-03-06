
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:10</date>
//</624456013424103424>

//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import common "github.com/hyperledger/fabric/protos/common"
import errors "github.com/hyperledger/fabric/common/errors"
import mock "github.com/stretchr/testify/mock"

//TransactionValidator是TransactionValidator类型的自动生成的模拟类型
type TransactionValidator struct {
	mock.Mock
}

//validate提供具有给定字段的模拟函数：block、namespace、txposition、actionposition、policy
func (_m *TransactionValidator) Validate(block *common.Block, namespace string, txPosition int, actionPosition int, policy []byte) errors.TxValidationError {
	ret := _m.Called(block, namespace, txPosition, actionPosition, policy)

	var r0 errors.TxValidationError
	if rf, ok := ret.Get(0).(func(*common.Block, string, int, int, []byte) errors.TxValidationError); ok {
		r0 = rf(block, namespace, txPosition, actionPosition, policy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.TxValidationError)
		}
	}

	return r0
}

