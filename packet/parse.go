package packet

import (
	"encoding/json"
)

type CQPacketType string

type CQPacket interface {
	ID() CQPacketType
	Layers() map[CQPacketType]CQPacket
}

func Parse(data []byte) (pk CQPacket, err error) {
	var echo *RequestEcho
	err = json.Unmarshal(data, &echo)
	if err == nil {
		if echo.Echo != "" {
			return echo, nil
		}
	}
	var meta *MetaPost
	err = json.Unmarshal(data, &meta)
	if err != nil {
		return nil, err
	}
	switch meta.PostType {
	case "meta_event":
		switch meta.MetaEventType {
		case "heartbeat":
			var heartbeat *HeartBeat
			err = json.Unmarshal(data, &heartbeat)
			return heartbeat, err
		case "lifecycle":
			var lifeCycle *LifeCycle
			err = json.Unmarshal(data, &lifeCycle)
			return lifeCycle, err
		}
	case "message":
		var message *Message
		err = json.Unmarshal(data, &message)
		if err != nil {
			return nil, err
		}
		switch message.MessageType {
		case "private":
			var privateMessage *PrivateMessage
			err = json.Unmarshal(data, &privateMessage)
			return privateMessage, err
		case "group":
			var groupMessage *GroupMessage
			err = json.Unmarshal(data, &groupMessage)
			return groupMessage, err
		case "guild":
			var guildMessage *GuildMessage
			err = json.Unmarshal(data, &guildMessage)
			return guildMessage, err
		default:
			return message, nil
		}
	default:
		return meta, nil
	}
	return nil, nil
}
