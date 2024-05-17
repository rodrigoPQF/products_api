package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodrigoPQF/products_api/internal/domain/entity/products"
)

func (c *MysqlDB) CreateProduct(ctx context.Context, product products.ProductEntity) (*products.ProductEntity, error){
	conn := c.session.WithContext(ctx)

	id := uuid.NewString()
	product.Id = id

	if _, err := conn.SQL().InsertInto("products").Values(map[string]interface{}{
		"id": product.Id,
		"name": product.Name,
		"description": product.Description,
		"price": product.Price,
	}).Exec(); err != nil {
		return nil, err
	}

	return &product, nil
}