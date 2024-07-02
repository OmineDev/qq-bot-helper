package packet

const (
	CQPacketTypeMessage = "Message"
)

type Message struct {
	PostBase
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int32  `json:"message_id"`
	UserID      int64  `json:"user_id"`
	Message     any    `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        int32  `json:"font"`
	PeerID      int64  `json:"peer_id"` // Shamrock
}

func (r *Message) ID() CQPacketType {
	return CQPacketTypeMessage
}

func (r *Message) Layers() map[CQPacketType]CQPacket {
	layers := r.PostBase.Layers()
	layers[CQPacketTypeMessage] = r
	return layers
}
