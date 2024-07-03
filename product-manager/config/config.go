package config

import (
	"fmt"
	"os"

	// jwt "pesto-product-manager/authorization"
	"pesto-product-manager/database"
	"pesto-product-manager/log"
	"pesto-product-manager/server"

	"github.com/spf13/viper"
)

type Config struct {
	Server   server.Config   `json:"server"`
	Logger   log.Config      `json:"logger"`
	Database database.Config `json:"database"`
}

var Cfg Config

func SetConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(os.Getenv("PRODUCT_CONFIG_YAML"))
	// viper.AddConfigPath(DEFAULT_CONFIG_PATH) // path to look for the config file in
	// viper.AddConfigPath(".")                 // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Logger.Error("unable to decode into struct,")
		log.Logger.Error(err.Error())
	}
}
