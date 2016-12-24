package tlp

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	imserverError "im/logicserver/error"
	"im/logicserver/server"
	"im/logicserver/service"
	tlpPb "im/logicserver/tlp/pb"
	"im/logicserver/util/key"
	"log"
)

func HandleLogin(ctx context.Context, message proto.Message) (proto.Message, error) {
	log.Println("Login")

	request := message.(*tlpPb.DeviceLoginRequest)
	serverContext := ctx.Value(key.Context).(*server.Context)

	protoMessage, err := service.HandleLogin(serverContext, request)

	if err != nil {
		log.Println(err.Error())
		reply := &tlpPb.DeviceLoginResponse{
			Rid:  request.GetRid(),
			Code: imserverError.CommonInternalServerError,
			Desc: imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
		}
		return reply, nil
	}
	return protoMessage, nil
}
