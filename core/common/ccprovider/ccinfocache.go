
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:06</date>
//</624455999486431232>

/*
版权所有IBM Corp.2017保留所有权利。

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


package ccprovider

import (
	"fmt"
	"sync"
)

//CCinfocacheimpl为链码数据实现内存缓存
//背书人需要验证本地实例化策略
//在兑现之前匹配通道上的实例化策略
//调用
type ccInfoCacheImpl struct {
	sync.RWMutex

	cache        map[string]*ChaincodeData
	cacheSupport CCCacheSupport
}

//newccinfo缓存在提供的ccinfo提供程序实例上返回新缓存
func NewCCInfoCache(cs CCCacheSupport) *ccInfoCacheImpl {
	return &ccInfoCacheImpl{
		cache:        make(map[string]*ChaincodeData),
		cacheSupport: cs,
	}
}

func (c *ccInfoCacheImpl) GetChaincodeData(ccname string, ccversion string) (*ChaincodeData, error) {
//c.cache保证非零

	key := ccname + "/" + ccversion

	c.RLock()
	ccdata, in := c.cache[key]
	c.RUnlock()

	if !in {
		var err error

//链码数据不在缓存中
//尝试从文件系统中查找
		ccpack, err := c.cacheSupport.GetChaincode(ccname, ccversion)
		if err != nil || ccpack == nil {
			return nil, fmt.Errorf("cannot retrieve package for chaincode %s/%s, error %s", ccname, ccversion, err)
		}

//我们有一个非零的chaincodedata，放在缓存中
		c.Lock()
		ccdata = ccpack.GetChaincodeData()
		c.cache[key] = ccdata
		c.Unlock()
	}

	return ccdata, nil
}

