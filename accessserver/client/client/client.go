package client

import (
	"github.com/golang/protobuf/proto"
	tlpPb "im/logicserver/tlp/pb"
	"im/tlp/coder"
	"log"
	"net"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
)

const (
	LoginStateNone      = 0
	LoginStateInRegiste = 1 //正在注册
	LoginStateRegisted  = 2 //已注册
	LoginStateInLogin   = 3 //正在登录
	LoginStateLogined   = 4 //已登录
)

type Client struct {
	rid        uint64
	recvCount  uint32
	conn       *net.TCPConn
	loginState uint32
	afterLogin func(c *Client)

	SsoToken string
	DeviceId string
	Token    string
	UserId   string
}

func (c *Client) GetRid() uint64 {
	atomic.AddUint64(&c.rid, 1)
	return c.rid
}

func (c *Client) GetConn() *net.TCPConn {
	return c.conn
}

func (c *Client) SetAfterLogin(afterLogin func(c *Client)) {
	c.afterLogin = afterLogin
}

func (c *Client) registe() {
	registeRequest := &tlpPb.DeviceRegisteRequest{
		Rid:      c.GetRid(),
		SsoToken: c.SsoToken,
		AppId:    "89897",
		DeviceId: c.DeviceId,
		Platform: "windows",
	}
	buffer, err := coder.EncoderProtoMessage(tlpPb.MessageTypeDeviceRegisteRequest, registeRequest)
	if err != nil {
		log.Println(err.Error())
	}
	c.conn.Write(buffer)
	c.loginState = LoginStateInRegiste
}

func (c *Client) login() {
	log.Println(c.Token)
	if runtime.GOOS == "windows" {
		loginRequest := &tlpPb.DeviceLoginRequest{
			Rid:      c.GetRid(),
			Token:    c.Token,
			UserId:   c.UserId,
			AppId:    "89897",
			DeviceId: c.DeviceId,
			Platform: "windows",
		}
		buffer, err := coder.EncoderProtoMessage(tlpPb.MessageTypeDeviceLoginRequest, loginRequest)
		if err != nil {
			log.Println(err.Error())
		}
		c.conn.Write(buffer)
		c.loginState = LoginStateInLogin
	} else {
		loginRequest := &tlpPb.DeviceLoginRequest{
			Rid:      c.GetRid(),
			Token:    c.Token,
			UserId:   c.UserId,
			AppId:    "89897",
			DeviceId: c.DeviceId,
			Platform: "windows",
		}

		buffer, err := coder.EncoderProtoMessage(tlpPb.MessageTypeDeviceLoginRequest, loginRequest)
		if err != nil {
			log.Println(err.Error())
		}
		c.conn.Write(buffer)
		c.loginState = LoginStateInLogin
	}
}

func (c *Client) toLogin() {
	log.Println(c.loginState)
	if c.loginState == LoginStateNone {
		if c.Token == "" {
			c.registe()
		} else {
			c.loginState = LoginStateRegisted
		}
	}
	if c.loginState == LoginStateRegisted {
		c.login()
	}
}

func (c *Client) LoginToAccessServer() {

	raddr, err := net.ResolveTCPAddr("tcp", "localhost:6000")
	if runtime.GOOS == "windows" {
		raddr, err = net.ResolveTCPAddr("tcp", "localhost:6000")
	}

	if err != nil {
		log.Println("net.ResolveTCPAddr fail.", err)
		return
	}
	connect, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Println("net.ListenTCP fail.", err.Error())
		return
	}

	connect.SetKeepAlive(true)
	connect.SetKeepAlivePeriod(10 * time.Second)
	c.conn = connect
	go c.handleConnection(connect)

	c.toLogin()
}

func (c *Client) handleConnection(conn *net.TCPConn) {

	decoder := coder.NEWDecoder()
	buf := make([]byte, 512)
	for true {
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
			go c.handleMessage(conn, message)
		}
	}
}

func (c *Client) handleMessage(conn *net.TCPConn, message *coder.Message) {

	protoMessage := tlpPb.Factory((tlpPb.MessageType)(message.Type))

	if protoMessage == nil {
		log.Println("未识别的消息")
		conn.Close()
		return
	}

	if err := proto.Unmarshal(message.ProtoBuf, protoMessage); err != nil {
		log.Println(err.Error())
		log.Println("消息格式错误")
		conn.Close()
		return
	}
	c.recvCount = atomic.AddUint32(&c.recvCount, 1)

	log.Println("userId = ", c.UserId, "token = ", c.Token, "recvMsg count = ", c.recvCount, "context:", proto.CompactTextString(protoMessage))

	if (tlpPb.MessageType)(message.Type) == tlpPb.MessageTypeDeviceLoginResponse {
		response := protoMessage.(*tlpPb.DeviceLoginResponse)
		if strings.EqualFold(response.Code, "00000001") {
			c.loginState = LoginStateLogined
			go c.afterLogin(c)
		} else {
			log.Print("登陆失败")
		}
	} else if (tlpPb.MessageType)(message.Type) == tlpPb.MessageTypeDeviceRegisteResponse {
		response := protoMessage.(*tlpPb.DeviceRegisteResponse)
		if strings.EqualFold(response.Code, "00000001") {
			c.loginState = LoginStateRegisted
			c.Token = response.Token
			c.toLogin()
		} else {
			log.Print("注册失败")
		}
	}
}
