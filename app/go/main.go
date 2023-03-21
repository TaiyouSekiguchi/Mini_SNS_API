package main

import (
	"problem1/database"
	"problem1/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.Init()
	defer database.Close()

	if err := server.Init(); err != nil {
		panic(err)
	}
}
