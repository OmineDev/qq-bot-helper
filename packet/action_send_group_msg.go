package packet

type SendedMsgID struct {
	MessageID int64 `json:"message_id"`
}

// 发送群消息
type RequestActionSendGroupMessage struct {
	GroupID    int64 `json:"group_id"`
	Message    any   `json:"message"`
	AutoEscape bool  `json:"auto_escape"`
}

func (r *RequestActionSendGroupMessage) GetAction() string {
	return "send_group_msg"
}

func (r *RequestActionSendGroupMessage) GetParams() interface{} {
	return r
}
