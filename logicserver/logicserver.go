package main

import (
	logicserverGrpc "im/logicserver/grpc"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/server"
	"log"
	//"runtime"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	s := server.NEWServer()

	rpc := &logicserverGrpc.Rpc{}

	//登陆注册
	rpc.AddHandleFunc(grpcPb.MessageTypeDeviceRegisteRequest, grpcPb.MessageTypeDeviceRegisteResponse, logicserverGrpc.HandleRegiste)
	rpc.AddHandleFunc(grpcPb.MessageTypeDeviceLoginRequest, grpcPb.MessageTypeDeviceLoginResponse, logicserverGrpc.HandleLogin)
	rpc.AddHandleFunc(grpcPb.MessageTypeCreateMessageRequest, grpcPb.MessageTypeCreateMessageResponse, logicserverGrpc.CreateMessage)

	grpcPb.RegisterRpcServer(s.GrpcServer(), rpc)

	//grpcPb.RegisterRegisteServer(s.GrpcServer(), &logicserverGrpc.Registe{})
	//grpcPb.RegisterLoginServer(s.GrpcServer(), &logicserverGrpc.Login{})
	grpcPb.RegisterSessionServer(s.GrpcServer(), &logicserverGrpc.Session{})
	//grpcPb.RegisterMessageServer(s.GrpcServer(), &logicserverGrpc.Message{})

	s.Run("localhost:6005")
}
