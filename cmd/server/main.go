package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/HellUpa/mediaboard/internal/config"
	"github.com/HellUpa/mediaboard/internal/logger"
)

func main() {
	//Load config
	cfg := config.MustLoad()

	//Setup logger
	log := logger.SetupLogger(cfg.Env)
	log.Debug("Logger setup complete")

	//Setup http server
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	err := router.Run(":" + cfg.HealthCheck.Port)
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}
}
