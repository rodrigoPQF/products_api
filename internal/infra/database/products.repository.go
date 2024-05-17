package database

import (
	"context"

	"github.com/rodrigoPQF/products_api/internal/domain/entity/products"
)

type RepositoryTransaction func(Repository) error

type Repository interface {
	Tx(RepositoryTransaction) error
	
	CreateProduct(ctx context.Context, product products.ProductEntity ) (*products.ProductEntity, error)
	DeleteProduct(ctx context.Context, productId string) (string, error)
	GetProduct(ctx context.Context, productId string) (*products.ProductEntity, error)
	UpdateProduct(ctx context.Context, productId string, product products.ProductEntity) (*products.ProductEntity, error)
	GetAllProduct(ctx context.Context) (*[]products.ProductEntity, error)
}
