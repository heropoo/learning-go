package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	r.Run(":8000")
}
