package database

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/rodrigoPQF/products_api/internal/domain/entity/products"
	"github.com/rodrigoPQF/products_api/internal/infra/database/models"
)

func (c *MysqlDB) GetAllProduct(ctx context.Context) (*[]products.ProductEntity, error){
	var product []products.ProductEntity
	var productModel []models.Product
	conn := c.session.WithContext(ctx)


	if err := conn.SQL().SelectFrom("products").All(&productModel); err != nil {
		return nil, err
	}

	copier.CopyWithOption(&product, &productModel, copier.Option{ DeepCopy: true})

	return  &product, nil
}