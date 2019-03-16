
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:05</date>
//</624455992536469504>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package statebased_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/statebased"
	"github.com/stretchr/testify/assert"
)

func TestAddOrg(t *testing.T) {
//
	ep, err := statebased.NewStateEP(nil)
	assert.NoError(t, err)
	err = ep.AddOrgs(statebased.RoleTypePeer, "Org1")
	assert.NoError(t, err)

//不良角色类型
	err = ep.AddOrgs("unknown", "Org1")
	assert.Equal(t, &statebased.RoleTypeDoesNotExistError{RoleType: statebased.RoleType("unknown")}, err)
	assert.EqualError(t, err, "role type unknown does not exist")

	epBytes, err := ep.Policy()
	assert.NoError(t, err)
	expectedEP := cauthdsl.SignedByMspPeer("Org1")
	expectedEPBytes, err := proto.Marshal(expectedEP)
	assert.NoError(t, err)
	assert.Equal(t, expectedEPBytes, epBytes)
}

func TestListOrgs(t *testing.T) {
	expectedEP := cauthdsl.SignedByMspPeer("Org1")
	expectedEPBytes, err := proto.Marshal(expectedEP)
	assert.NoError(t, err)

//
	ep, err := statebased.NewStateEP(expectedEPBytes)
	orgs := ep.ListOrgs()
	assert.Equal(t, []string{"Org1"}, orgs)
}

func TestDelAddOrg(t *testing.T) {
	expectedEP := cauthdsl.SignedByMspPeer("Org1")
	expectedEPBytes, err := proto.Marshal(expectedEP)
	assert.NoError(t, err)
	ep, err := statebased.NewStateEP(expectedEPBytes)

//检索组织
	orgs := ep.ListOrgs()
	assert.Equal(t, []string{"Org1"}, orgs)

//代言政策
	ep.AddOrgs(statebased.RoleTypePeer, "Org2")
	ep.DelOrgs("Org1")

//检查存储的内容是否正确
	epBytes, err := ep.Policy()
	assert.NoError(t, err)
	expectedEP = cauthdsl.SignedByMspPeer("Org2")
	expectedEPBytes, err = proto.Marshal(expectedEP)
	assert.NoError(t, err)
	assert.Equal(t, expectedEPBytes, epBytes)
}

