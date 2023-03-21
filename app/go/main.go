package main

import (
	"problem1/database"
	"problem1/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.Init()
	defer database.Close()

	if err := server.Init(); err != nil {
		panic(err)
	}
}

/*
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "minimal_sns_app")
	})

	e.GET("/get_friend_list", func(c echo.Context) error {
		// FIXME
		return nil
	})

	e.GET("/get_friend_of_friend_list", func(c echo.Context) error {
		// FIXME
		return nil
	})

	e.GET("/get_friend_of_friend_list_paging", func(c echo.Context) error {
		// FIXME
		return nil
	})
*/
