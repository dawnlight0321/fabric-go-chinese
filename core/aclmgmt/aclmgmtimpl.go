
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455976296124416>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package aclmgmt

import (
	"github.com/hyperledger/fabric/common/flogging"
)

var aclMgmtLogger = flogging.MustGetLogger("aclmgmt")

//实施ACLMgmt。结构中的checkacl调用导致以下流
//如果资源提供程序[资源名称]
//返回resourceprovider[resourcename].checkacl（…）
//其他的
//返回默认提供程序[resourcename].checkacl（…）
//使用rescfgprovider封装resourceprovider和defaultprovider
type aclMgmtImpl struct {
//资源提供程序从配置获取资源信息
	rescfgProvider ACLProvider
}

//checkacl使用
//IDFIN。IDinfo是一个对象，如SignedProposal，其中
//可以提取ID以根据策略进行测试
func (am *aclMgmtImpl) CheckACL(resName string, channelID string, idinfo interface{}) error {
//使用基于资源的配置提供程序（这将反过来默认为1.0提供程序）
	return am.rescfgProvider.CheckACL(resName, channelID, idinfo)
}

//aclprovider由两个提供程序组成，分别提供了一个和一个默认的提供程序（1.0 acl管理
//使用ChannelReader和ChannelWriter）。如果提供的提供程序为nil，则基于资源
//已创建ACL提供程序。
func NewACLProvider(rg ResourceGetter) ACLProvider {
	return &aclMgmtImpl{
		rescfgProvider: newResourceProvider(rg, NewDefaultACLProvider()),
	}
}

