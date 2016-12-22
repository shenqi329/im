package dao

import (
	"im/logicserver/bean"
	"im/logicserver/mysql"
)

func MessageInsert(message *bean.Message) (int64, error) {

	var err error
	syncKey, err := MessageMaxIndexByUserId(message.UserId)

	if err != nil {
		return 0, err
	}
	message.SyncKey = syncKey + 1
	_, err = NewDao().Insert(message)

	if err == nil {
		return 1, nil
	}

	if ErrorIsTooManyConnections(err) {
		return 0, err
	}

	return 0, err
}

func MessageMaxIndexByUserId(userId string) (uint64, error) {

	engine := mysql.GetXormEngine()
	sqlQuery := "select max(`sync_key`) from t_message where user_id = ?"

	var syncKey interface{}
	err := engine.DB().QueryRow(sqlQuery, userId).Scan(&syncKey)
	if err != nil {
		return 0, err
	}
	ret, ok := syncKey.(int64)
	if ok {
		return (uint64)(ret), nil
	}
	return 0, err
}
