package main

import (
	"pesto-product-manager/config"
	"pesto-product-manager/database"
	"pesto-product-manager/log"
	"pesto-product-manager/server"
)

var DEFAULT_CONFIG_PATH = "$HOME/.pesto"

func main() {
	config.SetConfig()
	log.CreateLogger(config.Cfg.Logger)
	log.Logger.Debug("Logger setup complete")
	// jwt.Init(config.Cfg.JWT)
	database.Init(config.Cfg.Database)
	defer database.HandleDBclose()
	server.StartServer(config.Cfg.Server)
}
