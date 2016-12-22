package bean

import (
	"time"
)

type (
	Token struct {
		Id         int64      `xorm:"'id'" json:"id,omitempty"`
		AppId      string     `xorm:"'app_id'" json:"appId,omitempty"`
		UserId     string     `xorm:"'user_id'" json:"userId,omitempty"`
		DeviceId   string     `xorm:"'device_id'" json:"deviceId,omitempty"`
		Platform   string     `xorm:"'platform'" json:"platform,omitempty"`
		CreateTime *time.Time `xorm:"'create_time'" json:"createTime,omitempty"`
		LoginTime  *time.Time `xorm:"'login_time'" json:"loginTime,omitempty"`
		//LoginoutTime *time.Time `xorm:"'t_token_logout_time'" json:"logoutTime,omitempty"`
	}
)

func (t Token) TableName() string {
	return "t_token"
}
