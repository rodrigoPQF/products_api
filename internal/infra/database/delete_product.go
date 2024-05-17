package database

import (
	"context"
	"errors"

	"github.com/rodrigoPQF/products_api/internal/infra/database/models"
)

func (c *MysqlDB) DeleteProduct(ctx context.Context, productId string) (string, error){
	conn := c.session.WithContext(ctx)
	var productModel models.Product

	
	if err := conn.SQL().SelectFrom("products").Where("id = ?", productId).One(&productModel); err != nil {
		return "", errors.New("not found product id")
	}

	if _ ,err := conn.SQL().DeleteFrom("products").Where("id = ?", productModel.Id).Exec(); err != nil {
		return "", err
	}

	return "deleted product with success", nil
}