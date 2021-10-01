package main

import (
	"net/http"

	"math/rand"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type CaptchaData struct {
	Vcode int    `json:"vcode"`
	Phone string `json:"phone"`
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/login-captcha", func(c *gin.Context) {
		var captcha CaptchaData
		captcha.Vcode = rand.Intn(99999) + 100000

		captcha.Phone = c.PostForm("phone")

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
			"data":    captcha,
		})
	})

	router.Run(":8000") // listen on 0.0.0.0:8000

}
