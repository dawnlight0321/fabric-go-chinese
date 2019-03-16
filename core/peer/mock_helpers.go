
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:16</date>
//</624456038967414784>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package peer

import (
	configtxtest "github.com/hyperledger/fabric/common/configtx/test"
	mockchannelconfig "github.com/hyperledger/fabric/common/mocks/config"
	mockconfigtx "github.com/hyperledger/fabric/common/mocks/configtx"
	mockpolicies "github.com/hyperledger/fabric/common/mocks/policies"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/ledgermgmt"
)

//mockinitialize为测试env重置链
func MockInitialize() {
	ledgermgmt.InitializeTestEnvWithInitializer(
		&ledgermgmt.Initializer{
			CustomTxProcessors: ConfigTxProcessors,
		},
	)
	chains.list = make(map[string]*chain)
	chainInitializer = func(string) { return }
}

//mockCreateChain用于为测试链创建分类帐
//不必加入
func MockCreateChain(cid string) error {
	var ledger ledger.PeerLedger
	var err error

	if ledger = GetLedger(cid); ledger == nil {
		gb, _ := configtxtest.MakeGenesisBlock(cid)
		if ledger, err = ledgermgmt.CreateLedger(gb); err != nil {
			return err
		}
	}

	chains.Lock()
	defer chains.Unlock()

	chains.list[cid] = &chain{
		cs: &chainSupport{
			Resources: &mockchannelconfig.Resources{
				PolicyManagerVal: &mockpolicies.Manager{
					Policy: &mockpolicies.Policy{},
				},
				ConfigtxValidatorVal: &mockconfigtx.Validator{},
				ApplicationConfigVal: &mockchannelconfig.MockApplication{CapabilitiesRv: &mockchannelconfig.MockApplicationCapabilities{}},
			},

			ledger: ledger,
		},
	}

	return nil
}

