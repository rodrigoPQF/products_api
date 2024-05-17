package database

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/rodrigoPQF/products_api/internal/domain/entity/products"
	"github.com/rodrigoPQF/products_api/internal/infra/database/models"
)

func (c *MysqlDB) GetProduct(ctx context.Context, productId string) (*products.ProductEntity, error){
	var product products.ProductEntity
	var productModel models.Product
	conn := c.session.WithContext(ctx)


	if err := conn.SQL().SelectFrom("products").Where("id = ?", productId).One(&productModel); err != nil {
		return nil, err
	}

	copier.Copy(&product, productModel)

	return  &product, nil
}