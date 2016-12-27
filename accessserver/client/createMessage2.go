package main

import (
	"fmt"
	client "im/accessserver/client/client"
	tlpPb "im/logicserver/tlp/pb"
	"im/logicserver/uuid"
	"im/tlp/coder"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	c := client.Client{
		//SsoToken: "a2ea6f80ff2748a39d89ecbd50e556fa",
		DeviceId: "034b36dc22425556bc01605d438f4d0c",
		Token:    "3",
		UserId:   "1",
	}

	c.SetAfterLogin(func(c *client.Client) {

		log.Println("登陆成功")
		for i := 0; i < 1; i++ {
			request := &tlpPb.CreateMessageRequest{
				Rid:       c.GetRid(),
				SessionId: 32,
				Type:      1,
				Id:        uuid.Rand().Hex(),
				Content:   fmt.Sprint("a message  from  token = ", c.Token, " userId = ", c.UserId),
			}

			buffer, err := coder.EncoderProtoMessage(tlpPb.MessageTypeCreateMessageRequest, request)
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
