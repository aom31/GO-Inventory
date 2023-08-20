package db

import (
	"context"
	"log"
	"time"

	"github.com/aom31/GO-Inventory/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DBConn(cfg *config.Config) *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Url))
	if err != nil {
		log.Fatalf("failed connect to mongodb:%s with url: %s", err.Error(), cfg.Db.Url)
	}

	//ping database
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("failed to ping mongodb:%s with url: %s", err.Error(), cfg.Db.Url)
	}

	log.Printf("successful connected mongodb with url: %s \n", cfg.Db.Url)
	return client
}
