package grpc

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	imserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/service"
	//"im/logicserver/util/key"
	"log"
)

type Login struct{}

func (m *Login) Login(ctx context.Context, request *grpcPb.DeviceLoginRequest) (*grpcPb.DeviceLoginResponse, error) {

	message, err := HandleLogin(ctx, request)

	response := message.(*grpcPb.DeviceLoginResponse)

	return response, err
}

func HandleLogin(ctx context.Context, message proto.Message) (proto.Message, error) {
	log.Println("Login")

	request := message.(*grpcPb.DeviceLoginRequest)

	protoMessage, err := service.HandleLogin(request)

	if err != nil {
		log.Println(err.Error())
		reply := &grpcPb.DeviceLoginResponse{
			Rid:  request.GetRid(),
			Code: imserverError.CommonInternalServerError,
			Desc: imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
		}
		return reply, nil
	}
	return protoMessage, nil
}
