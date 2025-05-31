package main

import (
	"github.com/rratchapol/isekai-shop-api/config"
	"github.com/rratchapol/isekai-shop-api/databases"
	"github.com/rratchapol/isekai-shop-api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectionGetting())

	server.Start()

}
