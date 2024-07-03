package server

import (
	"github.com/gin-gonic/gin"
)

func StartServer(cfg Config) {
	server := gin.Default()
	// add routes here

	// end routes
	server.Run(cfg.Host + ":" + cfg.Port)
	// log.Logger.Info("Server Started : ", zap.String("host", cfg.Host))
}
