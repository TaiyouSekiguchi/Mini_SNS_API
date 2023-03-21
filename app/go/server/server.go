package server

import (
	"problem1/configs"
	"strconv"
)

// Init initialize server
func Init() error {
	conf := configs.Get()

	r, err := NewRouter()
	if err != nil {
		return err
	}
	r.Logger.Fatal(r.Start(":" + strconv.Itoa(conf.Server.Port)))
	return nil
}
