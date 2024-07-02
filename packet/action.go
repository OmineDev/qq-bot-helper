package packet

import (
	"encoding/json"

	"github.com/google/uuid"
)

type RequestMessage struct {
	Action string `json:"action"`
	Params any    `json:"params"`
	Echo   string `json:"echo"`
}

type RequestAction interface {
	GetAction() string
	GetParams() any
}

func MakeRequestMessage(action RequestAction) (m *RequestMessage, id string) {
	id = uuid.New().String()
	return &RequestMessage{
		Action: action.GetAction(),
		Echo:   id,
		Params: action.GetParams(),
	}, id
}

func MakeRequestMessageByte(action RequestAction) (data []byte, id string, err error) {
	m, id := MakeRequestMessage(action)
	data, err = json.Marshal(m)
	return data, id, err
}
