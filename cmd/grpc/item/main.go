package main

import (
	"fmt"

	"github.com/aom31/GO-Inventory/config"
)

func main() {
	cfg := config.NewConfig("./.env.grpc.item")
	fmt.Println(cfg)
}
