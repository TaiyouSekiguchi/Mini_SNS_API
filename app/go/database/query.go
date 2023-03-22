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

func CreateFriendOfFriendListQuery(id int) string {
	query := fmt.Sprintf(`
		WITH friend_list AS (
			SELECT
				U.user_id AS id,
				U.name AS name,
				CASE
						WHEN U.user_id = FL.user1_id THEN FL.user2_id
						ELSE FL.user1_id
				END AS friend_id_1hop,
				BL.user1_id AS block_src,
				BL.user2_id AS block_dst
			FROM
				users AS U
				INNER JOIN friend_link AS FL
					ON U.user_id = FL.user1_id
					OR U.user_id = FL.user2_id
				LEFT OUTER JOIN block_list AS BL
					ON U.user_id = BL.user1_id
					AND (
						FL.user1_id = BL.user2_id
						OR FL.user2_id = BL.user2_id
					)
			WHERE
				BL.user1_id IS NULL
				AND BL.user2_id IS NULL
		)

		, friend_of_friend_list AS (
			SELECT
				FL1.id AS id,
				FL1.name AS name,
				FL1.friend_id_1hop AS friend_id_1hop,
				FL2.friend_id_1hop AS friend_id_2hop,
				BL.user1_id AS block_src,
				BL.user2_id AS block_dst
			FROM
				friend_list AS FL1
				INNER JOIN friend_list AS FL2
					ON FL1.friend_id_1hop = FL2.id
				LEFT OUTER JOIN block_list AS BL
					ON FL1.id = BL.user1_id
					AND FL2.friend_id_1hop = BL.user2_id
			WHERE
				FL1.id != FL2.friend_id_1hop
				AND (
					BL.user1_id IS NULL
					AND BL.user2_id IS NULL
				)
			ORDER BY
				id, friend_id_1hop, friend_id_2hop
		)

		, friend_of_friend_2hop AS (
			SELECT
				friend_id_2hop
			FROM
				friend_of_friend_list
			WHERE
				id = %d
				AND NOT friend_id_2hop IN (
					SELECT
						friend_id_1hop
					FROM
						friend_of_friend_list
					WHERE
						id = %d
				)
		)

		SELECT
			id,
			name
		FROM
			users
		WHERE
			id IN (
				SELECT
					friend_id_2hop
				FROM
					friend_of_friend_2hop
			)
		ORDER BY
			id
	;
	`, id, id)

	return query
}
