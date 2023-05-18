package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SomeJSON(c *gin.Context) {
	names := []string{"lena", "austin", "foo"}
	c.SecureJSON(http.StatusOK, names)
}
func Json(c *gin.Context) {
	c.JSON(200, gin.H{
		"html": "<b>Hello,world!</b>",
	})
}

func PureJson(c *gin.Context) {
	c.PureJSON(200, gin.H{
		"html": "<b>Hello,world!</b>",
	})
}
