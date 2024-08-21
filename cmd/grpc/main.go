package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/dhuki/go-template/internal/adapter/grpc"
	v1 "github.com/dhuki/go-template/internal/adapter/grpc/v1"
	"github.com/dhuki/go-template/internal/infra/configloader"
	db "github.com/dhuki/go-template/internal/infra/database"
	"github.com/dhuki/go-template/internal/infra/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	env := os.Getenv("ENV")
	if env == "" {
		env = "LOCAL"
	}
	flag.StringVar(&configloader.Conf.App.Env, "env", env, "define environment")
	flag.Parse()

	// init config
	configloader.InitConsul(ctx, configloader.Conf.App.Env)

	// init db
	connDb := db.NewConnectionDBClient(&configloader.Conf.ConnDatabase)
	pgRepository, err := connDb.NewPgRepository()
	if err != nil {
		logger.Fatal(ctx, "postgre.ConnectDatabase", "Error connect to database postgre, err : %v", err)
		// return
	}

	// init dependency
	handler := grpc.NewHandler(pgRepository)

	// init router
	v1Server := v1.NewGRPCHandlerV1(handler, configloader.Conf.App.Port)

	idleConsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		logrus.Infof("We received an interrupt signal, shutting down service")
		v1Server.Stop()
		logger.Info(ctx, "route stop", "Success stopping service go-rest-example")
		close(idleConsClosed)
	}()

	logger.Info(ctx, "route start", "Success start service go-rest-example listening on port :%d", configloader.Conf.App.Port)
	if err := v1Server.Start(ctx); err != nil {
		logger.Fatal(ctx, "route start", "Error starting service go-rest-example, err : %v", err)
	}
	<-idleConsClosed
}
