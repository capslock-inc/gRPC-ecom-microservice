package core

import (
	"context"
	"log"

	mongoclientmodel "github.com/capslock-inc/gprc-demo/Protos/database/mongoclient"
	model "github.com/capslock-inc/gprc-demo/Servers/mongoclient/Model"
	"go.mongodb.org/mongo-driver/mongo"
)

// rpc GetCartItemByUserId(userid) returns (cart);
// rpc CreateNewCart(userid) returns (status);
// rpc AddProductToCart(productid) returns (status);
// rpc DeleteCartByUserId(userid) returns (status);
// rpc DeleteCartItemByProductId(productid) returns (status);
// rpc CheckoutCart(emptyparam) returns (status);
// rpc Healthy(emptyparam) returns (status);

type MongoClientServer struct {
	mongoclientmodel.UnimplementedMongoClientServiceServer
	Client *mongo.Client
}

// rpc GetCartItemByUserId(userid) returns (cart);
func (s *MongoClientServer) GetCartItemByUserId(ctx context.Context, incomingdata *mongoclientmodel.Userid) (*mongoclientmodel.Cart, error) {
	log.Printf("received data: %s", incomingdata.GetId())
	item, err := GetCartItemByUserId(incomingdata.GetId(), s.Client)
	if err != nil {
		log.Printf("error geting cart : %v", err)
	}
	return &mongoclientmodel.Cart{Id: item.UserId, Productidlist: item.ProductId}, nil
}

// rpc CreateNewCart(userid) returns (status);
func (s *MongoClientServer) CreateNewCart(ctx context.Context, incomingdata *mongoclientmodel.Userid) (*mongoclientmodel.Status, error) {
	log.Println(incomingdata.GetId())
	arr, err := CreateNewCart(&model.CartItem{
		UserId:    incomingdata.GetId(),
		ProductId: []string{},
	}, s.Client)
	if err != nil {
		return &mongoclientmodel.Status{Value: string("cant create new cart")}, err
	}
	return &mongoclientmodel.Status{Value: arr}, nil
}

// rpc AddProductToCart(productid) returns (status);

func (s *MongoClientServer) AddProductToCart(ctx context.Context, incomingdata *mongoclientmodel.Addproduct) (*mongoclientmodel.Status, error) {
	log.Println(incomingdata.GetProductid(), incomingdata.GetUserid())
	arr, err := AddProductToCart(incomingdata.GetUserid(), incomingdata.GetProductid(), s.Client)
	if err != nil {
		return &mongoclientmodel.Status{Value: string("unable to add product")}, err
	}
	return &mongoclientmodel.Status{Value: arr}, nil
}

// rpc DeleteCartByUserId(userid) returns (status);
func (s *MongoClientServer) DeleteCartByUserId(ctx context.Context, incomingdata *mongoclientmodel.Userid) (*mongoclientmodel.Status, error) {
	arr, err := DeleteCartByUserId(incomingdata.GetId(), s.Client)
	if err != nil {
		log.Printf("error deleting userid: %v", err)
		return &mongoclientmodel.Status{Value: arr}, err
	}
	return &mongoclientmodel.Status{Value: arr}, nil
}

// rpc DeleteCartItemByProductId(productid) returns (status);
func (s *MongoClientServer) DeleteCartItemByProductId(ctx context.Context, incomingdata *mongoclientmodel.Addproduct) (*mongoclientmodel.Status, error) {
	arr, err := DeleteCartItemByProductId(incomingdata.Userid, incomingdata.Productid, s.Client)
	if err != nil {
		log.Printf("error deleting product: %v", err)
		return &mongoclientmodel.Status{Value: arr}, err
	}
	return &mongoclientmodel.Status{Value: arr}, nil
}
