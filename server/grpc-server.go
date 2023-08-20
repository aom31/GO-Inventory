package server

import (
	"log"
	"net"

	"github.com/aom31/GO-Inventory/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type grpcServer struct {
	cfg      *config.Config
	dbClient *mongo.Client
}

func NewGrpcServer(cfg *config.Config, dbClient *mongo.Client) *grpcServer {
	return &grpcServer{
		cfg:      cfg,
		dbClient: dbClient,
	}
}

func (serv *grpcServer) StartGrpcServer() {
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	log.Printf("successful starting grpc server:%v with app utl:%v", serv.cfg.App.AppName, serv.cfg.App.Url)

	//start
	lis, err := net.Listen("tcp", serv.cfg.App.Url)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(lis)

}
