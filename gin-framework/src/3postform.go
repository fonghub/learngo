package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Func3() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "3postform.html", gin.H{})
	})
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")

}
