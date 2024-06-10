package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/cmd/api"
	config "github.com/leetcode-golang-classroom/golang-ecom-sample/config"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/db"
)

func main() {
	log.Println(config.C)
	db, err := db.NewMYSQLStorage(mysql.Config{
		User:                 config.C.MYSQL_USER,
		Passwd:               config.C.MYSQL_PASSWORD,
		Addr:                 config.C.MYSQL_ADDR,
		DBName:               config.C.MYSQL_DATABASE,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)
	server := api.NewAPIServer(fmt.Sprintf(":%s", config.C.PORT), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Sucessfully Connected")
}
