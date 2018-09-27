package main

import (
	"github.com/aurumbot/cmd-utils/src/lib"
	f "github.com/aurumbot/lib/foundation"
	"github.com/bwmarrin/discordgo"
)

var Commands = make(map[string]*f.Command)

func init() {
	Commands["getchannels"] = &f.Command{
		Name:    "Get a List of Channels and Their IDs",
		Help:    "",
		Perms:   discordgo.PermissionManageChannels,
		Action:  utils.GetChannels,
		Version: "v2.0",
	}
	Commands["getperms"] = &f.Command{
		Name: "Get Permissions for Users (per channel)",
		Help: `Gets the value of permission integer for users in each of the server's channels.
The permissions are set as 53-bit integers calculated using bitwise operations.
For more info see https://discordapp.com/developers/docs/topics/permissions and
https://discordapi.com/permissions.html.
User mentions should be passed as arguments. Multiple users at a time are supported.
Usage: getperms @someuser @otheruser`,
		Perms:   discordgo.PermissionManageRoles,
		Version: "v2.0",
		Action:  utils.GetPerms,
	}
	Commands["getroles"] = &f.Command{
		Name:    "Get Server Roles",
		Help:    "Goes through all of the server's roles and posts them and their IDs.",
		Perms:   discordgo.PermissionManageChannels,
		Version: "v1.0",
		Action:  utils.GetRoles,
	}
}
