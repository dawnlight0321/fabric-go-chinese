
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:30</date>
//</624456096806866944>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package server

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/configtx"
	"github.com/hyperledger/fabric/common/crypto"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/common/ledger/blockledger/ram"
	"github.com/hyperledger/fabric/core/comm"
	"github.com/hyperledger/fabric/core/config/configtest"
	"github.com/hyperledger/fabric/orderer/common/cluster/mocks"
	"github.com/hyperledger/fabric/orderer/common/localconfig"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/orderer"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newServerNode(t *testing.T, key, cert []byte) *deliverServer {
	srv, err := comm.NewGRPCServer("127.0.0.1:0", comm.ServerConfig{
		SecOpts: &comm.SecureOptions{
			Key:         key,
			Certificate: cert,
			UseTLS:      true,
		},
	})
	if err != nil {
		panic(err)
	}
	ds := &deliverServer{
		t:              t,
		blockResponses: make(chan *orderer.DeliverResponse, 100),
		srv:            srv,
	}
	orderer.RegisterAtomicBroadcastServer(srv.Server(), ds)
	go srv.Start()
	return ds
}

type deliverServer struct {
	isConnected    int32
	t              *testing.T
	srv            *comm.GRPCServer
	blockResponses chan *orderer.DeliverResponse
}

func (*deliverServer) Broadcast(orderer.AtomicBroadcast_BroadcastServer) error {
	panic("implement me")
}

func (ds *deliverServer) Deliver(stream orderer.AtomicBroadcast_DeliverServer) error {
	atomic.StoreInt32(&ds.isConnected, 1)
	seekInfo, err := readSeekEnvelope(stream)
	if err != nil {
		panic(err)
	}
	if seekInfo.GetStart().GetSpecified() != nil {
		return ds.deliverBlocks(stream)
	}
	if seekInfo.GetStart().GetNewest() != nil {
		resp := <-ds.blockResponses
		return stream.Send(resp)
	}
	panic(fmt.Sprintf("expected either specified or newest seek but got %v", seekInfo.GetStart()))
}

func readSeekEnvelope(stream orderer.AtomicBroadcast_DeliverServer) (*orderer.SeekInfo, error) {
	env, err := stream.Recv()
	if err != nil {
		return nil, err
	}
	payload, err := utils.UnmarshalPayload(env.Payload)
	if err != nil {
		return nil, err
	}
	seekInfo := &orderer.SeekInfo{}
	if err = proto.Unmarshal(payload.Data, seekInfo); err != nil {
		return nil, err
	}
	return seekInfo, nil
}

func (ds *deliverServer) deliverBlocks(stream orderer.AtomicBroadcast_DeliverServer) error {
	for {
		blockChan := ds.blockResponses
		response := <-blockChan
		if response == nil {
			return nil
		}
		if err := stream.Send(response); err != nil {
			return err
		}
	}
}

func loadPEM(suffix string, t *testing.T) []byte {
	b, err := ioutil.ReadFile(filepath.Join("testdata", "tls", suffix))
	assert.NoError(t, err)
	return b
}

func TestReplicateIfNeeded(t *testing.T) {
	t.Parallel()

	caCert := loadPEM("ca.crt", t)
	key := loadPEM("server.key", t)
	cert := loadPEM("server.crt", t)

	deliverServer := newServerNode(t, key, cert)
	defer deliverServer.srv.Stop()

	flogging.ActivateSpec("testReplicateIfNeeded=debug")

	cleanup := configtest.SetDevFabricConfigPath(t)
	defer cleanup()

	blockBytes, err := ioutil.ReadFile(filepath.Join("testdata", "genesis.block"))
	assert.NoError(t, err)

	bootBlock := &common.Block{}
	assert.NoError(t, proto.Unmarshal(blockBytes, bootBlock))
	bootBlock.Header.Number = 10
	injectOrdererEndpoint(t, bootBlock, deliverServer.srv.Address())

	copyBlock := func(block *common.Block, seq uint64) *common.Block {
		res := &common.Block{}
		proto.Unmarshal(utils.MarshalOrPanic(block), res)
		res.Header.Number = seq
		return res
	}

	bootBlockWithCorruptedPayload := copyBlock(bootBlock, 100)
	env := &common.Envelope{}
	proto.Unmarshal(bootBlockWithCorruptedPayload.Data.Data[0], env)
	payload := &common.Payload{}
	proto.Unmarshal(env.Payload, payload)
	payload.Data = []byte{1, 2, 3}

	deliverServer.blockResponses <- &orderer.DeliverResponse{
		Type: &orderer.DeliverResponse_Block{Block: bootBlock},
	}

	blocks := make([]*common.Block, 11)
//创世块可以是任何东西…对于通道遍历不重要
//因为它被跳过了。
	blocks[0] = &common.Block{Header: &common.BlockHeader{}}
	for seq := uint64(1); seq <= uint64(10); seq++ {
		block := copyBlock(bootBlock, seq)
		block.Header.PreviousHash = blocks[seq-1].Header.Hash()
		blocks[seq] = block
		deliverServer.blockResponses <- &orderer.DeliverResponse{
			Type: &orderer.DeliverResponse_Block{Block: block},
		}
	}
//我们关闭块响应以标记要从中返回的服务器端
//方法调度。
	close(deliverServer.blockResponses)

//我们需要确保哈希链对于引导程序块是有效的。
//当我们遍历通道时，将注意验证哈希链本身
//在Fab-12926中。
	bootBlock.Header.PreviousHash = blocks[9].Header.Hash()

	var hooksActivated bool

	for _, testCase := range []struct {
		name               string
		panicValue         string
		systemLedgerHeight uint64
		bootBlock          *common.Block
		secOpts            *comm.SecureOptions
		conf               *localconfig.TopLevel
		ledgerFactoryErr   error
		signer             crypto.LocalSigner
		zapHooks           []func(zapcore.Entry) error
		shouldConnect      bool
	}{
		{
			name:               "Genesis block makes replication be skipped",
			bootBlock:          &common.Block{Header: &common.BlockHeader{Number: 0}},
			systemLedgerHeight: 10,
			zapHooks: []func(entry zapcore.Entry) error{
				func(entry zapcore.Entry) error {
					hooksActivated = true
					assert.Equal(t, entry.Message, "Booted with a genesis block, replication isn't an option")
					return nil
				},
			},
		},
		{
			name:               "Block puller initialization failure panics",
			systemLedgerHeight: 10,
			panicValue:         "Failed creating puller config from bootstrap block: unable to decode TLS certificate PEM: ",
			bootBlock:          bootBlockWithCorruptedPayload,
			conf:               &localconfig.TopLevel{},
			secOpts:            &comm.SecureOptions{},
		},
		{
			name:               "Extraction of sytem channel name fails",
			systemLedgerHeight: 10,
			panicValue:         "Failed extracting system channel name from bootstrap block: failed to retrieve channel id - block is empty",
			bootBlock:          &common.Block{Header: &common.BlockHeader{Number: 100}},
			conf:               &localconfig.TopLevel{},
			secOpts:            &comm.SecureOptions{},
		},
		{
			name:               "Is Replication needed fails",
			systemLedgerHeight: 10,
			ledgerFactoryErr:   errors.New("I/O error"),
			panicValue:         "Failed determining whether replication is needed: I/O error",
			bootBlock:          bootBlock,
			conf:               &localconfig.TopLevel{},
			secOpts: &comm.SecureOptions{
				Certificate: cert,
				Key:         key,
			},
		},
		{
			name:               "Replication isn't needed",
			systemLedgerHeight: 11,
			bootBlock:          bootBlock,
			conf:               &localconfig.TopLevel{},
			secOpts: &comm.SecureOptions{
				Certificate: cert,
				Key:         key,
			},
			zapHooks: []func(entry zapcore.Entry) error{
				func(entry zapcore.Entry) error {
					hooksActivated = true
					assert.Equal(t, entry.Message, "Replication isn't needed")
					return nil
				},
			},
		},
		{
			name: "Replication is needed, but pulling fails",
			panicValue: "Failed pulling system channel: " +
				"failed obtaining the latest block for channel testchainid",
			shouldConnect:      true,
			systemLedgerHeight: 10,
			bootBlock:          bootBlock,
			conf: &localconfig.TopLevel{
				General: localconfig.General{
					SystemChannel: "system",
					Cluster: localconfig.Cluster{
						ReplicationPullTimeout:  time.Millisecond * 100,
						DialTimeout:             time.Millisecond * 100,
						RPCTimeout:              time.Millisecond * 100,
						ReplicationRetryTimeout: time.Millisecond * 100,
						ReplicationBufferSize:   1,
					},
				},
			},
			secOpts: &comm.SecureOptions{
				Certificate:   cert,
				Key:           key,
				UseTLS:        true,
				ServerRootCAs: [][]byte{caCert},
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			lw := &mocks.LedgerWriter{}
			lw.On("Height").Return(testCase.systemLedgerHeight).Once()

			lf := &mocks.LedgerFactory{}
			lf.On("GetOrCreate", mock.Anything).Return(lw, testCase.ledgerFactoryErr).Once()
			lf.On("Close")

			r := &replicationInitiator{
				lf:     lf,
				logger: flogging.MustGetLogger("testReplicateIfNeeded"),
				signer: testCase.signer,

				conf:           testCase.conf,
				bootstrapBlock: testCase.bootBlock,
				secOpts:        testCase.secOpts,
			}

			if testCase.panicValue != "" {
				assert.PanicsWithValue(t, testCase.panicValue, r.replicateIfNeeded)
				return
			}

//否则，我们就不会惊慌失措。
			r.logger = r.logger.WithOptions(zap.Hooks(testCase.zapHooks...))

//这是我们正在测试的方法。
			r.replicateIfNeeded()

//确保我们为一个不惊慌的测试用例运行钩子
			assert.True(t, hooksActivated)
//恢复下一次迭代的标志
			defer func() {
				hooksActivated = false
			}()

			assert.Equal(t, testCase.shouldConnect, atomic.LoadInt32(&deliverServer.isConnected) == int32(1))
		})
	}
}

func TestLedgerFactory(t *testing.T) {
	lf := &ledgerFactory{ramledger.New(1)}
	lw, err := lf.GetOrCreate("mychannel")
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), lw.Height())
}

func injectOrdererEndpoint(t *testing.T, block *common.Block, endpoint string) {
	ordererAddresses := channelconfig.OrdererAddressesValue([]string{endpoint})
//展开层直到到达订购者地址
	env, err := utils.ExtractEnvelope(block, 0)
	assert.NoError(t, err)
	payload, err := utils.ExtractPayload(env)
	assert.NoError(t, err)
	confEnv, err := configtx.UnmarshalConfigEnvelope(payload.Data)
	assert.NoError(t, err)
//替换订购者地址
	confEnv.Config.ChannelGroup.Values[ordererAddresses.Key()].Value = utils.MarshalOrPanic(ordererAddresses.Value())
//把它放回街区
	payload.Data = utils.MarshalOrPanic(confEnv)
	env.Payload = utils.MarshalOrPanic(payload)
	block.Data.Data[0] = utils.MarshalOrPanic(env)
	block.Header.DataHash = block.Data.Hash()
}

