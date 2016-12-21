package server

import (
	"github.com/golang/protobuf/proto"
	grpcPb "im/logicserver/grpc/pb"
	"log"
)

func Handle(context Context) error {
	//log.Println("handle")
	protoMessage := grpcPb.Factory((grpcPb.MessageType)(context.Message().Type))

	if protoMessage == nil {
		log.Println("未识别的消息")
		context.Close()
		return nil
	}
	if err := proto.Unmarshal(context.Message().Body, protoMessage); err != nil {
		log.Println(err.Error())
		context.Close()
		return nil
	}

	//只检查消息的合法性,然后将消息转发出去
	context.Request().message = context.Message()
	context.Request().protoMessage = protoMessage
	context.Request().messageType = (grpcPb.MessageType)(context.Message().Type)

	context.ReqChan() <- context.Request()

	return nil
}
