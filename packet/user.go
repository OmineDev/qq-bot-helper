package packet

type User struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
}

type GroupUser struct {
	User
	Card    string `json:"card"`
	Area    string `json:"area"`
	Level   string `json:"level"`
	Role    string `json:"role"`
	Title   string `json:"title"`
	GroupId int64  `json:"group_id"`
}

type GuildUser struct {
	User
	TinyID string `json:"tiny_id"`
}
