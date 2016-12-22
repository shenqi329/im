package bean

import ()

type (
	SessionMap struct {
		Id        int64  `xorm:"'id'" json:"id,omitempty"`
		SessionId uint64 `xorm:"'session_id'" json:"sessionId,omitempty"`
		UserId    string `xorm:"'user_id'" json:"userId,omitempty"`
	}
)

func (s SessionMap) TableName() string {
	return "t_session_map"
}
