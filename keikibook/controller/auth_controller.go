package controller


import (
	"keikibook/services"
	"keikibook/models"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter,r *http.Request){
	requestUser:=new(models.User)
	decoder:=json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)
}