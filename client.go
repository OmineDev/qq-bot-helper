package cqhttp_helper

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/OmineDev/qq-bot-helper/packet"
)

type CQHttpClient struct {
	conn             *CQHttpConnection
	cbs              sync.Map
	packetNoBlockCBs []func(packet.CQPacket, []byte)
}

func NewCQHttpClient(cqHttpAddr, cqHttpBearerToken string, log func(string)) *CQHttpClient {
	c := &CQHttpClient{
		conn:             NewCQHttpConnection(cqHttpAddr, cqHttpBearerToken, log),
		cbs:              sync.Map{},
		packetNoBlockCBs: []func(packet.CQPacket, []byte){},
	}
	return c
}

func (c *CQHttpClient) RegisterPacketNoBlockCB(cb func(pk packet.CQPacket, data []byte)) {
	c.packetNoBlockCBs = append(c.packetNoBlockCBs, cb)
}

func (c *CQHttpClient) StartRoutine(ctx context.Context) {
	c.conn.startRoutine(ctx)
	c.conn.RegisterReaderNoBlockCB(c.onMsg)
}

func (c *CQHttpClient) setCB(cbID string, cb func(*packet.RequestEcho, []byte)) {
	c.cbs.Store(cbID, cb)
}

func (c *CQHttpClient) SendGroupMessage(groupID int64, message string, onCb func(ok bool, msgID int64)) {
	act := packet.RequestActionSendGroupMessage{
		GroupID: groupID,
		Message: message,
	}
	data, id, err := packet.MakeRequestMessageByte(&act)
	if err != nil {
		panic("should not happen")
	}
	if onCb != nil {
		c.setCB(id, func(re *packet.RequestEcho, data []byte) {
			var msgID packet.SendedMsgID
			json.Unmarshal([]byte(re.Data), &msgID)
			go onCb(re.Status == "ok", msgID.MessageID)
		})
	}
	c.conn.SendNoBlock(data)
}

func (c *CQHttpClient) SendPrivateMessage(userID int64, message string, onCb func(ok bool, msgID int64)) {
	act := packet.RequestActionSendPrivateMessage{
		UserID:  userID,
		Message: message,
	}
	data, id, err := packet.MakeRequestMessageByte(&act)
	if err != nil {
		panic("should not happen")
	}
	if onCb != nil {
		c.setCB(id, func(re *packet.RequestEcho, data []byte) {
			var msgID packet.SendedMsgID
			json.Unmarshal([]byte(re.Data), &msgID)
			go onCb(re.Status == "ok", msgID.MessageID)
		})
	}
	c.conn.SendNoBlock(data)
}

func (c *CQHttpClient) GetGroupMember(groupID int64, onCb func(packet.GroupMemberCards)) {
	act := packet.RequestActionGetGroupMemberList{
		GroupID: groupID,
		Nocache: true,
	}
	data, id, err := packet.MakeRequestMessageByte(&act)
	if err != nil {
		panic("should not happen")
	}
	if onCb != nil {
		c.setCB(id, func(re *packet.RequestEcho, data []byte) {
			retMsg := re.Data
			var cards packet.GroupMemberCards
			json.Unmarshal([]byte(retMsg), &cards)
			go onCb(cards)
		})
	}
	c.conn.SendNoBlock(data)
}

func (c *CQHttpClient) GetGuildList(onCb func(guilds packet.GuildList)) {
	act := packet.RequestActionGetGuildList{}
	data, id, err := packet.MakeRequestMessageByte(&act)
	if err != nil {
		panic("should not happen")
	}
	if onCb != nil {
		c.setCB(id, func(re *packet.RequestEcho, data []byte) {
			retMsg := re.Data
			var guilds packet.GuildList
			json.Unmarshal([]byte(retMsg), &guilds)
			go onCb(guilds)
		})
	}
	c.conn.SendNoBlock(data)
}

func (c *CQHttpClient) GetGuildChannels(guildID string, onCb func(channels packet.GuildChannels)) {
	act := packet.RequestActionGetGuildChannelList{
		GuildID: guildID,
	}
	data, id, err := packet.MakeRequestMessageByte(&act)
	if err != nil {
		panic("should not happen")
	}
	if onCb != nil {
		c.setCB(id, func(re *packet.RequestEcho, data []byte) {
			retMsg := re.Data
			var channels packet.GuildChannels
			json.Unmarshal([]byte(retMsg), &channels)
			go onCb(channels)
		})
	}
	c.conn.SendNoBlock(data)
}

func (c *CQHttpClient) GetGuildMemberProfile(guildID, userID string, onCB func(packet.GuildMemberProfile)) {
	act := packet.RequestActionGetGuildMemberProfile{
		GuildID: guildID,
		UserID:  userID,
	}
	data, id, err := packet.MakeRequestMessageByte(&act)
	if err != nil {
		panic("should not happen")
	}
	if onCB != nil {
		c.setCB(id, func(re *packet.RequestEcho, data []byte) {
			retMsg := re.Data
			var profile packet.GuildMemberProfile
			// fmt.Println(string(retMsg))
			json.Unmarshal([]byte(retMsg), &profile)
			go onCB(profile)
		})
	}
	c.conn.SendNoBlock(data)
}

func (c *CQHttpClient) SendGuildMessage(guildID, channelID string, message string, onCb func(ok bool, msgID int64)) {
	act := packet.RequestActionSendGuildMessage{
		GuildID:   guildID,
		ChannelID: channelID,
		Message:   message,
	}
	data, id, err := packet.MakeRequestMessageByte(&act)
	if err != nil {
		panic("should not happen")
	}
	if onCb != nil {
		c.setCB(id, func(re *packet.RequestEcho, data []byte) {
			var msgID packet.SendedMsgID
			json.Unmarshal([]byte(re.Data), &msgID)
			go onCb(re.Status == "ok", msgID.MessageID)
		})
	}
	c.conn.SendNoBlock(data)
}

func (c *CQHttpClient) onMsg(data []byte) {
	pk, err := packet.Parse(data)
	if err != nil {
		return
	}
	switch pk.ID() {
	case packet.CQPacketTypeRequestEcho:
		p := pk.(*packet.RequestEcho)
		cbID := p.Echo
		cb, ok := c.cbs.Load(cbID)
		if ok {
			go cb.(func(*packet.RequestEcho, []byte))(p, data)
		}
	}
	for _, cb := range c.packetNoBlockCBs {
		cb(pk, data)
	}
}
