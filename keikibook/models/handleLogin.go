package models

import (
	"github.com/gorilla/securecookie"
	"keikibook/db"
	"log"
	"net/http"
)

func Login(username,password string) {
	username,password,err:=db.FindDataAccount()
	if err!=true{
		log.Fatal("error when connect account to database keikibook")
	}
}

func loginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		// .. check credentials ..
		setSession(name, response)
		Login(name,pass)
		redirectTarget = "/internal"
	}
	http.Redirect(response, request, redirectTarget, 302)
}


func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))
