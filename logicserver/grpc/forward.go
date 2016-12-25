package grpc

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	logicserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	tlpPb "im/logicserver/tlp/pb"
	"im/logicserver/util/key"
	"log"
)

//protoc --go_out=plugins=grpc:. *.proto

type HandleFunc func(ctx context.Context, request proto.Message) (proto.Message, error)

type HandleFuncInfo struct {
	handle       HandleFunc
	responseType tlpPb.MessageType
}

type Forward struct {
	handleFuncMap map[tlpPb.MessageType]*HandleFuncInfo
}

func (r *Forward) AddHandleFunc(messageType tlpPb.MessageType, responseType tlpPb.MessageType, handle HandleFunc) {
	if r.handleFuncMap == nil {
		r.handleFuncMap = make(map[tlpPb.MessageType]*HandleFuncInfo)
	}
	r.handleFuncMap[messageType] = &HandleFuncInfo{
		handle:       handle,
		responseType: responseType,
	}
}

func (r *Forward) ForwardTLP(ctx context.Context, request *grpcPb.ForwardTLPRequest) (*grpcPb.ForwardTLPResponse, error) {

	rpcResponse := &grpcPb.ForwardTLPResponse{
		Code:   logicserverError.CommonInternalServerError,
		Desc:   logicserverError.ErrorCodeToText(logicserverError.CommonInternalServerError),
		ConnId: request.RpcInfo.ConnId,
	}

	log.Println(request.String())
	ctx = context.WithValue(ctx, key.UserId, request.RpcInfo.UserId)
	ctx = context.WithValue(ctx, key.Token, request.RpcInfo.Token)
	ctx = context.WithValue(ctx, key.ConnId, request.RpcInfo.ConnId)

	handleFuncInfo := r.handleFuncMap[(tlpPb.MessageType)(request.MessageType)]
	if handleFuncInfo == nil {
		log.Println("不支持的类型")
		return rpcResponse, nil
	}

	protoMessage := tlpPb.Factory((tlpPb.MessageType)(request.MessageType))
	err := proto.Unmarshal(request.ProtoBuf, protoMessage)
	if err != nil {
		log.Println(err.Error())
		return rpcResponse, nil
	}

	response, err := handleFuncInfo.handle(ctx, protoMessage)
	if err != nil {
		log.Println(err.Error())
		return rpcResponse, nil
	}
	if response == nil {
		rpcResponse = &grpcPb.ForwardTLPResponse{
			Code: logicserverError.CommonSuccess,
			Desc: logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
		}
		return rpcResponse, nil
	}

	protoBuf, err := proto.Marshal(response)
	if err != nil {
		log.Println(err.Error())
		return rpcResponse, nil
	}

	rpcResponse = &grpcPb.ForwardTLPResponse{
		Code:        logicserverError.CommonSuccess,
		Desc:        logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
		MessageType: (uint32)(handleFuncInfo.responseType),
		ProtoBuf:    protoBuf,
		ConnId:      request.RpcInfo.ConnId,
	}

	return rpcResponse, nil
}
