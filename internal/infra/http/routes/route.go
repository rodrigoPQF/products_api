package routes

import (
	"github.com/labstack/echo/v4"
	handlers "github.com/rodrigoPQF/products_api/internal/infra/http/handlers/products"
	usecase "github.com/rodrigoPQF/products_api/internal/usecase/products"
)

type productsRoutes struct {
	products  *usecase.ProductsUseCase
	app           *echo.Echo
}

func NewRouteProvider(server *echo.Echo, products *usecase.ProductsUseCase) *productsRoutes {
	return &productsRoutes{
		products:  products,
		app:           server,
	}
}

func (pr *productsRoutes) Routes() {

	productHandler := handlers.NewProductsHandler(*pr.products)

	productRouter := pr.app.Group("/v1/products")
	productRouter.POST("", productHandler.CreateProduct)
	productRouter.GET("", productHandler.ListAllProduct)
	productRouter.GET("/:id", productHandler.ListProduct)
	productRouter.PUT("/update/:id", productHandler.UpdateProduct)
	productRouter.DELETE("/:id", productHandler.DeleteProduct)

}
