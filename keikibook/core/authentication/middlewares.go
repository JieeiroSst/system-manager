package authentication

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
)

func RequireTokenAuthentication(w http.ResponseWriter,r *http.Request,next http.HandlerFunc){
	token,err:=jwt.ParseFrom
}