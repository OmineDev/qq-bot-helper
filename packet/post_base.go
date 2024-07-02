package packet

const (
	CQPacketTypePostBase CQPacketType = "PostBase"
)

type PostBase struct {
	Time     int64  `json:"time"`
	SelfID   int    `json:"self_id"`
	PostType string `json:"post_type"`
}

func (r *PostBase) ID() CQPacketType {
	return CQPacketTypeMetaPost
}

func (r *PostBase) Layers() map[CQPacketType]CQPacket {
	return map[CQPacketType]CQPacket{
		CQPacketTypeMetaPost: r,
	}
}
