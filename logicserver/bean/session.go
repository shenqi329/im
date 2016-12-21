package bean

import ()

type (
	Session struct {
		Id           int64  `xorm:"'t_session_id'" json:"id,omitempty"`
		AppId        string `xorm:"'t_session_app_id'" json:"appId,omitempty"`
		CreateUserId string `xorm:"'t_session_create_user_id'" json:"create_user_id,omitempty"`
	}
)

func (s Session) TableName() string {
	return "t_session"
}
