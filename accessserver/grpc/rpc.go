package grpc

import (
	//"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	accessserverError "im/accessserver/error"
	accessserverGrpcPb "im/accessserver/grpc/pb"
	"log"
)

type Rpc struct {
}

func (r *Rpc) Rpc(ctx context.Context, request *accessserverGrpcPb.RpcRequest) (*accessserverGrpcPb.RpcResponse, error) {

	rpcResponse := &accessserverGrpcPb.RpcResponse{
		Code: accessserverError.CommonInternalServerError,
		Desc: accessserverError.ErrorCodeToText(accessserverError.CommonInternalServerError),
	}

	log.Println(request.String())

	rpcResponse = &accessserverGrpcPb.RpcResponse{
		Code: accessserverError.CommonSuccess,
		Desc: accessserverError.ErrorCodeToText(accessserverError.CommonSuccess),
	}

	return rpcResponse, nil
}
