
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:22</date>
//</624456064015798272>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"time"

	"github.com/hyperledger/fabric/common/crypto"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/common/localmsp"
	genesisconfig "github.com/hyperledger/fabric/common/tools/configtxgen/localconfig"
	"github.com/hyperledger/fabric/common/tools/protolator"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/comm"
	config2 "github.com/hyperledger/fabric/core/config"
	"github.com/hyperledger/fabric/msp"
	common2 "github.com/hyperledger/fabric/peer/common"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/orderer"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/spf13/viper"
)

var (
	channelID        string
	serverAddr       string
	clientKeyPath    string
	clientCertPath   string
	serverRootCAPath string
	seek             int
	quiet            bool
	filtered         bool
	tlsEnabled       bool
	mTlsEnabled      bool

	oldest  = &orderer.SeekPosition{Type: &orderer.SeekPosition_Oldest{Oldest: &orderer.SeekOldest{}}}
	newest  = &orderer.SeekPosition{Type: &orderer.SeekPosition_Newest{Newest: &orderer.SeekNewest{}}}
	maxStop = &orderer.SeekPosition{Type: &orderer.SeekPosition_Specified{Specified: &orderer.SeekSpecified{Number: math.MaxUint64}}}
)

const (
	OLDEST = -2
	NEWEST = -1

	ROOT = "core"
)

var logger = flogging.MustGetLogger("eventsclient")

//DeliverClient抽象公共接口
//用于传递和传递筛选的服务
type deliverClient interface {
	Send(*common.Envelope) error
	Recv() (*peer.DeliverResponse, error)
}

//要连接到的事件客户端
//事件对等传递系统
type eventsClient struct {
	client      deliverClient
	signer      crypto.LocalSigner
	tlsCertHash []byte
}

func (r *eventsClient) seekOldest() error {
	return r.client.Send(r.seekHelper(oldest, maxStop))
}

func (r *eventsClient) seekNewest() error {
	return r.client.Send(r.seekHelper(newest, maxStop))
}

func (r *eventsClient) seekSingle(blockNumber uint64) error {
	specific := &orderer.SeekPosition{Type: &orderer.SeekPosition_Specified{Specified: &orderer.SeekSpecified{Number: blockNumber}}}
	return r.client.Send(r.seekHelper(specific, specific))
}

func (r *eventsClient) seekHelper(start *orderer.SeekPosition, stop *orderer.SeekPosition) *common.Envelope {
	env, err := utils.CreateSignedEnvelopeWithTLSBinding(common.HeaderType_DELIVER_SEEK_INFO, channelID, r.signer, &orderer.SeekInfo{
		Start:    start,
		Stop:     stop,
		Behavior: orderer.SeekInfo_BLOCK_UNTIL_READY,
	}, 0, 0, r.tlsCertHash)
	if err != nil {
		panic(err)
	}
	return env
}

func (r *eventsClient) readEventsStream() {
	for {
		msg, err := r.client.Recv()
		if err != nil {
			logger.Info("Error receiving:", err)
			return
		}

		switch t := msg.Type.(type) {
		case *peer.DeliverResponse_Status:
			logger.Info("Got status ", t)
			return
		case *peer.DeliverResponse_Block:
			if !quiet {
				logger.Info("Received block: ")
				err := protolator.DeepMarshalJSON(os.Stdout, t.Block)
				if err != nil {
					fmt.Printf("  Error pretty printing block: %s", err)
				}
			} else {
				logger.Info("Received block: ", t.Block.Header.Number)
			}
		case *peer.DeliverResponse_FilteredBlock:
			if !quiet {
				logger.Info("Received filtered block: ")
				err := protolator.DeepMarshalJSON(os.Stdout, t.FilteredBlock)
				if err != nil {
					fmt.Printf("  Error pretty printing filtered block: %s", err)
				}
			} else {
				logger.Info("Received filtered block: ", t.FilteredBlock.Number)
			}
		}
	}
}

func (r *eventsClient) seek(s int) error {
	var err error
	switch seek {
	case OLDEST:
		err = r.seekOldest()
	case NEWEST:
		err = r.seekNewest()
	default:
		err = r.seekSingle(uint64(seek))
	}
	return err
}

func main() {
	initConfig()
	initMSP()
	readCLInputs()

	if seek < OLDEST {
		logger.Info("Invalid seek value")
		flag.PrintDefaults()
		return
	}

	clientConfig := comm.ClientConfig{
		KaOpts:  comm.DefaultKeepaliveOptions,
		SecOpts: &comm.SecureOptions{},
		Timeout: 5 * time.Minute,
	}

	if tlsEnabled {
		clientConfig.SecOpts.UseTLS = true
		rootCert, err := ioutil.ReadFile(serverRootCAPath)
		if err != nil {
			logger.Info("error loading TLS root certificate", err)
			return
		}
		clientConfig.SecOpts.ServerRootCAs = [][]byte{rootCert}
		if mTlsEnabled {
			clientConfig.SecOpts.RequireClientCert = true
			clientKey, err := ioutil.ReadFile(clientKeyPath)
			if err != nil {
				logger.Info("error loading client TLS key from", clientKeyPath)
				return
			}
			clientConfig.SecOpts.Key = clientKey

			clientCert, err := ioutil.ReadFile(clientCertPath)
			if err != nil {
				logger.Info("error loading client TLS certificate from path", clientCertPath)
				return
			}
			clientConfig.SecOpts.Certificate = clientCert
		}
	}

	grpcClient, err := comm.NewGRPCClient(clientConfig)
	if err != nil {
		logger.Info("Error creating grpc client:", err)
		return
	}
	conn, err := grpcClient.NewConnection(serverAddr, "")
	if err != nil {
		logger.Info("Error connecting:", err)
		return
	}

	var client deliverClient
	if filtered {
		client, err = peer.NewDeliverClient(conn).DeliverFiltered(context.Background())
	} else {
		client, err = peer.NewDeliverClient(conn).Deliver(context.Background())
	}

	if err != nil {
		logger.Info("Error connecting:", err)
		return
	}

	events := &eventsClient{
		client: client,
		signer: localmsp.NewSigner(),
	}

	if mTlsEnabled {
		events.tlsCertHash = util.ComputeSHA256(grpcClient.Certificate().Certificate[0])
	}

	events.seek(seek)
	if err != nil {
		logger.Info("Received error:", err)
		return
	}

	events.readEventsStream()
}

func readCLInputs() {
	flag.StringVar(&serverAddr, "server", "localhost:7051", "The RPC server to connect to.")
	flag.StringVar(&channelID, "channelID", genesisconfig.TestChainID, "The channel ID to deliver from.")
	flag.BoolVar(&quiet, "quiet", false, "Only print the block number, will not attempt to print its block contents.")
	flag.BoolVar(&filtered, "filtered", true, "Whenever to read filtered events from the peer delivery service or get regular blocks.")
	flag.BoolVar(&tlsEnabled, "tls", false, "TLS enabled/disabled")
	flag.BoolVar(&mTlsEnabled, "mTls", false, "Mutual TLS enabled/disabled (whenever server side validates clients TLS certificate)")
	flag.StringVar(&clientKeyPath, "clientKey", "", "Specify path to the client TLS key")
	flag.StringVar(&clientCertPath, "clientCert", "", "Specify path to the client TLS certificate")
	flag.StringVar(&serverRootCAPath, "rootCert", "", "Specify path to the server root CA certificate")
	flag.IntVar(&seek, "seek", OLDEST, "Specify the range of requested blocks."+
		"Acceptable values:"+
		"-2 (or -1) to start from oldest (or newest) and keep at it indefinitely."+
		"N >= 0 to fetch block N only.")
	flag.Parse()
}

func initMSP() {
//初始化MSP
	var mspMgrConfigDir = config2.GetPath("peer.mspConfigPath")
	var mspID = viper.GetString("peer.localMspId")
	var mspType = viper.GetString("peer.localMspType")
	if mspType == "" {
		mspType = msp.ProviderTypeToString(msp.FABRIC)
	}
	err := common2.InitCrypto(mspMgrConfigDir, mspID, mspType)
if err != nil { //处理读取配置文件时的错误
		panic(fmt.Sprintf("Cannot run client because %s", err.Error()))
	}
}

func initConfig() {
//对于环境变量。
	viper.SetEnvPrefix(ROOT)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	err := common2.InitConfig(ROOT)
if err != nil { //处理读取配置文件时的错误
		panic(fmt.Errorf("fatal error when initializing %s config : %s", ROOT, err))
	}
}

