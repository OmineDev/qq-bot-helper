package packet

const (
	CQPacketTypeHeartBeat CQPacketType = "HeartBeat"
)

type HeartBeatStatusStat struct {
	PacketReceived  int `json:"packet_received"`
	PacketSent      int `json:"packet_sent"`
	PacketLost      int `json:"packet_lost"`
	MessageReceived int `json:"message_received"`
	MessageSent     int `json:"message_sent"`
	DisconnectTimes int `json:"disconnect_times"`
	LostTimes       int `json:"lost_times"`
	LastMessageTime int `json:"last_message_time"`
}

type HeartBeatStatus struct {
	AppEnabled     bool                `json:"app_enabled"`
	AppGood        bool                `json:"app_good"`
	AppInitialized bool                `json:"app_initialized"`
	Good           bool                `json:"good"`
	Online         bool                `json:"online"`
	PluginsGood    any                 `json:"plugins_good"`
	Stat           HeartBeatStatusStat `json:"stat"`
}

type HeartBeat struct {
	MetaPost
	Status   HeartBeatStatus `json:"status"`
	Interval int64           `json:"interval"`
}

func (r *HeartBeat) ID() CQPacketType {
	return CQPacketTypeHeartBeat
}

func (r *HeartBeat) Layers() map[CQPacketType]CQPacket {
	layers := r.MetaPost.Layers()
	layers[CQPacketTypeHeartBeat] = r
	return layers
}
