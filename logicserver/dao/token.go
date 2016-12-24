package dao

import (
	//"im/logicserver/bean"
	"im/logicserver/mysql"
)

func SyncKeyByUserId(userId string) (uint64, error) {

	engine := mysql.GetXormEngine()
	sqlQuery := "select (`sync_key`) from t_token where user_id = ?"

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

func UpdateSyncKey(newSyncKey int64, oldSyncKey int64, userId string) error {
	engine := mysql.GetXormEngine()

	sql := "update `t_token` set `sync_key`=? where `id`=? and `sync_key` = ?"

	_, err := engine.DB().Exec(sql, newSyncKey, userId, oldSyncKey)

	return err
}
