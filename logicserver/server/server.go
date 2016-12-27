package server

import (
	netContext "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"im/logicserver/bean"
	logicserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	tlpPb "im/logicserver/tlp/pb"
	"im/logicserver/util/key"
	"log"
	"net"
	"reflect"
)

type Server struct {
	grpcServer       *grpc.Server
	loginInfoManager *LoginInfoManager
	accessClient     *AccessClient
}

func NEWServer() *Server {
	s := &Server{
		loginInfoManager: NEWLoginInfoManager(),
		accessClient:     NEWAccessClient(),
	}
	return s
}

func (s *Server) SendMessageToUser(message *bean.Message) {
	userInfo := s.loginInfoManager.SafeGetLoginInfoWithUserId(message.UserId)
	if userInfo == nil {
		return
	}
	for _, token := range userInfo.Tokens {
		s.accessClient.SendMessage(message, token)
	}
}

func (s *Server) SendSyncMessageToUser(userId string) {
	userInfo := s.loginInfoManager.SafeGetLoginInfoWithUserId(userId)
	log.Println(bean.StructToJsonString(userInfo))
	if userInfo == nil {
		return
	}
	for _, token := range userInfo.Tokens {
		//s.accessClient.SendMessage(message, token)
		s.accessClient.SendSyncMessageToUser(userId, token)
	}
}
func (s *Server) HandleOffline(request *grpcPb.DeviceOfflineRequest) (err error) {

	s.loginInfoManager.SafeRemoveLoginInfo(request.Token, request.UserId)

	//log.Println(bean.StructToJsonString(s.loginInfoManager.SafeGetLoginInfoWithUserId(request.UserId)))
	//log.Println(bean.StructToJsonString(s.loginInfoManager.SafeGetLoginInfoWithToken(request.Token)))
	return nil
}

func (s *Server) HandleLogin(request *tlpPb.DeviceLoginRequest) (err error) {

	s.loginInfoManager.SafeAddLoginInfo(request.Token, request.UserId)
	log.Println(bean.StructToJsonString(s.loginInfoManager.SafeGetLoginInfoWithUserId(request.UserId)))
	log.Println(bean.StructToJsonString(s.loginInfoManager.SafeGetLoginInfoWithToken(request.Token)))
	s.SendSyncMessageToUser(request.UserId)
	return nil
}

func (s *Server) GrpcServer() *grpc.Server {
	if s.grpcServer == nil {
		s.grpcServer = s.newGrpcServer()
	}
	return s.grpcServer
}

func (s *Server) Run(grpcTcpPort string, accessAddr string) {
	s.accessClient.Connect(accessAddr)
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
