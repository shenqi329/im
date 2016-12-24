package main

import (
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
		SsoToken: "",
		DeviceId: "124b36dc22425556bc01605d438f4d0d",
		Token:    "2",
		UserId:   "2",
	}

	c.SetAfterLogin(func(c *client.Client) {

		log.Println("登陆成功")
		for i := 0; i < 1; i++ {
			request := &tlpPb.CreateMessageRequest{
				Rid:       c.GetRid(),
				SessionId: 32,
				Type:      1,
				Id:        uuid.Rand().Hex(),
				Content:   "a message from push",
			}

			buffer, err := coder.EncoderProtoMessage(tlpPb.MessageTypeCreateMessageRequest, request)
			if err != nil {
				log.Println(err.Error())
			}
			c.GetConn().Write(buffer)
		}
	})

	c.LoginToAccessServer()
	time.Sleep(60 * time.Minute)
}
