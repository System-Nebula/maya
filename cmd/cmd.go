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
	}
)
