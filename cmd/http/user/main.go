package main

import (
	"context"
	"fmt"

	"github.com/aom31/GO-Inventory/config"
	"github.com/aom31/GO-Inventory/pkg/db"
	"github.com/aom31/GO-Inventory/server"
)

func main() {
	cfg := config.NewConfig("./.env.http.user")
	//connect database
	dbClient := db.DBConn(cfg)
	defer dbClient.Disconnect(context.Background())

	//start server
	server.NewHttpServer(cfg, dbClient).StartHttpServer()
	fmt.Println(cfg, dbClient)

}
