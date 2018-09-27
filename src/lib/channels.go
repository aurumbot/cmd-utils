package utils

import (
	"github.com/aurumbot/lib/dat"
	f "github.com/aurumbot/lib/foundation"
	dsg "github.com/bwmarrin/discordgo"
)

func GetChannels(session *dsg.Session, message *dsg.Message) {
	s := session
	m := message

	guild, err := f.GetGuild(s, m)
	if err != nil {
		dat.Log.Println(err.Error())
		return
	}
	channels, err := s.GuildChannels(guild.ID)
	if err != nil {
		dat.Log.Println(err.Error())
		return
	}

	multimsg := false
	msg := "**Channel List:**\n```"
	for _, channel := range channels {
		if multimsg {
			msg = "```\n"
			multimsg = false
		}
		msg += "Channel: " + channel.Name + " (ID :" + channel.ID + ") .\n"

		if len(msg) > 1900 {
			msg += "```"
			s.ChannelMessageSend(m.ChannelID, msg)
			msg = ""
			multimsg = true
		}
	}
	if len(msg) > 0 {
		msg += "```"
	}
	s.ChannelMessageSend(m.ChannelID, msg)
}
