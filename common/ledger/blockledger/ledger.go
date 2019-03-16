
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:57</date>
//</624455960475209728>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package blockledger

import (
	cb "github.com/hyperledger/fabric/protos/common"
	ab "github.com/hyperledger/fabric/protos/orderer"
)

//
type Factory interface {
//
//
	GetOrCreate(chainID string) (ReadWriter, error)

//chainIds返回工厂知道的链ID
	ChainIDs() []string

//
	Close()
}

//迭代器对于链读取器在创建块时流式处理块很有用。
type Iterator interface {
//
//下一个块不再可检索
	Next() (*cb.Block, cb.Status)
//
	Close()
}

//
type Reader interface {
//
//
	Iterator(startType *ab.SeekPosition) (Iterator, uint64)
//height返回分类帐上的块数
	Height() uint64
}

//
type Writer interface {
//将新块追加到分类帐
	Append(block *cb.Block) error
}

//

//
type ReadWriter interface {
	Reader
	Writer
}

