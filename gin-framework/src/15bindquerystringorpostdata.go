package src

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Person15 struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func Func15() {
	route := gin.Default()
	route.GET("/testing", startPage15)
	route.Run(":8085")
}

func startPage15(c *gin.Context) {
	var Person15 Person15
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&Person15) == nil {
		log.Println(Person15.Name)
		log.Println(Person15.Address)
		log.Println(Person15.Birthday)
		log.Println(Person15.CreateTime)
		log.Println(Person15.UnixTime)
	}

	c.String(200, "Success")
}
