package controllers

import (
	"net/http"
	myhttp "problem1/http"
	"problem1/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FriendListController struct{}

func NewFriendListController() *FriendListController {
	return new(FriendListController)
}

/*
	GetFriendList userIdからそのユーザーの友達一覧を取得
*/
func (fc *FriendListController) GetFriendList(c echo.Context) error {
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

/*
	GetFriendOfFriendList userIdからそのユーザーの友達の友達一覧を取得
*/
func (fc *FriendListController) GetFriendOfFriendList(c echo.Context) error {
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

	// idのユーザーの友達の友達一覧取得
	list, err := models.GetFriendOfFriendList(id)
	if err != nil {
		c.Logger().Fatal(err)
	}

	return c.JSON(http.StatusOK, myhttp.NewResponse(
		"friend of friend list",
		list,
		len(list),
	))
}
