package models

import (
	"keikibook/db"
	"log"
)

func Create(username,password string){
	err:=db.CreateAccount(username,password)
	if err!=true{
		log.Fatal("create account fail")
	}
}