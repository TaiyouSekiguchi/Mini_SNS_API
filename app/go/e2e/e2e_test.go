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
					{Id: 5, Name: "test5"},
					{Id: 6, Name: "test6"},
					{Id: 7, Name: "test7"},
					{Id: 8, Name: "test8"},
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
func TestGetFriendOfFriendList(t *testing.T) {
	endpoint := "/get_friend_of_friend_list"
	content := "friend of friend list"

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
			"正常系2", // 友達はいるけど、その友達に友達がいない userId = 1
			"1",
			&NormalResponse{
				StatusCode: http.StatusOK,
				Content:    content,
				Result:     models.FriendList{},
			},
		},
		{
			"正常系3", // 友達がいて、その友達に友達がいる userId = 5
			"5",
			&NormalResponse{
				StatusCode: http.StatusOK,
				Content:    content,
				Result: models.FriendList{
					{Id: 2, Name: "test2"},
					{Id: 3, Name: "test3"},
					{Id: 4, Name: "test4"},
					{Id: 6, Name: "test6"},
					{Id: 7, Name: "test7"},
					{Id: 8, Name: "test8"},
				},
			},
		},
		{
			"正常系4", // 友達がいて、その友達に友達がいるけど、1hopの友達がいる userId = 6
			"6",
			&NormalResponse{
				StatusCode: http.StatusOK,
				Content:    content,
				Result: models.FriendList{
					{Id: 4, Name: "test4"},
					{Id: 5, Name: "test5"},
					{Id: 7, Name: "test7"},
					{Id: 8, Name: "test8"},
				},
			},
		},
		{
			"正常系5", // 友達がいて、その友達に友達がいるけど、1hopの友達がブロック対象 userId = 7
			"7",
			&NormalResponse{
				StatusCode: http.StatusOK,
				Content:    content,
				Result:     models.FriendList{},
			},
		},
		{
			"正常系6", // 友達がいて、その友達に友達がいるけど、2hopの友達がブロック対象 userId = 8
			"8",
			&NormalResponse{
				StatusCode: http.StatusOK,
				Content:    content,
				Result: models.FriendList{
					{Id: 2, Name: "test2"},
					{Id: 3, Name: "test3"},
					{Id: 4, Name: "test4"},
					{Id: 5, Name: "test5"},
					{Id: 6, Name: "test6"},
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
