
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:55</date>
//</624455953416196096>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package metrics

import (
	"github.com/hyperledger/fabric/common/metrics"
	"go.uber.org/zap/zapcore"
)

var (
	CheckedCountOpts = metrics.CounterOpts{
		Namespace:    "logging",
		Name:         "entries_checked",
		Help:         "Number of log entries checked against the active logging level",
		LabelNames:   []string{"level"},
		StatsdFormat: "%{#fqname}.%{level}",
	}

	WriteCountOpts = metrics.CounterOpts{
		Namespace:    "logging",
		Name:         "entries_written",
		Help:         "Number of log entries that are written",
		LabelNames:   []string{"level"},
		StatsdFormat: "%{#fqname}.%{level}",
	}
)

type Observer struct {
	CheckedCounter metrics.Counter
	WrittenCounter metrics.Counter
}

func NewObserver(provider metrics.Provider) *Observer {
	return &Observer{
		CheckedCounter: provider.NewCounter(CheckedCountOpts),
		WrittenCounter: provider.NewCounter(WriteCountOpts),
	}
}

func (m *Observer) Check(e zapcore.Entry, ce *zapcore.CheckedEntry) {
	m.CheckedCounter.With("level", e.Level.String()).Add(1)
}

func (m *Observer) WriteEntry(e zapcore.Entry, fields []zapcore.Field) {
	m.WrittenCounter.With("level", e.Level.String()).Add(1)
}

