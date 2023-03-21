package server

import (
	"problem1/configs"
	"strconv"
)

/*
	Start サーバー起動
*/
func Start() error {
	conf := configs.Get()

	s, err := NewServer()
	if err != nil {
		return err
	}
	s.Logger.Fatal(s.Start(":" + strconv.Itoa(conf.Server.Port)))
	return nil
}
