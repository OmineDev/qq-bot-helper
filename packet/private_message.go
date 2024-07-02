package packet

const (
	CQPacketTypePrivateMessage CQPacketType = "PrivateMessage"
)

type PrivateMessage struct {
	Message
	Sender     User  `json:"sender"`
	TargetID   int64 `json:"target_id"`
	TempSource int   `json:"temp_source"`
}

func (r *PrivateMessage) ID() CQPacketType {
	return CQPacketTypePrivateMessage
}

func (r *PrivateMessage) Layers() map[CQPacketType]CQPacket {
	layers := r.Message.Layers()
	layers[CQPacketTypePrivateMessage] = r
	return layers
}
