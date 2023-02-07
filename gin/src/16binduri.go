package src

import "github.com/gin-gonic/gin"

type Person16 struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func Func16() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var Person16 Person16
		if err := c.ShouldBindUri(&Person16); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": Person16.Name, "uuid": Person16.ID})
	})
	route.Run(":8088")
}
