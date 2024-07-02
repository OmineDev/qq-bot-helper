package packet

// 发送私聊消息
type RequestActionSendPrivateMessage struct {
	UserID     int64 `json:"user_id"`
	GroupID    int64 `json:"group_id"` // go-cqhttp
	Message    any   `json:"message"`
	AutoEscape bool  `json:"auto_escape"`
}

func (r *RequestActionSendPrivateMessage) GetAction() string {
	return "send_private_msg"
}

func (r *RequestActionSendPrivateMessage) GetParams() interface{} {
	return r
}
