package grpc

import (
	//"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	accessserverError "im/accessserver/error"
	accessserverGrpcPb "im/accessserver/grpc/pb"
	"im/accessserver/server"
	"im/accessserver/util/key"
	"log"
)

type Forward struct {
}

func (r *Forward) ForwardTLP(ctx context.Context, request *accessserverGrpcPb.ForwardTLPRequest) (*accessserverGrpcPb.ForwardTLPResponse, error) {

	rpcResponse := &accessserverGrpcPb.ForwardTLPResponse{
		Code: accessserverError.CommonInternalServerError,
		Desc: accessserverError.ErrorCodeToText(accessserverError.CommonInternalServerError),
	}
	s := ctx.Value(key.Server).(*server.Server)
	log.Println(request.String())

	s.ForwardTLP(request)

	rpcResponse = &accessserverGrpcPb.ForwardTLPResponse{
		Code: accessserverError.CommonSuccess,
		Desc: accessserverError.ErrorCodeToText(accessserverError.CommonSuccess),
	}

	return rpcResponse, nil
}
