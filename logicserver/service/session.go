package service

import (
	//"github.com/golang/protobuf/proto"
	logicserverBean "im/logicserver/bean"
	dao "im/logicserver/dao"
	logicserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"log"
	"strings"
)

func CreateSession(request *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionResponse, error) {
	log.Println(request.String())

	session := &logicserverBean.Session{
		AppId:        request.RpcInfo.AppId,
		CreateUserId: request.RpcInfo.UserId,
	}

	_, err := dao.NewDao().Insert(session)
	if err != nil {
		log.Println(err.Error())
		err = logicserverError.ErrorInternalServerError
		return nil, err
	}

	sessionMap := &logicserverBean.SessionMap{
		SessionId: (uint64)(session.Id),
		UserId:    request.RpcInfo.UserId,
	}
	sessionMaps := []interface{}{sessionMap}

	for i := 0; i < len(request.UserIds); i++ {
		if strings.EqualFold(request.UserIds[i], request.RpcInfo.UserId) {
			continue
		}
		sessionMap = &logicserverBean.SessionMap{
			SessionId: (uint64)(session.Id),
			UserId:    request.UserIds[i],
		}
		sessionMaps = append(sessionMaps, sessionMap)
	}
	_, err = dao.NewDao().Insert(sessionMaps...)

	if err != nil {
		log.Println(err.Error())
		err = logicserverError.ErrorInternalServerError
		return nil, err
	}

	response := &grpcPb.CreateSessionResponse{
		Rid:       (uint64)(request.Rid),
		Code:      logicserverError.CommonSuccess,
		Desc:      logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
		SessionId: (uint64)(session.Id),
	}

	return response, nil
}

func DeleteSessionUsers(request *grpcPb.DeleteSessionUsersRequest) (*grpcPb.Response, error) {

	log.Println(request.String())

	for i := 0; i < len(request.DeleteUserIds); i++ {
		sessionMap := &logicserverBean.SessionMap{
			SessionId: request.SessionId,
			UserId:    request.DeleteUserIds[i],
		}
		_, _ = dao.NewDao().Delete(sessionMap)
	}

	response := &grpcPb.Response{
		Rid:  (uint64)(request.Rid),
		Code: logicserverError.CommonSuccess,
		Desc: logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
	}
	return response, nil
}

func AddSessionUsers(request *grpcPb.AddSessionUsersRequest) (*grpcPb.Response, error) {

	log.Println(request.String())

	sessionMaps := []interface{}{}

	for i := 0; i < len(request.AddUserIds); i++ {
		sessionMap := &logicserverBean.SessionMap{
			SessionId: request.SessionId,
			UserId:    request.AddUserIds[i],
		}
		sessionMaps = append(sessionMaps, sessionMap)
	}

	_, err := dao.NewDao().Insert(sessionMaps...)

	if err != nil {
		log.Println(err.Error())
		err = logicserverError.ErrorInternalServerError
		return nil, err
	}

	response := &grpcPb.Response{
		Rid:  (uint64)(request.Rid),
		Code: logicserverError.CommonSuccess,
		Desc: logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
	}
	return response, nil
}
