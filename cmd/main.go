package main

import (
	"github.com/zakariawahyu/go-iris-news/cmd/server"
	"github.com/zakariawahyu/go-iris-news/config"
	"github.com/zakariawahyu/go-iris-news/utils"
)

func main() {
	cfg := config.NewConfig()
	db := utils.NewDbConnection(cfg)

	repo := server.NewRepository(db)
	serv := server.NewServices(repo, cfg.App.ContextTimeout)
	server.NewHandler(cfg, serv)
}
