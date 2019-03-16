
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:11</date>
//</624456017786179584>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package customtx

import (
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/protos/common"
)

//自定义事务处理器（接口“processor”的实现）应引发InvalidTxError。
//如果它希望分类帐将特定交易记录为无效
type InvalidTxError struct {
	Msg string
}

func (e *InvalidTxError) Error() string {
	return e.Msg
}

//处理器允许在提交期间为自定义事务生成模拟结果。
//自定义处理器可以以适当的方式表示信息，并可以使用此过程来翻译
//信息以“txSimulationResults”的形式显示。因为原始信息是在
//自定义表示，“processor”的一个实现应注意自定义表示
//用于确定性方式的模拟，应注意跨结构版本的兼容性。
//“initializingledger”true表示正在处理的事务来自Genesis块或分类帐是
//同步状态（如果发现statedb落后于区块链，则可能在对等启动期间发生）。
//在前一种情况下，处理的交易预期是有效的，在后一种情况下，只有有效的交易
//重新处理，因此可以跳过任何验证。
type Processor interface {
	GenerateSimulationResults(txEnvelop *common.Envelope, simulator ledger.TxSimulator, initializingLedger bool) error
}

