package main

import (
	"database/sql"
	"fmt"
	"hola/adress"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS WALLET (id INTEGER PRIMARY KEY, privatekey TEXT,publickey TEXT)")
	checkErr(err)
	statement.Exec()

	var a adress.Wallet

	a.Genwallet(db)
	a.PrintPriv()

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
