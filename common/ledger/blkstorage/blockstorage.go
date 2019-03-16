
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:56</date>
//</624455957409173504>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package blkstorage

import (
	"github.com/hyperledger/fabric/common/ledger"
	l "github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)

//indexableattr表示可索引属性
type IndexableAttr string

//可索引属性的常量
const (
	IndexableAttrBlockNum         = IndexableAttr("BlockNum")
	IndexableAttrBlockHash        = IndexableAttr("BlockHash")
	IndexableAttrTxID             = IndexableAttr("TxID")
	IndexableAttrBlockNumTranNum  = IndexableAttr("BlockNumTranNum")
	IndexableAttrBlockTxID        = IndexableAttr("BlockTxID")
	IndexableAttrTxValidationCode = IndexableAttr("TxValidationCode")
)

//indexconfig-包含应索引的属性列表的配置
type IndexConfig struct {
	AttrsToIndex []IndexableAttr
}

var (
//errnotfoundinindex用于指示索引中缺少条目。
	ErrNotFoundInIndex = l.NotFoundInIndexErr("")

//ErratrNotIndexed用于指示未对属性进行索引
	ErrAttrNotIndexed = errors.New("attribute not indexed")
)

//BlockstoreProvider提供Blockstore的句柄
type BlockStoreProvider interface {
	CreateBlockStore(ledgerid string) (BlockStore, error)
	OpenBlockStore(ledgerid string) (BlockStore, error)
	Exists(ledgerid string) (bool, error)
	List() ([]string, error)
	Close()
}

//blockstore-用于持久化和检索块的接口
//此接口的实现应采用参数
//'indexconfig'类型，用于配置应为哪些项编制索引的块存储
type BlockStore interface {
	AddBlock(block *common.Block) error
	GetBlockchainInfo() (*common.BlockchainInfo, error)
	RetrieveBlocks(startNum uint64) (ledger.ResultsIterator, error)
	RetrieveBlockByHash(blockHash []byte) (*common.Block, error)
RetrieveBlockByNumber(blockNum uint64) (*common.Block, error) //math.maxuint64的blocknum将返回最后一个块
	RetrieveTxByID(txID string) (*common.Envelope, error)
	RetrieveTxByBlockNumTranNum(blockNum uint64, tranNum uint64) (*common.Envelope, error)
	RetrieveBlockByTxID(txID string) (*common.Block, error)
	RetrieveTxValidationCodeByTxID(txID string) (peer.TxValidationCode, error)
	Shutdown()
}

