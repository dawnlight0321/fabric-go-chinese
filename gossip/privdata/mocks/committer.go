
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:23</date>
//</624456070265311232>

//Code generated by mockery v1.0.0. 不要编辑。

package mocks

import common "github.com/hyperledger/fabric/protos/common"
import ledger "github.com/hyperledger/fabric/core/ledger"
import mock "github.com/stretchr/testify/mock"

//committer是committer类型的自动生成的模拟类型
type Committer struct {
	mock.Mock
}

//close为给定字段提供模拟函数：
func (_m *Committer) Close() {
	_m.Called()
}

//commitpvtdataofoldblocks为给定字段提供模拟函数：blockpvdtdata
func (_m *Committer) CommitPvtDataOfOldBlocks(blockPvtData []*ledger.BlockPvtData) ([]*ledger.PvtdataHashMismatch, error) {
	ret := _m.Called(blockPvtData)

	var r0 []*ledger.PvtdataHashMismatch
	if rf, ok := ret.Get(0).(func([]*ledger.BlockPvtData) []*ledger.PvtdataHashMismatch); ok {
		r0 = rf(blockPvtData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ledger.PvtdataHashMismatch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*ledger.BlockPvtData) error); ok {
		r1 = rf(blockPvtData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//commitWithPvtData为给定字段提供模拟函数：blockAndPvtData
func (_m *Committer) CommitWithPvtData(blockAndPvtData *ledger.BlockAndPvtData) error {
	ret := _m.Called(blockAndPvtData)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ledger.BlockAndPvtData) error); ok {
		r0 = rf(blockAndPvtData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//getBlocks提供了一个具有给定字段的模拟函数：blockseq
func (_m *Committer) GetBlocks(blockSeqs []uint64) []*common.Block {
	ret := _m.Called(blockSeqs)

	var r0 []*common.Block
	if rf, ok := ret.Get(0).(func([]uint64) []*common.Block); ok {
		r0 = rf(blockSeqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*common.Block)
		}
	}

	return r0
}

//getconfigHistoryRetriever提供具有给定字段的模拟函数：
func (_m *Committer) GetConfigHistoryRetriever() (ledger.ConfigHistoryRetriever, error) {
	ret := _m.Called()

	var r0 ledger.ConfigHistoryRetriever
	if rf, ok := ret.Get(0).(func() ledger.ConfigHistoryRetriever); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ledger.ConfigHistoryRetriever)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//GetMissingPvtDataTracker提供了一个具有给定字段的模拟函数：
func (_m *Committer) GetMissingPvtDataTracker() (ledger.MissingPvtDataTracker, error) {
	ret := _m.Called()

	var r0 ledger.MissingPvtDataTracker
	if rf, ok := ret.Get(0).(func() ledger.MissingPvtDataTracker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ledger.MissingPvtDataTracker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//getpvtdata和blockbynum提供了一个具有给定字段的模拟函数：seqnum
func (_m *Committer) GetPvtDataAndBlockByNum(seqNum uint64) (*ledger.BlockAndPvtData, error) {
	ret := _m.Called(seqNum)

	var r0 *ledger.BlockAndPvtData
	if rf, ok := ret.Get(0).(func(uint64) *ledger.BlockAndPvtData); ok {
		r0 = rf(seqNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ledger.BlockAndPvtData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(seqNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//getpvtdatabynum提供了一个具有给定字段的模拟函数：blocknum、filter
func (_m *Committer) GetPvtDataByNum(blockNum uint64, filter ledger.PvtNsCollFilter) ([]*ledger.TxPvtData, error) {
	ret := _m.Called(blockNum, filter)

	var r0 []*ledger.TxPvtData
	if rf, ok := ret.Get(0).(func(uint64, ledger.PvtNsCollFilter) []*ledger.TxPvtData); ok {
		r0 = rf(blockNum, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ledger.TxPvtData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, ledger.PvtNsCollFilter) error); ok {
		r1 = rf(blockNum, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//LedgerHeight为给定字段提供模拟函数：
func (_m *Committer) LedgerHeight() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

