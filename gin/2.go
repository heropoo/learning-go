package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/some-json", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "<br />",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
