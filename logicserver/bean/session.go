package bean

import ()

type (
	Session struct {
		Id           int64  `xorm:"'id'" json:"id,omitempty"`
		AppId        string `xorm:"'app_id'" json:"appId,omitempty"`
		CreateUserId string `xorm:"'create_user_id'" json:"createUserId,omitempty"`
	}
)

func (s Session) TableName() string {
	return "t_session"
}
