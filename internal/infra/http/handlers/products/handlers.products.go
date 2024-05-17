package products

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rodrigoPQF/products_api/internal/infra/http/handlers"
	"github.com/rodrigoPQF/products_api/internal/usecase/products"
	"github.com/rodrigoPQF/products_api/pkg/validators"
)

type ProductsHandlers interface {
	CreateProduct(ctx echo.Context) error
	ListProduct(ctx echo.Context) error
	DeleteProduct(ctx echo.Context) error
	UpdateProduct(ctx echo.Context) error
	ListAllProduct(ctx echo.Context) error 
}

type ProductHandler struct {
	productsUseCase products.ProductsUseCase
}

func NewProductsHandler(productsUseCase products.ProductsUseCase) ProductsHandlers {
	return &ProductHandler{
		productsUseCase: productsUseCase,
	}
}

func (ph *ProductHandler) CreateProduct(ctx echo.Context) error {
	var input products.CreateProductsInputDto

	if err := ctx.Bind(&input); err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	if err := validators.Validate(&input); err != nil {

		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	output, err := ph.productsUseCase.Create(ctx.Request().Context(), &input)
	if err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	return handlers.NewSuccessResponse(ctx, http.StatusCreated, "success register product", output)
}

func (ph *ProductHandler) ListAllProduct(ctx echo.Context) error {

	output, err := ph.productsUseCase.GetAll(ctx.Request().Context())
	if err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	return handlers.NewSuccessResponse(ctx, http.StatusOK, "list product", output)
}

func (ph *ProductHandler) ListProduct(ctx echo.Context) error {
	var input products.GetProductInputDto

	input.Id = ctx.Param("id")

	if err := validators.Validate(&input); err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	output, err := ph.productsUseCase.Get(ctx.Request().Context(), &input)
	if err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	return handlers.NewSuccessResponse(ctx, http.StatusOK, "list product", output)
}

func (ph *ProductHandler) DeleteProduct(ctx echo.Context) error {
	var input products.DeleteProductInputDto

	input.Id = ctx.Param("id")

	if err := validators.Validate(&input); err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	output, err := ph.productsUseCase.Delete(ctx.Request().Context(), &input)
	if err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	return handlers.NewSuccessResponse(ctx, http.StatusOK, "success delete product", output)
}

func (ph *ProductHandler) UpdateProduct(ctx echo.Context) error {
	var input products.UpdateProductsInputDto


	if err := ctx.Bind(&input); err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	input.Id = ctx.Param("id")

	if err := validators.Validate(&input); err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	output, err := ph.productsUseCase.Update(ctx.Request().Context(), &input)
	if err != nil {
		return handlers.NewErrorReponse(ctx, http.StatusBadRequest, err.Error())
	}

	return handlers.NewSuccessResponse(ctx, http.StatusOK, "success update product", output)
}