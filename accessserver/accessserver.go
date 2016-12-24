package main

import (
	accessserverGrpc "im/accessserver/grpc"
	accessserverGrpcPb "im/accessserver/grpc/pb"
	accessserver "im/accessserver/server"
	"log"
	"runtime"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if runtime.GOOS == "windows" {
		localTcpAddr := "localhost:6000"
		proxyUdpAddr := "localhost:6001"
		s := accessserver.NEWServer(localTcpAddr, proxyUdpAddr)

		forward := &accessserverGrpc.Forward{}
		accessserverGrpcPb.RegisterForwardToAccessServer(s.GrpcServer(), forward)

		s.Run("localhost:6004")
	} else {
		localTcpAddr := "localhost:6000"
		proxyUdpAddr := "localhost:6001"
		s := accessserver.NEWServer(localTcpAddr, proxyUdpAddr)
		s.Run("localhost:6004")

		forward := &accessserverGrpc.Forward{}
		accessserverGrpcPb.RegisterForwardToAccessServer(s.GrpcServer(), forward)
	}
}
