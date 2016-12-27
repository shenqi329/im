package server

import (
	"github.com/golang/protobuf/proto"
	netContext "golang.org/x/net/context"
	"google.golang.org/grpc"
	grpcPb "im/accessserver/grpc/pb"
	"im/logicserver/bean"
	"im/logicserver/dao"
	tlpPb "im/logicserver/tlp/pb"
	"log"
)

type AccessClient struct {
	conn *grpc.ClientConn
}

func NEWAccessClient() *AccessClient {
	a := &AccessClient{}
	return a
}

func (a *AccessClient) Connect(tcpAddr string) {
	log.Println("Connect")
	conn, err := grpc.Dial(tcpAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	a.conn = conn
}

func (a *AccessClient) SendSyncMessageToUser(userId string, token string) error {

	syncKey, err := dao.SyncKeyByUserId(userId)
	if err != nil {
		return err
	}
	maxSyncKey, err := dao.MessageMaxIndexByUserId(userId)
	if err != nil {
		return err
	}
	if syncKey >= maxSyncKey {
		return nil
	}

	fromSyncKey := syncKey + 1
	toSyncKey := maxSyncKey
	return a.SendSyncMessageFromKeyToKeyToUser(fromSyncKey, toSyncKey, userId, token)
}

func (a *AccessClient) SendSyncMessageFromKeyToKeyToUser(fromSyncKey uint64, toSyncKey uint64, userId string, token string) error {

	forwardToAccessClient := grpcPb.NewForwardToAccessClient(a.conn)
	for syncKey := fromSyncKey; syncKey <= toSyncKey; syncKey++ {
		message := &bean.Message{
			UserId:  userId,
			SyncKey: syncKey,
		}
		_, err := dao.NewDao().Get(message)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(bean.StructToJsonString(message))
		syncMessage := &tlpPb.SyncMessage{
			Type:       (int32)(message.Type),
			Id:         message.Id,
			SessionId:  message.SessionId,
			UserId:     message.UserId,
			Content:    message.Content,
			SyncKey:    message.SyncKey,
			CreateTime: message.CreateTime.Unix(),
		}
		protobuf, _ := proto.Marshal(syncMessage)
		request := &grpcPb.ForwardTLPRequest{
			UserId:   message.UserId,
			Token:    token,
			Type:     tlpPb.MessageTypeSyncMessage,
			ProtoBuf: protobuf,
		}
		_, err = forwardToAccessClient.ForwardTLP(netContext.Background(), request)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	syncFinMessage := &tlpPb.SyncFinInform{
		SyncKey: toSyncKey,
	}
	protobuf, _ := proto.Marshal(syncFinMessage)
	request := &grpcPb.ForwardTLPRequest{
		UserId:   userId,
		Token:    token,
		Type:     tlpPb.MessageTypeSyncFinInform,
		ProtoBuf: protobuf,
	}
	_, err := forwardToAccessClient.ForwardTLP(netContext.Background(), request)
	return err
}

func (a *AccessClient) SendMessage(message *bean.Message, token string) error {

	log.Println(bean.StructToJsonString(message), "token:", token)
	syncKey, err := dao.SyncKeyByUserId(message.UserId)
	if err != nil {
		log.Println(err)
		return err
	}
	return a.SendSyncMessageFromKeyToKeyToUser(syncKey+1, message.SyncKey, message.UserId, token)
}
