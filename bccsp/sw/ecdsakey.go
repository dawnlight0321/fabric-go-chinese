
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:51</date>
//</624455935925948416>

/*
版权所有IBM Corp.2016保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；

您可以在以下网址获得许可证副本：

   http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/

package sw

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/bccsp"
)

type ecdsaPrivateKey struct {
	privKey *ecdsa.PrivateKey
}

//字节将此键转换为其字节表示形式，
//如果允许此操作。
func (k *ecdsaPrivateKey) Bytes() ([]byte, error) {
	return nil, errors.New("Not supported.")
}

//ski返回此密钥的主题密钥标识符。
func (k *ecdsaPrivateKey) SKI() []byte {
	if k.privKey == nil {
		return nil
	}

//整理公钥
	raw := elliptic.Marshal(k.privKey.Curve, k.privKey.PublicKey.X, k.privKey.PublicKey.Y)

//散列
	hash := sha256.New()
	hash.Write(raw)
	return hash.Sum(nil)
}

//如果此密钥是对称密钥，则对称返回true，
//如果此密钥是非对称的，则为false
func (k *ecdsaPrivateKey) Symmetric() bool {
	return false
}

//如果此密钥是私钥，则private返回true，
//否则为假。
func (k *ecdsaPrivateKey) Private() bool {
	return true
}

//public key返回非对称公钥/私钥对的相应公钥部分。
//此方法返回对称密钥方案中的错误。
func (k *ecdsaPrivateKey) PublicKey() (bccsp.Key, error) {
	return &ecdsaPublicKey{&k.privKey.PublicKey}, nil
}

type ecdsaPublicKey struct {
	pubKey *ecdsa.PublicKey
}

//字节将此键转换为其字节表示形式，
//如果允许此操作。
func (k *ecdsaPublicKey) Bytes() (raw []byte, err error) {
	raw, err = x509.MarshalPKIXPublicKey(k.pubKey)
	if err != nil {
		return nil, fmt.Errorf("Failed marshalling key [%s]", err)
	}
	return
}

//ski返回此密钥的主题密钥标识符。
func (k *ecdsaPublicKey) SKI() []byte {
	if k.pubKey == nil {
		return nil
	}

//整理公钥
	raw := elliptic.Marshal(k.pubKey.Curve, k.pubKey.X, k.pubKey.Y)

//散列
	hash := sha256.New()
	hash.Write(raw)
	return hash.Sum(nil)
}

//如果此密钥是对称密钥，则对称返回true，
//如果此密钥是非对称的，则为false
func (k *ecdsaPublicKey) Symmetric() bool {
	return false
}

//如果此密钥是私钥，则private返回true，
//否则为假。
func (k *ecdsaPublicKey) Private() bool {
	return false
}

//public key返回非对称公钥/私钥对的相应公钥部分。
//此方法返回对称密钥方案中的错误。
func (k *ecdsaPublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}

