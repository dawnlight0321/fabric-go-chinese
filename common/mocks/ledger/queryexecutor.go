
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455966221406208>

/*
版权所有IBM Corp.2017保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

   http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
*/


package ledger

import (
	"fmt"

	commonledger "github.com/hyperledger/fabric/common/ledger"
	"github.com/hyperledger/fabric/core/ledger"
)

type MockQueryExecutor struct {
//
	State map[string]map[string][]byte
}

func NewMockQueryExecutor(state map[string]map[string][]byte) *MockQueryExecutor {
	return &MockQueryExecutor{
		State: state,
	}
}

func (m *MockQueryExecutor) GetState(namespace string, key string) ([]byte, error) {
	ns := m.State[namespace]
	if ns == nil {
		return nil, fmt.Errorf("Could not retrieve namespace %s", namespace)
	}

	return ns[key], nil
}

func (m *MockQueryExecutor) GetStateMultipleKeys(namespace string, keys []string) ([][]byte, error) {
	res, err := m.GetState(namespace, keys[0])
	if err != nil {
		return nil, err
	}
	return [][]byte{res}, nil
}

func (m *MockQueryExecutor) GetStateRangeScanIterator(namespace string, startKey string, endKey string) (commonledger.ResultsIterator, error) {
	return nil, nil
}

func (m *MockQueryExecutor) GetStateRangeScanIteratorWithMetadata(namespace string, startKey, endKey string, metadata map[string]interface{}) (ledger.QueryResultsIterator, error) {
	return nil, nil
}

func (m *MockQueryExecutor) ExecuteQuery(namespace, query string) (commonledger.ResultsIterator, error) {
	return nil, nil
}

func (m *MockQueryExecutor) ExecuteQueryWithMetadata(namespace, query string, metadata map[string]interface{}) (ledger.QueryResultsIterator, error) {
	return nil, nil
}

func (m *MockQueryExecutor) GetPrivateData(namespace, collection, key string) ([]byte, error) {
	return nil, nil
}

func (m *MockQueryExecutor) GetPrivateDataMetadataByHash(namespace, collection string, keyhash []byte) (map[string][]byte, error) {
	return nil, nil
}

func (m *MockQueryExecutor) GetPrivateDataMultipleKeys(namespace, collection string, keys []string) ([][]byte, error) {
	return nil, nil
}

func (m *MockQueryExecutor) GetPrivateDataRangeScanIterator(namespace, collection, startKey, endKey string) (commonledger.ResultsIterator, error) {
	return nil, nil
}

func (m *MockQueryExecutor) ExecuteQueryOnPrivateData(namespace, collection, query string) (commonledger.ResultsIterator, error) {
	return nil, nil
}

func (m *MockQueryExecutor) Done() {
}

func (m *MockQueryExecutor) GetStateMetadata(namespace, key string) (map[string][]byte, error) {
	return nil, nil
}

func (m *MockQueryExecutor) GetPrivateDataMetadata(namespace, collection, key string) (map[string][]byte, error) {
	return nil, nil
}

