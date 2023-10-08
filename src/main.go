package main

import (
	"fmt"

	"github.com/m-posluszny/go-ynab/src/config"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/server"
)

func main() {
	cfg := config.Init()
	fmt.Println("DB host:", cfg.Db.Host)
	db.InitDbs(cfg.Db, cfg.Db)
	server.Init(cfg.Server, cfg.Auth)
}
