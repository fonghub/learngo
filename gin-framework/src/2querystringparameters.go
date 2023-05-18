package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Func2() {
	router := gin.Default()
	// /welcome?firstname=Tom&lastname=James
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")

}
