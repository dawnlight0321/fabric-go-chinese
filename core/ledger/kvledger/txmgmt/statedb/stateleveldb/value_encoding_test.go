
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:13</date>
//</624456026506137600>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package stateleveldb

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
	"github.com/stretchr/testify/assert"
)

//testencodestring使用旧格式测试字符串值的编码和解码
func TestEncodeDecodeStringOldFormat(t *testing.T) {
	bytesString1 := []byte("value1")
	version1 := version.NewHeight(1, 1)
	encodedValue := encodeValueOldFormat(bytesString1, version1)
	decodedValue, err := decodeValue(encodedValue)
	assert.NoError(t, err)
	assert.Equal(t, &statedb.VersionedValue{Version: version1, Value: bytesString1}, decodedValue)
}

//testencodedecodejsonoldform使用旧格式测试JSON值的编码和解码
func TestEncodeDecodeJSONOldFormat(t *testing.T) {
	bytesJSON2 := []byte(`{"asset_name":"marble1","color":"blue","size":"35","owner":"jerry"}`)
	version2 := version.NewHeight(1, 1)
	encodedValue := encodeValueOldFormat(bytesJSON2, version2)
	decodedValue, err := decodeValue(encodedValue)
	assert.NoError(t, err)
	assert.Equal(t, &statedb.VersionedValue{Version: version2, Value: bytesJSON2}, decodedValue)
}

func TestEncodeDecodeOldAndNewFormat(t *testing.T) {
	testdata := []*statedb.VersionedValue{
		{
			Value:   []byte("value0"),
			Version: version.NewHeight(0, 0),
		},
		{
			Value:   []byte("value1"),
			Version: version.NewHeight(1, 2),
		},

		{
			Value:   []byte{},
			Version: version.NewHeight(50, 50),
		},
		{
			Value:    []byte{},
			Version:  version.NewHeight(50, 50),
			Metadata: []byte("sample-metadata"),
		},
	}

	for i, testdatum := range testdata {
		t.Run(fmt.Sprintf("testcase-newfmt-%d", i),
			func(t *testing.T) { testEncodeDecodeNewFormat(t, testdatum) },
		)
	}

	for i, testdatum := range testdata {
		t.Run(fmt.Sprintf("testcase-oldfmt-%d", i),
			func(t *testing.T) {
				testdatum.Metadata = nil
				testEncodeDecodeOldFormat(t, testdatum)
			},
		)
	}

}

func testEncodeDecodeNewFormat(t *testing.T, v *statedb.VersionedValue) {
	encodedNewFmt, err := encodeValue(v)
	assert.NoError(t, err)
//使用新格式的编码解码应返回相同版本的值
	decodedFromNewFmt, err := decodeValue(encodedNewFmt)
	assert.NoError(t, err)
	assert.Equal(t, v, decodedFromNewFmt)
}

func testEncodeDecodeOldFormat(t *testing.T, v *statedb.VersionedValue) {
	encodedOldFmt := encodeValueOldFormat(v.Value, v.Version)
//decodeValue应该能够处理旧格式
	decodedFromOldFmt, err := decodeValue(encodedOldFmt)
	assert.NoError(t, err)
	assert.Equal(t, v, decodedFromOldFmt)
}

