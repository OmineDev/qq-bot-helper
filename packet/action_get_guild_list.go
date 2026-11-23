package packet

import "github.com/OmineDev/neomega-plugin-sdk/qq"

type GuildInfo = qq.GuildInfo

type GuildList = qq.GuildList

// 获取频道列表 (全部)
type RequestActionGetGuildList struct {
}

func (r *RequestActionGetGuildList) GetAction() string {
	return "get_guild_list"
}

func (r *RequestActionGetGuildList) GetParams() interface{} {
	return r
}
