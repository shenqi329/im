package dao

import (
	//"im/logicserver/bean"
	"database/sql"
	"im/logicserver/mysql"
	"log"
)

func SyncKeyByUserId(userId string) (uint64, error) {

	engine := mysql.GetXormEngine()
	sqlQuery := "select (`sync_key`) from `t_user_sync_key` where user_id = ?"

	var syncKey interface{}
	err := engine.DB().QueryRow(sqlQuery, userId).Scan(&syncKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		log.Println(err)
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

	if oldSyncKey == 0 {
		sql := "insert into `t_user_sync_key` (`user_id`,`sync_key`) values(?,?)"
		_, err := engine.DB().Exec(sql, userId, newSyncKey)
		return err
	}
	sql := "update `t_user_sync_key` set `sync_key`=? where `user_id`=? and `sync_key` = ?"
	_, err := engine.DB().Exec(sql, newSyncKey, userId, oldSyncKey)

	return err
}
