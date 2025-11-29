/*
main is the entrypoint to maya
this package should be as tidy as possible containing the minimal amount of logic
we want to use other packages inside the project to handle the more complex operations
*/
package main

import (
	"github.com/System-Nebula/maya/client"
	"github.com/System-Nebula/maya/cmd"
	log "github.com/System-Nebula/maya/logger"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
)

func main() {
	log.Info("starting bot")
	c, err := client.Create()

	if err != nil {
		log.ErrorFmt("Caught error initializing client: %v", err)
		panic(err)
	}
	log.Info("Starting connection to discord API")

	commands := cmd.Commands
	_, err = c.ApplicationCommandBulkOverwrite("", "", commands)
	if err != nil {
		log.ErrorFmt("Unable to override commands %v", commands)
		return
	}
	c.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.ApplicationCommandData()
		switch data.Name {
		case "ping":
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: cmd.Ping(),
				},
			})
			if err != nil {
				log.ErrorFmt("Error while getting interaction: %v", err)
			}
		}

	})

	err = c.Open()
	if err != nil {
		log.ErrorFmt("Unable to connect to discord API: %v", err)
		return
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Info("Press Ctrl+C to exit")
	<-stop

	err = c.Close()
	if err != nil {
		panic(err)
	}
}
