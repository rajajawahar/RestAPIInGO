package main

import (
	"database/sql"
	"log"

	"gopkg.in/gorp.v1"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

}

var dbmap = initDatabase()

func initDatabase() *gorp.DbMap {
	db, error := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/")
	checkErr(error, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Employee{}, "Employee").SetKeys(true, "Id")
	error = dbmap.CreateTablesIfNotExists()
	checkErr(error, "Create tables failed")
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
