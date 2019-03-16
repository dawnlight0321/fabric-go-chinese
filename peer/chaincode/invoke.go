
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:32</date>
//</624456106210496512>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package chaincode

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var chaincodeInvokeCmd *cobra.Command

//invokeCmd返回chaincode invoke的cobra命令
func invokeCmd(cf *ChaincodeCmdFactory) *cobra.Command {
	chaincodeInvokeCmd = &cobra.Command{
		Use:       "invoke",
		Short:     fmt.Sprintf("Invoke the specified %s.", chainFuncName),
		Long:      fmt.Sprintf("Invoke the specified %s. It will try to commit the endorsed transaction to the network.", chainFuncName),
		ValidArgs: []string{"1"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return chaincodeInvoke(cmd, cf)
		},
	}
	flagList := []string{
		"name",
		"ctor",
		"channelID",
		"peerAddresses",
		"tlsRootCertFiles",
		"connectionProfile",
		"waitForEvent",
		"waitForEventTimeout",
	}
	attachFlags(chaincodeInvokeCmd, flagList)

	return chaincodeInvokeCmd
}

func chaincodeInvoke(cmd *cobra.Command, cf *ChaincodeCmdFactory) error {
	if channelID == "" {
		return errors.New("The required parameter 'channelID' is empty. Rerun the command with -C flag")
	}
//对命令行的分析已完成，因此沉默命令用法
	cmd.SilenceUsage = true

	var err error
	if cf == nil {
		cf, err = InitCmdFactory(cmd.Name(), true, true)
		if err != nil {
			return err
		}
	}
	defer cf.BroadcastClient.Close()

	return chaincodeInvokeOrQuery(cmd, true, cf)
}

