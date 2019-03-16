
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:51</date>
//</624455936911609856>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package sw

import (
	"encoding/hex"
	"sync"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/pkg/errors"
)

//NewInMemoryKeyStore实例化内存中的临时密钥库
func NewInMemoryKeyStore() bccsp.KeyStore {
	eks := &inmemoryKeyStore{}
	eks.keys = make(map[string]bccsp.Key)
	return eks
}

type inmemoryKeyStore struct {
//keys将十六进制编码的ski映射到keys
	keys map[string]bccsp.Key
	m    sync.RWMutex
}

//read only返回false-密钥存储不是只读的
func (ks *inmemoryKeyStore) ReadOnly() bool {
	return false
}

//getkey返回一个key对象，其ski是通过的。
func (ks *inmemoryKeyStore) GetKey(ski []byte) (bccsp.Key, error) {
	if len(ski) == 0 {
		return nil, errors.New("ski is nil or empty")
	}

	skiStr := hex.EncodeToString(ski)

	ks.m.RLock()
	defer ks.m.RUnlock()
	if key, found := ks.keys[skiStr]; found {
		return key, nil
	}
	return nil, errors.Errorf("no key found for ski %x", ski)
}

//storekey将密钥k存储在此密钥库中。
func (ks *inmemoryKeyStore) StoreKey(k bccsp.Key) error {
	if k == nil {
		return errors.New("key is nil")
	}

	ski := hex.EncodeToString(k.SKI())

	ks.m.Lock()
	defer ks.m.Unlock()

	if _, found := ks.keys[ski]; found {
		return errors.Errorf("ski %x already exists in the keystore", k.SKI())
	}
	ks.keys[ski] = k

	return nil
}

