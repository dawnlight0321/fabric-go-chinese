
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:55</date>
//</624455952166293504>

/*
版权所有IBM公司。保留所有权利。


*/


package flogging

import (
	"strings"

	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
)

const (
	defaultFormat = "%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}"
	defaultLevel  = zapcore.InfoLevel
)

var Global *Logging
var logger *FabricLogger

func init() {
	logging, err := New(Config{})
	if err != nil {
		panic(err)
	}

	Global = logging
	logger = Global.Logger("flogging")
	grpcLogger := Global.ZapLogger("grpc")
	grpclog.SetLogger(NewGRPCLogger(grpcLogger))
}

//
func Init(config Config) {
	err := Global.Apply(config)
	if err != nil {
		panic(err)
	}
}

//重置将日志记录设置为此包中定义的默认值。
//
//在测试和包初始化中使用
func Reset() {
	Global.Apply(Config{})
}

//GetLoggerLevel获取具有
//提供名称。
func GetLoggerLevel(loggerName string) string {
	return strings.ToUpper(Global.Level(loggerName).String())
}

//MustGetLogger创建具有指定名称的记录器。如果名称无效
//如有提供，操作会恐慌。
func MustGetLogger(loggerName string) *FabricLogger {
	return Global.Logger(loggerName)
}

//activatespec用于激活日志规范。
func ActivateSpec(spec string) {
	err := Global.ActivateSpec(spec)
	if err != nil {
		panic(err)
	}
}

