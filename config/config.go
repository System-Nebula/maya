package config

import (
	"errors"
	log "github.com/System-Nebula/maya/logger"
	"os"
)

// Config struct is used to propagate the config for the bot
// Config will also hold the endpoints for the other microservices in play
type Config struct {
	Token string // Token is the discord token
	Otlp  string // OTLP endpoint to send traces, metrics and logs
}

// Set creates the config loading the given envars
func (c *Config) Set() *Config {
	c.Token = os.Getenv("MAYA_TOKEN")
	c.Otlp = os.Getenv("MAYA_OTLP") // The OTLP token can be empty since we dont necessarily want to panic on not having distributed traces

	// Checking for each envar if its empty or not, if so either panic or log
	switch {
	case c.Token == "":
		log.ErrorFmt("TOKEN_NOT_SET: %v", (errors.New("MAYA_TOKEN needs to be set in order for the bot to run")))
	case c.Otlp == "":
		log.Info("MAYA_OTLP envar not set, wont send traces")
	}

	return c
}
