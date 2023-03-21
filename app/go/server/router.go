package server

import (
	"net/http"
	"problem1/controllers"

	"github.com/labstack/echo/v4"
)

func NewRouter() (*echo.Echo, error) {

	router := echo.New()

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "minimal_sns_app")
	})

	friendController := controllers.NewFriendController()
	router.GET("/get_friend_list", friendController.GetFriendList)

	// FIXME
	// router.GET("/get_friend_of_friend_list", )

	// FIXME
	// router.GET("/get_friend_of_friend_list_paging", )

	router.GET("/get_friend_list", friendController.GetFriendList)

	return router, nil
}
