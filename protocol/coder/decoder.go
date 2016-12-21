package coder

import (
	"errors"
)

type Decoder struct {
	buffer        []byte
	messageType   *int
	messageLength *int
}

type Message struct {
	Type int    //消息类型
	Body []byte //消息的具体内容
}

func (m *Message) Encode() []byte {
	msg, _ := EncoderMessage(m.Type, m.Body)
	return msg
}

func NEWDecoder() *Decoder {

	decoder := &Decoder{}

	decoder.Reset()

	return decoder
}

func (d *Decoder) Reset() {
	d.buffer = make([]byte, 0)
	d.messageType = nil
	d.messageLength = nil
}

func (d *Decoder) decoder(buf []byte) (messages []*Message, err error) {

	d.buffer = append(d.buffer, buf...)
	//log.Println(d.buffer)

	if d.messageType == nil {
		retType, retCountByte, retNeedMore := decodeByteToType(d.buffer)
		if retNeedMore {
			return
		}
		if retCountByte > 4 {
			err = errors.New("message type too long , illegal")
			return
		}
		//log.Println(retType, "+", retCountByte, "+", retNeedMore)
		d.messageType = &retType
		d.buffer = d.buffer[retCountByte:]
	}
	if d.messageLength == nil {
		retLength, retCountByte, retNeedMore := decodeByteToLength(d.buffer)
		if retNeedMore {
			return
		}
		if retCountByte > 5 {
			err = errors.New("message content too long ,illegal")
			return
		}
		//log.Println(retLength, "+", retCountByte, "+", retNeedMore)
		d.messageLength = &retLength
		d.buffer = d.buffer[retCountByte:]
	}

	if len(d.buffer) < *d.messageLength {
		return
	}

	message := &Message{
		Type: *d.messageType,
		Body: d.buffer[:*d.messageLength],
	}
	//log.Println(message.MessageBuf)

	messages = append(messages, message)

	d.buffer = d.buffer[*d.messageLength:]
	d.messageType = nil
	d.messageLength = nil

	temp, err := d.decoder(make([]byte, 0))

	messages = append(messages, temp...)

	return
}

func (d *Decoder) Decode(buf []byte) (messages []*Message, err error) {
	return d.decoder(buf)
}

func decodeByteToType(buffer []byte) (retType int, retCountByte int, retNeedMore bool) {
	return decodeByteToLength(buffer)
}

func decodeByteToLength(buffer []byte) (retLength int, retCountByte int, retNeedMore bool) {

	retLength = 0
	retNeedMore = true
	retCountByte = 0

	for retCountByte = 0; retCountByte < len(buffer); retCountByte++ {
		byt := buffer[retCountByte]
		retLength |= int(byt & 0x7F)
		byt = buffer[retCountByte]
		if (byt & 0x80) == 0x80 {
			retLength = retLength << 7
		} else {
			retNeedMore = false
			retCountByte++
			break
		}
	}
	if retNeedMore {
		retLength = 0
	}

	return retLength, retCountByte, retNeedMore
}
