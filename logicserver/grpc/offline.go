package grpc

import (
	"golang.org/x/net/context"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/server"
	"im/logicserver/service"
	"im/logicserver/util/key"
)

type Offline struct{}

func (m *Offline) Offline(ctx context.Context, request *grpcPb.DeviceOfflineRequest) (*grpcPb.DeviceOfflineResponse, error) {

	serverContext := ctx.Value(key.Context).(*server.Context)

	response, err := service.HandleOffline(serverContext, request)

	return response, err
}
