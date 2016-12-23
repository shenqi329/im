package server

import (
	"im/logicserver/bean"
	//"im/logicserver/dao"
	grpcPb "im/logicserver/grpc/pb"
	//"im/logicserver/service"
	"log"
)

type Context struct {
	Server *Server
}

func (c *Context) PushMessageToClient(message *bean.Message) error {
	log.Println(bean.StructToJsonString(message))

	// userInfo := c.Server.SafeGetLoginInfoWithUserId(message.UserId)

	// for value, count := range userInfo.Tokens {

	// }
	return nil
}

func (c *Context) HandleOffline(request *grpcPb.DeviceOfflineRequest) (err error) {
	log.Println(bean.StructToJsonString(request))

	c.Server.SafeRemoveLoginInfo(request.Token, request.UserId)

	log.Println(bean.StructToJsonString(c.Server.SafeGetLoginInfoWithUserId(request.UserId)))
	log.Println(bean.StructToJsonString(c.Server.SafeGetLoginInfoWithToken(request.Token)))

	return nil
}

func (c *Context) HandleLogin(request *grpcPb.DeviceLoginRequest) (err error) {
	log.Println(bean.StructToJsonString(request))
	c.Server.SafeAddLoginInfo(request.Token, request.UserId)

	log.Println(bean.StructToJsonString(c.Server.SafeGetLoginInfoWithUserId(request.UserId)))
	log.Println(bean.StructToJsonString(c.Server.SafeGetLoginInfoWithToken(request.Token)))

	return nil
}
