package database

import "fmt"

/*
	CreateGetUserQuery userIdからそのユーザー情報を取得するクエリ作成
*/
func CreateGetUserQuery(id int) string {
	query := fmt.Sprintf(
		`
			SELECT
				user_id, name
			FROM
				users
			WHERE
				user_id = %d
			;
		`, id)

	return query
}

/*
	CreateFriendListQuery userIdから特定のユーザーの友達一覧を取得するクエリ作成
*/
func CreateFriendListQuery(id int) string {
	query := fmt.Sprintf(`
		SELECT
			U.user_id AS id,
			U.name AS name
		FROM
			users AS U
		INNER JOIN friend_link AS FL
			ON U.user_id = FL.user1_id
			OR U.user_id = FL.user2_id
		WHERE
			(FL.user1_id = %d OR FL.user2_id = %d)
			AND NOT U.user_id = %d
		;
	`, id, id, id)

	return query
}
