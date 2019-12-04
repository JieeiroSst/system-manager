package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	db *sql.DB
)
func init(){
	db,err:=sql.Open("mysql","root:manhquan@/keikibook")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()

	err= db.Ping()
	if err!=nil{
		log.Fatal("connect to database fail")
	}
}

