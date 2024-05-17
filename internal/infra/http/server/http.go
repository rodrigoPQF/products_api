package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rodrigoPQF/products_api/internal/infra/http/routes"
	"github.com/rodrigoPQF/products_api/internal/usecase/products"
)

type HttpServerInterface interface {
	Shutdown(ctx context.Context) error
	Start(port string) error
}
type HttpServer struct {
	app           *echo.Echo
	products  *products.ProductsUseCase

}

func NewHttpServer(products *products.ProductsUseCase) HttpServerInterface {
	return &HttpServer{
		app:           echo.New(),
		products: products,
	}
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.app.Shutdown(ctx)
}

func (s *HttpServer) Start(port string) error {
	s.app.Use(middleware.Logger())
	s.app.Use(middleware.CORS())

	routes.NewRouteProvider(s.app, s.products).Routes()

	return s.app.Start(fmt.Sprintf(":%s", port))
}
