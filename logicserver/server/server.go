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
	"strings"
	"sync"
)

type Request interface {
	GetRid() uint64
}

type Server struct {
	grpcServer              *grpc.Server
	tokenLoginInfoMap       map[string]*LoginInfo
	tokenLoginInfoMapMutex  sync.Mutex
	userIdLoginInfoMap      map[string]*LoginInfo
	userIdLoginInfoMapMutex sync.Mutex
}

func NEWServer() *Server {
	s := &Server{
		tokenLoginInfoMap:  make(map[string]*LoginInfo),
		userIdLoginInfoMap: make(map[string]*LoginInfo),
	}
	return s
}

func (s *Server) SafeGetLoginInfoWithToken(token string) *LoginInfo {

	s.tokenLoginInfoMapMutex.Lock()
	loginInfo := s.tokenLoginInfoMap[token]
	s.tokenLoginInfoMapMutex.Unlock()

	return loginInfo
}

func (s *Server) SafeAddLoginInfo(token string, userId string) bool {

	s.tokenLoginInfoMapMutex.Lock()
	loginInfo := s.tokenLoginInfoMap[token]
	if loginInfo == nil {
		s.tokenLoginInfoMap[token] = &LoginInfo{
			UserId: userId,
			Tokens: []string{token},
		}
	} else {
		s.tokenLoginInfoMap[token] = &LoginInfo{
			UserId: userId,
			Tokens: loginInfo.Tokens,
		}
	}
	s.tokenLoginInfoMapMutex.Unlock()

	s.userIdLoginInfoMapMutex.Lock()
	loginInfo = s.userIdLoginInfoMap[userId]
	if loginInfo == nil {
		s.userIdLoginInfoMap[token] = &LoginInfo{
			UserId: userId,
			Tokens: []string{token},
		}
	} else {
		temp := []string{token}
		for _, val := range loginInfo.Tokens {
			if !strings.EqualFold(val, token) {
				temp = append(temp, val)
			}
		}
		loginInfo.Tokens = temp
	}
	s.userIdLoginInfoMapMutex.Unlock()
	return true
}

func (s *Server) SafeRemoveLoginInfo(token string, userId string) bool {

	s.tokenLoginInfoMapMutex.Lock()
	delete(s.tokenLoginInfoMap, token)
	s.tokenLoginInfoMapMutex.Unlock()

	s.userIdLoginInfoMapMutex.Lock()

	loginInfo := s.userIdLoginInfoMap[userId]
	if loginInfo != nil {
		temp := []string{}
		for _, val := range loginInfo.Tokens {
			if !strings.EqualFold(val, token) {
				temp = append(temp, val)
			}
		}
		loginInfo.Tokens = temp
		if len(loginInfo.Tokens) == 0 {
			delete(s.userIdLoginInfoMap, userId)
		}
	}
	s.userIdLoginInfoMapMutex.Unlock()

	return true
}

func (s *Server) SafeGetLoginInfoWithUserId(userId string) *LoginInfo {
	s.userIdLoginInfoMapMutex.Lock()
	loginInfo := s.userIdLoginInfoMap[userId]
	s.userIdLoginInfoMapMutex.Unlock()
	return loginInfo
}

func (s *Server) GrpcServer() *grpc.Server {
	if s.grpcServer == nil {
		s.grpcServer = s.newServer()
	}
	return s.grpcServer
}

func (s *Server) Run(grpcTcpPort string) {
	s.grpcServerServe(grpcTcpPort)
}

func (s *Server) newServer() *grpc.Server {

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
