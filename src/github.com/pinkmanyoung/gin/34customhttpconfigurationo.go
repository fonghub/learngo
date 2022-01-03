package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)


func main() {
	router := gin.Default()
	router.GET("/hello",func(c *gin.Context){
		c.JSON(200,gin.H{"hello":"world"})
	})
	http.ListenAndServe(":8080",router)
}

