package packet

type GroupMemberCard struct {
	GroupUser
	JoinTime        int32 `json:"join_time"`
	LastSentTime    int32 `json:"last_sent_time"`
	Unfriendly      bool  `json:"unfriendly"`
	TitleExpireTime int64 `json:"title_expire_time"`
	CardChangeable  bool  `json:"card_changeable"`
	ShutUpTimestamp int64 `json:"shut_up_timestamp"`
}

type GroupMemberCards []GroupMemberCard

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
