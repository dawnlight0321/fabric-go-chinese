
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:57</date>
//</624455959908978688>

/*
版权所有IBM Corp.2016保留所有权利。

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


package blockledger_test

import (
	"io/ioutil"
	"os"

	. "github.com/hyperledger/fabric/common/ledger/blockledger"
	fileledger "github.com/hyperledger/fabric/common/ledger/blockledger/file"
	genesisconfig "github.com/hyperledger/fabric/common/tools/configtxgen/localconfig"
)

func init() {
	testables = append(testables, &fileLedgerTestEnv{})
}

type fileLedgerTestFactory struct {
	location string
}

type fileLedgerTestEnv struct {
}

func (env *fileLedgerTestEnv) Initialize() (ledgerTestFactory, error) {
	var err error
	location, err := ioutil.TempDir("", "hyperledger")
	if err != nil {
		return nil, err
	}
	return &fileLedgerTestFactory{location: location}, nil
}

func (env *fileLedgerTestEnv) Name() string {
	return "fileledger"
}

func (env *fileLedgerTestEnv) Close(lf Factory) {
	lf.Close()
}

func (env *fileLedgerTestFactory) Destroy() error {
	err := os.RemoveAll(env.location)
	return err
}

func (env *fileLedgerTestFactory) Persistent() bool {
	return true
}

func (env *fileLedgerTestFactory) New() (Factory, ReadWriter) {
	flf := fileledger.New(env.location)
	fl, err := flf.GetOrCreate(genesisconfig.TestChainID)
	if err != nil {
		panic(err)
	}
	if fl.Height() == 0 {
		if err = fl.Append(genesisBlock); err != nil {
			panic(err)
		}
	}
	return flf, fl
}

