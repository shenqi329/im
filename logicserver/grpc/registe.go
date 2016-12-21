package grpc

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/service"
	"log"
)

type Registe struct{}

func (r *Registe) Registe(ctx context.Context, request *grpcPb.DeviceRegisteRequest) (*grpcPb.DeviceRegisteResponse, error) {

	message, err := HandleRegiste(ctx, request)
	response := message.(*grpcPb.DeviceRegisteResponse)

	return response, err
}

func HandleRegiste(ctx context.Context, message proto.Message) (proto.Message, error) {

	log.Println("Registe")
	request, ok := message.(*grpcPb.DeviceRegisteRequest)
	if !ok {
		return nil, nil
	}
	protoMessage, err := service.HandleRegiste(request)

	return protoMessage, err
}
