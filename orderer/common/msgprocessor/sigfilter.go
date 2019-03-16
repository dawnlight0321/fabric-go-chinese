
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:29</date>
//</624456094290284544>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package msgprocessor

import (
	"fmt"

	"github.com/hyperledger/fabric/common/policies"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/pkg/errors"
)

//sigfiltersupport提供签名筛选器所需的资源
type SigFilterSupport interface {
//policyManager返回对当前策略管理器的引用
	PolicyManager() policies.Manager
}

//sigfilter存储应用于将请求传递到的策略的名称
//确定客户是否被授权
type SigFilter struct {
	policyName string
	support    SigFilterSupport
}

//newsigfilter创建一个新的签名筛选器，每次评估时都调用策略管理器
//检索策略的最新版本
func NewSigFilter(policyName string, support SigFilterSupport) *SigFilter {
	return &SigFilter{
		policyName: policyName,
		support:    support,
	}
}

//应用应用应用给定的策略，导致拒绝或转发，从不接受
func (sf *SigFilter) Apply(message *cb.Envelope) error {
	signedData, err := message.AsSignedData()

	if err != nil {
		return fmt.Errorf("could not convert message to signedData: %s", err)
	}

	policy, ok := sf.support.PolicyManager().GetPolicy(sf.policyName)
	if !ok {
		return fmt.Errorf("could not find policy %s", sf.policyName)
	}

	err = policy.Evaluate(signedData)
	if err != nil {
		return errors.Wrap(errors.WithStack(ErrPermissionDenied), err.Error())
	}
	return nil
}

