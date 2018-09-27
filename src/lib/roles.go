package utils

import (
	"github.com/aurumbot/lib/dat"
	f "github.com/aurumbot/lib/foundation"
	dsg "github.com/bwmarrin/discordgo"
)

/* # Get server roles
* A g-d impossibility.
*
* Parameters/return values:
* This function complies with the foundation's action function protocol.
* For documentation on that, please see https://github.com/whitman-colm/discord-public
*
* TODO: Make a godoc for our nonsence.
*
* NOTE: If you print this into a discord chat, it WILL mention @everyone
 */
func GetRoles(session *dsg.Session, message *dsg.Message) {
	s := session
	m := message

	// Retrieves roles, puts them into a slice
	guild, err := f.GetGuild(s, m)
	if err != nil {
		dat.Log.Println(err.Error())
		return
	}
	roles, err := s.GuildRoles(guild.ID)
	if err != nil {
		dat.Log.Println(err.Error())
		return
	}

	// Sends out the info, this is complex in case the character count is
	// over 2000, which won't send.
	rolemsg := "Role list for server:\n```\n"
	multimsg := false //in case message trails over the 2k char limit
	for _, role := range roles {
		if multimsg {
			rolemsg = "```\n"
			multimsg = false
		}
		rolemsg += "Role: " + role.Name + ", ID: " + role.ID + "\n"

		if len(rolemsg) > 1900 {
			rolemsg += "```"
			s.ChannelMessageSend(m.ChannelID, rolemsg)
			rolemsg = ""
			multimsg = true
		}
	}
	if len(rolemsg) > 0 {
		rolemsg += "```"
		s.ChannelMessageSend(m.ChannelID, rolemsg)
	}
}
