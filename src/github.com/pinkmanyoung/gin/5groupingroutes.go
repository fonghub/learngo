package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.GET("/login",func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{
				"status":"get ok",
			})
		})
		v1.GET("/logout",func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{
				"status":"logout ok",
			})
		})
	}
	router.Run(":8080")

}
