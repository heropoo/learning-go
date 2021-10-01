package main

import (
	"math/rand"
)

type CaptchaData struct {
	Vcode int    `json:"vcode"`
	Phone string `json:"phone"`
}

func getLoginCaptcha(phone string) CaptchaData {
	var captcha CaptchaData
	captcha.Vcode = rand.Intn(99999) + 100000
	captcha.Phone = phone
	return captcha
}
