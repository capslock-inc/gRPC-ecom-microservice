package core

import (
	"context"
	"log"

	postgresclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/postgresclient"
	productmodel "github.com/capslock-inc/gprc-demo/Protos/product"
)

type ProductServer struct {
	productmodel.UnimplementedProductServiceServer
	PostgresRPC postgresclientmodel.PostgresClientServiceClient
}

// rpc GetAllProduct(emptyparam) returns (productlist);
// rpc GetProductById(productid) returns (product);
// rpc AddNewProduct(product) returns (status);
// rpc UpdateProduct (product) returns (status);
// rpc DeleteProduct(productid) returns (status);

func (s *ProductServer) GetAllProduct(ctx context.Context, _ *productmodel.Emptyparam) (*productmodel.Productlist, error) {
	result, err := s.PostgresRPC.AllProduct(ctx, &postgresclientmodel.Emptyparam{})
	if err != nil {
		log.Println("error getting all product")
		return nil, err
	}
	var list []*productmodel.Product
	for _, l := range result.List {
		list = append(list, &productmodel.Product{
			Id:          l.Id,
			Name:        l.Name,
			Price:       l.Price,
			Description: l.Description,
		})
	}
	log.Println("getallproduct sucessfull 😇")
	return &productmodel.Productlist{
		List: list,
	}, nil
}

func (s *ProductServer) GetProductById(ctx context.Context, id *productmodel.Productid) (*productmodel.Product, error) {
	result, err := s.PostgresRPC.ProductById(
		ctx,
		&postgresclientmodel.Productid{
			Id: id.Id,
		},
	)
	if err != nil {
		log.Println("error getproductbyid")
		return nil, err
	}
	log.Println("getproductbyid sucessful 😇")
	return &productmodel.Product{
		Id:          result.Id,
		Name:        result.Name,
		Price:       result.Price,
		Description: result.Description,
	}, nil

}

func (s *ProductServer) AddNewProduct(ctx context.Context, data *productmodel.Product) (*productmodel.Status, error) {
	_, err := s.PostgresRPC.AddProduct(
		ctx,
		&postgresclientmodel.Product{
			Id:          data.Id,
			Name:        data.Name,
			Price:       data.Price,
			Description: data.Description,
		},
	)
	if err != nil {
		log.Println("error addnewproduct")
		return nil, err
	}
	log.Println("addnewproduct sucessful 😇")
	return &productmodel.Status{
		Value: "✔️ added successfully 😃",
	}, nil
}

func (s *ProductServer) UpdateProduct(ctx context.Context, data *productmodel.Product) (*productmodel.Status, error) {
	_, err := s.PostgresRPC.UpdateProduct(
		ctx,
		&postgresclientmodel.Product{
			Id:          data.Id,
			Name:        data.Name,
			Price:       data.Price,
			Description: data.Description,
		},
	)
	if err != nil {
		log.Println("error updateproduct")
		return nil, err
	}
	log.Println("updateproduct sucessful 😇")
	return &productmodel.Status{
		Value: "updated successfully 😃",
	}, nil
}

func (s *ProductServer) DeleteProduct(ctx context.Context, data *productmodel.Productid) (*productmodel.Status, error) {
	_, err := s.PostgresRPC.DeleteProduct(
		ctx,
		&postgresclientmodel.Productid{
			Id: data.Id,
		},
	)
	if err != nil {
		log.Println("error deleteproduct")
		return nil, err
	}
	log.Println("deleteproduct successful 😇")
	return &productmodel.Status{
		Value: "deleted successfully 😃",
	}, nil
}
