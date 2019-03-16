
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:07</date>
//</624456002342752256>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package privdata

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/cauthdsl"
	lm "github.com/hyperledger/fabric/common/mocks/ledger"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

type mockStoreSupport struct {
	Qe   *lm.MockQueryExecutor
	QErr error
}

func (c *mockStoreSupport) GetQueryExecutorForLedger(cid string) (ledger.QueryExecutor, error) {
	return c.Qe, c.QErr
}

func (c *mockStoreSupport) GetIdentityDeserializer(chainID string) msp.IdentityDeserializer {
	return &mockDeserializer{}
}

func TestCollectionStore(t *testing.T) {
	wState := make(map[string]map[string][]byte)
	support := &mockStoreSupport{Qe: &lm.MockQueryExecutor{State: wState}}
	cs := NewSimpleCollectionStore(support)
	assert.NotNil(t, cs)

	support.QErr = errors.New("")
	_, err := cs.RetrieveCollection(common.CollectionCriteria{})
	assert.Error(t, err)

	support.QErr = nil
	wState["lscc"] = make(map[string][]byte)

	_, err = cs.RetrieveCollection(common.CollectionCriteria{})
	assert.Error(t, err)

	ccr := common.CollectionCriteria{Channel: "ch", Namespace: "cc", Collection: "mycollection"}

	wState["lscc"][BuildCollectionKVSKey(ccr.Namespace)] = []byte("barf")

	_, err = cs.RetrieveCollection(ccr)
	assert.Error(t, err)

	cc := &common.CollectionConfig{Payload: &common.CollectionConfig_StaticCollectionConfig{
		StaticCollectionConfig: &common.StaticCollectionConfig{Name: "mycollection"}},
	}
	ccp := &common.CollectionConfigPackage{Config: []*common.CollectionConfig{cc}}
	ccpBytes, err := proto.Marshal(ccp)
	assert.NoError(t, err)
	assert.NotNil(t, ccpBytes)

	wState["lscc"][BuildCollectionKVSKey(ccr.Namespace)] = ccpBytes

	_, err = cs.RetrieveCollection(ccr)
	assert.Error(t, err)

	var signers = [][]byte{[]byte("signer0"), []byte("signer1")}
	policyEnvelope := cauthdsl.Envelope(cauthdsl.Or(cauthdsl.SignedBy(0), cauthdsl.SignedBy(1)), signers)
	accessPolicy := createCollectionPolicyConfig(policyEnvelope)

	cc = &common.CollectionConfig{Payload: &common.CollectionConfig_StaticCollectionConfig{
		StaticCollectionConfig: &common.StaticCollectionConfig{
			Name:             "mycollection",
			MemberOrgsPolicy: accessPolicy,
			MemberOnlyRead:   false,
		},
	}}
	ccp = &common.CollectionConfigPackage{Config: []*common.CollectionConfig{cc}}
	ccpBytes, err = proto.Marshal(ccp)
	assert.NoError(t, err)
	assert.NotNil(t, ccpBytes)

	wState["lscc"][BuildCollectionKVSKey(ccr.Namespace)] = ccpBytes

	c, err := cs.RetrieveCollection(ccr)
	assert.NoError(t, err)
	assert.NotNil(t, c)

	ca, err := cs.RetrieveCollectionAccessPolicy(ccr)
	assert.NoError(t, err)
	assert.NotNil(t, ca)

	c, err = cs.RetrieveCollection(common.CollectionCriteria{Channel: "ch", Namespace: "cc", Collection: "asd"})
	assert.Error(t, err)
	assert.Nil(t, c)

	ccc, err := cs.RetrieveCollectionConfigPackage(ccr)
	assert.NoError(t, err)
	assert.NotNil(t, ccc)

	cc = &common.CollectionConfig{Payload: &common.CollectionConfig_StaticCollectionConfig{
		StaticCollectionConfig: &common.StaticCollectionConfig{
			Name:             "mycollection",
			MemberOrgsPolicy: accessPolicy,
			MemberOnlyRead:   true,
		},
	}}
	ccp = &common.CollectionConfigPackage{Config: []*common.CollectionConfig{cc}}
	ccpBytes, err = proto.Marshal(ccp)
	assert.NoError(t, err)
	assert.NotNil(t, ccpBytes)

	wState["lscc"][BuildCollectionKVSKey(ccr.Namespace)] = ccpBytes

	signedProp, _ := utils.MockSignedEndorserProposalOrPanic("A", &peer.ChaincodeSpec{}, []byte("signer0"), []byte("msg1"))
	allowedAccess, err := cs.HasReadAccess(ccr, signedProp, &lm.MockQueryExecutor{State: wState})
	assert.NoError(t, err)
	assert.True(t, allowedAccess)

//只有Signer0和Signer1是成员
	signedProp, _ = utils.MockSignedEndorserProposalOrPanic("A", &peer.ChaincodeSpec{}, []byte("signer2"), []byte("msg1"))
	allowedAccess, err = cs.HasReadAccess(ccr, signedProp, &lm.MockQueryExecutor{State: wState})
	assert.NoError(t, err)
	assert.False(t, allowedAccess)
}

