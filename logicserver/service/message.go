package service

import (
	//"github.com/golang/protobuf/proto"
	logicserverBean "im/logicserver/bean"
	dao "im/logicserver/dao"
	logicserverError "im/logicserver/error"
	"im/logicserver/server"
	tlpPb "im/logicserver/tlp/pb"
	"log"
	"strings"
	"time"
)

func messageInsert(message *logicserverBean.Message) error {
	var err error
	for i := 1; i <= 10; i++ {
		_, err = dao.MessageInsert(message)
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

func HandleCreateMessage(ctx *server.Context, request *tlpPb.CreateMessageRequest, userId string) (*tlpPb.CreateMessageResponse, error) {

	timeNow := time.Now()
	message := &logicserverBean.Message{
		Id:         request.Id,
		SessionId:  request.SessionId,
		UserId:     userId,
		Type:       (int)(request.Type),
		Content:    request.Content,
		CreateTime: timeNow,
	}

	err := messageInsert(message)
	if err != nil {
		log.Println(err)
		return nil, logicserverError.ErrorInternalServerError
	}
	ctx.SendMessageToUser(message)

	go insertMessageToUserInSession(ctx, request, userId, &timeNow)

	response := &tlpPb.CreateMessageResponse{
		Rid:  (uint64)(request.Rid),
		Code: logicserverError.CommonSuccess,
		Desc: logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
	}

	return response, nil
}

//因为消息发送发已经发送成功了,这里一定要确保插入成功
func insertMessageToUserInSession(ctx *server.Context, request *tlpPb.CreateMessageRequest, userId string, timeNow *time.Time) {
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
				CreateTime: *timeNow,
			}
			err := messageInsert(message)
			if err == nil {
				ctx.SendMessageToUser(message)
			}
		}
	}
}

func HandleSyncFinResponse(ctx *server.Context, response *tlpPb.SyncFinResponse, userId string) error {

	oldSyncKey, err := dao.SyncKeyByUserId(userId)
	if err != nil {
		return err
	}
	if (uint64)(oldSyncKey) >= (uint64)(response.SyncKey) {
		return nil
	}
	return dao.UpdateSyncKey((int64)(response.SyncKey), (int64)(oldSyncKey), userId)
}
