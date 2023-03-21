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
			"正常系1", // userId = 1 の friend_list取得
			"1",
			&NormalResponse{
				StatusCode: http.StatusOK,
				Content:    content,
				Result: []models.Friend{
					{ID: 2, Name: "test2"},
					{ID: 4, Name: "test4"},
					{ID: 6, Name: "test6"},
				},
			},
		},
		{
			"正常系2", // userId = 2 の friend_list取得
			"2",
			&NormalResponse{
				StatusCode: http.StatusOK,
				Content:    content,
				Result: []models.Friend{
					{ID: 1, Name: "test1"},
					{ID: 3, Name: "test3"},
					{ID: 5, Name: "test5"},
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
			"10",
			&ErrorResponse{
				StatusCode: http.StatusOK,
				Id:         "10",
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			Story(t, endpoint, tc.userId, tc.expected)
		})
	}
}
