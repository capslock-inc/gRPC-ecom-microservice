package core

import (
	"context"
	"encoding/json"
	"log"

	cartmodel "github.com/capslock-inc/gprc-demo/Protos/cart"
	mongoclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/mongoclient"
	"github.com/capslock-inc/gprc-demo/logmodel"
)

type CartServer struct {
	cartmodel.UnimplementedCartServiceServer
	Mongorpc mongoclientmodel.MongoClientServiceClient
	Logs     *log.Logger
}

// rpc GetCartItem(userid) returns (cart);
// rpc CreateNewCart(userid) returns (status);
// rpc AddToCart(productid) returns (status);
// rpc DeleteCart(userid) returns (status);
// rpc DeleteCartItem(productid) returns (status);

func (s *CartServer) GetCartItem(ctx context.Context, data *cartmodel.Userid) (*cartmodel.Cart, error) {

	pid := &logmodel.Userid{
		UID: data.Id,
	}
	jdata, err := json.Marshal(pid)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("101-cartserver", "😥", err.Error(), "error marshalling json in GetCartItem"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("101-cartserver", "😃", string(jdata), "data retrived at GetCartItem"))

	response, err := s.Mongorpc.GetCartItemByUserId(
		ctx,
		&mongoclientmodel.Userid{
			Id: data.Id,
		},
	)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("101-cartserver", "😥", err.Error(), "error getting item to cart in GetCartItem rpc connection"))
		return nil, err
	}

	jsdata, err := json.Marshal(&cartmodel.Cart{
		Id:            data.Id,
		Productidlist: response.Productidlist,
	})
	if err != nil {
		s.Logs.Println(logmodel.Logdata("101-cartserver", "😥", err.Error(), "error marshalling json in GetCartItem"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("101-cartserver", "😃", string(jsdata), "successfully got items in cart at GetCartItem"))

	return &cartmodel.Cart{
		Id:            data.Id,
		Productidlist: response.Productidlist,
	}, nil
}
func (s *CartServer) CreateNewCart(ctx context.Context, data *cartmodel.Userid) (*cartmodel.Status, error) {

	pid := &logmodel.Userid{
		UID: data.Id,
	}
	jdata, err := json.Marshal(pid)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("103-cartserver", "😥", err.Error(), "error marshalling json in CreateNewCart"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("103-cartserver", "😃", string(jdata), "data retrived at CreateNewCart"))

	response, err := s.Mongorpc.CreateNewCart(ctx, &mongoclientmodel.Userid{
		Id: data.Id,
	})
	if err != nil {
		s.Logs.Println(logmodel.Logdata("103-cartserver", "😥", err.Error(), "error creating cart in CreateNewCart"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("103-cartserver", "😃", response.Value, "successfully created cart by CreateNewCart"))

	return &cartmodel.Status{
		Value: response.Value,
	}, nil
}
func (s *CartServer) AddToCart(ctx context.Context, data *cartmodel.Productid) (*cartmodel.Status, error) {

	pid := &logmodel.Productid{
		PID: data.Id,
		UID: data.Userid,
	}
	jdata, err := json.Marshal(pid)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("104-cartserver", "😥", err.Error(), "error marshalling json in AddToCart"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("104-cartserver", "😃", string(jdata), "data retrived at AddToCart"))

	response, err := s.Mongorpc.AddProductToCart(ctx, &mongoclientmodel.Addproduct{
		Userid:    data.Userid,
		Productid: data.Id,
	})
	if err != nil {
		s.Logs.Println(logmodel.Logdata("104-cartserver", "😥", err.Error(), "error adding item to cart in AddToCart"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("104-cartserver", "😃", response.Value, "successfully added item to cart at AddToCart"))

	return &cartmodel.Status{
		Value: response.Value,
	}, nil
}
func (s *CartServer) DeleteCart(ctx context.Context, data *cartmodel.Userid) (*cartmodel.Status, error) {

	pid := &logmodel.Userid{
		UID: data.Id,
	}
	jdata, err := json.Marshal(pid)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("102-cartserver", "😥", err.Error(), "error marshalling json in DeleteCart"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("102-cartserver", "😃", string(jdata), "data retrived at DeleteCart"))

	response, err := s.Mongorpc.DeleteCartByUserId(
		ctx,
		&mongoclientmodel.Userid{
			Id: data.Id,
		},
	)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("102-cartserver", "😥", err.Error(), "error deleting cart in DeleteCart"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("102-cartserver", "😃", response.Value, "successfully deleted cart by DeleteCart"))

	return &cartmodel.Status{
		Value: response.Value,
	}, nil
}
func (s *CartServer) DeleteCartItem(ctx context.Context, data *cartmodel.Productid) (*cartmodel.Status, error) {

	pid := &logmodel.Productid{
		PID: data.Id,
		UID: data.Userid,
	}
	jdata, err := json.Marshal(pid)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("105-cartserver", "😥", err.Error(), "error marshalling json in DeleteCartItem"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("105-cartserver", "😃", string(jdata), "data retrived at DeleteCartItem"))

	response, err := s.Mongorpc.DeleteCartItemByProductId(ctx, &mongoclientmodel.Addproduct{
		Userid:    data.Userid,
		Productid: data.Id,
	})
	if err != nil {
		s.Logs.Println(logmodel.Logdata("105-cartserver", "😥", err.Error(), "error deleting item in cart in DeleteCartItem"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("105-cartserver", "😃", response.Value, "successfully deleted item in cart at DeleteCartItem"))
	return &cartmodel.Status{
		Value: response.Value,
	}, nil
}
