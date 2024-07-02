package packet

type RoleInfo struct {
	RoleID   string `json:"role_id"`
	RoleName string `json:"role_name"`
}

type GuildMemberProfile struct {
	TinyID    string     `json:"tiny_id"`
	Nickname  string     `json:"nickname"`
	AvatarURL string     `json:"avatar_url"`
	JoinTime  int64      `json:"join_time"`
	Roles     []RoleInfo `json:"roles"`
}

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
