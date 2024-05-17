package products

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/rodrigoPQF/products_api/internal/domain/entity/products"
	"github.com/rodrigoPQF/products_api/internal/infra/database"
)

type ProductsUseCase interface {
	Create(ctx context.Context, input *CreateProductsInputDto) (*CreateProductsOutputDto, error)
	Update(ctx context.Context, input *UpdateProductsInputDto) (*UpdateProductsOutputDto, error)
	Get(ctx context.Context, input *GetProductInputDto) (*GetProductOutputDto, error)
	Delete(ctx context.Context, input *DeleteProductInputDto) (*DeleteProductOutputDto, error)
	GetAll(ctx context.Context) (*[]GetProductOutputDto, error)
}

type productsUseCase struct {
	ProductRepository database.Repository
}

func NewProductsUseCase(repository database.Repository) ProductsUseCase {
	return &productsUseCase{
		ProductRepository: repository,
	}
}

func (pr *productsUseCase) Create(ctx context.Context, input *CreateProductsInputDto) (*CreateProductsOutputDto, error){
	res, err := pr.ProductRepository.CreateProduct(ctx, products.ProductEntity{
		Name: input.Name,
		Price: float64(input.Price),
		Description: input.Description,
	})

	if err != nil {
		return nil, err
	}

	var output CreateProductsOutputDto

	copier.CopyWithOption(&output, &res, copier.Option{
		IgnoreEmpty: true,
	})
	return &output, nil	
}

func (pr *productsUseCase) Update(ctx context.Context, input *UpdateProductsInputDto) (*UpdateProductsOutputDto, error){


	res, err := pr.ProductRepository.UpdateProduct(ctx, input.Id , products.ProductEntity{
		Name: input.Name,
		Price: float64(input.Price),
		Description: input.Description,
	})

	if err != nil {
		return nil, err
	}

	var output UpdateProductsOutputDto
	
	copier.CopyWithOption(&output, &res, copier.Option{
		IgnoreEmpty: true,
	})

	return &output, nil	
}

func (pr *productsUseCase) Get(ctx context.Context, input *GetProductInputDto) (*GetProductOutputDto, error){
	res, err := pr.ProductRepository.GetProduct(ctx, input.Id)

	if err != nil {
		return nil, err
	}

	var output GetProductOutputDto
	
	copier.CopyWithOption(&output, &res, copier.Option{
		IgnoreEmpty: true,
	})

	return &output, nil	
}

func (pr *productsUseCase) GetAll(ctx context.Context) (*[]GetProductOutputDto, error){
	res, err := pr.ProductRepository.GetAllProduct(ctx)

	if err != nil {
		return nil, err
	}

	var output []GetProductOutputDto
	
	copier.CopyWithOption(&output, &res, copier.Option{
		IgnoreEmpty: true,
	})

	return &output, nil	
}

func (pr *productsUseCase) 	Delete(ctx context.Context, input *DeleteProductInputDto) (*DeleteProductOutputDto, error){
	res, err := pr.ProductRepository.DeleteProduct(ctx, input.Id)

	if err != nil {
		return nil, err
	}

	var output DeleteProductOutputDto
	
	copier.CopyWithOption(&output, &res, copier.Option{
		IgnoreEmpty: true,
	})

	return &output, nil	
}
