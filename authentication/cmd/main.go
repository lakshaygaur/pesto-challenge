package main

import (
	"fmt"
	jwt "pesto-auth/authorization"
	"pesto-auth/config"
	"pesto-auth/log"
	"pesto-auth/server"
)

var DEFAULT_CONFIG_PATH = "$HOME/.pesto"

func main() {
	config.SetConfig()
	log.CreateLogger(config.Cfg.Logger)
	fmt.Println("config", config.Cfg)
	log.Logger.Debug("Logger setup complete")
	jwt.Init(config.Cfg.JWT)
	server.StartServer(config.Cfg.Server)
}
