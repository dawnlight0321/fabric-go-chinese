
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:50</date>
//</624455929227644928>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package bridge

import (
	"github.com/hyperledger/fabric-amcl/amcl"
	"github.com/hyperledger/fabric/bccsp/idemix/handlers"

	"github.com/golang/protobuf/proto"
	cryptolib "github.com/hyperledger/fabric/idemix"
	"github.com/pkg/errors"
)

//nymsignaturescheme封装IDemix算法，使用IDemix进行签名和验证
//假名。
type NymSignatureScheme struct {
	NewRand func() *amcl.RAND
}

//签名在传递的摘要上生成签名。它接受输入，用户密钥（sk）。
//假名公钥（NYM）和密钥（RNYM）以及颁发者公钥（IPK）。
func (n *NymSignatureScheme) Sign(sk handlers.Big, Nym handlers.Ecp, RNym handlers.Big, ipk handlers.IssuerPublicKey, digest []byte) (res []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			res = nil
			err = errors.Errorf("failure [%s]", r)
		}
	}()

	isk, ok := sk.(*Big)
	if !ok {
		return nil, errors.Errorf("invalid user secret key, expected *Big, got [%T]", sk)
	}
	inym, ok := Nym.(*Ecp)
	if !ok {
		return nil, errors.Errorf("invalid nym public key, expected *Ecp, got [%T]", Nym)
	}
	irnym, ok := RNym.(*Big)
	if !ok {
		return nil, errors.Errorf("invalid nym secret key, expected *Big, got [%T]", RNym)
	}
	iipk, ok := ipk.(*IssuerPublicKey)
	if !ok {
		return nil, errors.Errorf("invalid issuer public key, expected *IssuerPublicKey, got [%T]", ipk)
	}

	sig, err := cryptolib.NewNymSignature(
		isk.E,
		inym.E,
		irnym.E,
		iipk.PK,
		digest,
		n.NewRand())
	if err != nil {
		return nil, errors.WithMessage(err, "failed creating new nym signature")
	}

	return proto.Marshal(sig)
}

//验证是否检查传递的签名对于传递的摘要、颁发者公钥有效，
//和假名公钥。
func (*NymSignatureScheme) Verify(ipk handlers.IssuerPublicKey, Nym handlers.Ecp, signature, digest []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("failure [%s]", r)
		}
	}()

	iipk, ok := ipk.(*IssuerPublicKey)
	if !ok {
		return errors.Errorf("invalid issuer public key, expected *IssuerPublicKey, got [%T]", ipk)
	}
	inym, ok := Nym.(*Ecp)
	if !ok {
		return errors.Errorf("invalid nym public key, expected *Ecp, got [%T]", Nym)
	}

	sig := &cryptolib.NymSignature{}
	err = proto.Unmarshal(signature, sig)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling signature")
	}

	return sig.Ver(inym.E, iipk.PK, digest)
}

