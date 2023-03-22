package models

import "problem1/database"

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

/*
	GetUser データベースからuserIdでユーザー情報を取得
*/
func GetUser(id int) (*User, error) {
	db := database.GetDB()
	query := database.CreateGetUserQuery(id) // クエリ作成

	var user User
	err := db.QueryRow(query).Scan(&user.Id, &user.Name) // db問い合わせ
	if err != nil {
		return nil, err
	}

	return &user, nil
}
