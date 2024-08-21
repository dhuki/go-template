package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	httpHandler "github.com/dhuki/go-template/internal/adapter/http"
	v1 "github.com/dhuki/go-template/internal/adapter/http/v1"
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

	// init database
	connDb := db.NewConnectionDBClient(&configloader.Conf.ConnDatabase)
	pgRepository, err := connDb.NewPgRepository()
	if err != nil {
		logger.Fatal(ctx, "postgre.ConnectDatabase", "Error connect to database postgre, err : %v", err)
		// return
	}
	// mysqlRepository, err := connDb.NewMySQLRepository()
	// if err != nil {
	// 	logger.Fatal(ctx, "mysql.ConnectDatabase", "Error connect to database mysql, err : %v", err)
	// 	// return
	// }

	// init dependency
	handler := httpHandler.NewHandler(pgRepository, nil)

	// init router
	v1Server := v1.NewHttpHandlerV1(handler, configloader.Conf.App.Port)

	// run server v1
	idleConsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		logrus.Infof("We received an interrupt signal, shutting down service")
		if err := v1Server.Stop(ctx); err != nil {
			logger.Fatal(ctx, "route stop", "Error stopping service go-date, err : %v", err)
		}
		logger.Info(ctx, "route stop", "Success stopping service go-rest-example")
		close(idleConsClosed)
	}()

	logger.Info(ctx, "route start", "Success start service go-rest-example listening on port :%d", configloader.Conf.App.Port)
	if err := v1Server.Start(ctx); err != nil {
		logger.Fatal(ctx, "route start", "Error starting service go-rest-example, err : %v", err)
	}
	<-idleConsClosed

}
