package core

import (
	"context"
	"log"

	cartmodel "github.com/capslock-inc/gprc-demo/Protos/cart"
	mongoclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/mongoclient"
)

type CartServer struct {
	cartmodel.UnimplementedCartServiceServer
	Mongorpc mongoclientmodel.MongoClientServiceClient
}

// rpc GetCartItem(userid) returns (cart);
// rpc CreateNewCart(userid) returns (status);
// rpc AddToCart(productid) returns (status);
// rpc DeleteCart(userid) returns (status);
// rpc DeleteCartItem(productid) returns (status);

func (s *CartServer) GetCartItem(ctx context.Context, data *cartmodel.Userid) (*cartmodel.Cart, error) {
	log.Printf("recived data : %s", data.GetId())
	response, err := s.Mongorpc.GetCartItemByUserId(
		ctx,
		&mongoclientmodel.Userid{
			Id: data.Id,
		},
	)
	if err != nil {
		return nil, err
	}
	return &cartmodel.Cart{
		Productidlist: response.Productidlist,
	}, nil
}
func (s *CartServer) CreateNewCart(ctx context.Context, data *cartmodel.Userid) (*cartmodel.Status, error) {
	response, err := s.Mongorpc.CreateNewCart(ctx, &mongoclientmodel.Userid{
		Id: data.Id,
	})
	if err != nil {
		return nil, err
	}
	return &cartmodel.Status{
		Value: response.Value,
	}, nil
}
func (s *CartServer) AddToCart(ctx context.Context, data *cartmodel.Productid) (*cartmodel.Status, error) {
	response, err := s.Mongorpc.AddProductToCart(ctx, &mongoclientmodel.Addproduct{
		Userid:    data.Userid,
		Productid: data.Id,
	})
	if err != nil {
		return nil, err
	}
	return &cartmodel.Status{
		Value: response.Value,
	}, nil
}
func (s *CartServer) DeleteCart(ctx context.Context, data *cartmodel.Userid) (*cartmodel.Status, error) {
	response, err := s.Mongorpc.DeleteCartByUserId(
		ctx,
		&mongoclientmodel.Userid{
			Id: data.Id,
		},
	)
	if err != nil {
		return nil, err
	}
	return &cartmodel.Status{
		Value: response.Value,
	}, nil
}
func (s *CartServer) DeleteCartItem(ctx context.Context, data *cartmodel.Productid) (*cartmodel.Status, error) {
	response, err := s.Mongorpc.DeleteCartItemByProductId(ctx, &mongoclientmodel.Addproduct{
		Userid:    data.Userid,
		Productid: data.Id,
	})
	if err != nil {
		return nil, err
	}
	return &cartmodel.Status{
		Value: response.Value,
	}, nil
}
