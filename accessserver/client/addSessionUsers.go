package main

import (
	client "im/accessserver/client/client"
	grpcPb "im/logicserver/grpc/pb"
	//"im/logicserver/uuid"
	"github.com/golang/protobuf/proto"
	"im/protocol/coder"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	c := client.Client{
		SsoToken: "",
		DeviceId: "024b36dc22425556bc01605d438f4d0c",
		Token:    "1",
		UserId:   "1",
	}

	c.SetAfterLogin(func(c *client.Client) {

		log.Println("登陆成功")
		for i := 0; i < 1; i++ {
			addSessionUserRequest := &grpcPb.AddSessionUsersRequest{
				Rid:        c.GetRid(),
				SessionId:  32,
				AddUserIds: []string{"1", "2"},
			}
			protoBuf, err := proto.Marshal(addSessionUserRequest)
			if err != nil {
				log.Print(err.Error())
			}

			rpcRequest := &grpcPb.RpcRequest{
				Rid:         c.GetRid(),
				MessageType: grpcPb.MessageTypeAddSessionUsersRequest,
				ProtoBuf:    protoBuf,
			}

			buffer, err := coder.EncoderProtoMessage(grpcPb.MessageTypeRPCRequest, rpcRequest)
			if err != nil {
				log.Println(err.Error())
			}
			c.GetConn().Write(buffer)
		}
	})

	c.LoginToAccessServer()
	time.Sleep(60 * time.Minute)
}
