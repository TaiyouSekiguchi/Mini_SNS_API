package controllers

import (
	"net/http"
	"problem1/database"
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

	db := database.GetDB()

	query := database.CreateGetUserQuery(id)

	var name string
	err = db.QueryRow(query).Scan(&name)
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

	query = database.CreateFriendListQuery(id)

	rows, err := db.Query(query)
	if err != nil {
		c.Logger().Fatal(err)
	}
	defer rows.Close()

	list := make([]models.Friend, 0)

	for rows.Next() {
		var friend models.Friend
		if err := rows.Scan(&friend.ID, &friend.Name); err != nil {
			c.Logger().Fatal(err)
		}
		list = append(list, friend)
	}

	err = rows.Err()
	if err != nil {
		c.Logger().Fatal(err)
	}

	return c.JSON(http.StatusOK, myhttp.NewResponse(
		"friend list",
		list,
		len(list),
	))
}
