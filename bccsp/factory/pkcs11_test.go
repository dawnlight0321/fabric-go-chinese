
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:49</date>
//</624455927231156224>

//+构建PKCS11

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
	"os"
	"testing"

	"github.com/hyperledger/fabric/bccsp/pkcs11"
	"github.com/stretchr/testify/assert"
)

func TestInitFactories(t *testing.T) {
//从以前的负测试运行重置错误
	factoriesInitError = nil

	err := InitFactories(nil)
	assert.NoError(t, err)
}

func TestSetFactories(t *testing.T) {
	err := setFactories(nil)
	assert.NoError(t, err)

	err = setFactories(&FactoryOpts{})
	assert.NoError(t, err)
}

func TestSetFactoriesInvalidArgs(t *testing.T) {
	err := setFactories(&FactoryOpts{
		ProviderName: "SW",
		SwOpts:       &SwOpts{},
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed initializing SW.BCCSP")

	err = setFactories(&FactoryOpts{
		ProviderName: "PKCS11",
		Pkcs11Opts:   &pkcs11.PKCS11Opts{},
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Failed initializing PKCS11.BCCSP")
}

func TestGetBCCSPFromOpts(t *testing.T) {
	opts := GetDefaultOpts()
	opts.SwOpts.FileKeystore = &FileKeystoreOpts{KeyStorePath: os.TempDir()}
	opts.SwOpts.Ephemeral = false
	csp, err := GetBCCSPFromOpts(opts)
	assert.NoError(t, err)
	assert.NotNil(t, csp)

	lib, pin, label := pkcs11.FindPKCS11Lib()
	csp, err = GetBCCSPFromOpts(&FactoryOpts{
		ProviderName: "PKCS11",
		Pkcs11Opts: &pkcs11.PKCS11Opts{
			SecLevel:   256,
			HashFamily: "SHA2",
			Ephemeral:  true,
			Library:    lib,
			Pin:        pin,
			Label:      label,
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, csp)

	csp, err = GetBCCSPFromOpts(&FactoryOpts{
		ProviderName: "BadName",
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Could not find BCCSP, no 'BadName' provider")
	assert.Nil(t, csp)
}

