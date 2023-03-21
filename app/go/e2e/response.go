package e2e

import (
	"encoding/json"
	"net/http"
	"problem1/controllers"
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
	Result     []models.Friend
}

func (nr *NormalResponse) CheckHTTPResponse(t *testing.T, resp *http.Response) {

	// レスポンスをデコード
	var tmpResp controllers.Response
	err := json.NewDecoder(resp.Body).Decode(&tmpResp)
	if err != nil {
		t.Fatal(err)
	}

	// ステータスコード検証
	assert.Equal(t, nr.StatusCode, resp.StatusCode, "status code should be OK")

	// コンテント検証
	assert.Equal(t, nr.Content, tmpResp.Content, "content should match")

	// データ検証
	for i := range nr.Result {
		assert.Equal(t, nr.Result[i].ID, tmpResp.Result[i].ID, "id should match")
		assert.Equal(t, nr.Result[i].Name, tmpResp.Result[i].Name, "name should match")
	}

	// トータル検証
	assert.Equal(t, len(nr.Result), tmpResp.Total, "content should match")
}

type ErrorResponse struct {
	StatusCode int
	Id         string
	Code       string
	Title      string
	Detail     string
	Info       string
}

func (nr *ErrorResponse) CheckHTTPResponse(t *testing.T, resp *http.Response) {

	// レスポンスをデコード
	var tmpResp controllers.ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&tmpResp)
	if err != nil {
		t.Fatal(err)
	}

	// ステータスコード検証
	assert.Equal(t, nr.StatusCode, resp.StatusCode, "status code should be OK")

	// パラメタ検証
	assert.Equal(t, nr.Id, tmpResp.Error.Parameter.Id, "parameters id should be OK")

	// エラーコード検証
	assert.Equal(t, nr.Code, tmpResp.Error.Code, "error code should match")

	// タイトル検証
	assert.Equal(t, nr.Title, tmpResp.Error.Title, "title should match")

	// 詳細メッセージ検証
	assert.Equal(t, nr.Detail, tmpResp.Error.Detail, "title should match")

	// Info URL検証
	assert.Equal(t, nr.Info, tmpResp.Error.Info, "info url should match")
}
