package productcore

import (
	"context"
	"log"

	product_model "github.com/capslock-inc/gprc-demo/ProductService/product-model"
	"github.com/gofrs/uuid"
)

type ProductServer struct {
	product_model.UnimplementedProductServiceServer
}

func (p *ProductServer) AddProduct(ctx context.Context, in *product_model.Product) (*product_model.ProductID, error) {
	log.Printf("received: %v", in.GetName())
	id, err := uuid.NewV4()
	if err != nil {
		return nil, error(err)
	}
	return &product_model.ProductID{Value: id.String()}, nil
}
