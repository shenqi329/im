package server

import (
	"im/logicserver/bean"
)

type Context struct {
	Server *Server
}

func (c *Context) PushMessageToClient(message *bean.Message) error {
	return nil
}
