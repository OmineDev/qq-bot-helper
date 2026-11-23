package packet

import "github.com/OmineDev/neomega-plugin-sdk/qq"

type RoleInfo = qq.RoleInfo

type GuildMemberProfile = qq.GuildMemberProfile

// 获取频道成员身份组
type RequestActionGetGuildMemberProfile struct {
	GuildID string `json:"guild_id"`
	UserID  string `json:"user_id"`
}

func (r *RequestActionGetGuildMemberProfile) GetAction() string {
	return "get_guild_member_profile"
}

func (r *RequestActionGetGuildMemberProfile) GetParams() interface{} {
	return r
}
