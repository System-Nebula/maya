package client

import (
	"github.com/System-Nebula/maya/config"
	log "github.com/System-Nebula/maya/logger"
	"github.com/bwmarrin/discordgo"
)

var cfg config.Config

// Create Initializes the client that is going to be used
func Create() (*discordgo.Session, error) {
	token := cfg.Set().Token
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.ErrorFmt("Unable to initialize client", err)
		return nil, err
	}
	log.Info("client succesfully created")
	return dg, nil
}
