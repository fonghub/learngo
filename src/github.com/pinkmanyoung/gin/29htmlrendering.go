package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
	router := gin.Default()
	router.Delims("{{{","}}}")
	router.LoadHTMLFiles("html/29htmlrendering.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "29htmlrendering.html", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}

