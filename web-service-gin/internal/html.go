package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "main/index.tmpl", gin.H{
		"title": "Main website",
		"body":  "Website body.",
	})
}
func GameIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "game/index.tmpl", gin.H{
		"title": "game",
		"body":  "game body",
	})
}

func UserIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "user/index.tmpl", gin.H{
		"title": "user",
		"body":  "user body",
	})
}

func UserAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "user/add.tmpl", gin.H{
		"title": "user",
	})
}

func UserSave(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")
	fmt.Fprintf(c.Writer, "id:%s;page:%s;name:%s;message:%s \n", id, page, name, message)
}
