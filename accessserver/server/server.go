package server

import (
	"github.com/golang/protobuf/proto"
	netContext "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	accessError "im/accessserver/error"
	accessserverGrpcPb "im/accessserver/grpc/pb"
	"im/accessserver/util/key"
	grpcPb "im/logicserver/grpc/pb"
	tlpPb "im/logicserver/tlp/pb"
	coder "im/tlp/coder"
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
	messageType  tlpPb.MessageType
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
	rid                    uint64 //请求流水号
	ridMutex               sync.Mutex
	connCount              int32
	connIdMutex            sync.Mutex
	connId                 uint64 //请求的id
	localTcpAddr           string
	proxyUdpAddr           string
	grpcEasynoteClientConn *grpc.ClientConn
	grpcLogicClientConn    *grpc.ClientConn
	grpcServer             *grpc.Server
	rpcRespChan            chan *grpcPb.ForwardTLPResponse
	protocolRespChan       chan *ProtocolResp
	forwardTLPRequestChan  chan *accessserverGrpcPb.ForwardTLPRequest
	handle                 func(context Context) error
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
		localTcpAddr:          localTcpAddr,
		proxyUdpAddr:          proxyUdpAddr,
		rpcRespChan:           make(chan *grpcPb.ForwardTLPResponse, 1000),
		protocolRespChan:      make(chan *ProtocolResp, 1000),
		forwardTLPRequestChan: make(chan *accessserverGrpcPb.ForwardTLPRequest, 1000),
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

func (s *Server) ForwardTLP(request *accessserverGrpcPb.ForwardTLPRequest) (*accessserverGrpcPb.ForwardTLPResponse, error) {
	log.Println("ForwardTLP", request.String())
	s.forwardTLPRequestChan <- request
	return nil, nil
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

	reflection.Register(s.GrpcServer())
	if err := s.GrpcServer().Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) newGrpcServer() *grpc.Server {

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(func(ctx netContext.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		ctx = netContext.WithValue(ctx, key.Server, s)
		response, err := handler(ctx, req)
		v := reflect.ValueOf(response)
		if !v.IsValid() || v.IsNil() {
			response = &grpcPb.Response{
				Code: accessError.CommonInternalServerError,
				Desc: accessError.ErrorCodeToText(accessError.CommonInternalServerError),
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

func (s *Server) transToLogicServer(request *grpcPb.ForwardTLPRequest, protocolRespChan chan<- *ProtocolResp) {
	//log.Println("转发给逻辑服务器")

	rpcClient := grpcPb.NewForwardToLogicClient(s.grpcLogicClientConn)
	response, err := rpcClient.ForwardTLP(netContext.Background(), request)

	if err != nil {
		log.Println(err)
		return
	}
	if !accessError.CodeIsSuccess(response.Code) {
		log.Println(response)
		return
	}

	var isLogin bool = false
	var isLogout bool = false
	if request.MessageType == tlpPb.MessageTypeDeviceLoginRequest {
		isLogin = true
	}
	if response.ProtoBuf == nil || response.MessageType <= 0 {
		//没有数据,或者消息类型不对,则不需要将消息在发给客户端
		return
	}

	protocolBuf, _ := coder.EncoderProtoBuf((int)(response.MessageType), response.ProtoBuf)

	protocolRespChan <- &ProtocolResp{
		protocolBuf: protocolBuf,
		connId:      request.RpcInfo.ConnId,
		isLogin:     isLogin,
		isLogout:    isLogout,
	}
}

func (s *Server) transToBusinessServer(rpcRequest *grpcPb.ForwardTLPRequest, rpcRespChan chan<- *grpcPb.ForwardTLPResponse) {
	//easynote业务id
	if rpcRequest.RpcInfo.AppId == "89897" {
		//log.Println("转发给业务服务器")
		rpcClient := grpcPb.NewForwardToLogicClient(s.grpcEasynoteClientConn)
		response, err := rpcClient.ForwardTLP(netContext.Background(), rpcRequest)
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
	tokenMap := make(map[string]*ConnectInfo)

	for {
		select {
		case connId := <-closeChan:
			{
				connInfo := connMap[connId]
				if connInfo != nil {
					delete(connMap, connId)
					delete(tokenMap, connInfo.token)
					//发送消息给逻辑服务器
					rpcClient := grpcPb.NewOfflineClient(s.grpcLogicClientConn)
					offlineRequest := &grpcPb.DeviceOfflineRequest{
						Token:  connInfo.token,
						UserId: connInfo.userId,
					}
					rpcClient.Offline(netContext.Background(), offlineRequest)
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
					if req.messageType == tlpPb.MessageTypeDeviceLoginRequest {
						loginRequest, ok := req.protoMessage.(*tlpPb.DeviceLoginRequest)
						log.Println(ok)
						if ok {
							log.Println(loginRequest.String())
							connInfo.token = loginRequest.Token
							connInfo.userId = loginRequest.UserId
							connInfo.appId = loginRequest.AppId
						}
					}
					connMap[req.connId] = connInfo
					tokenMap[connInfo.token] = connInfo
				}
				if req.message.Type == tlpPb.MessageTypeForwardTLPRequest {
					//转发给具体的业务服务器
					if !connInfo.isLogin {
						log.Println("没有登录,不转发消息")
						connInfo.conn.Close()
						delete(connMap, req.connId)
						delete(tokenMap, connInfo.token)
						break
					}
					request, ok := req.protoMessage.(*grpcPb.ForwardTLPRequest)
					log.Println(request.String())
					if !ok {
						break
					}
					request.RpcInfo = &grpcPb.RpcInfo{
						AppId:  connInfo.appId,
						ConnId: (uint64)(req.connId),
						UserId: connInfo.userId,
						Token:  connInfo.token,
					}
					go s.transToBusinessServer(request, s.rpcRespChan)
				} else {
					//转发给im逻辑服务器
					protoBuf, err := proto.Marshal(req.protoMessage)
					if err == nil {
						request := &grpcPb.ForwardTLPRequest{}
						request.RpcInfo = &grpcPb.RpcInfo{
							AppId:  connInfo.appId,
							ConnId: (uint64)(req.connId),
							UserId: connInfo.userId,
							Token:  connInfo.token,
						}
						request.MessageType = (uint32)(req.messageType)
						request.ProtoBuf = protoBuf
						go s.transToLogicServer(request, s.protocolRespChan)
					}
				}
			}
		case rpcResp := <-s.rpcRespChan:
			{
				connInfo := connMap[rpcResp.ConnId]
				if connInfo == nil {
					break
				}
				buffer, err := coder.EncoderProtoMessage(tlpPb.MessageTypeForwardTLPResponse, rpcResp)
				if err != nil {
					log.Println(err)
				}
				go connInfo.conn.Write(buffer)
			}
		case protocolBufChan := <-s.protocolRespChan:
			{
				connInfo := connMap[protocolBufChan.connId]
				if connInfo == nil {
					break
				}
				connInfo.conn.Write(protocolBufChan.protocolBuf)
				if protocolBufChan.isLogin {
					connInfo.isLogin = true
				} else if protocolBufChan.isLogout {
					connInfo.isLogin = false
				}
			}
		case request := <-s.forwardTLPRequestChan:
			{
				connInfo := tokenMap[request.Token]
				if connInfo == nil {
					break
				}
				buffer, err := coder.EncoderProtoBuf((int)(request.Type), request.ProtoBuf)
				if err != nil {
					log.Println(err)
				}
				go connInfo.conn.Write(buffer)
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
