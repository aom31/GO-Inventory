package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aom31/GO-Inventory/config"
	httphandler "github.com/aom31/GO-Inventory/handler/httpHandler"
	"github.com/aom31/GO-Inventory/src/repository"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type IHttpServer interface {

	//StartHttpServer is func to start server app with echo and config
	StartHttpServer()
}

type httpServer struct {
	app      *echo.Echo
	cfg      *config.Config
	dbClient *mongo.Client
}

// constructure httpServer struct
func NewHttpServer(cfg *config.Config, dbClient *mongo.Client) IHttpServer {
	return &httpServer{
		app:      echo.New(),
		cfg:      cfg,
		dbClient: dbClient,
	}
}

func (serv *httpServer) StartHttpServer() {
	//setup
	serv.RouteHttpHandle()

	log.Printf("starting server http with app url: %s", serv.cfg.App.Url)

	// Start server
	go func() {
		if err := serv.app.Start(serv.cfg.App.Url); err != nil && err != http.ErrServerClosed {
			serv.app.Logger.Fatal("shutting down the server")
		}
	}()

	//when shutdown server from any interupt , will recovery resouce
	serv.graceFullShutdown()

}
func (serv *httpServer) graceFullShutdown() {
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := serv.app.Shutdown(ctx); err != nil {
		serv.app.Logger.Fatal(err)
	}
}
func (serv *httpServer) RouteHttpHandle() {
	//init handler
	userHandler := &httphandler.UserHttpHandler{
		Cfg: serv.cfg,
		UserRepository: &repository.UserRepository{
			Client: serv.dbClient,
		},
	}

	//init router
	serv.app.GET("/api/user/:userId", userHandler.FindOneUser)
}
