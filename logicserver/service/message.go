package service

import (
	//"github.com/golang/protobuf/proto"
	logicserverBean "im/logicserver/bean"
	dao "im/logicserver/dao"
	logicserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/server"
	"log"
	"strings"
	"time"
)

func messageInsert(message *logicserverBean.Message) error {
	var err error
	for i := 1; i <= 10; i++ {
		_, err := dao.MessageInsert(message)
		if err == nil {
			return nil
		}
		if dao.ErrorIsDuplicate(err) {
			time.Sleep((time.Duration)(i) * 200 * time.Millisecond)
			continue
		}
		if dao.ErrorIsTooManyConnections(err) {
			time.Sleep((time.Duration)(i) * 200 * time.Millisecond)
			continue
		}
		break
	}
	return err
}

func HandleCreateMessage(ctx *server.Context, request *grpcPb.CreateMessageRequest, userId string) (*grpcPb.CreateMessageResponse, error) {

	timeNow := time.Now()
	message := &logicserverBean.Message{
		Id:         request.Id,
		SessionId:  request.SessionId,
		UserId:     userId,
		Type:       (int)(request.Type),
		Content:    request.Content,
		CreateTime: &timeNow,
	}

	err := messageInsert(message)
	if err != nil {
		log.Println(err)
		return nil, logicserverError.ErrorInternalServerError
	}

	go insertMessageToUserInSession(ctx, request, userId, &timeNow)

	response := &grpcPb.CreateMessageResponse{
		Rid:  (uint64)(request.Rid),
		Code: logicserverError.CommonSuccess,
		Desc: logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
	}

	return response, nil
}

//因为消息发送发已经发送成功了,这里一定要确保插入成功
func insertMessageToUserInSession(ctx *server.Context, request *grpcPb.CreateMessageRequest, userId string, timeNow *time.Time) {
	if request.SessionId > 0 {
		var sessionMaps []*logicserverBean.SessionMap
		dao.NewDao().Find(&sessionMaps, &logicserverBean.SessionMap{
			SessionId: request.SessionId,
		})

		for _, sessionMap := range sessionMaps {
			log.Println(sessionMap.UserId)
			if strings.EqualFold(sessionMap.UserId, userId) {
				continue
			}
			message := &logicserverBean.Message{
				Id:         request.Id,
				SessionId:  request.SessionId,
				UserId:     sessionMap.UserId,
				Type:       (int)(request.Type),
				Content:    request.Content,
				CreateTime: timeNow,
			}
			err := messageInsert(message)
			if err == nil {
				ctx.PushMessageToClient(message)
			}
		}
	}
}

// func xxxxxxxxxxxxxxxxxxx(tokenConnInfoChan chan<- int64, sessionId int64) {

// 	var sessionMaps []*logicserverBean.SessionMap

// 	err := dao.NewDao().Find(&sessionMaps,
// 		&logicserverBean.SessionMap{
// 			SessionId: sessionId,
// 		})
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	for _, sessionMap := range sessionMaps {
// 		xxx(tokenConnInfoChan, sessionMap)
// 	}
// }

// func xxx(tokenConnInfoChan chan<- int64, sessionMap *logicserverBean.SessionMap) {

// 	var tokens []*logicserverBean.Token

// 	err := dao.NewDao().Find(&tokens,
// 		&logicserverBean.Token{
// 			UserId: sessionMap.UserId,
// 		})
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	for _, token := range tokens {
// 		//log.Println(token.Id)
// 		tokenConnInfoChan <- token.Id
// 		//token.Id 根据登录的id去发送
// 	}
// }
