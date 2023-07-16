package main

import (
	"github.com/zakariawahyu/go-iris-news/cmd/server"
	"github.com/zakariawahyu/go-iris-news/config"
	"github.com/zakariawahyu/go-iris-news/pkg/db"
)

func main() {
	cfg := config.NewConfig()
	conn := db.NewDbConnection(cfg)

	repo := server.NewRepository(conn)
	serv := server.NewServices(repo, cfg.App.ContextTimeout)
	server.NewHandler(cfg, serv)
}
