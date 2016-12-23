package server

import (
	netContext "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	logicserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/util/key"
	"log"
	"net"
	"reflect"
)

type Server struct {
	grpcServer       *grpc.Server
	loginInfoManager *LoginInfoManager
}

func NEWServer() *Server {
	s := &Server{
		loginInfoManager: NEWLoginInfoManager(),
	}
	return s
}

func (s *Server) SafeGetLoginInfoWithToken(token string) *LoginInfo {
	return s.loginInfoManager.SafeGetLoginInfoWithToken(token)
}

func (s *Server) SafeAddLoginInfo(token string, userId string) bool {

	return s.loginInfoManager.SafeAddLoginInfo(token, userId)
}

func (s *Server) SafeRemoveLoginInfo(token string, userId string) bool {
	return s.SafeRemoveLoginInfo(token, userId)
}

func (s *Server) SafeGetLoginInfoWithUserId(userId string) *LoginInfo {
	return s.SafeGetLoginInfoWithUserId(userId)
}

func (s *Server) GrpcServer() *grpc.Server {
	if s.grpcServer == nil {
		s.grpcServer = s.newGrpcServer()
	}
	return s.grpcServer
}

func (s *Server) Run(grpcTcpPort string) {
	s.grpcServerServe(grpcTcpPort)
}

func (s *Server) newGrpcServer() *grpc.Server {
	type Request interface {
		GetRid() uint64
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(func(ctx netContext.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		serverContext := &Context{
			Server: s,
		}
		ctx = netContext.WithValue(ctx, key.Context, serverContext)
		response, err := handler(ctx, req)
		v := reflect.ValueOf(response)
		if !v.IsValid() || v.IsNil() {
			request, ok := req.(Request)
			if ok {
				response = &grpcPb.Response{
					Rid:  request.GetRid(),
					Code: logicserverError.CommonInternalServerError,
					Desc: logicserverError.ErrorCodeToText(logicserverError.CommonInternalServerError),
				}
			} else {
				response = &grpcPb.Response{
					Code: logicserverError.CommonInternalServerError,
					Desc: logicserverError.ErrorCodeToText(logicserverError.CommonInternalServerError),
				}
			}
		}
		return response, nil
	}))
	return grpcServer
}

func (s *Server) grpcServerServe(addr string) {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("grpcServerServe addr:", addr)

	reflection.Register(s.grpcServer)
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
