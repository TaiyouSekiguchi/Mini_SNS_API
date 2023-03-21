package e2e

import (
	"net/http"
	"net/url"
	"testing"
)

// serverURL サーバーのURL
const serverURL = "http://localhost:1323"

// createRequest エンドポイントとuserIdを指定してリクエストを作成
func createRequest(endpoint, id string) (*http.Request, error) {

	// URLの作成
	u, err := url.Parse(serverURL + endpoint)
	if err != nil {
		return nil, err
	}

	// リクエスト作成
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	// クエリパラメータ付与
	q := req.URL.Query()
	q.Add("id", id)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

// Story
// 	エンドポイントとuserIdを指定してリクエストを作成。
// 	作成したリクエストに対するレスポンスを受信。
// 	レスポンスの中身が期待通りか検証。
func Story(t *testing.T, endpoint string, id string, expectedResponse response) {
	// リクエスト作成
	req, err := createRequest(endpoint, id)
	if err != nil {
		t.Fatal(err)
	}

	// リクエスト送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// レスポンス検証
	expectedResponse.CheckHTTPResponse(t, resp)
}
