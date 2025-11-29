// Package cmd contains all the logic for the discord bot commands
package cmd

import "github.com/bwmarrin/discordgo"

var (
	// Commands global command slice that registers all the available commands
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "healthcheck, replies pong to the user",
		},
		{
			// This command should have several subcommand maybe?
			Name:        "play",
			Description: "Streams music from a given url",
		},
		{
			// This command should have several subcommand maybe?
			Name:        "comfy",
			Description: "Genearate media from comfyui",
		},
	}
)

// AppendCommands just appends commands to the Commands slice
func AppendCommands(cmds []*discordgo.ApplicationCommand, newItems []*discordgo.ApplicationCommand) []*discordgo.ApplicationCommand {
	cmds = append(cmds, newItems...)
	return cmds
}

/*
/play <url>
/queue list -> this lists the songs in the queue
/queue add <url> -> this adds a new a new song to the queue
/queue remove <song/id> -> this removes a song or playlist from the queue
*/
