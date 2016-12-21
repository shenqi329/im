package bean

import (
	"time"
)

type (
	Token struct {
		Id         int64      `xorm:"'t_token_id'" json:"id,omitempty"`
		AppId      string     `xorm:"'t_token_app_id'" json:"appId,omitempty"`
		UserId     string     `xorm:"'t_token_user_id'" json:"userId,omitempty"`
		DeviceId   string     `xorm:"'t_token_device_id'" json:"deviceId,omitempty"`
		Platform   string     `xorm:"'t_token_platform'" json:"platform,omitempty"`
		CreateTime *time.Time `xorm:"'t_token_create_time'" json:"createTime,omitempty"`
		LoginTime  *time.Time `xorm:"'t_token_login_time'" json:"loginTime,omitempty"`
		//LoginoutTime *time.Time `xorm:"'t_token_logout_time'" json:"logoutTime,omitempty"`
	}
)

func (t Token) TableName() string {
	return "t_token"
}
