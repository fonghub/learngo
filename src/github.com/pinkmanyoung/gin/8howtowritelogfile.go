package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
)
func main() {
	accesslog()
    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
    router.Run(":8080")
}
// 写日志到文件
func accesslog() {
	gin.DisableConsoleColor()
	f,_ := os.Create("log/8gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}
