package packet

import "github.com/OmineDev/neomega-plugin-sdk/qq"

type User = qq.User

type GroupUser = qq.GroupUser

type GuildUser struct {
	User
	TinyID string `json:"tiny_id"`
}
