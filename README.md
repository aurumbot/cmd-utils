# cmd-utils
Extraneous utility commands for getting obscure server info

## Important:

Commands defined:
- getchannels
- getperms
- getroles

Notes:
- [x] multi-server support
- [ ] requires additional system dependencies
- [ ] creates autonomous background processes

## About

cmd-utils provides a kit of utilities for getting obscure server info
such as user's permissions per channel, list of roles and IDs, or list
of channels and IDs.

## Commands

getchannels: gets all channels in a server and their IDs, requires 
PermissionManageChannels perms
getperms: gets the perms of mentioned @users in each channel, requires 
PermissionManageRoles perms
getroles: gets the list of all the server's roles and their IDs,
requires PermissionManageChannels perms

# EOF
