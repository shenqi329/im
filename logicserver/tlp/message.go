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

func CreateMessage(context context.Context, message proto.Message) (proto.Message, error) {

	log.Println("CreateMessage")

	request := message.(*tlpPb.CreateMessageRequest)
	userId := context.Value(key.UserId).(string)
	serverContext := context.Value(key.Context).(*server.Context)

	protoMessage, err := service.HandleCreateMessage(serverContext, request, userId)

	if err != nil {
		log.Println(err.Error())
		reply := &tlpPb.CreateMessageResponse{
			Rid:  request.GetRid(),
			Code: imserverError.CommonInternalServerError,
			Desc: imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
		}
		return reply, nil
	}
	return protoMessage, nil
}
