package pb

import (
	proto "github.com/golang/protobuf/proto"
)

type MessageType int32

const (
	MessageTypeCreateSessionRequest       = 10001
	MessageTypeCreateSessionResponse      = 10002
	MessageTypeAddSessionUsersRequest     = 10003
	MessageTypeAddSessionUsersResponse    = 10004
	MessageTypeDeleteSessionUsersRequest  = 10005
	MessageTypeDeleteSessionUsersResponse = 10006
)

var kinds = map[MessageType]func() proto.Message{
	MessageTypeCreateSessionRequest:       func() proto.Message { return &CreateSessionRequest{} },
	MessageTypeCreateSessionResponse:      func() proto.Message { return &CreateSessionResponse{} },
	MessageTypeAddSessionUsersRequest:     func() proto.Message { return &AddSessionUsersRequest{} },
	MessageTypeAddSessionUsersResponse:    func() proto.Message { return &Response{} },
	MessageTypeDeleteSessionUsersRequest:  func() proto.Message { return &DeleteSessionUsersRequest{} },
	MessageTypeDeleteSessionUsersResponse: func() proto.Message { return &Response{} },
}

func Factory(messageType MessageType) proto.Message {

	createFunc := kinds[messageType]
	if createFunc != nil {
		return createFunc()
	}
	return nil
}
