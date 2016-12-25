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
	MessageTypeSyncFinInform         = 10
	MessageTypeSyncFinResponse       = 11
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
	MessageTypeSyncFinInform:         func() proto.Message { return &SyncFinInform{} },
	MessageTypeSyncFinResponse:       func() proto.Message { return &SyncFinResponse{} },
}

func Factory(messageType MessageType) proto.Message {

	createFunc := kinds[messageType]
	if createFunc != nil {
		return createFunc()
	}
	return nil
}
