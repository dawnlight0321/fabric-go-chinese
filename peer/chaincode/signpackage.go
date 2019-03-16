
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:32</date>
//</624456107208740864>

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


package chaincode

import (
	"fmt"
	"io/ioutil"

	"github.com/hyperledger/fabric/core/common/ccpackage"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/spf13/cobra"
)

//signpackageCmd返回用于对包进行签名的COBRA命令
func signpackageCmd(cf *ChaincodeCmdFactory) *cobra.Command {
	spCmd := &cobra.Command{
		Use:       "signpackage",
		Short:     "Sign the specified chaincode package",
		Long:      "Sign the specified chaincode package",
		ValidArgs: []string{"2"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("peer chaincode signpackage <inputpackage> <outputpackage>")
			}
			return signpackage(cmd, args[0], args[1], cf)
		},
	}

	return spCmd
}

func signpackage(cmd *cobra.Command, ipackageFile string, opackageFile string, cf *ChaincodeCmdFactory) error {
//对命令行的分析已完成，因此沉默命令用法
	cmd.SilenceUsage = true

	var err error
	if cf == nil {
		cf, err = InitCmdFactory(cmd.Name(), false, false)
		if err != nil {
			return err
		}
	}

	b, err := ioutil.ReadFile(ipackageFile)
	if err != nil {
		return err
	}

	env := utils.UnmarshalEnvelopeOrPanic(b)

	env, err = ccpackage.SignExistingPackage(env, cf.Signer)
	if err != nil {
		return err
	}

	b = utils.MarshalOrPanic(env)
	err = ioutil.WriteFile(opackageFile, b, 0700)
	if err != nil {
		return err
	}

	fmt.Printf("Wrote signed package to %s successfully\n", opackageFile)

	return nil
}

