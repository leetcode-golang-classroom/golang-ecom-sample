package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/cmd/api"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/config"
	"github.com/leetcode-golang-classroom/golang-ecom-sample/db"
)

func main() {
	log.Println(config.C)
	dbInstance, err := db.NewMYSQLStorage(mysql.Config{
		User:                 config.C.MysqlUser,
		Passwd:               config.C.MysqlPassword,
		Addr:                 config.C.MysqlAddr,
		DBName:               config.C.MysqlDatabase,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(dbInstance)
	server := api.NewAPIServer(fmt.Sprintf(":%s", config.C.Port), dbInstance)
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
