package grpc

import (
	"golang.org/x/net/context"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/service"
)

type Session struct{}

func (s *Session) CreateSession(ctx context.Context, request *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionResponse, error) {

	response, err := service.CreateSession(request)
	return response, err
}

func (s *Session) DeleteUsers(ctx context.Context, request *grpcPb.DeleteSessionUsersRequest) (*grpcPb.Response, error) {

	response, err := service.DeleteSessionUsers(request)
	return response, err
}

func (s *Session) AddUsers(ctx context.Context, request *grpcPb.AddSessionUsersRequest) (*grpcPb.Response, error) {

	response, err := service.AddSessionUsers(request)
	return response, err
}
