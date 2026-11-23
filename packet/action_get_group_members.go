package packet

import "github.com/OmineDev/neomega-plugin-sdk/qq"

type GroupMemberCard = qq.GroupMemberCard

type GroupMemberCards = qq.GroupMemberCards

// 获取群成员列表
type RequestActionGetGroupMemberList struct {
	GroupID int64 `json:"group_id"`
	Nocache bool  `json:"no_cache"`
}

func (r *RequestActionGetGroupMemberList) GetAction() string {
	return "get_group_member_list"
}

func (r *RequestActionGetGroupMemberList) GetParams() interface{} {
	return r
}
