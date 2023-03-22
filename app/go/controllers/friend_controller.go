package controllers

import (
	"net/http"
	myhttp "problem1/http"
	"problem1/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FriendController struct{}

func NewFriendController() *FriendController {
	return new(FriendController)
}

/*
	GetFriendList userIdからそのユーザーの友達一覧を取得
*/
func (fc *FriendController) GetFriendList(c echo.Context) error {
	// クエリパラメータ
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusBadRequest,
			myhttp.NewErrorResponse(
				c.QueryParam("id"),
				myhttp.IdErrorCode,
				myhttp.InvalidRequest,
				myhttp.IdErrorDetail,
				myhttp.InfoUrl,
			),
		)
	}

	// idのユーザーが存在するか確認
	_, err = models.GetUser(id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(
			http.StatusOK,
			myhttp.NewErrorResponse(
				c.QueryParam("id"),
				myhttp.NotFoundCode,
				myhttp.NotFound,
				myhttp.NotFoundDetail,
				myhttp.InfoUrl,
			),
		)
	}

	// idのユーザーの友達一覧取得
	list, err := models.GetFriendList(id)
	if err != nil {
		c.Logger().Fatal(err)
	}

	return c.JSON(http.StatusOK, myhttp.NewResponse(
		"friend list",
		list,
		len(list),
	))
}
