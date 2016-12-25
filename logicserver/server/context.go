package server

import (
	"im/logicserver/bean"
	grpcPb "im/logicserver/grpc/pb"
	tlpPb "im/logicserver/tlp/pb"
	//"log"
)

type Context struct {
	Server *Server
}

func (c *Context) PushMessageToClient(message *bean.Message) error {
	c.Server.SendMessageToUser(message)
	return nil
}

func (c *Context) HandleOffline(request *grpcPb.DeviceOfflineRequest) (err error) {
	return c.Server.HandleOffline(request)
}

func (c *Context) HandleLogin(request *tlpPb.DeviceLoginRequest) (err error) {
	return c.Server.HandleLogin(request)
}
