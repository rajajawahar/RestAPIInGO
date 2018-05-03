package main

import (
	"database/sql"
	"log"

	"gopkg.in/gorp.v1"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.GET("/employees", GetEmployees)
	}

	r.Run(":8080")
}

var dbmap = initDatabase()

func initDatabase() *gorp.DbMap {
	db, error := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/")
	if error != nil {
		panic(error)
	}
	defer db.Close()

	_, error = db.Exec("CREATE DATABASE " + "SalesApp")
	if error != nil {
		panic(error)
	}
	_, error = db.Exec("USE " + "SalesApp")
	if error != nil {
		panic(error)
	}

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

type Employee struct {
	Id        int64  `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
	Salary    string `db:"salary" json:"salary"`
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func GetEmployees(c *gin.Context) {
	var employees []Employee
	_, err := dbmap.Select(&employees, "SELECT * FROM employee")

	if err == nil {
		c.JSON(200, employees)
	} else {
		c.JSON(404, gin.H{"error": "no employee(s) into the table"})
	}

	// curl -i http://localhost:8080/api/v1/users
}
