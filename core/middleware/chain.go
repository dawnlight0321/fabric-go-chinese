
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456036333391872>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

//链是用于HTTP请求处理的中间件链。
type Chain struct {
	mw []Middleware
}

//new chain创建了一个新的中间件链。链将调用中间件
//按照提供的顺序。
func NewChain(middlewares ...Middleware) Chain {
	return Chain{
		mw: append([]Middleware{}, middlewares...),
	}
}

//handler返回此链的http.handler。
func (c Chain) Handler(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}

	for i := len(c.mw) - 1; i >= 0; i-- {
		h = c.mw[i](h)
	}
	return h
}

