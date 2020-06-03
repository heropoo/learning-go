package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("SESSIONID")
		if err != nil {
			//fmt.Println("XERR:", err)

			http.SetCookie(w, &http.Cookie{
				Name:     "SESSIONID",
				Value:    url.QueryEscape("123"),
				MaxAge:   3600,
				Path:     "/",
				Domain:   "localhost",
				SameSite: 0,
				Secure:   false,
				HttpOnly: true,
			})
		} else {
			fmt.Println("sessionid:", sessionCookie.Value)
			sessionCookie.Name = "SESSIONID"
			sessionCookie.Value = "123123"
			sessionCookie.HttpOnly = true
			sessionCookie.Domain = "localhost"
			sessionCookie.MaxAge = 3600
			sessionCookie.Path = "/"
			sessionCookie.SameSite = 0
			sessionCookie.Secure = false
			http.SetCookie(w, sessionCookie)
		}

		w.Write([]byte("hello 世界"))
	})

	http.ListenAndServe("localhost:8080", nil)
}
