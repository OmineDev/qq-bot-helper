package main

import (
	"context"
	"fmt"
	"time"

	cqhttp_helper "github.com/OmineDev/qq-bot-helper"
	"github.com/OmineDev/qq-bot-helper/packet"
)

func main() {
	client := cqhttp_helper.NewCQHttpClient("localhost:6701", "", func(s string) {
		fmt.Println(s)
	})
	client.StartRoutine(context.Background())
	client.RegisterPacketNoBlockCB(func(pk packet.CQPacket, data []byte) {
		if pk.ID() == packet.CQPacketTypeRequestEcho {
			return
		}
		fmt.Println(pk)
	})

	// send group message
	client.SendGroupMessage(548589654, fmt.Sprintf("hello: %v", time.Now()), func(ok bool, msgID int64) {
		fmt.Println(ok)
	})

	// get group member cards
	client.GetGroupMember(548589654, func(cards packet.GroupMemberCards) {
		for _, card := range cards {
			fmt.Printf("card: %v\n", card)
		}
	})

	// get guild info
	client.GetGuildList(func(guilds packet.GuildList) {
		for _, guild := range guilds {
			fmt.Printf("guild: displayID=%v, name=%v, guildID=%v\n", guild.GuildDisplayID, guild.GuildName, guild.GuildID)
			// displayID=4294649252, name=FastBuilder, guildID=95958611651917951
		}
	})

	// get guild channel list 95958611651917951->FastBuilder
	client.GetGuildChannels("95958611651917951", func(channels packet.GuildChannels) {
		for _, channel := range channels {
			fmt.Printf("channel: type=%v, name=%v, channelID=%v\n", channel.ChannelType, channel.ChannelName, channel.ChannelID)
		}
	})

	// get guild member profile
	client.GetGuildMemberProfile("95958611651917951", "144115218680105518", func(profile packet.GuildMemberProfile) {
		fmt.Printf("profile: %v\n", profile)
	})

	// send guild message
	client.SendGuildMessage("95958611651917951", "580655711", "啊巴", func(ok bool, msgID int64) {
		fmt.Printf("ok=%v, msgID=%v\n", ok, msgID)
	})

	time.Sleep(time.Hour)
}
