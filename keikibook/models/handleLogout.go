package models

import "net/http"

func LogoutHandle(w http.ResponseWriter,r*http.Request){
	clearSession(w)
	http.Redirect(w,r,"/",http.StatusFound)
}

func clearSession(w http.ResponseWriter){
	cookie:=&http.Cookie{
		Name:       "session",
		Value:      "",
		Path:       "/",
		MaxAge:     -1,
	}
	http.SetCookie(w,cookie)
}