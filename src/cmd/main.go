package main

import (
	"log"

	"github.com/melika0-0/bookstore-project/api"
	"github.com/melika0-0/bookstore-project/api/config"
	"github.com/melika0-0/bookstore-project/data/cache"
	"github.com/melika0-0/bookstore-project/data/cache/db"
)




func main() {

	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal("Redis init failed:", err)
	}
	defer cache.CloseRedis()

	err = db.InitDB(cfg)
	if err != nil {
		log.Fatal("DB init failed:", err)
	}
	defer db.CloseDB()

	api.InitServer()
}



//infrastructure bootstrapper (Config, Redis, Database).
//This keeps your network protocol layer completely isolated from your database layer we handle rest in init serve api as manager 
//api owns the web layer