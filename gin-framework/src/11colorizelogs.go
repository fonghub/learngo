package src

import (
	"github.com/gin-gonic/gin"
)

func Func11() {

	gin.ForceConsoleColor()
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
