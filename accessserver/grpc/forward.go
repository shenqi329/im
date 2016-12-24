package grpc

import (
	//"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	accessserverError "im/accessserver/error"
	accessserverGrpcPb "im/accessserver/grpc/pb"
	"log"
)

type Forward struct {
}

func (r *Forward) ForwardTLP(ctx context.Context, request *accessserverGrpcPb.ForwardTLPRequest) (*accessserverGrpcPb.ForwardTLPResponse, error) {

	rpcResponse := &accessserverGrpcPb.ForwardTLPResponse{
		Code: accessserverError.CommonInternalServerError,
		Desc: accessserverError.ErrorCodeToText(accessserverError.CommonInternalServerError),
	}

	log.Println(request.String())

	rpcResponse = &accessserverGrpcPb.ForwardTLPResponse{
		Code: accessserverError.CommonSuccess,
		Desc: accessserverError.ErrorCodeToText(accessserverError.CommonSuccess),
	}

	return rpcResponse, nil
}
