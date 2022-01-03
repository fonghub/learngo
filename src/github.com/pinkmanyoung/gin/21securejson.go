package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
	r := gin.Default()
	// while(1);["lena","austin","foo"]
	r.GET("/someJSONarr",func(c *gin.Context){
		names := []string{"lena","austin","foo"}
		c.SecureJSON(http.StatusOK,names)
	})
	// {"0":"tody","1":"andy"}
	r.GET("/someJSONmap",func(c *gin.Context){
		m := map[int]string{0:"tody",1:"andy"}
		c.SecureJSON(http.StatusOK,m)
	})
	r.Run(":8080")


}
