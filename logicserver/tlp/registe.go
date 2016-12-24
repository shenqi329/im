package tlp

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"im/logicserver/service"
	tlpPb "im/logicserver/tlp/pb"
	"log"
)

func HandleRegiste(ctx context.Context, message proto.Message) (proto.Message, error) {

	log.Println("Registe")
	request, ok := message.(*tlpPb.DeviceRegisteRequest)
	if !ok {
		return nil, nil
	}
	protoMessage, err := service.HandleRegiste(request)

	return protoMessage, err
}
