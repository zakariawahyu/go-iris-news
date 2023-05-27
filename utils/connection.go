package utils

import (
	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zakariawahyu/go-iris-news/config"
	"log"
)

type Conn struct {
	Mysql rel.Repository
}

func NewDbConnection(cfg *config.Config) *Conn {
	return &Conn{
		Mysql: InitMysql(cfg),
	}
}

func InitMysql(cfg *config.Config) rel.Repository {
	adapter, err := mysql.Open(cfg.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	repo := rel.New(adapter)

	return repo
}
