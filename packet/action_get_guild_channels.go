package packet

type SlowModeInfo struct {
	SlowModeKey    int32  `json:"slow_mode_key"`
	SlowModeText   string `json:"slow_mode_text"`
	SpeakFrequency int32  `json:"speak_frequency"`
	SlowModeCircle int32  `json:"slow_mode_circle"`
}

type ChannelInfo struct {
	OwnerGuildID    string         `json:"owner_guild_id"`
	ChannelID       string         `json:"channel_id"`
	ChannelType     int32          `json:"channel_type"`
	ChannelName     string         `json:"channel_name"`
	CreateTime      int64          `json:"create_time"`
	CreatorTinyID   string         `json:"creator_tiny_id"`
	TalkPermission  int32          `json:"talk_permission"`
	VisibleType     int32          `json:"visible_type"`
	CurrentSLowMode int32          `json:"current_slow_mode"`
	SlowModes       []SlowModeInfo `json:"slow_modes"`
}

type GuildChannels []ChannelInfo

// 获取子频道列表 (指定频道-Guild)
type RequestActionGetGuildChannelList struct {
	GuildID string `json:"guild_id"`
	NoCache bool   `json:"no_cache"`
}

func (r *RequestActionGetGuildChannelList) GetAction() string {
	return "get_guild_channel_list"
}

func (r *RequestActionGetGuildChannelList) GetParams() interface{} {
	return r
}
