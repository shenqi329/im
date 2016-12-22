package service

import (
	logicserverBean "im/logicserver/bean"
	dao "im/logicserver/dao"
	logicserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/server"
	"strconv"
	"time"
)

func HandleOffline(ctx *server.Context, deviceOfflineRequest *grpcPb.DeviceOfflineRequest) (protoMessage *grpcPb.Response, err error) {

	protoMessage = &grpcPb.Response{
		Code: logicserverError.CommonSuccess,
		Desc: logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
	}

	err = ctx.HandleOffline(deviceOfflineRequest)
	if err != nil {
		protoMessage = &grpcPb.Response{
			Code: logicserverError.CommonInternalServerError,
			Desc: logicserverError.ErrorCodeToText(logicserverError.CommonInternalServerError),
		}
	}
	return
}

func HandleLogin(ctx *server.Context, deviceLoginRequest *grpcPb.DeviceLoginRequest) (protoMessage *grpcPb.DeviceLoginResponse, err error) {

	id, _ := strconv.ParseUint(deviceLoginRequest.Token, 10, 64)

	tokenBean := &logicserverBean.Token{
		Id:       (int64)(id),
		AppId:    deviceLoginRequest.AppId,
		DeviceId: deviceLoginRequest.DeviceId,
		Platform: deviceLoginRequest.Platform,
	}
	has, err := dao.NewDao().Get(tokenBean)

	if err != nil {
		protoMessage = &grpcPb.DeviceLoginResponse{
			Rid:  deviceLoginRequest.Rid,
			Code: logicserverError.CommonInternalServerError,
			Desc: logicserverError.ErrorCodeToText(logicserverError.CommonInternalServerError),
		}
		return
	}
	if !has {
		protoMessage = &grpcPb.DeviceLoginResponse{
			Rid:  deviceLoginRequest.Rid,
			Code: logicserverError.CommonResourceNoExist,
			Desc: logicserverError.ErrorCodeToText(logicserverError.CommonResourceNoExist),
		}
		//err = logicserverError.ErrorNotFound
		return
	}
	if tokenBean.LoginTime == nil {
		timeNow := time.Now()
		tokenBean.LoginTime = &timeNow
	}

	protoMessage = &grpcPb.DeviceLoginResponse{
		Rid:  deviceLoginRequest.Rid,
		Code: logicserverError.CommonSuccess,
		Desc: logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
	}
	ctx.HandleLogin(deviceLoginRequest)

	return
}

// //发送同步通知
// func sendSyncInform(c server.Context, deviceLoginRequest *protocolClient.DeviceLoginRequest, userId string) {

// 	var sessionMaps []*logicserverBean.SessionMap

// 	err := dao.NewDao().Find(&sessionMaps, &logicserverBean.SessionMap{
// 		UserId: userId,
// 	})
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	for _, sessionMap := range sessionMaps {
// 		sendSyncInformWithSessionMap(c, sessionMap)
// 	}
// }

// func sendSyncInformWithSessionMap(c server.Context, sessionMap *logicserverBean.SessionMap) {

// 	var messages []*logicserverBean.Message

// 	err := dao.NewDao().Find(&messages, &logicserverBean.Message{
// 		SessionId: sessionMap.SessionId,
// 	})

// 	latestIndex, err := dao.MessageMaxIndex(sessionMap.SessionId)

// 	if sessionMap.ReadIndex >= latestIndex {
// 		//log.Println("不需发送同步通知")
// 		return
// 	}

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	log.Println(latestIndex)

// 	syncInfo := &protocolClient.SyncInform{
// 		SessionId:   sessionMap.SessionId,
// 		LatestIndex: latestIndex,
// 		ReadIndex:   sessionMap.ReadIndex,
// 	}

// 	c.SendProtoMessage(protocolClient.MessageTypeSyncInform, syncInfo)
// }
