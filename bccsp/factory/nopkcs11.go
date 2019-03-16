
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:49</date>
//</624455926253883392>

//+建设！PKCS11

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

package factory

import (
	"github.com/hyperledger/fabric/bccsp"
	"github.com/pkg/errors"
)

//FactoryOpts保存用于初始化工厂实现的配置信息
type FactoryOpts struct {
	ProviderName string      `mapstructure:"default" json:"default" yaml:"Default"`
	SwOpts       *SwOpts     `mapstructure:"SW,omitempty" json:"SW,omitempty" yaml:"SwOpts"`
	PluginOpts   *PluginOpts `mapstructure:"PLUGIN,omitempty" json:"PLUGIN,omitempty" yaml:"PluginOpts"`
}

//必须在使用工厂接口之前调用InitFactories
//可以使用config=nil调用，在这种情况下
//将使用某些默认值
//只有在找不到defaultbccsp时才会返回错误
func InitFactories(config *FactoryOpts) error {
	factoriesInitOnce.Do(func() {
//对默认选项采取一些预防措施
		if config == nil {
			config = GetDefaultOpts()
		}

		if config.ProviderName == "" {
			config.ProviderName = "SW"
		}

		if config.SwOpts == nil {
			config.SwOpts = GetDefaultOpts().SwOpts
		}

//初始化工厂映射
		bccspMap = make(map[string]bccsp.BCCSP)

//基于软件的BCCSP
		if config.SwOpts != nil {
			f := &SWFactory{}
			err := initBCCSP(f, config)
			if err != nil {
				factoriesInitError = errors.Wrapf(err, "Failed initializing BCCSP.")
			}
		}

//BCCSP插件
		if config.PluginOpts != nil {
			f := &PluginFactory{}
			err := initBCCSP(f, config)
			if err != nil {
				factoriesInitError = errors.Wrapf(err, "Failed initializing PKCS11.BCCSP %s", factoriesInitError)
			}
		}

		var ok bool
		defaultBCCSP, ok = bccspMap[config.ProviderName]
		if !ok {
			factoriesInitError = errors.Errorf("%s\nCould not find default `%s` BCCSP", factoriesInitError, config.ProviderName)
		}
	})

	return factoriesInitError
}

//getbccspfromopts返回根据输入中传递的选项创建的bccsp。
func GetBCCSPFromOpts(config *FactoryOpts) (bccsp.BCCSP, error) {
	var f BCCSPFactory
	switch config.ProviderName {
	case "SW":
		f = &SWFactory{}
	case "PLUGIN":
		f = &PluginFactory{}
	default:
		return nil, errors.Errorf("Could not find BCCSP, no '%s' provider", config.ProviderName)
	}

	csp, err := f.Get(config)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not initialize BCCSP %s", f.Name())
	}
	return csp, nil
}

