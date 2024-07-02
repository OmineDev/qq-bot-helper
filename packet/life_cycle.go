package packet

const (
	CQPacketTypeLifeCycle CQPacketType = "LifeCycle"
)

type LifeCycle struct {
	MetaPost
	SubType string `json:"sub_type"`
}

func (r *LifeCycle) ID() CQPacketType {
	return CQPacketTypeLifeCycle
}

func (r *LifeCycle) Layers() map[CQPacketType]CQPacket {
	layers := r.MetaPost.Layers()
	layers[CQPacketTypeLifeCycle] = r
	return layers
}
