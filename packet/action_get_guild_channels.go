package packet

import "github.com/OmineDev/neomega-plugin-sdk/qq"

type SlowModeInfo = qq.SlowModeInfo

type ChannelInfo = qq.ChannelInfo

type GuildChannels = qq.GuildChannels

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
