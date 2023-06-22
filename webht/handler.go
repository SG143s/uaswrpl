package webht

import (
	"fmt"
	"strings"

	sq "github.com/SG143s/uaswrpl/sqlop"
	str "github.com/SG143s/uaswrpl/structsn"
	"github.com/gin-gonic/gin"
)

func Webstart() {
	r := gin.Default()

	r.GET("/search/:param", func(c *gin.Context) {
		sparam := c.Param("param")
		params := strings.Split(sparam, "&1")
		fmt.Println(params)
		params[0] = strings.ReplaceAll(params[0], "key=", "")
		params[1] = strings.ReplaceAll(params[1], "min=", "")
		params[2] = strings.ReplaceAll(params[2], "max=", "")
		var res []str.Items
		res = sq.GetSearch(params[0], params[1], params[2])
		c.JSON(200, res)
	})

	r.Run(":8000")
}
