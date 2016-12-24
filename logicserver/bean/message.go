package bean

import (
	"time"
)

type (
	Message struct {
		Id         string    `xorm:"'id'" json:"id,omitempty"`
		SessionId  uint64    `xorm:"'session_id'" json:"sessionId,omitempty"`
		UserId     string    `xorm:"'user_id'" json:"userId,omitempty"`
		Type       int       `xorm:"'type'" json:"type,omitempty"`
		Content    string    `xorm:"'content'" json:"content,omitempty"`
		SyncKey    uint64    `xorm:"'sync_key'" json:"syncKey,omitempty"`
		CreateTime time.Time `xorm:"'create_time'" json:"createTime,omitempty"`
	}
)

func (t Message) TableName() string {
	return "t_message"
}
