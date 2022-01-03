package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
	r := gin.Default()
	r.GET("jsonp",func(c *gin.Context){
		data := gin.H{
			"foo":"bar",
		}
		c.JSONP(http.StatusOK,data)
	})



	r.Run(":8080")

}
// curl http://127.0.0.1:8080/jsonp?callback=x

// x({"foo":"bar"})
