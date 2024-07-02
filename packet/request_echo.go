package packet

import "encoding/json"

const (
	CQPacketTypeRequestEcho CQPacketType = "RequestEcho"
)

type RequestEcho struct {
	Status  string          `json:"status"`
	RetCode int             `json:"retcode"`
	Msg     string          `json:"msg,omitempty"`
	Wording string          `json:"wording,omitempty"`
	Data    json.RawMessage `json:"data"`
	Echo    string          `json:"echo"`
}

func (r *RequestEcho) ID() CQPacketType {
	return CQPacketTypeRequestEcho
}

func (r *RequestEcho) Layers() map[CQPacketType]CQPacket {
	return map[CQPacketType]CQPacket{
		CQPacketTypeRequestEcho: r,
	}
}
