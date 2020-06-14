package services

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandService handle a command event.
func CommandService(s *discordgo.Session, m *discordgo.MessageCreate, botPrefix string) {
	// var commandMsg []string = strings.Split(m.Content, " ")
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println(err)
	}
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println(err)
	}
	if strings.HasPrefix(m.Content, botPrefix+"help") {
		MessageSender(s, m.ChannelID, "ยังทำไม่เสร็จ กำลังทำอยู่ค้าบ")
	} else if strings.HasPrefix(m.Content, botPrefix+"join") {
		connectToVoiceChannel(s, m, guild)
	} else if strings.HasPrefix(m.Content, botPrefix+"play") {

	} else {
		MessageSender(s, m.ChannelID, botPrefix+"help เพื่อดูคำสั่งทั้งหมดนะค้าบ")
	}
}

func findVoiceChannelID(guild *discordgo.Guild, m *discordgo.MessageCreate) string {
	for _, voiceState := range guild.VoiceStates {
		if voiceState.UserID == m.Author.ID {
			return voiceState.ChannelID
		}
	}
	return ""
}

func connectToVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild) {
	voiceChannelID := findVoiceChannelID(guild, m)
	if voiceChannelID == "" {
		MessageSender(s, m.ChannelID, "กรุณาเข้าห้องก่อนนะค้าบ")
	}

	_, err := s.ChannelVoiceJoin(guild.ID, voiceChannelID, false, false)

	if err != nil {
		fmt.Println(err)
	}
}