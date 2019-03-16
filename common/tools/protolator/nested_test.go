
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455975029444608>

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
**/


package protolator

import (
	"bytes"
	"testing"

	"github.com/hyperledger/fabric/common/tools/protolator/testprotos"
	"github.com/stretchr/testify/assert"
)

func TestPlainNestedMsg(t *testing.T) {
	fromPrefix := "from"
	toPrefix := "to"
	tppff := &testProtoPlainFieldFactory{
		fromPrefix: fromPrefix,
		toPrefix:   toPrefix,
	}

	fieldFactories = []protoFieldFactory{tppff}

	pfValue := "foo"
	startMsg := &testprotos.NestedMsg{
		PlainNestedField: &testprotos.SimpleMsg{
			PlainField: pfValue,
		},
	}

	var buffer bytes.Buffer
	assert.NoError(t, DeepMarshalJSON(&buffer, startMsg))
	newMsg := &testprotos.NestedMsg{}
	assert.NoError(t, DeepUnmarshalJSON(bytes.NewReader(buffer.Bytes()), newMsg))
	assert.NotEqual(t, fromPrefix+toPrefix+startMsg.PlainNestedField.PlainField, newMsg.PlainNestedField.PlainField)

	fieldFactories = []protoFieldFactory{tppff, nestedFieldFactory{}}

	buffer.Reset()
	assert.NoError(t, DeepMarshalJSON(&buffer, startMsg))
	assert.NoError(t, DeepUnmarshalJSON(bytes.NewReader(buffer.Bytes()), newMsg))
	assert.Equal(t, fromPrefix+toPrefix+startMsg.PlainNestedField.PlainField, newMsg.PlainNestedField.PlainField)
}

func TestMapNestedMsg(t *testing.T) {
	fromPrefix := "from"
	toPrefix := "to"
	tppff := &testProtoPlainFieldFactory{
		fromPrefix: fromPrefix,
		toPrefix:   toPrefix,
	}

	fieldFactories = []protoFieldFactory{tppff}

	pfValue := "foo"
	mapKey := "bar"
	startMsg := &testprotos.NestedMsg{
		MapNestedField: map[string]*testprotos.SimpleMsg{
			mapKey: {
				PlainField: pfValue,
			},
		},
	}

	var buffer bytes.Buffer
	assert.NoError(t, DeepMarshalJSON(&buffer, startMsg))
	newMsg := &testprotos.NestedMsg{}
	assert.NoError(t, DeepUnmarshalJSON(bytes.NewReader(buffer.Bytes()), newMsg))
	assert.NotEqual(t, fromPrefix+toPrefix+startMsg.MapNestedField[mapKey].PlainField, newMsg.MapNestedField[mapKey].PlainField)

	fieldFactories = []protoFieldFactory{tppff, nestedMapFieldFactory{}}

	buffer.Reset()
	assert.NoError(t, DeepMarshalJSON(&buffer, startMsg))
	assert.NoError(t, DeepUnmarshalJSON(bytes.NewReader(buffer.Bytes()), newMsg))
	assert.Equal(t, fromPrefix+toPrefix+startMsg.MapNestedField[mapKey].PlainField, newMsg.MapNestedField[mapKey].PlainField)
}

func TestSliceNestedMsg(t *testing.T) {
	fromPrefix := "from"
	toPrefix := "to"
	tppff := &testProtoPlainFieldFactory{
		fromPrefix: fromPrefix,
		toPrefix:   toPrefix,
	}

	fieldFactories = []protoFieldFactory{tppff}

	pfValue := "foo"
	startMsg := &testprotos.NestedMsg{
		SliceNestedField: []*testprotos.SimpleMsg{
			{
				PlainField: pfValue,
			},
		},
	}

	var buffer bytes.Buffer
	assert.NoError(t, DeepMarshalJSON(&buffer, startMsg))
	newMsg := &testprotos.NestedMsg{}
	assert.NoError(t, DeepUnmarshalJSON(bytes.NewReader(buffer.Bytes()), newMsg))
	assert.NotEqual(t, fromPrefix+toPrefix+startMsg.SliceNestedField[0].PlainField, newMsg.SliceNestedField[0].PlainField)

	fieldFactories = []protoFieldFactory{tppff, nestedSliceFieldFactory{}}

	buffer.Reset()
	assert.NoError(t, DeepMarshalJSON(&buffer, startMsg))
	assert.NoError(t, DeepUnmarshalJSON(bytes.NewReader(buffer.Bytes()), newMsg))
	assert.Equal(t, fromPrefix+toPrefix+startMsg.SliceNestedField[0].PlainField, newMsg.SliceNestedField[0].PlainField)
}

