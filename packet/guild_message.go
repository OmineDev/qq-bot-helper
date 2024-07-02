package packet

const (
	CQPacketTypeGuildMessage CQPacketType = "GuildMessage"
)

// CQ supported only, 27/12/2023
type GuildMessage struct {
	Message
	Sender    GuildUser `json:"sender"`
	GuildID   string    `json:"guild_id"`
	ChannelID string    `json:"channel_id"`
}

func (r *GuildMessage) ID() CQPacketType {
	return CQPacketTypeGuildMessage
}

func (r *GuildMessage) Layers() map[CQPacketType]CQPacket {
	layers := r.Message.Layers()
	layers[CQPacketTypeGuildMessage] = r
	return layers
}
