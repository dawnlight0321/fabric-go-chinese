
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:00</date>
//</624455974635180032>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package integration

import (
	"bytes"
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/tools/configtxgen/configtxgentest"
	"github.com/hyperledger/fabric/common/tools/configtxgen/encoder"
	genesisconfig "github.com/hyperledger/fabric/common/tools/configtxgen/localconfig"
	"github.com/hyperledger/fabric/common/tools/protolator"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/msp"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/stretchr/testify/assert"
)

func bidirectionalMarshal(t *testing.T, doc proto.Message) {
	var buffer bytes.Buffer

	assert.NoError(t, protolator.DeepMarshalJSON(&buffer, doc))

	newRoot := proto.Clone(doc)
	newRoot.Reset()
	assert.NoError(t, protolator.DeepUnmarshalJSON(bytes.NewReader(buffer.Bytes()), newRoot))

//注意，我们不能在newroot和sampleDoc之间进行相等检查
//因为二元协议编组的不确定性
//所以我们改为重新编组到JSON，这是一个确定性编组
//把那里的平等比作

//T.Logo（DOC）
//T.Logo（NeWoRoT）

	var remarshaled bytes.Buffer
	assert.NoError(t, protolator.DeepMarshalJSON(&remarshaled, newRoot))
	assert.Equal(t, string(buffer.Bytes()), string(remarshaled.Bytes()))
//t.log（字符串（buffer.bytes（））
//t.log（字符串（remarkshaled.bytes（））
}

func TestConfigUpdate(t *testing.T) {
	cg, err := encoder.NewChannelGroup(configtxgentest.Load(genesisconfig.SampleSingleMSPSoloProfile))
	assert.NoError(t, err)

	bidirectionalMarshal(t, &cb.ConfigUpdateEnvelope{
		ConfigUpdate: utils.MarshalOrPanic(&cb.ConfigUpdate{
			ReadSet:  cg,
			WriteSet: cg,
		}),
	})
}

func TestIdemix(t *testing.T) {
	bidirectionalMarshal(t, &msp.MSPConfig{
		Type: 1,
		Config: utils.MarshalOrPanic(&msp.IdemixMSPConfig{
			Name: "fooo",
		}),
	})
}

func TestGenesisBlock(t *testing.T) {
	p := encoder.New(configtxgentest.Load(genesisconfig.SampleSingleMSPSoloProfile))
	gb := p.GenesisBlockForChannel("foo")

	bidirectionalMarshal(t, gb)
}

func TestEmitDefaultsBug(t *testing.T) {
	block := &cb.Block{
		Header: &cb.BlockHeader{
			PreviousHash: []byte("foo"),
		},
		Data: &cb.BlockData{
			Data: [][]byte{
				utils.MarshalOrPanic(&cb.Envelope{
					Payload: utils.MarshalOrPanic(&cb.Payload{
						Header: &cb.Header{
							ChannelHeader: utils.MarshalOrPanic(&cb.ChannelHeader{
								Type: int32(cb.HeaderType_CONFIG),
							}),
						},
					}),
					Signature: []byte("bar"),
				}),
			},
		},
	}

	err := protolator.DeepMarshalJSON(os.Stdout, block)
	assert.NoError(t, err)
}

func TestProposalResponsePayload(t *testing.T) {
	prp := &pb.ProposalResponsePayload{}
	assert.NoError(t, protolator.DeepUnmarshalJSON(bytes.NewReader([]byte(`{
            "extension": {
              "chaincode_id": {
                "name": "test",
                "path": "",
                "version": "1.0"
              },
              "events": {
                  "chaincode_id": "test"
              },
              "response": {
                "message": "",
                "payload": null,
                "status": 200
              },
              "results": {
                "data_model": "KV",
                "ns_rwset": [
                  {
                    "collection_hashed_rwset": [],
                    "namespace": "lscc",
                    "rwset": {
                      "metadata_writes": [],
                      "range_queries_info": [],
                      "reads": [
                        {
                          "key": "cc1",
                          "version": {
                            "block_num": "3",
                            "tx_num": "0"
                          }
                        },
                        {
                          "key": "cc2",
                          "version": {
                            "block_num": "4",
                            "tx_num": "0"
                          }
                        }
                      ],
                      "writes": []
                    }
                  },
                  {
                    "collection_hashed_rwset": [],
                    "namespace": "cc1",
                    "rwset": {
                      "metadata_writes": [],
                      "range_queries_info": [],
                      "reads": [
                        {
                          "key": "key1",
                          "version": {
                            "block_num": "8",
                            "tx_num": "0"
                          }
                        }
                      ],
                      "writes": [
                        {
                          "is_delete": false,
                          "key": "key2"
                        }
                      ]
                    }
                  },
                  {
                    "collection_hashed_rwset": [],
                    "namespace": "cc2",
                    "rwset": {
                      "metadata_writes": [],
                      "range_queries_info": [],
                      "reads": [
                        {
                          "key": "key1",
                          "version": {
                            "block_num": "9",
                            "tx_num": "0"
                          }
                        },
                        {
                          "key": "key2",
                          "version": {
                            "block_num": "10",
                            "tx_num": "0"
                          }
                        }
                      ],
                      "writes": [
                        {
                          "is_delete": false,
                          "key": "key1"
                        },
                        {
                          "is_delete": true,
                          "key": "key2"
                        }
                      ]
                    }
                  }
                ]
              }
            }
        }`)), prp))
	bidirectionalMarshal(t, prp)
}

func TestChannelCreationPolicy(t *testing.T) {
	cu := &cb.ConfigUpdate{
		WriteSet: &cb.ConfigGroup{
			Groups: map[string]*cb.ConfigGroup{
				"Consortiums": {
					Groups: map[string]*cb.ConfigGroup{
						"SampleConsortium": {
							Values: map[string]*cb.ConfigValue{
								"ChannelCreationPolicy": {
									Version: 0,
								},
							},
						},
					},
				},
			},
		},
	}

	bidirectionalMarshal(t, cu)
}

