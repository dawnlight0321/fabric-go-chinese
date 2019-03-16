
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:02</date>
//</624455979261497344>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package chaincode

import (
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/common/flogging"
	logging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

const (
	defaultExecutionTimeout = 30 * time.Second
	minimumStartupTimeout   = 5 * time.Second
)

type Config struct {
	TLSEnabled     bool
	Keepalive      time.Duration
	ExecuteTimeout time.Duration
	StartupTimeout time.Duration
	LogFormat      string
	LogLevel       string
	ShimLogLevel   string
}

func GlobalConfig() *Config {
	c := &Config{}
	c.load()
	return c
}

func (c *Config) load() {
	viper.SetEnvPrefix("CORE")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	c.TLSEnabled = viper.GetBool("peer.tls.enabled")

	c.Keepalive = toSeconds(viper.GetString("chaincode.keepalive"), 0)
	c.ExecuteTimeout = viper.GetDuration("chaincode.executetimeout")
	if c.ExecuteTimeout < time.Second {
		c.ExecuteTimeout = defaultExecutionTimeout
	}
	c.StartupTimeout = viper.GetDuration("chaincode.startuptimeout")
	if c.StartupTimeout < minimumStartupTimeout {
		c.StartupTimeout = minimumStartupTimeout
	}

	c.LogFormat = viper.GetString("chaincode.logging.format")
	c.LogLevel = getLogLevelFromViper("chaincode.logging.level")
	c.ShimLogLevel = getLogLevelFromViper("chaincode.logging.shim")
}

func toSeconds(s string, def int) time.Duration {
	seconds, err := strconv.Atoi(s)
	if err != nil {
		return time.Duration(def) * time.Second
	}

	return time.Duration(seconds) * time.Second
}

//GetLogLevelFromViper从Viper获取链码容器日志级别
func getLogLevelFromViper(key string) string {
	levelString := viper.GetString(key)
	_, err := logging.LogLevel(levelString)
	if err != nil {
		chaincodeLogger.Warningf("%s has invalid log level %s. defaulting to %s", key, levelString, flogging.DefaultLevel())
		levelString = flogging.DefaultLevel()
	}

	return levelString
}

//devmodeuserrunschaincode支持在开发中执行链码
//环境
const DevModeUserRunsChaincode string = "dev"

//如果对等机配置了开发模式，则isdevmode返回true。
//启用。
func IsDevMode() bool {
	mode := viper.GetString("chaincode.mode")

	return mode == DevModeUserRunsChaincode
}

