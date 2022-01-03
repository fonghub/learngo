package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)


func main() {
	router := gin.Default()
	router.GET("/hello",func(c *gin.Context){
		c.JSON(200,gin.H{"hello":"world"})
	})
	s := &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go s.ListenAndServe()


	router1 := gin.Default()
	router1.GET("/hello",func(c *gin.Context){
		c.JSON(200,gin.H{"hello":"world"})
	})
	s1 := &http.Server{
		Addr: ":8088",
		Handler: router1,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s1.ListenAndServe()
}

