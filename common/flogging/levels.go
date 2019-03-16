
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:55</date>
//</624455952883519488>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package flogging

import (
	"fmt"
	"math"

	"go.uber.org/zap/zapcore"
)

const (
//DisabledLevel表示已禁用的日志级别。此级别的日志应
//永远不要排放。
	DisabledLevel = zapcore.Level(math.MinInt8)

//payloadLevel用于记录非常详细的消息级别调试
//信息。
	PayloadLevel = zapcore.Level(zapcore.DebugLevel - 1)
)

//nametoLevel将级别名称转换为zapcore.level。如果级别名称为
//未知，返回zapcore.infoLevel。
func NameToLevel(level string) zapcore.Level {
	l, err := nameToLevel(level)
	if err != nil {
		return zapcore.InfoLevel
	}
	return l
}

func nameToLevel(level string) (zapcore.Level, error) {
	switch level {
	case "PAYLOAD", "payload":
		return PayloadLevel, nil
	case "DEBUG", "debug":
		return zapcore.DebugLevel, nil
	case "INFO", "info":
		return zapcore.InfoLevel, nil
	case "WARNING", "WARN", "warning", "warn":
		return zapcore.WarnLevel, nil
	case "ERROR", "error":
		return zapcore.ErrorLevel, nil
	case "DPANIC", "dpanic":
		return zapcore.DPanicLevel, nil
	case "PANIC", "panic":
		return zapcore.PanicLevel, nil
	case "FATAL", "fatal":
		return zapcore.FatalLevel, nil

	case "NOTICE", "notice":
return zapcore.InfoLevel, nil //未来
	case "CRITICAL", "critical":
return zapcore.ErrorLevel, nil //未来

	default:
		return DisabledLevel, fmt.Errorf("invalid log level: %s", level)
	}
}

func IsValidLevel(level string) bool {
	_, err := nameToLevel(level)
	return err == nil
}

