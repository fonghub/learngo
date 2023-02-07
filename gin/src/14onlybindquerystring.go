package src

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person14 struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func Func14() {
	route := gin.Default()
	route.Any("/testing", startPage14)
	route.Run(":8085")
}

func startPage14(c *gin.Context) {
	var Person14 Person14
	if c.ShouldBindQuery(&Person14) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(Person14.Name)
		log.Println(Person14.Address)
	}
	c.String(200, "Success")
}
