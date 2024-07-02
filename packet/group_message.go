package packet

const (
	CQPacketTypeGroupMessage CQPacketType = "GroupMessage"
)

type Anonymous struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

type GroupMessage struct {
	Message
	Sender    GroupUser `json:"sender"`
	Anonymous Anonymous `json:"anonymous"`
	GroupID   int       `json:"group_id"`
}

func (r *GroupMessage) ID() CQPacketType {
	return CQPacketTypeGroupMessage
}

func (r *GroupMessage) Layers() map[CQPacketType]CQPacket {
	layers := r.Message.Layers()
	layers[CQPacketTypeGroupMessage] = r
	return layers
}
