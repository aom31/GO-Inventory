package main

import (
	"fmt"

	"github.com/aom31/GO-Inventory/config"
)

func main() {
	cfg := config.NewConfig("./.env.http.user")
	fmt.Println(cfg)

}
