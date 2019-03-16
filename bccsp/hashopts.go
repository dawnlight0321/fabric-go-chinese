
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:49</date>
//</624455928019685376>

/*
版权所有IBM Corp.2016保留所有权利。

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


package bccsp

import "fmt"

//sha256opts包含与sha-256相关的选项。
type SHA256Opts struct {
}

//Algorithm返回哈希算法标识符（要使用）。
func (opts *SHA256Opts) Algorithm() string {
	return SHA256
}

//sha 384 opts包含与sha-384相关的选项。
type SHA384Opts struct {
}

//Algorithm返回哈希算法标识符（要使用）。
func (opts *SHA384Opts) Algorithm() string {
	return SHA384
}

//sha3_256opts包含与sha3-256相关的选项。
type SHA3_256Opts struct {
}

//Algorithm返回哈希算法标识符（要使用）。
func (opts *SHA3_256Opts) Algorithm() string {
	return SHA3_256
}

//sha3_384opts包含与sha3-384相关的选项。
type SHA3_384Opts struct {
}

//Algorithm返回哈希算法标识符（要使用）。
func (opts *SHA3_384Opts) Algorithm() string {
	return SHA3_384
}

//GetHashOpt返回与传递的哈希函数对应的哈希值
func GetHashOpt(hashFunction string) (HashOpts, error) {
	switch hashFunction {
	case SHA256:
		return &SHA256Opts{}, nil
	case SHA384:
		return &SHA384Opts{}, nil
	case SHA3_256:
		return &SHA3_256Opts{}, nil
	case SHA3_384:
		return &SHA3_384Opts{}, nil
	}
	return nil, fmt.Errorf("hash function not recognized [%s]", hashFunction)
}

