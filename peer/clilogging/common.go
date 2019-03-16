
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:33</date>
//</624456109230395392>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package clilogging

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/crypto"
	"github.com/hyperledger/fabric/peer/common"
	common2 "github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type envelopeWrapper func(msg proto.Message) *common2.Envelope

//loggingCmdFactory保留loggingCmd使用的客户端
type LoggingCmdFactory struct {
	AdminClient      pb.AdminClient
	wrapWithEnvelope envelopeWrapper
}

//initcmdfactory用默认管理客户端初始化loggingCmdfactory
func InitCmdFactory() (*LoggingCmdFactory, error) {
	var err error
	var adminClient pb.AdminClient

	adminClient, err = common.GetAdminClient()
	if err != nil {
		return nil, err
	}

	signer, err := common.GetDefaultSignerFnc()
	if err != nil {
		return nil, errors.Errorf("failed obtaining default signer: %v", err)
	}

	localSigner := crypto.NewSignatureHeaderCreator(signer)
	wrapEnv := func(msg proto.Message) *common2.Envelope {
		env, err := utils.CreateSignedEnvelope(common2.HeaderType_PEER_ADMIN_OPERATION, "", localSigner, msg, 0, 0)
		if err != nil {
			logger.Panicf("Failed signing: %v", err)
		}
		return env
	}

	return &LoggingCmdFactory{
		AdminClient:      adminClient,
		wrapWithEnvelope: wrapEnv,
	}, nil
}

func checkLoggingCmdParams(cmd *cobra.Command, args []string) error {
	var err error
	if cmd.Name() == "revertlevels" || cmd.Name() == "getlogspec" {
		if len(args) > 0 {
			err = errors.Errorf("more parameters than necessary were provided. Expected 0, received %d", len(args))
			return err
		}
	} else {
//检查是否至少传入了一个参数
		if len(args) == 0 {
			err = errors.New("no parameters provided")
			return err
		}
	}

	if cmd.Name() == "setlevel" {
//检查是否提供了日志级别参数
		if len(args) == 1 {
			err = errors.New("no log level provided")
		} else {
//检查日志级别是否有效。如果没有，则设置err
			err = common.CheckLogLevel(args[1])
		}
	}

	return err
}

