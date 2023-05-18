package main

import (
	"web-service-gin/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//albums
	router.GET("/albums", internal.GetAlbums)
	router.GET("/albums/:id", internal.GetAlbumByID)
	router.POST("/albums/", internal.PostAlbums)

	// html
	router.LoadHTMLGlob("html/**/*")
	router.GET("index", internal.MainIndex)
	router.GET("/game/index", internal.GameIndex)
	router.GET("user/index", internal.UserIndex)
	router.GET("user/add", internal.UserAdd)
	router.POST("user/save", internal.UserSave)

	//json
	router.GET("someJSON", internal.SomeJSON)
	router.GET("json", internal.Json)
	router.GET("pureJSON", internal.PureJson)

	router.Run("localhost:8080")
}
