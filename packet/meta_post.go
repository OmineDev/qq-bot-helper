package packet

const (
	CQPacketTypeMetaPost CQPacketType = "MetaPost"
)

type MetaPost struct {
	PostBase
	MetaEventType string `json:"meta_event_type"`
}

func (r *MetaPost) ID() CQPacketType {
	return CQPacketTypeMetaPost
}

func (r *MetaPost) Layers() map[CQPacketType]CQPacket {
	layers := r.PostBase.Layers()
	layers[CQPacketTypeMetaPost] = r
	return layers
}
