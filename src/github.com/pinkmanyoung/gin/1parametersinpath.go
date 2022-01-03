package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/name/:name",func(c *gin.Context){
		name := c.Param("name")
		c.String(http.StatusOK,"Hello,I am %s",name)
	})

	router.GET("/name/:name/*action",func(c *gin.Context){
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK,message)
	})
	router.GET("/name/groups",func(c *gin.Context){
		c.String(http.StatusOK,"The avaiable groups art [...]")
	})
	router.Run(":8080")

}
