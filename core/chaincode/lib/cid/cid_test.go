
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:02</date>
//</624455981358649344>

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

package cid_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/stretchr/testify/assert"
)

const certWithOutAttrs = `-----BEGIN CERTIFICATE-----
MIICXTCCAgSgAwIBAgIUeLy6uQnq8wwyElU/jCKRYz3tJiQwCgYIKoZIzj0EAwIw
eTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNh
biBGcmFuY2lzY28xGTAXBgNVBAoTEEludGVybmV0IFdpZGdldHMxDDAKBgNVBAsT
A1dXVzEUMBIGA1UEAxMLZXhhbXBsZS5jb20wHhcNMTcwOTA4MDAxNTAwWhcNMTgw
OTA4MDAxNTAwWjBdMQswCQYDVQQGEwJVUzEXMBUGA1UECBMOTm9ydGggQ2Fyb2xp
bmExFDASBgNVBAoTC0h5cGVybGVkZ2VyMQ8wDQYDVQQLEwZGYWJyaWMxDjAMBgNV
BAMTBWFkbWluMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFq/90YMuH4tWugHa
oyZtt4Mbwgv6CkBSDfYulVO1CVInw1i/k16DocQ/KSDTeTfgJxrX1Ree1tjpaodG
1wWyM6OBhTCBgjAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAdBgNVHQ4E
FgQUhKs/VJ9IWJd+wer6sgsgtZmxZNwwHwYDVR0jBBgwFoAUIUd4i/sLTwYWvpVr
TApzcT8zv/kwIgYDVR0RBBswGYIXQW5pbHMtTWFjQm9vay1Qcm8ubG9jYWwwCgYI
KoZIzj0EAwIDRwAwRAIgCoXaCdU8ZiRKkai0QiXJM/GL5fysLnmG2oZ6XOIdwtsC
IEmCsI8Mhrvx1doTbEOm7kmIrhQwUVDBNXCWX1t3kJVN
-----END CERTIFICATE-----
`
const certWithAttrs = `-----BEGIN CERTIFICATE-----
MIIB6TCCAY+gAwIBAgIUHkmY6fRP0ANTvzaBwKCkMZZPUnUwCgYIKoZIzj0EAwIw
GzEZMBcGA1UEAxMQZmFicmljLWNhLXNlcnZlcjAeFw0xNzA5MDgwMzQyMDBaFw0x
ODA5MDgwMzQyMDBaMB4xHDAaBgNVBAMTE015VGVzdFVzZXJXaXRoQXR0cnMwWTAT
BgcqhkjOPQIBBggqhkjOPQMBBwNCAATmB1r3CdWvOOP3opB3DjJnW3CnN8q1ydiR
dzmuA6A2rXKzPIltHvYbbSqISZJubsy8gVL6GYgYXNdu69RzzFF5o4GtMIGqMA4G
A1UdDwEB/wQEAwICBDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBTYKLTAvJJK08OM
VGwIhjMQpo2DrjAfBgNVHSMEGDAWgBTEs/52DeLePPx1+65VhgTwu3/2ATAiBgNV
HREEGzAZghdBbmlscy1NYWNCb29rLVByby5sb2NhbDAmBggqAwQFBgcIAQQaeyJh
dHRycyI6eyJhdHRyMSI6InZhbDEifX0wCgYIKoZIzj0EAwIDSAAwRQIhAPuEqWUp
svTTvBqLR5JeQSctJuz3zaqGRqSs2iW+QB3FAiAIP0mGWKcgSGRMMBvaqaLytBYo
9v3hRt1r8j8vN0pMcg==
-----END CERTIFICATE-----
`

func TestClient(t *testing.T) {
	stub, err := getMockStub()
	assert.NoError(t, err, "Failed to get mock submitter")
	sinfo, err := cid.New(stub)
	assert.NoError(t, err, "Error getting submitter of the transaction")
	id, err := cid.GetID(stub)
	assert.NoError(t, err, "Error getting ID of the submitter of the transaction")
	assert.NotEmpty(t, id, "Transaction submitter ID should not be empty")
	t.Logf("The client's ID is: %s", id)
	cert, err := cid.GetX509Certificate(stub)
	assert.NoError(t, err, "Error getting X509 certificate of the submitter of the transaction")
	assert.NotNil(t, cert, "Transaction submitter certificate should not be nil")
	mspid, err := cid.GetMSPID(stub)
	assert.NoError(t, err, "Error getting MSP ID of the submitter of the transaction")
	assert.NotEmpty(t, mspid, "Transaction submitter MSP ID should not be empty")
	_, found, err := sinfo.GetAttributeValue("foo")
	assert.NoError(t, err, "Error getting Unique ID of the submitter of the transaction")
	assert.False(t, found, "Attribute 'foo' should not be found in the submitter cert")
	err = cid.AssertAttributeValue(stub, "foo", "")
	assert.Error(t, err, "AssertAttributeValue should have returned an error with no attribute")

	stub, err = getMockStubWithAttrs()
	assert.NoError(t, err, "Failed to get mock submitter")
	sinfo, err = cid.New(stub)
	assert.NoError(t, err, "Failed to new client")
	attrVal, found, err := sinfo.GetAttributeValue("attr1")
	assert.NoError(t, err, "Error getting Unique ID of the submitter of the transaction")
	assert.True(t, found, "Attribute 'attr1' should be found in the submitter cert")
	assert.Equal(t, attrVal, "val1", "Value of attribute 'attr1' should be 'val1'")
	attrVal, found, err = cid.GetAttributeValue(stub, "attr1")
	assert.NoError(t, err, "Error getting Unique ID of the submitter of the transaction")
	assert.True(t, found, "Attribute 'attr1' should be found in the submitter cert")
	assert.Equal(t, attrVal, "val1", "Value of attribute 'attr1' should be 'val1'")
	err = cid.AssertAttributeValue(stub, "attr1", "val1")
	assert.NoError(t, err, "Error in AssertAttributeValue")
	err = cid.AssertAttributeValue(stub, "attr1", "val2")
	assert.Error(t, err, "Assert should have failed; value was val1, not val2")

//错误案例1
	stub, err = getMockStubWithNilCreator()
	assert.NoError(t, err, "Failed to get mock submitter")
	sinfo, err = cid.New(stub)
	assert.Error(t, err, "NewSubmitterInfo should have returned an error when submitter with nil creator is passed")

//错误案例2
	stub, err = getMockStubWithFakeCreator()
	assert.NoError(t, err, "Failed to get mock submitter")
	sinfo, err = cid.New(stub)
	assert.Error(t, err, "NewSubmitterInfo should have returned an error when submitter with fake creator is passed")
}

func getMockStub() (cid.ChaincodeStubInterface, error) {
	stub := &mockStub{}
	sid := &msp.SerializedIdentity{Mspid: "SampleOrg",
		IdBytes: []byte(certWithOutAttrs)}
	b, err := proto.Marshal(sid)
	if err != nil {
		return nil, err
	}
	stub.creator = b
	return stub, nil
}

func getMockStubWithAttrs() (cid.ChaincodeStubInterface, error) {
	stub := &mockStub{}
	sid := &msp.SerializedIdentity{Mspid: "SampleOrg",
		IdBytes: []byte(certWithAttrs)}
	b, err := proto.Marshal(sid)
	if err != nil {
		return nil, err
	}
	stub.creator = b
	return stub, nil
}

func getMockStubWithNilCreator() (cid.ChaincodeStubInterface, error) {
	c := &mockStub{}
	c.creator = nil
	return c, nil
}

func getMockStubWithFakeCreator() (cid.ChaincodeStubInterface, error) {
	c := &mockStub{}
	c.creator = []byte("foo")
	return c, nil
}

type mockStub struct {
	creator []byte
}

func (s *mockStub) GetCreator() ([]byte, error) {
	return s.creator, nil
}

