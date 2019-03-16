
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:23</date>
//</624456070558912512>

//Code generated by mockery v1.0.0. 不要编辑。

package mocks

import common "github.com/hyperledger/fabric/gossip/privdata/common"
import mock "github.com/stretchr/testify/mock"

//ReconciliationFetcher是为ReconciliationFetcher类型自动生成的模拟类型。
type ReconciliationFetcher struct {
	mock.Mock
}

//FetchReconciledItems为给定字段提供模拟函数：dig2collectionconfig
func (_m *ReconciliationFetcher) FetchReconciledItems(dig2collectionConfig common.Dig2CollectionConfig) (*common.FetchedPvtDataContainer, error) {
	ret := _m.Called(dig2collectionConfig)

	var r0 *common.FetchedPvtDataContainer
	if rf, ok := ret.Get(0).(func(common.Dig2CollectionConfig) *common.FetchedPvtDataContainer); ok {
		r0 = rf(dig2collectionConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.FetchedPvtDataContainer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Dig2CollectionConfig) error); ok {
		r1 = rf(dig2collectionConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

