package pb

import (
	proto "github.com/golang/protobuf/proto"
)

type MessageType int32

const (
	MessageTypeDeviceRegisteRequest  = 1
	MessageTypeDeviceRegisteResponse = 2
	MessageTypeDeviceLoginRequest    = 3
	MessageTypeDeviceLoginResponse   = 4
	MessageTypeForwardTLPRequest     = 5
	MessageTypeForwardTLPResponse    = 6
	MessageTypeCreateMessageRequest  = 7
	MessageTypeCreateMessageResponse = 8
	MessageTypeSyncMessage           = 9

	// MessageTypeCreateSessionRequest       = 10001
	// MessageTypeCreateSessionResponse      = 10002
	// MessageTypeAddSessionUsersRequest     = 10003
	// MessageTypeAddSessionUsersResponse    = 10004
	// MessageTypeDeleteSessionUsersRequest  = 10005
	// MessageTypeDeleteSessionUsersResponse = 10006
)

var kinds = map[MessageType]func() proto.Message{
	MessageTypeDeviceRegisteRequest:  func() proto.Message { return &DeviceRegisteRequest{} },
	MessageTypeDeviceRegisteResponse: func() proto.Message { return &DeviceRegisteResponse{} },
	MessageTypeDeviceLoginRequest:    func() proto.Message { return &DeviceLoginRequest{} },
	MessageTypeDeviceLoginResponse:   func() proto.Message { return &DeviceLoginResponse{} },
	MessageTypeForwardTLPRequest:     func() proto.Message { return &ForwardTLPRequest{} },
	MessageTypeForwardTLPResponse:    func() proto.Message { return &ForwardTLPResponse{} },
	MessageTypeCreateMessageRequest:  func() proto.Message { return &CreateMessageRequest{} },
	MessageTypeCreateMessageResponse: func() proto.Message { return &CreateMessageResponse{} },
	MessageTypeSyncMessage:           func() proto.Message { return &SyncMessage{} },

	// MessageTypeCreateSessionRequest:       func() proto.Message { return &CreateSessionRequest{} },
	// MessageTypeCreateSessionResponse:      func() proto.Message { return &CreateSessionResponse{} },
	// MessageTypeAddSessionUsersRequest:     func() proto.Message { return &AddSessionUsersRequest{} },
	// MessageTypeAddSessionUsersResponse:    func() proto.Message { return &Response{} },
	// MessageTypeDeleteSessionUsersRequest:  func() proto.Message { return &DeleteSessionUsersRequest{} },
	// MessageTypeDeleteSessionUsersResponse: func() proto.Message { return &Response{} },
}

func Factory(messageType MessageType) proto.Message {

	createFunc := kinds[messageType]
	if createFunc != nil {
		return createFunc()
	}
	return nil
}
