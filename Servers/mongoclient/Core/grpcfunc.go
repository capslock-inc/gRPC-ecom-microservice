package core

import (
	"context"
	"encoding/json"
	"log"

	mongoclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/mongoclient"
	model "github.com/capslock-inc/gprc-demo/Servers/mongoclient/Model"
	"github.com/capslock-inc/gprc-demo/logmodel"
	"go.mongodb.org/mongo-driver/mongo"
)

// rpc GetCartItemByUserId(logmodel.Userid) returns (cart);
// rpc CreateNewCart(logmodel.Userid) returns (status);
// rpc AddProductToCart(logmodel.Productid) returns (status);
// rpc DeleteCartByUserId(logmodel.Userid) returns (status);
// rpc DeleteCartItemByProductId(logmodel.Productid) returns (status);
// rpc CheckoutCart(emptyparam) returns (status);
// rpc Healthy(emptyparam) returns (status);

type MongoClientServer struct {
	mongoclientmodel.UnimplementedMongoClientServiceServer
	Client *mongo.Client
	Logs   *log.Logger
}

// rpc GetCartItemByUserId(logmodel.Userid) returns (cart);
func (s *MongoClientServer) GetCartItemByUserId(ctx context.Context, incomingdata *mongoclientmodel.Userid) (*mongoclientmodel.Cart, error) {
	uid := &logmodel.Userid{
		UID: incomingdata.GetId(),
	}
	jdata, err := json.Marshal(uid)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("101-mongoclientserver", "😥", err.Error(), "error marshalling json in GetCartItemByUserId"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("101-mongoclientserver", "😃", string(jdata), "data retrived at GetCartItemByUserId"))

	item, err := GetCartItemByUserId(incomingdata.GetId(), s.Client)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("101-mongoclientserver", "😥", err.Error(), "error getting data in GetCartItemByUserId helper function"))
		return nil, err
	}
	out, err := json.Marshal(&mongoclientmodel.Cart{Id: item.UserId, Productidlist: item.ProductId})
	if err != nil {
		s.Logs.Println(logmodel.Logdata("101-mongoclientserver", "😥", err.Error(), "error getting data in GetCartItemByUserId helper function"))
		return nil, err
	}
	s.Logs.Println(logmodel.Logdata("101-mongoclientserver", "😃", string(out), "successfully got cartitem by GetCartItemByUserId"))
	return &mongoclientmodel.Cart{Id: item.UserId, Productidlist: item.ProductId}, nil
}

// rpc CreateNewCart(logmodel.Userid) returns (status);
func (s *MongoClientServer) CreateNewCart(ctx context.Context, incomingdata *mongoclientmodel.Userid) (*mongoclientmodel.Status, error) {
	uid := &logmodel.Userid{
		UID: incomingdata.GetId(),
	}
	jdata, err := json.Marshal(uid)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("103-mongoclientserver", "😥", err.Error(), "error marshalling json in CreateNewCart"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("103-mongoclientserver", "😃", string(jdata), "data retrived at CreateNewCart"))

	arr, err := CreateNewCart(&model.CartItem{
		UserId:    incomingdata.GetId(),
		ProductId: []string{},
	}, s.Client)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("103-mongoclientserver", "😥", err.Error(), "error creating cart in CreateNewCart helper function"))
		return &mongoclientmodel.Status{Value: string("cant create new cart")}, err
	}

	s.Logs.Println(logmodel.Logdata("103-mongoclientserver", "😃", arr, "successfully created cart by CreateNewCart"))
	return &mongoclientmodel.Status{Value: arr}, nil
}

// rpc AddProductToCart(logmodel.Productid) returns (status);

func (s *MongoClientServer) AddProductToCart(ctx context.Context, incomingdata *mongoclientmodel.Addproduct) (*mongoclientmodel.Status, error) {

	pid := &logmodel.Productid{
		PID: incomingdata.Productid,
		UID: incomingdata.Userid,
	}
	jdata, err := json.Marshal(pid)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("104-mongoclientserver", "😥", err.Error(), "error marshalling json in AddProductToCart"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("104-mongoclientserver", "😃", string(jdata), "data retrived at AddProductToCart"))

	log.Println(incomingdata.GetProductid(), incomingdata.GetUserid())
	arr, err := AddProductToCart(incomingdata.GetUserid(), incomingdata.GetProductid(), s.Client)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("104-mongoclientserver", "😥", err.Error(), "error adding item to cart in AddProductToCart helper function"))
		return &mongoclientmodel.Status{Value: string("unable to add product")}, err
	}

	s.Logs.Println(logmodel.Logdata("104-mongoclientserver", "😃", arr, "successfully added item to the cart by AddProductToCart"))
	return &mongoclientmodel.Status{Value: arr}, nil

}

// rpc DeleteCartByUserId(logmodel.Userid) returns (status);
func (s *MongoClientServer) DeleteCartByUserId(ctx context.Context, incomingdata *mongoclientmodel.Userid) (*mongoclientmodel.Status, error) {

	uid := &logmodel.Userid{
		UID: incomingdata.GetId(),
	}
	jdata, err := json.Marshal(uid)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("102-mongoclientserver", "😥", err.Error(), "error marshalling json in DeleteCartByUserId"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("102-mongoclientserver", "😃", string(jdata), "data retrived at DeleteCartByUserId"))

	arr, err := DeleteCartByUserId(incomingdata.GetId(), s.Client)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("102-mongoclientserver", "😥", err.Error(), "error deleting cart in DeleteCartByUserId helper function"))
		return &mongoclientmodel.Status{Value: arr}, err
	}

	s.Logs.Println(logmodel.Logdata("102-mongoclientserver", "😃", arr, "successfully deleted cart by DeleteCartByUserId"))
	return &mongoclientmodel.Status{Value: arr}, nil

}

// rpc DeleteCartItemByProductId(logmodel.Productid) returns (status);
func (s *MongoClientServer) DeleteCartItemByProductId(ctx context.Context, incomingdata *mongoclientmodel.Addproduct) (*mongoclientmodel.Status, error) {

	pid := &logmodel.Productid{
		PID: incomingdata.Productid,
		UID: incomingdata.Userid,
	}
	jdata, err := json.Marshal(pid)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("105-mongoclientserver", "😥", err.Error(), "error marshalling json in DeleteCartItemByProductId"))
		return nil, err
	}

	s.Logs.Println(logmodel.Logdata("105-mongoclientserver", "😃", string(jdata), "data retrived at DeleteCartItemByProductId"))

	arr, err := DeleteCartItemByProductId(incomingdata.Userid, incomingdata.Productid, s.Client)
	if err != nil {
		s.Logs.Println(logmodel.Logdata("105-mongoclientserver", "😥", err.Error(), "error deleting cartitem in DeleteCartItemByProductId helper function"))
		log.Printf("error deleting product: %v", err)
		return &mongoclientmodel.Status{Value: arr}, err
	}

	s.Logs.Println(logmodel.Logdata("105-mongoclientserver", "😃", arr, "successfully deleted cart item by DeleteCartItemByProductId"))

	return &mongoclientmodel.Status{Value: arr}, nil

}
