package e2e

import (
	"net/http"
	myhttp "problem1/http"
	"problem1/models"
	"testing"
)

/*
	TestGetFriendList /get_friend_list エンドポイント用テスト
*/
func TestGetFriendList(t *testing.T) {
	endpoint := "/get_friend_list"
	content := "friend list"

	tests := []struct {
		name     string
		userId   string
		expected response
	}{
		{
			"正常系1", // 友達がいない userId = 0
			"0",
			&NormalResponse{
				StatusCode: http.StatusOK,
				Content:    content,
				Result:     models.FriendList{},
			},
		},
		{
			"正常系2", // 友達がいる userId = 1
			"1",
			&NormalResponse{
				StatusCode: http.StatusOK,
				Content:    content,
				Result: models.FriendList{
					{Id: 2, Name: "test2"},
					{Id: 3, Name: "test3"},
					{Id: 4, Name: "test4"},
				},
			},
		},
		{
			"異常系1", // id が空文字列
			"",
			&ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Id:         "",
				Code:       myhttp.IdErrorCode,
				Title:      myhttp.InvalidRequest,
				Detail:     myhttp.IdErrorDetail,
				Info:       myhttp.InfoUrl,
			},
		},
		{
			"異常系2", // id が数値ではない文字列
			"abc",
			&ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Id:         "abc",
				Code:       myhttp.IdErrorCode,
				Title:      myhttp.InvalidRequest,
				Detail:     myhttp.IdErrorDetail,
				Info:       myhttp.InfoUrl,
			},
		},
		{
			"異常系3", // id が存在しないid
			"1000",
			&ErrorResponse{
				StatusCode: http.StatusOK,
				Id:         "1000",
				Code:       myhttp.NotFoundCode,
				Title:      myhttp.NotFound,
				Detail:     myhttp.NotFoundDetail,
				Info:       myhttp.InfoUrl,
			},
		},
		{
			"異常系4", // id が存在しないidで異常に大きい
			"99999999999999999999999999999999999999999999999999999999999999999999999",
			&ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Id:         "99999999999999999999999999999999999999999999999999999999999999999999999",
				Code:       myhttp.IdErrorCode,
				Title:      myhttp.InvalidRequest,
				Detail:     myhttp.IdErrorDetail,
				Info:       myhttp.InfoUrl,
			},
		},
		{
			"異常系5", // id が負の数
			"-42",
			&ErrorResponse{
				StatusCode: http.StatusOK,
				Id:         "-42",
				Code:       myhttp.NotFoundCode,
				Title:      myhttp.NotFound,
				Detail:     myhttp.NotFoundDetail,
				Info:       myhttp.InfoUrl,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			Story(t, endpoint, tc.userId, tc.expected)
		})
	}
}

/*
	TestGetFriendOfFriendList /get_friend_of_friend_list エンドポイント用テスト
*/
// func TestGetFriendOfFriendList(t *testing.T) {
// 	endpoint := "/get_friend_of_friend_list"
// 	content := "friend of friend list"

// 	tests := []struct {
// 		name     string
// 		userId   string
// 		expected response
// 	}{
// 		{
// 			/*
// 				userId(1)のテスト
// 				userId(1)のfriend[2, 4, 6]
// 				4はブロックのため除外
// 				userId(2)のfriend[1, 3, 5]
// 				userId(6)のfriend[1, 3, 4]
// 				候補[3, 4, 5]
// 				4はブロックのため除外
// 				6は1hopのfriendのため除外
// 				answer[3, 5]
// 			*/
// 			"正常系1",
// 			"1",
// 			&NormalResponse{
// 				StatusCode: http.StatusOK,
// 				Content:    content,
// 				Result: models.FriendList{
// 					{Id: 3, Name: "test3"},
// 					{Id: 5, Name: "test5"},
// 				},
// 			},
// 		},
// 		{
// 			/*
// 				userId(2)のテスト
// 				userId(2)のfriend[1, 3, 5]
// 				userId(1)のfriend[2, 4, 6] (4は1のブロック対象)
// 				userId(3)のfriend[2, 4, 6] (4は3のブロック対象)
// 				userId(5)のfriend[2]

// 				候補[6]
// 				1hop、ブロック該当なし
// 				answer[6]
// 			*/
// 			"正常系2",
// 			"2",
// 			&NormalResponse{
// 				StatusCode: http.StatusOK,
// 				Content:    content,
// 				Result: models.FriendList{
// 					{Id: 6, Name: "test6"},
// 				},
// 			},
// 		},
// 		{
// 			/*
// 				userId(4)のテスト
// 				userId(4)のfriend[1, 3, 6]
// 				userId(1)のfriend[2, 4, 6] (4は1のブロック対象)
// 				userId(3)のfriend[2, 4, 6] (4は3のブロック対象)
// 				userId(6)のfriend[1, 3, 4]

// 				候補[1, 2, 3, 6]
// 				1, 3, 6は1hopにより除外
// 				ブロックは該当なし
// 				answer[2]
// 			*/
// 			"正常系3",
// 			"4",
// 			&NormalResponse{
// 				StatusCode: http.StatusOK,
// 				Content:    content,
// 				Result: models.FriendList{
// 					{Id: 2, Name: "test2"},
// 				},
// 			},
// 		},
// 		{
// 			/*
// 				userId(7)のテスト
// 				userId(7)のfriend[8]
// 				userId(8)のfriend[7]

// 				answer[]
// 			*/
// 			"正常系4",
// 			"7",
// 			&NormalResponse{
// 				StatusCode: http.StatusOK,
// 				Content:    content,
// 				Result:     models.FriendList{},
// 			},
// 		},
// 		{
// 			/*
// 				userId(9)のテスト
// 				userId(9)のfriend[]

// 				answer[]
// 			*/
// 			"正常系5",
// 			"9",
// 			&NormalResponse{
// 				StatusCode: http.StatusOK,
// 				Content:    content,
// 				Result:     models.FriendList{},
// 			},
// 		},
// 		{
// 			"異常系1", // id が空文字列
// 			"",
// 			&ErrorResponse{
// 				StatusCode: http.StatusBadRequest,
// 				Id:         "",
// 				Code:       myhttp.IdErrorCode,
// 				Title:      myhttp.InvalidRequest,
// 				Detail:     myhttp.IdErrorDetail,
// 				Info:       myhttp.InfoUrl,
// 			},
// 		},
// 		{
// 			"異常系2", // id が数値ではない文字列
// 			"abc",
// 			&ErrorResponse{
// 				StatusCode: http.StatusBadRequest,
// 				Id:         "abc",
// 				Code:       myhttp.IdErrorCode,
// 				Title:      myhttp.InvalidRequest,
// 				Detail:     myhttp.IdErrorDetail,
// 				Info:       myhttp.InfoUrl,
// 			},
// 		},
// 		{
// 			"異常系3", // id が存在しないid
// 			"10",
// 			&ErrorResponse{
// 				StatusCode: http.StatusOK,
// 				Id:         "10",
// 				Code:       myhttp.NotFoundCode,
// 				Title:      myhttp.NotFound,
// 				Detail:     myhttp.NotFoundDetail,
// 				Info:       myhttp.InfoUrl,
// 			},
// 		},
// 		{
// 			"異常系4", // id が存在しないidで異常に大きい
// 			"99999999999999999999999999999999999999999999999999999999999999999999999",
// 			&ErrorResponse{
// 				StatusCode: http.StatusBadRequest,
// 				Id:         "99999999999999999999999999999999999999999999999999999999999999999999999",
// 				Code:       myhttp.IdErrorCode,
// 				Title:      myhttp.InvalidRequest,
// 				Detail:     myhttp.IdErrorDetail,
// 				Info:       myhttp.InfoUrl,
// 			},
// 		},
// 		{
// 			"異常系5", // id が負の数
// 			"-42",
// 			&ErrorResponse{
// 				StatusCode: http.StatusOK,
// 				Id:         "-42",
// 				Code:       myhttp.NotFoundCode,
// 				Title:      myhttp.NotFound,
// 				Detail:     myhttp.NotFoundDetail,
// 				Info:       myhttp.InfoUrl,
// 			},
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			Story(t, endpoint, tc.userId, tc.expected)
// 		})
// 	}
// }
