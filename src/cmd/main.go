package main

import (
	"github.com/melika0-0/bookstore-project/api"
	"github.com/melika0-0/bookstore-project/api/config"
	"github.com/melika0-0/bookstore-project/data/cache"
	"github.com/melika0-0/bookstore-project/data/cache/db"

)

func main(){
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		panic(err)
	}
	err = db.InitDB(cfg)
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()

	api.InitServer()

}