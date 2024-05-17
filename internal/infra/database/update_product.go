package database

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/rodrigoPQF/products_api/internal/domain/entity/products"
	"github.com/rodrigoPQF/products_api/internal/infra/database/models"
)

func (c *MysqlDB) UpdateProduct(ctx context.Context, productId string, product products.ProductEntity) (*products.ProductEntity, error){
	conn := c.session.WithContext(ctx)
	var productModel models.Product

	if err := conn.SQL().SelectFrom("products").Where("id = ?", productId).One(&productModel); err != nil {
		return nil, errors.New("not found product id")
	}

	copier.Copy(&productModel, product)

	if _, err := conn.SQL().Update("products").Set(&productModel).Where("id = ?", productId).Exec(); err != nil {
		return nil, err
	}

	return  &product, nil
}