package src

import (
	"github.com/gin-gonic/gin"
)

func Func25() {
	router := gin.Default()

	router.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})

	//var fs http.FileSystem = // ...
	//router.GET("/fs/file", func(c *gin.Context) {
	//	c.FileFromFS("fs/file.go", fs)
	//})
}
