package server

import (
	"github.com/golang/protobuf/proto"
	netContext "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	accessError "im/accessserver/error"
	accessserverGrpc "im/accessserver/grpc"
	accessserverGrpcPb "im/accessserver/grpc/pb"
	grpcPb "im/logicserver/grpc/pb"
	coder "im/protocol/coder"
	"log"
	"net"
	"os"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

type Request struct {
	message      *coder.Message
	protoMessage proto.Message
	messageType  grpcPb.MessageType
	connId       uint64
	conn         *net.TCPConn
}

type ProtocolResp struct {
	protocolBuf []byte
	connId      uint64
	isLogin     bool
	isLogout    bool
}

type ConnectInfo struct {
	conn    *net.TCPConn
	isLogin bool
	token   string
	userId  string
	appId   string
}

type Server struct {
	rid         uint64 //请求流水号
	ridMutex    sync.Mutex
	connCount   int32
	connIdMutex sync.Mutex
	connId      uint64 //请求的id

	localTcpAddr string
	proxyUdpAddr string

	grpcEasynoteClientConn *grpc.ClientConn
	grpcLogicClientConn    *grpc.ClientConn
	grpcServer             *grpc.Server

	rpcRespChan      chan *grpcPb.RpcResponse
	protocolRespChan chan *ProtocolResp
	rpcReqChan       chan *accessserverGrpcPb.RpcRequest

	handle func(context Context) error
}

func (s *Server) createRID() uint64 {
	s.ridMutex.Lock()
	s.rid++
	s.ridMutex.Unlock()
	return s.rid
}

func (s *Server) createConnId() uint64 {
	s.connIdMutex.Lock()
	s.connId++
	s.connIdMutex.Unlock()
	return s.connId
}

func NEWServer(localTcpAddr string, proxyUdpAddr string) (s *Server) {

	return &Server{
		localTcpAddr:     localTcpAddr,
		proxyUdpAddr:     proxyUdpAddr,
		rpcRespChan:      make(chan *grpcPb.RpcResponse, 1000),
		protocolRespChan: make(chan *ProtocolResp, 1000),
		rpcReqChan:       make(chan *accessserverGrpcPb.RpcRequest, 1000),
	}
}

func (s *Server) GrpcServer() *grpc.Server {
	if s.grpcServer == nil {
		s.grpcServer = s.newGrpcServer()
	}
	return s.grpcServer
}

func (s *Server) Run(grpcServerAddr string) {
	if s.handle == nil {
		s.handle = Handle
	}

	//grpcPb.RegisterRpcServer(s.GrpcServer(), &serverGrpc.Rpc{})

	go s.grpcServerServe(grpcServerAddr)
	s.grpcEasynoteClientConn = s.grpcConnectServer("localhost:6006")
	s.grpcLogicClientConn = s.grpcConnectServer("localhost:6005")
	s.ListenOnTcpPort(s.localTcpAddr)
}

func (s *Server) grpcConnectServer(tcpAddr string) *grpc.ClientConn {
	//Set up a connection to the server.
	conn, err := grpc.Dial(tcpAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func (s *Server) grpcServerServe(addr string) {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpc := &accessserverGrpc.Rpc{}

	accessserverGrpcPb.RegisterRpcServer(s.GrpcServer(), rpc)

	reflection.Register(s.GrpcServer())
	if err := s.GrpcServer().Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) newGrpcServer() *grpc.Server {

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(func(ctx netContext.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = netContext.WithValue(ctx, "", s.rpcRespChan)
		response, err := handler(ctx, req)
		v := reflect.ValueOf(response)
		if !v.IsValid() || v.IsNil() {
			request, ok := req.(accessserverGrpc.Request)
			if ok {
				response = &grpcPb.Response{
					Rid:  request.GetRid(),
					Code: accessError.CommonInternalServerError,
					Desc: accessError.ErrorCodeToText(accessError.CommonInternalServerError),
				}
			} else {
				response = &grpcPb.Response{
					Code: accessError.CommonInternalServerError,
					Desc: accessError.ErrorCodeToText(accessError.CommonInternalServerError),
				}
			}
		}
		return response, nil
	}))
	return grpcServer
}

func (s *Server) ListenOnTcpPort(localTcpAddr string) {

	addr, err := net.ResolveTCPAddr("tcp", localTcpAddr)

	if err != nil {
		log.Println("net.ResolveTCPAddr fail.", err)
		os.Exit(1)
	}

	listen, err := net.ListenTCP("tcp", addr)
	defer listen.Close()

	if err != nil {
		log.Println("net.ListenTCP fail.", err)
		os.Exit(1)
	}
	log.Println("net.ListenTCP", addr)

	//
	reqChan := make(chan *Request, 1000)
	closeChan := make(chan uint64, 1000)

	go s.connectIMServer(reqChan, closeChan)

	//
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Println("accept tcp fail", err.Error())
			continue
		}

		go s.handleTcpConnection(conn, reqChan, closeChan)
	}
}

func (s *Server) transToLogicServer(rpcRequest *grpcPb.RpcRequest, protocolRespChan chan<- *ProtocolResp) {
	//log.Println("转发给逻辑服务器")

	rpcClient := grpcPb.NewRpcClient(s.grpcLogicClientConn)
	response, err := rpcClient.Rpc(netContext.Background(), rpcRequest)

	if err != nil {
		s.sendErrorProtocolToChan(accessError.CommonInternalServerError, accessError.ErrorCodeToText(accessError.CommonInternalServerError), rpcRequest, protocolRespChan)
		return
	}

	if !accessError.CodeIsSuccess(response.Code) {
		s.sendErrorProtocolToChan(response.Code, response.Desc, rpcRequest, protocolRespChan)
		return
	}

	var isLogin bool = false
	var isLogout bool = false
	if rpcRequest.MessageType == grpcPb.MessageTypeDeviceLoginRequest {
		isLogin = true
	}
	// if rpcRequest.MessageType == grpcPb.MessageTypeDeviceLoginRequest {

	// }

	protocolBuf, _ := coder.EncoderMessage((int)(response.MessageType), response.ProtoBuf)

	protocolRespChan <- &ProtocolResp{
		protocolBuf: protocolBuf,
		connId:      rpcRequest.RpcInfo.ConnId,
		isLogin:     isLogin,
		isLogout:    isLogout,
	}
}

func (s *Server) sendErrorProtocolToChan(code string, desc string, rpcRequest *grpcPb.RpcRequest, protocolRespChan chan<- *ProtocolResp) {
	response := &grpcPb.Response{
		Rid:  rpcRequest.Rid,
		Code: code,
		Desc: desc,
	}
	protoBuf, _ := proto.Marshal(response)
	protocolBuf, _ := coder.EncoderMessage((int)(rpcRequest.MessageType+1), protoBuf)

	protocolRespChan <- &ProtocolResp{
		protocolBuf: protocolBuf,
		connId:      rpcRequest.RpcInfo.ConnId,
	}
}

func (s *Server) transToBusinessServer(rpcRequest *grpcPb.RpcRequest, rpcRespChan chan<- *grpcPb.RpcResponse) {
	//easynote业务id
	if rpcRequest.RpcInfo.AppId == "89897" {
		//log.Println("转发给业务服务器")
		rpcClient := grpcPb.NewRpcClient(s.grpcEasynoteClientConn)
		response, err := rpcClient.Rpc(netContext.Background(), rpcRequest)
		if err != nil {
			//直接返回错误给调用者[刘俊仕]
			log.Println(err.Error())
			return
		}
		if response != nil {
			log.Println(response.String())
		}
		rpcRespChan <- response
	}
}

//连接到逻辑服务器
func (s *Server) connectIMServer(reqChan <-chan *Request, closeChan <-chan uint64) {

	//rpcRespChan := make(chan *grpcPb.RpcResponse, 1000)
	//protocolRespChan := make(chan *ProtocolResp, 1000)
	connMap := make(map[uint64]*ConnectInfo)

	for {
		select {
		case connId := <-closeChan:
			{
				if connMap[connId] != nil {
					delete(connMap, connId)
				}
			}

		case req := <-reqChan:
			{
				connInfo := connMap[req.connId]
				if connInfo == nil {
					connInfo = &ConnectInfo{
						conn:    req.conn,
						isLogin: false,
					}
					log.Println(req.messageType)
					if req.messageType == grpcPb.MessageTypeDeviceLoginRequest {
						loginRequest, ok := req.protoMessage.(*grpcPb.DeviceLoginRequest)
						log.Println(ok)
						if ok {
							log.Println(loginRequest.String())
							connInfo.token = loginRequest.Token
							connInfo.userId = loginRequest.UserId
							connInfo.appId = loginRequest.AppId
						}
					}
					connMap[req.connId] = connInfo
				}
				if req.message.Type == grpcPb.MessageTypeRPCRequest {
					//转发给具体的业务服务器
					if !connInfo.isLogin {
						log.Println("没有登录,不转发消息")
						connInfo.conn.Close()
						delete(connMap, req.connId)
						break
					}
					rpcRequest, ok := req.protoMessage.(*grpcPb.RpcRequest)
					log.Println(rpcRequest.String())
					if !ok {
						break
					}
					rpcRequest.RpcInfo = &grpcPb.RpcInfo{
						AppId:  connInfo.appId,
						ConnId: (uint64)(req.connId),
						UserId: connInfo.userId,
						Token:  connInfo.token,
					}
					go s.transToBusinessServer(rpcRequest, s.rpcRespChan)
				} else {
					//转发给im逻辑服务器
					protoBuf, err := proto.Marshal(req.protoMessage)
					if err == nil {
						rpcRequest := &grpcPb.RpcRequest{}
						rpcRequest.RpcInfo = &grpcPb.RpcInfo{
							AppId:  connInfo.appId,
							ConnId: (uint64)(req.connId),
							UserId: connInfo.userId,
							Token:  connInfo.token,
						}
						rpcRequest.MessageType = (uint32)(req.messageType)
						rpcRequest.ProtoBuf = protoBuf
						go s.transToLogicServer(rpcRequest, s.protocolRespChan)
					}
				}
			}
		case rpcResp := <-s.rpcRespChan:
			{
				connInfo := connMap[rpcResp.ConnId]
				if connInfo == nil {
					break
				}
				buffer, err := coder.EncoderProtoMessage(grpcPb.MessageTypeRPCResponse, rpcResp)
				if err != nil {
					log.Println(err)
				}
				go connInfo.conn.Write(buffer)
			}
		case protocolBufChan := <-s.protocolRespChan:
			{
				connInfo := connMap[protocolBufChan.connId]
				connInfo.conn.Write(protocolBufChan.protocolBuf)
				if protocolBufChan.isLogin {
					connInfo.isLogin = true
				} else if protocolBufChan.isLogout {
					connInfo.isLogin = false
				}
			}
		case rpcReq := <-s.rpcReqChan:
			{
				rpcReq = rpcReq
				// connInfo := connMap[rpcReq.ConnId]
				// if connInfo == nil {
				// 	break
				// }
				// buffer, err := coder.EncoderProtoMessage(rpcReq.MessageType, rpcReq.ProtoBuf)
				// if err != nil {
				// 	log.Println(err)
				// }
				// go connInfo.conn.Write(buffer)
			}
		}
	}
}

func (s *Server) handleTcpConnection(conn *net.TCPConn, reqChan chan<- *Request, closeChan chan<- uint64) {

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(10 * time.Second)

	decoder := coder.NEWDecoder()

	atomic.AddInt32(&s.connCount, 1)
	connId := s.createConnId() //生成连接的id
	log.Println("connCount=", s.connCount)

	defer func() {
		atomic.AddInt32(&s.connCount, -1)
		log.Println("connCount=", s.connCount)
		conn.Close()
		closeChan <- connId
	}()

	for true {
		buf := make([]byte, 1024)
		count, err := conn.Read(buf)
		if err != nil {
			log.Println(err.Error())
			break
		}
		messages, err := decoder.Decode(buf[0:count])
		if err != nil {
			log.Println(err.Error())
			break
		}
		for _, message := range messages {
			request := &Request{
				connId: connId,
				conn:   conn,
			}
			//log.Println(message.Type)
			context := &context{
				reqChan:   reqChan,
				message:   message,
				closeChan: closeChan,
				conn:      conn,
				server:    s,
				request:   request,
			}
			if s.handle != nil {
				s.handle(context)
			}
		}
	}
}
