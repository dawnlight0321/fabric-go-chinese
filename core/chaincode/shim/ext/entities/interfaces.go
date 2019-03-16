
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:05</date>
//</624455992091873280>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package entities

//实体是所有加密实体的基本接口
//库用于获取CC级加密的
type Entity interface {
//id返回实体的标识符；
//
//
//
	ID() string

//
//
//
//
	Equals(Entity) bool

//
//如果使用非对称加密。如果不是，
//公共收益本身
	Public() (Entity, error)
}

//签名者是一个提供基本签名/验证功能的接口
type Signer interface {
//sign返回所提供消息的签名（或错误）
	Sign(msg []byte) (signature []byte, err error)

//验证是否检查提供的签名
//根据此接口，over提供的消息有效
	Verify(signature, msg []byte) (valid bool, err error)
}

//Encrypter是一个提供基本加密/解密功能的接口。
type Encrypter interface {
//
	Encrypt(plaintext []byte) (ciphertext []byte, err error)

//
	Decrypt(ciphertext []byte) (plaintext []byte, err error)
}

//
type EncrypterEntity interface {
	Entity
	Encrypter
}

//
type SignerEntity interface {
	Entity
	Signer
}

//
//
type EncrypterSignerEntity interface {
	Entity
	Encrypter
	Signer
}

