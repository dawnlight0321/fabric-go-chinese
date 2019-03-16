
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456036677324800>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package middleware_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/hyperledger/fabric/core/middleware"
	"github.com/hyperledger/fabric/core/middleware/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("WithRequestID", func() {
	var (
		requestID middleware.Middleware
		handler   *fakes.HTTPHandler
		chain     http.Handler

		req  *http.Request
		resp *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		handler = &fakes.HTTPHandler{}
		requestID = middleware.WithRequestID(
			middleware.GenerateIDFunc(func() string { return "generated-id" }),
		)
		chain = requestID(handler)

		req = httptest.NewRequest("GET", "/", nil)
		resp = httptest.NewRecorder()
	})

	It("propagates the generated request ID in the request context", func() {
		chain.ServeHTTP(resp, req)
		_, r := handler.ServeHTTPArgsForCall(0)
		requestID := middleware.RequestID(r.Context())
		Expect(requestID).To(Equal("generated-id"))
	})

	It("returns the generated request ID in a header", func() {
		chain.ServeHTTP(resp, req)
		Expect(resp.Header().Get("X-Request-Id")).To(Equal("generated-id"))
	})

	Context("when a request ID is already present", func() {
		BeforeEach(func() {
			req.Header.Set("X-Request-Id", "received-id")
		})

		It("sets the received ID into the context", func() {
			chain.ServeHTTP(resp, req)
			_, r := handler.ServeHTTPArgsForCall(0)
			requestID := middleware.RequestID(r.Context())
			Expect(requestID).To(Equal("received-id"))
		})

		It("sets the received ID into the request", func() {
			chain.ServeHTTP(resp, req)
			_, r := handler.ServeHTTPArgsForCall(0)
			Expect(r.Header.Get("X-Request-Id")).To(Equal("received-id"))
		})

		It("propagates the request ID to the response", func() {
			chain.ServeHTTP(resp, req)
			Expect(resp.Header().Get("X-Request-Id")).To(Equal("received-id"))
		})
	})
})

