package main

import (
    "net/http"
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/**/*")
	router.GET("index",index)
	router.GET("game/index",gameindex)
	router.GET("user/index",userindex)
	router.GET("user/add",useradd)
	router.POST("user/save",usersave)
	router.GET("someJSON",someJSON)
	router.GET("json",json)
	router.GET("purejson",purejson)
	router.Run(":8080")
}
func index(c *gin.Context) {
	c.HTML(http.StatusOK,"main/index.tmpl",gin.H{
		"title":"Main website",
		"body":"Website body.",
	})
}
func gameindex(c *gin.Context) {
	c.HTML(http.StatusOK,"game/index.tmpl",gin.H{
		"title":"game",
		"body":"game body",
	})
}
func userindex(c *gin.Context) {
	c.HTML(http.StatusOK,"user/index.tmpl",gin.H{
		"title":"user",
		"body":"user body",
	})
}
func useradd(c *gin.Context) {
	c.HTML(http.StatusOK,"user/add.tmpl",gin.H{
		"title":"user",
	})
}
func usersave(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page","0")
	name := c.PostForm("name")
	message := c.PostForm("message")
	fmt.Printf("id:%s;page:%s;name:%s;message:%s",id,page,name,message)
}

func someJSON(c *gin.Context) {
	names := []string{"lena","austin","foo"}
	c.SecureJSON(http.StatusOK,names)
}
func json(c *gin.Context) {
	c.JSON(200,gin.H{
		"html":"<b>Hello,world!</b>",
	})
}
func purejson(c *gin.Context) {
	c.PureJSON(200,gin.H{
		"html":"<b>Hello,world!</b>",
	})
}
