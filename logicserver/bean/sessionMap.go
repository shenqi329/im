package bean

import ()

type (
	SessionMap struct {
		Id        int64  `xorm:"'t_session_map_id'" json:"id,omitempty"`
		SessionId uint64 `xorm:"'t_session_map_session_id'" json:"sessionId,omitempty"`
		UserId    string `xorm:"'t_session_map_user_id'" json:"userId,omitempty"`
	}
)

func (s SessionMap) TableName() string {
	return "t_session_map"
}
