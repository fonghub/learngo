package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
	r := gin.Default()
	r.GET("/somejson",func(c *gin.Context){
		data := gin.H{
			"lang":"Go 语言",
			"tag":"<br>",
			"num":23,
		}
		c.AsciiJSON(http.StatusOK,data)
		c.JSON(http.StatusOK,data)
		c.PureJSON(http.StatusOK,data)
	})

	r.Run(":8080")

}
// AsciiJSON 把非ASCII码字符转义成只有ASCII字符
// JSON 把HTML字符转义成只有ASCII字符
// PureJSON 不转易字符，输出JSON格式
