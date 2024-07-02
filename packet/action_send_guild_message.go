package packet

// 发送频道消息
type RequestActionSendGuildMessage struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	Message   any    `json:"message"`
}

func (r *RequestActionSendGuildMessage) GetAction() string {
	return "send_guild_channel_msg"
}

func (r *RequestActionSendGuildMessage) GetParams() interface{} {
	return r
}
