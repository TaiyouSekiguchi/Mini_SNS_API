package main

import (
	"problem1/database"
	"problem1/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// データベース接続
	database.Init()
	defer database.Close()

	// サーバー起動
	if err := server.Start(); err != nil {
		panic(err)
	}
}
