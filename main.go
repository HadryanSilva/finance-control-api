package main

import (
	"github.com/HadryanSilva/finance-control-api/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errf("Error initializing config: %v", err)
		return
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
