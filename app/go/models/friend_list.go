package models

import "problem1/database"

type FriendList []User

/*
	GetFriendList データベースからuserIdで友達一覧を取得
*/
func GetFriendList(id int) (FriendList, error) {

	db := database.GetDB()

	// クエリ作成
	query := database.CreateFriendListQuery(id)

	// db問い合わせ
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	friendList := make(FriendList, 0)

	// クエリの返り値を一行ずつ確認、変数に格納
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return nil, err
		}
		friendList = append(friendList, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return friendList, nil
}

/*
	GetFriendOfFriendList データベースからuserIdで友達の友達一覧を取得
*/
func GetFriendOfFriendList(id int) (FriendList, error) {

	db := database.GetDB()

	// クエリ作成
	query := database.CreateFriendOfFriendListQuery(id)

	// db問い合わせ
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	friendList := make(FriendList, 0)

	// クエリの返り値を一行ずつ確認、変数に格納
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return nil, err
		}
		friendList = append(friendList, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return friendList, nil
}
