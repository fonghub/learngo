package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Func28() {
	router := gin.Default()
	router.LoadHTMLGlob("html/28htmlrendering/**/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.Run(":8080")
}
