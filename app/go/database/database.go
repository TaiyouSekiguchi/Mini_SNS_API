package database

import (
	"database/sql"
	"problem1/configs"
)

var db *sql.DB

/*
	Init データベース接続処理
*/
func Init() {
	conf := configs.Get()

	var err error
	db, err = sql.Open(conf.DB.Driver, conf.DB.DataSource)
	if err != nil {
		panic(err)
	}
}

/*
	GetDB db変数取得
*/
func GetDB() *sql.DB {
	return db
}

/*
	Close db接続切断
*/
func Close() {
	db.Close()
}
