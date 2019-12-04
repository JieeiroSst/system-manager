package db

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func FindDataAccount()(string,string,bool){
	var id int
	sql,err:=db.Query("select Account.UserName,Account.Password from Account where id=?" ,id)
	if err!=nil{
		log.Fatal(err)
	}

	defer sql.Close()

	var username string
	var password string

	for sql.Next(){
		err:=sql.Scan(username,password)
		if err!=nil{
			log.Fatal(err)
		}
	}
	return username,password,true
}

func CreateAccount(username,password string)bool{
	sql,err:=db.Query("insert into Account(UserName,Password) values (?,?)",username,password)
	if err!=nil{
		log.Fatal("query fail")
	}
	defer sql.Close()

	return true
}

func UpdateAccount(password string)bool{
	sql,err:=db.Query("update Account set Password=? where UserName=?",password)
	if err!=nil{
		log.Fatal(err)
	}
	defer  sql.Close()

	return true
}

func DeleteAccount() bool{
	var id int
	sql,err:=db.Query("delete from Account where  id=?",id)
	if err!=nil{
		log.Fatal("no account")
	}
	defer sql.Close()
	return true
}
