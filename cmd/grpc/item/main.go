package main

import (
	"fmt"

	"github.com/aom31/GO-Inventory/config"
	"github.com/aom31/GO-Inventory/pkg/db"
	"github.com/aom31/GO-Inventory/server"
)

func main() {
	cfg := config.NewConfig("./.env.grpc.item")
	//connect database
	dbClient := db.DBConn(cfg)

	//start server grpc
	server.NewGrpcServer(cfg, dbClient).StartGrpcServer()

	fmt.Println(cfg, dbClient)
}
