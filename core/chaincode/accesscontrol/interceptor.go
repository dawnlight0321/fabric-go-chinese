
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455978049343488>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package accesscontrol

import (
	"fmt"

	pb "github.com/hyperledger/fabric/protos/peer"
	"google.golang.org/grpc"
)

type interceptor struct {
	next pb.ChaincodeSupportServer
	auth authorization
}

//chaincodestream定义用于发送的grpc流
//接收链码信息
type ChaincodeStream interface {
//发送发送一条链码消息
	Send(*pb.ChaincodeMessage) error
//接收一条链码消息
	Recv() (*pb.ChaincodeMessage, error)
}

type authorization func(message *pb.ChaincodeMessage, stream grpc.ServerStream) error

func newInterceptor(srv pb.ChaincodeSupportServer, auth authorization) pb.ChaincodeSupportServer {
	return &interceptor{
		next: srv,
		auth: auth,
	}
}

//
func (i *interceptor) Register(stream pb.ChaincodeSupport_RegisterServer) error {
	is := &interceptedStream{
		incMessages:  make(chan *pb.ChaincodeMessage, 1),
		stream:       stream,
		ServerStream: stream,
		auth:         i.auth,
	}
	msg, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("Recv() error: %v, closing connection", err)
	}
	err = is.auth(msg, is.ServerStream)
	if err != nil {
		return err
	}
	is.incMessages <- msg
	close(is.incMessages)
	return i.next.Register(is)
}

type interceptedStream struct {
	incMessages chan *pb.ChaincodeMessage
	stream      ChaincodeStream
	grpc.ServerStream
	auth authorization
}

//发送发送一条链码消息
func (is *interceptedStream) Send(msg *pb.ChaincodeMessage) error {
	return is.stream.Send(msg)
}

//接收一条链码消息
func (is *interceptedStream) Recv() (*pb.ChaincodeMessage, error) {
	msg, ok := <-is.incMessages
	if !ok {
		return is.stream.Recv()
	}
	return msg, nil
}

