package main

import (
	"fmt"

	"github.com/aom31/GO-Inventory/config"
	"github.com/aom31/GO-Inventory/pkg/db"
)

func main() {
	cfg := config.NewConfig("./.env.http.item")

	//connect database
	dbClient := db.DBConn(cfg)

	fmt.Println(cfg, dbClient)
}
