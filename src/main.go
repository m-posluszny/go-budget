package main

import (
	"fmt"

	"github.com/m-posluszny/go-ynab/src/config"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/server"
	"github.com/m-posluszny/go-ynab/src/stores"
)

func main() {
	cfg := config.Init()
	fmt.Println("DB host:", cfg.Db.Host)
	db.InitDbs(cfg.Db, cfg.Db)
	s := server.Init(cfg, stores.GetRedisSessionStore(cfg.Auth, cfg.Redis), "./src/templates/**/*")
	s.Run(cfg.Server.Host)
}
