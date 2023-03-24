package server

import (
	"net/http"
	"problem1/controllers"

	"github.com/labstack/echo/v4"
)

/*
	NewServer echoの起動、ルーティング設定
*/
func NewServer() (*echo.Echo, error) {

	// echo起動
	server := echo.New()

	// ルーティング処理
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "minimal_sns_app")
	})

	friendController := controllers.NewFriendListController()
	server.GET("/get_friend_list", friendController.GetFriendList)
	server.GET("/get_friend_of_friend_list", friendController.GetFriendOfFriendList)

	// FIXME
	// server.GET("/get_friend_of_friend_list_paging", )

	return server, nil
}
