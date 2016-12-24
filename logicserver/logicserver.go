package main

import (
	logicserverGrpc "im/logicserver/grpc"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/server"
	logicserverTlp "im/logicserver/tlp"
	tlpPb "im/logicserver/tlp/pb"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	s := server.NEWServer()

	forward := &logicserverGrpc.Forward{}

	//登陆注册
	forward.AddHandleFunc(tlpPb.MessageTypeDeviceRegisteRequest, tlpPb.MessageTypeDeviceRegisteResponse, logicserverTlp.HandleRegiste)
	forward.AddHandleFunc(tlpPb.MessageTypeDeviceLoginRequest, tlpPb.MessageTypeDeviceLoginResponse, logicserverTlp.HandleLogin)
	forward.AddHandleFunc(tlpPb.MessageTypeCreateMessageRequest, tlpPb.MessageTypeCreateMessageResponse, logicserverTlp.CreateMessage)

	grpcPb.RegisterForwardToLogicServer(s.GrpcServer(), forward)

	grpcPb.RegisterOfflineServer(s.GrpcServer(), &logicserverGrpc.Offline{})
	grpcPb.RegisterSessionServer(s.GrpcServer(), &logicserverGrpc.Session{})

	s.Run("localhost:6005", "localhost:6004")
}
