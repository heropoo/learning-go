package main

import (
	//"log"

	"net/http"

	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/autotls"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "welcome api-demo",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/login-captcha", func(c *gin.Context) {

		phone := c.PostForm("phone")
		captcha := getLoginCaptcha(phone)

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
			"data":    captcha,
		})
	})

	router.Run(":8000") // listen on 0.0.0.0:8000
	//log.Fatal(autotls.Run(router, "demo.gotoo.ml"))

}
