package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
	router := gin.Default()
//	router.LoadHTMLGlob("templates/*")
	router.LoadHTMLFiles("html/27htmlrendering.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "27htmlrendering.html", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}

