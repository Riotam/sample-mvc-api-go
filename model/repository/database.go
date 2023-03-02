package repository

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func init() {
	var err error

	username := "todo-app"
	password := "todo-password"
	protocol := "tcp"
	address := "sample-mvc-api-go-db:3306"
	dbName := "todo"
	dataSourceName := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8", username, password, protocol, address, dbName)
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
}
