
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456126821306368>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package client

import (
	"context"
	"io"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric/protos/token"
	tk "github.com/hyperledger/fabric/token"
	"github.com/pkg/errors"
)

type TimeFunc func() time.Time

type ProverPeer struct {
	ChannelID        string
	ProverClient     token.ProverClient
	RandomnessReader io.Reader
	Time             TimeFunc
}

func (prover *ProverPeer) RequestImport(tokensToIssue []*token.TokenToIssue, signingIdentity tk.SigningIdentity) ([]byte, error) {
	ir := &token.ImportRequest{
		TokensToIssue: tokensToIssue,
	}
	payload := &token.Command_ImportRequest{ImportRequest: ir}

	sc, err := prover.CreateSignedCommand(payload, signingIdentity)

	if err != nil {
		return nil, err
	}

	scr, err := prover.ProverClient.ProcessCommand(context.Background(), sc)
	if err != nil {
		return nil, err
	}
	return scr.Response, nil
}

func (prover *ProverPeer) RequestTransfer(
	tokenIDs [][]byte,
	shares []*token.RecipientTransferShare,
	signingIdentity tk.SigningIdentity) ([]byte, error) {

	tr := &token.TransferRequest{
		Shares:   shares,
		TokenIds: tokenIDs,
	}
	payload := &token.Command_TransferRequest{TransferRequest: tr}

	sc, err := prover.CreateSignedCommand(payload, signingIdentity)
	if err != nil {
		return nil, err
	}
	scr, err := prover.ProverClient.ProcessCommand(context.Background(), sc)
	if err != nil {
		return nil, err
	}

	return scr.Response, nil
}

func (prover *ProverPeer) CreateSignedCommand(payload interface{}, signingIdentity tk.SigningIdentity) (*token.SignedCommand, error) {

	command, err := commandFromPayload(payload)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, 32)
	_, err = io.ReadFull(prover.RandomnessReader, nonce)
	if err != nil {
		return nil, err
	}

	ts, err := ptypes.TimestampProto(prover.Time())
	if err != nil {
		return nil, err
	}

	creator, err := signingIdentity.GetPublicVersion().Serialize()
	if err != nil {
		return nil, err
	}

	header := &token.Header{Timestamp: ts,
		Nonce:     nonce,
		Creator:   creator,
		ChannelId: prover.ChannelID,
	}
	command.Header = header

	raw, err := proto.Marshal(command)
	if err != nil {
		return nil, err
	}

	signature, err := signingIdentity.Sign(raw)
	if err != nil {
		return nil, err
	}

	sc := &token.SignedCommand{
		Command:   raw,
		Signature: signature,
	}
	return sc, nil
}

func commandFromPayload(payload interface{}) (*token.Command, error) {
	switch t := payload.(type) {
	case *token.Command_ImportRequest:
		return &token.Command{Payload: t}, nil
	case *token.Command_TransferRequest:
		return &token.Command{Payload: t}, nil
	default:
		return nil, errors.Errorf("command type not recognized: %T", t)
	}
}

