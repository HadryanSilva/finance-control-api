package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	initRoutes(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
