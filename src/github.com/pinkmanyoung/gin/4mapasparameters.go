package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.GET("/get",func(c *gin.Context){
		ids := c.QueryMap("ids")
		c.JSON(200,ids)
	})
	router.GET("/index",func(c *gin.Context){
		c.HTML(http.StatusOK,"4mapasparameters.html",gin.H{})
	})
	router.POST("/form_post",func(c *gin.Context){
		names := c.PostFormMap("names")
		c.JSON(200,names)
	})
	router.Run(":8080")

}
