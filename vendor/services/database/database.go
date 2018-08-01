package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB
var err error

func Connect(Host string,dbName string)*sql.DB{
	DB , err = sql.Open(Host,"root:ilker123@/"+dbName+"?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println(err)

	}
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
	}
	return DB
}
