package main

import (
	"fmt"
	client "im/accessserver/client/client"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/uuid"
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
			request := &grpcPb.CreateMessageRequest{
				Rid:       c.GetRid(),
				SessionId: 32,
				Type:      1,
				Id:        uuid.Rand().Hex(),
				Content:   fmt.Sprint("a message from push ", i+1),
			}

			buffer, err := coder.EncoderProtoMessage(grpcPb.MessageTypeCreateMessageRequest, request)
			if err != nil {
				log.Println(err.Error())
			}
			c.GetConn().Write(buffer)
			//time.Sleep(25 * time.Millisecond)
		}
	})

	c.LoginToAccessServer()
	time.Sleep(60 * time.Minute)
}
