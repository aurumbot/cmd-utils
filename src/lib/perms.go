package utils

import (
	"github.com/aurumbot/lib/dat"
	f "github.com/aurumbot/lib/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"strconv"
)

func GetPerms(session *dsg.Session, message *dsg.Message) {
	s := session
	m := message

	// Fetches guild info
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

	for _, mention := range m.Mentions {
		permsMSG := "Permissions info for user " + mention.Mention() + ":\n```\n"
		multimsg := false //in case message trails over the 2k char limit

		for _, channel := range channels {
			if multimsg {
				permsMSG = "```\n"
				multimsg = false
			}
			cid := channel.ID
			perms, err := s.UserChannelPermissions(mention.ID, cid)
			if err != nil {
				dat.Log.Println(err.Error())
				// Doesn't quit function. best idea? (fix later if
				// issues arise)
			}
			permsMSG += "Channel: " + channel.Name + " (ID " + channel.ID + ") : " + strconv.Itoa(perms) + ".\n"

			if len(permsMSG) > 1900 {
				permsMSG += "```"
				s.ChannelMessageSend(m.ChannelID, permsMSG)
				permsMSG = ""
				multimsg = true
			}
		}
		if len(permsMSG) > 0 {
			permsMSG += "```"
		}
		s.ChannelMessageSend(m.ChannelID, permsMSG)
	}
}
