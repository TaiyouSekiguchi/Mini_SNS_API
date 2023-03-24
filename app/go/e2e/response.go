package e2e

import (
	"encoding/json"
	"net/http"
	myhttp "problem1/http"
	"problem1/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

type response interface {
	CheckHTTPResponse(t *testing.T, resp *http.Response)
}

type NormalResponse struct {
	StatusCode int
	Content    string
	Result     models.FriendList
}

/*
	CheckHTTPResponse APIの返り値が期待値通りか確認
*/
func (nr *NormalResponse) CheckHTTPResponse(t *testing.T, resp *http.Response) {

	// レスポンスをデコード
	var tmpResp myhttp.Response
	err := json.NewDecoder(resp.Body).Decode(&tmpResp)
	if err != nil {
		t.Fatal(err)
	}

	// 検証パート
	assert.Equal(t, nr.StatusCode, resp.StatusCode, "status code should be OK") // ステータスコード
	assert.Equal(t, nr.Content, tmpResp.Content, "content should match")        // コンテント

	for i := range nr.Result {
		assert.Equal(t, nr.Result[i].Id, tmpResp.Result[i].Id, "id should match")       // id
		assert.Equal(t, nr.Result[i].Name, tmpResp.Result[i].Name, "name should match") // name
	}

	assert.Equal(t, len(nr.Result), tmpResp.Total, "total should match") // トータル
}

type ErrorResponse struct {
	StatusCode int
	Id         string
	Code       string
	Title      string
	Detail     string
	Info       string
}

/*
	CheckHTTPResponse APIの返り値が期待値通りか確認
*/
func (nr *ErrorResponse) CheckHTTPResponse(t *testing.T, resp *http.Response) {

	// レスポンスをデコード
	var tmpResp myhttp.ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&tmpResp)
	if err != nil {
		t.Fatal(err)
	}

	// 検証パート
	assert.Equal(t, nr.StatusCode, resp.StatusCode, "status code should be OK")      // ステータスコード
	assert.Equal(t, nr.Id, tmpResp.Error.Parameter.Id, "parameters id should be OK") // パラメタ
	assert.Equal(t, nr.Code, tmpResp.Error.Code, "error code should match")          // エラーコード
	assert.Equal(t, nr.Title, tmpResp.Error.Title, "title should match")             // タイトル
	assert.Equal(t, nr.Detail, tmpResp.Error.Detail, "title should match")           // 詳細メッセージ
	assert.Equal(t, nr.Info, tmpResp.Error.Info, "info url should match")            // Info URL
}
