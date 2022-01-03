package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
	router := gin.Default()
// 请求重定向，支持内部和外部重定向
	router.GET("/test",func(c *gin.Context){
		c.Redirect(http.StatusMovedPermanently,"https://www.google.com")
	})
// 支持POST重定向
	router.POST("/test",func(c *gin.Context){
		c.Redirect(http.StatusMovedPermanently,"https://www.google.com")
	})
// 支持路由重定向
	router.GET("/test2",func(c *gin.Context){
		c.Request.URL.Path = "/test3"
		router.HandleContext(c)
	})
	router.GET("/test3",func(c *gin.Context){
		c.JSON(200,gin.H{"Hello":"world"})
	})
	router.Run(":8080")
}

