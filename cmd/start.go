package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/rodrigoPQF/products_api/internal/config"
	"github.com/rodrigoPQF/products_api/internal/infra/database"
	"github.com/rodrigoPQF/products_api/internal/infra/http/server"
	products_usecase "github.com/rodrigoPQF/products_api/internal/usecase/products"
)


func StartApp() {
	if err := config.InitAppConfig(); err != nil {
		log.Fatal("ocurred an error on initalizing app config", err)
	}


	conn, err := database.New()
	if err != nil {
		log.Fatal("ocurred an error on initalizing database config", err)
	}

	productUsecase := products_usecase.NewProductsUseCase(conn)

	httpServer := server.NewHttpServer(&productUsecase)

	sigInterruptChannel := make(chan os.Signal, 1)
	signal.Notify(sigInterruptChannel, os.Interrupt)

	go func() {
		err = httpServer.Start(config.GetEnvs().PORT)
		if err != nil {
			fmt.Println("Http server error", err)
		}
	}()
	<-sigInterruptChannel

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	<-ctx.Done()
}
