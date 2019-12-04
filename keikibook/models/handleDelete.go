package models

import (
	"keikibook/db"
	"log"
)


func Delete(){
	err:=db.DeleteAccount()
	if err!=true{
		log.Fatal("no account")
	}
}