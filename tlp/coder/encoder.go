package coder

import (
	proto "github.com/golang/protobuf/proto"
)

func EncoderMessage(messageType int, message []byte) ([]byte, error) {

	length := len(message)

	messageByte := encodeTypeToByte(messageType)
	lengthByte := encodeLengthToByte(length)

	messageByte = append(messageByte, lengthByte...)
	messageByte = append(messageByte, message...)

	return messageByte, nil
}

func EncoderProtoMessage(messageType int, message proto.Message) ([]byte, error) {

	b, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	return EncoderMessage(messageType, b)
}

func encodeLengthToByte(length int) []byte {

	needByte := 1
	tempLen := length

	for tempLen >= 128 {
		tempLen = tempLen >> 7
		needByte++
	}

	tempLen = length
	buffer := make([]byte, needByte)

	for index := 0; index < needByte; index++ {
		buffer[needByte-index-1] = (byte)(tempLen & 0x7F)
		if index != 0 {
			buffer[needByte-index-1] |= 0x80
		}
		tempLen = tempLen >> 7
	}
	return buffer
}

func encodeTypeToByte(typ int) []byte {

	return encodeLengthToByte(typ)

}
