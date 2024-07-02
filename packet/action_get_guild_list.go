package packet

type GuildInfo struct {
	GuildID        string `json:"guild_id"`
	GuildName      string `json:"guild_name"`
	GuildDisplayID int64  `json:"guild_display_id"`
}

type GuildList []GuildInfo

// 获取频道列表 (全部)
type RequestActionGetGuildList struct {
}

func (r *RequestActionGetGuildList) GetAction() string {
	return "get_guild_list"
}

func (r *RequestActionGetGuildList) GetParams() interface{} {
	return r
}
