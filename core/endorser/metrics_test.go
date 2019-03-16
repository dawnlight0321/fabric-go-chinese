
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456008801980416>

/*
版权所有State Street Corp.保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package endorser

import (
	"testing"

	"github.com/hyperledger/fabric/common/metrics/metricsfakes"
	. "github.com/onsi/gomega"
)

func TestNewEndorserMetrics(t *testing.T) {
	gt := NewGomegaWithT(t)

	provider := &metricsfakes.Provider{}
	provider.NewHistogramReturns(&metricsfakes.Histogram{})
	provider.NewCounterReturns(&metricsfakes.Counter{})

	endorserMetrics := NewEndorserMetrics(provider)
	gt.Expect(endorserMetrics).To(Equal(&EndorserMetrics{
		ProposalDuration:         &metricsfakes.Histogram{},
		ProposalsReceived:        &metricsfakes.Counter{},
		SuccessfulProposals:      &metricsfakes.Counter{},
		ProposalValidationFailed: &metricsfakes.Counter{},
		ProposalACLCheckFailed:   &metricsfakes.Counter{},
		InitFailed:               &metricsfakes.Counter{},
		EndorsementsFailed:       &metricsfakes.Counter{},
		DuplicateTxsFailure:      &metricsfakes.Counter{},
	}))

	gt.Expect(provider.NewHistogramCallCount()).To(Equal(1))
	gt.Expect(provider.Invocations()["NewHistogram"]).To(ConsistOf([][]interface{}{
		{proposalDurationHistogramOpts},
	}))

	gt.Expect(provider.NewCounterCallCount()).To(Equal(7))
	gt.Expect(provider.Invocations()["NewCounter"]).To(ConsistOf([][]interface{}{
		{receivedProposalsCounterOpts},
		{successfulProposalsCounterOpts},
		{proposalValidationFailureCounterOpts},
		{proposalChannelACLFailureOpts},
		{initFailureCounterOpts},
		{endorsementFailureCounterOpts},
		{duplicateTxsFailureCounterOpts},
	}))
}

